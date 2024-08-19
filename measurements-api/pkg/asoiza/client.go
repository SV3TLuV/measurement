package asoiza

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"io"
	"measurements-api/internal/model"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	baseUrl           = "https://asoiza.voeikovmgo.ru"
	sessionCookieName = "PHPSESSID"
)

var _ Client = (*client)(nil)

type client struct {
	client    http.Client
	sessionID *string
	loginWg   sync.WaitGroup
	factory   ConfigurationFactory
}

func NewClient(factory ConfigurationFactory) *client {
	return &client{
		client:  http.Client{},
		factory: factory,
		loginWg: sync.WaitGroup{},
	}
}

func (c *client) GetMeasurements(
	ctx context.Context,
	objectID uint64,
	limit int,
	offset int) ([]*Measurement, error) {
	year, month, day := time.Now().Date()
	endOfDay := time.Date(year, month, day, 23, 59, 59, 0, time.UTC)

	queryParams := url.Values{}
	queryParams.Add("task", "read")
	queryParams.Add("mtype", "auto")
	queryParams.Add("period", "month")
	queryParams.Add("end", endOfDay.Format("2006-01-02 15:04:05"))
	queryParams.Add("titletype", "long")
	queryParams.Add("pdktype", "sanpin21")
	queryParams.Add("presstype", "gpa")
	queryParams.Add("object", strconv.FormatUint(objectID, 10))
	queryParams.Add("page", "1")
	queryParams.Add("start", strconv.Itoa(offset))
	queryParams.Add("limit", strconv.Itoa(limit))
	queryParams.Add("sort", "[{\"property\":\"date_time\",\"direction\":\"DESC\"},{\"property\":\"obj_name\",\"direction\":\"ASC\"}]")

	requestUrl := fmt.Sprintf("%s/data/MeasurementsStore.php?%s", baseUrl, queryParams.Encode())
	request, err := http.NewRequestWithContext(ctx, "GET", requestUrl, nil)
	if err != nil {
		return nil, errors.Wrap(err, "create model")
	}

	request.Header.Set("Content-Type", "application/json")
	bytes, err := c.baseRequestWithReAuth(ctx, request, 1)
	if err != nil {
		return nil, errors.Wrap(err, "do model")
	}

	regex := regexp.MustCompile("\"v_\\d{6,}\":\"\\d{1,},\\d{1,}\"")
	jsonStr := regex.ReplaceAllStringFunc(string(bytes), func(m string) string {
		return strings.Replace(m, ",", ".", 1)
	})

	var response MeasurementResponse
	err = json.Unmarshal([]byte(jsonStr), &response)
	if err != nil || response.Data == nil {
		return make([]*Measurement, 0), nil
	}

	for _, measurement := range response.Data {
		if measurement.WindDir != nil {
			winDirStr := WindDirToRmb16(*measurement.WindDir)
			measurement.WindDirStr = &winDirStr
		}
		if measurement.DateTime != nil {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", *measurement.DateTime)
			if err != nil {
				return nil, errors.Wrap(err, "parse date_time")
			}

			measurement.RealDateTime = &parsedTime
		}
	}

	return response.Data, nil
}

func (c *client) GetNewestMeasurement(ctx context.Context, objectID uint64) (*Measurement, error) {
	measurements, err := c.GetMeasurements(ctx, objectID, 1, 0)
	if err != nil {
		return nil, errors.Wrap(err, "get measurements")
	}
	if len(measurements) == 0 {
		return nil, nil
	}
	return measurements[0], nil
}

func (c *client) GetNavTree(ctx context.Context, node string) ([]*Node, error) {
	requestUrl := fmt.Sprintf("%s/data/NavTreeStore.php?task=read&sort=title&node=%s", baseUrl, node)
	request, err := http.NewRequestWithContext(ctx, "GET", requestUrl, nil)
	if err != nil {
		return nil, errors.Wrap(err, "create model")
	}

	request.Header.Set("Content-Type", "application/json")
	bytes, err := c.baseRequestWithReAuth(ctx, request, 1)
	if err != nil {
		return nil, errors.Wrap(err, "do model")
	}

	var response NavStoreResponse
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}

	if response.Data == nil {
		return make([]*Node, 0), nil
	}

	return response.Data, nil
}

func (c *client) login(ctx context.Context) error {
	configuration := c.factory.Get()
	if configuration == nil {
		return errors.New("configuration is nil")
	}

	requestUrl := fmt.Sprintf("%s/data/auth.php?action=login", baseUrl)
	formData := url.Values{}
	formData.Add("username", configuration.Username)
	formData.Add("password", configuration.Password)
	encodedFormData := formData.Encode()
	request, err := http.NewRequestWithContext(ctx, "POST", requestUrl, strings.NewReader(encodedFormData))
	if err != nil {
		return errors.Wrap(err, "create request")
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r, err := c.client.Do(request)
	if err != nil {
		return errors.Wrap(err, "base model")
	}

	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == sessionCookieName {
			c.sessionID = &cookie.Value
			break
		}
	}

	if c.sessionID == nil {
		return errors.Wrap(model.NotFound, "session id")
	}

	return nil
}

func (c *client) sessionIsAlive(ctx context.Context) (bool, error) {
	requestUrl := fmt.Sprintf("%s/data/auth.php?action=check&timeoffset=-180", baseUrl)
	request, err := http.NewRequestWithContext(ctx, "GET", requestUrl, nil)
	if err != nil {
		return false, errors.Wrap(err, "create model")
	}

	request.Header.Set("Content-Type", "application/json")
	r, err := c.client.Do(request)
	if err != nil {
		return false, errors.Wrap(err, "do model")
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return false, errors.Wrap(err, "read response body")
	}

	var response BaseResponse
	if err = json.Unmarshal(bytes, &response); err != nil {
		return false, errors.Wrap(err, "unmarshal")
	}

	return response.Success, nil
}

func (c *client) baseRequestWithReAuth(
	ctx context.Context,
	request *http.Request,
	attemptCount uint) ([]byte, error) {
	if attemptCount == 0 {
		return nil, errors.New("number of attempts has been exceeded")
	}

	c.loginWg.Wait()

	if c.sessionID == nil {
		err := func() error {
			c.loginWg.Add(1)
			defer c.loginWg.Done()
			return c.login(ctx)
		}()
		if err != nil {
			return nil, errors.Wrap(err, "login")
		}
	}

	request.AddCookie(&http.Cookie{Name: sessionCookieName, Value: *c.sessionID})

	r, err := c.client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "do model")
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}

	var response BaseResponse
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}

	if !response.Success {
		isAlive, err := c.sessionIsAlive(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "check session is alive")
		}

		if !isAlive {
			c.sessionID = nil
		}

		time.Sleep(5 * time.Second)
		return c.baseRequestWithReAuth(ctx, request, attemptCount-1)
	}

	return bytes, nil
}

package token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"measurements-api/internal/config"
	"measurements-api/internal/model"
	def "measurements-api/internal/service"
	"time"
)

var _ def.TokenService = (*service)(nil)

type service struct {
	config *config.JwtConfig
}

func NewService(config *config.JwtConfig) *service {
	return &service{
		config: config,
	}
}

func (s *service) CreateAccessToken(claims *model.Claims) (*string, error) {
	claims.Audience = []string{s.config.Audience}
	claims.Issuer = s.config.Issuer
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.ExpiresAt = jwt.NewNumericDate(claims.IssuedAt.Add(time.Hour))
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(s.config.Key)
	if err != nil {
		return nil, errors.Wrap(err, "signed string")
	}

	return &token, nil
}

func (s *service) CreateRefreshToken(claims *model.Claims) (*string, error) {
	claims.Audience = []string{s.config.Audience}
	claims.Issuer = s.config.Issuer
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.ExpiresAt = jwt.NewNumericDate(claims.IssuedAt.Add(time.Hour * 24 * 30))
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(s.config.Key)
	if err != nil {
		return nil, errors.Wrap(err, "signed string")
	}

	return &token, nil
}

func (s *service) GetClaims(token string) (*model.Claims, error) {
	var claims model.Claims
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return s.config.Key, nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "invalid token")
	}

	return &claims, nil
}

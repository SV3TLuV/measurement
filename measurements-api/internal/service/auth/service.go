package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"measurements-api/internal/model"
	"measurements-api/internal/repository"
	def "measurements-api/internal/service"
	"time"
)

var _ def.AuthService = (*service)(nil)

type service struct {
	userRepo        repository.UserRepository
	sessionRepo     repository.SessionRepository
	tokenService    def.TokenService
	passwordService def.PasswordService
}

func NewService(
	userRepo repository.UserRepository,
	sessionRepo repository.SessionRepository,
	tokenService def.TokenService,
	passwordService def.PasswordService) *service {
	return &service{
		userRepo:        userRepo,
		sessionRepo:     sessionRepo,
		tokenService:    tokenService,
		passwordService: passwordService,
	}
}

func (u *service) Login(ctx context.Context, request *model.LoginData) (*model.AuthResult, error) {
	user, err := u.userRepo.GetByLogin(ctx, request.Login)
	if err != nil {
		return nil, errors.Wrap(err, "search user by login")
	}
	if !u.passwordService.CheckPasswordHash(user.Password, request.Password) {
		return nil, errors.New("incorrect login or password")
	}

	session := &model.Session{
		ID:           uuid.New(),
		UserID:       user.ID,
		Created:      time.Now(),
		Updated:      nil,
		RefreshToken: "",
	}

	claims := &model.Claims{
		ClientID: session.UserID,
		SID:      session.ID,
		Role:     user.Role.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: user.Login,
		},
	}

	accessToken, err := u.tokenService.CreateAccessToken(claims)
	if err != nil {
		return nil, errors.Wrap(err, "generate access token")
	}

	refreshToken, err := u.tokenService.CreateRefreshToken(claims)
	if err != nil {
		return nil, errors.Wrap(err, "generate refresh token")
	}

	session.RefreshToken = *refreshToken
	err = u.sessionRepo.Save(ctx, session)
	if err != nil {
		return nil, errors.Wrap(err, "save session")
	}

	return &model.AuthResult{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}, nil
}

func (u *service) Logout(ctx context.Context, token *jwt.Token) error {
	claims, ok := token.Claims.(*model.Claims)
	if !ok {
		return errors.New("get claims")
	}

	_, err := u.sessionRepo.GetById(ctx, claims.SID)
	if err != nil {
		return errors.Wrap(err, "token invalid")
	}

	err = u.sessionRepo.DeleteOne(ctx, claims.SID)
	if err != nil {
		return errors.Wrap(err, "logout")
	}

	return nil
}

func (u *service) Refresh(ctx context.Context, token string) (*model.AuthResult, error) {
	claims, err := u.tokenService.GetClaims(token)
	if err != nil {
		return nil, errors.Wrap(err, "get claims")

	}

	session, err := u.sessionRepo.GetById(ctx, claims.SID)
	if err != nil {
		return nil, errors.Wrap(err, "get session")
	}

	if session.RefreshToken != token {
		return nil, errors.Wrap(err, "token invalid")
	}

	user, err := u.userRepo.GetById(ctx, session.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "get user")
	}

	claims.Role = user.Role.Name
	claims.RegisteredClaims.Subject = user.Login
	accessToken, err := u.tokenService.CreateAccessToken(claims)
	if err != nil {
		return nil, errors.Wrap(err, "generate access token")
	}

	refreshToken, err := u.tokenService.CreateRefreshToken(claims)
	if err != nil {
		return nil, errors.Wrap(err, "generate refresh token")
	}

	updated := time.Now()
	session.Updated = &updated
	session.RefreshToken = *refreshToken
	err = u.sessionRepo.Save(ctx, session)
	if err != nil {
		return nil, errors.Wrap(err, "save session")
	}

	return &model.AuthResult{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}, nil
}

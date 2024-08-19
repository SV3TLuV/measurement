package password

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	def "measurements-api/internal/service"
)

var _ def.PasswordService = (*service)(nil)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) HashPassword(password string) (*string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "hash password")
	}
	hashStr := string(hash)
	return &hashStr, nil
}

func (s *service) CheckPasswordHash(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

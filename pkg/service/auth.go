package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/Liopun/notes-app"
	"github.com/Liopun/notes-app/pkg/repository"
	"github.com/golang-jwt/jwt/v4"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo         repository.Authorization
	passwordSalt string
	signingKey   string
	tokenTTL     time.Duration
}

func NewAuthService(repo repository.Authorization, passwordSalt, signingKey string, tokenTTL time.Duration) *AuthService {
	return &AuthService{
		repo:         repo,
		passwordSalt: passwordSalt,
		signingKey:   signingKey,
		tokenTTL:     tokenTTL,
	}
}

func (s *AuthService) CreateUser(user notes.User) (int, error) {
	user.Password = generateHashedPasswword(user.Password, s.passwordSalt)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generateHashedPasswword(password, s.passwordSalt))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		user.Id,
	})

	return token.SignedString([]byte(s.signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(s.signingKey), nil
	})
	if err != nil {
		return -1, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return -1, errors.New("token claims invalid type")
	}

	return claims.UserId, nil
}

func generateHashedPasswword(password, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

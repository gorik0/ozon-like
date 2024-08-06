package jwter

import (
	uuid "github.com/satori/go.uuid"
	"ozone/internal/pkg/config"
	"time"
)

type JWTer interface {
	EncodeAuthToken(uuid.UUID) (string, time.Time, error)
	DecodeAuthToken(string) (uuid.UUID, error)
	EncodeCSRFToken(string) (string, time.Time, error)
	DecodeCSRFToken(string) (string, error)
}

type jwtManager struct {
	issuer string
	ttl    time.Duration
	secret string
}

func NewJWTManager(cfg *config.Config) JWTer {
	return &jwtManager{
		issuer: cfg.GetIssuer(),
		ttl:    cfg.GetTTL(),
		secret: cfg.GetSecret(),
	}
}

func (j *jwtManager) EncodeAuthToken(u uuid.UUID) (string, time.Time, error) {
	//TODO implement me
	panic("implement me")
}

func (j *jwtManager) DecodeAuthToken(s string) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (j *jwtManager) EncodeCSRFToken(s string) (string, time.Time, error) {
	//TODO implement me
	panic("implement me")
}

func (j *jwtManager) DecodeCSRFToken(s string) (string, error) {
	//TODO implement me
	panic("implement me")
}

var _ JWTer = &jwtManager{}

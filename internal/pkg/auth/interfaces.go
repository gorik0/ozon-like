package auth

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"ozone/internal/models"
	"time"
)

type AuthUsecase interface {
	SignIn(context.Context, *models.SignInPayload) (*models.Profile, string, time.Time, error)
	SignUp(context.Context, *models.SignUpPayload) (*models.Profile, string, time.Time, error)
	CheckAuth(context.Context, uuid.UUID) (*models.Profile, error)
}

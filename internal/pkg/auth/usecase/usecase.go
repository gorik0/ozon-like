package usecase

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"ozone/internal/models"
	"ozone/internal/pkg/auth"
	"ozone/internal/pkg/config"
	"ozone/internal/pkg/profile/repo"
	"ozone/internal/utils/jwter"
	"time"
)

type AuthUsecase struct {
	repo    *repo.ProfileRepo
	authJWT jwter.JWTer
}

func NewAuthUsecase(repo *repo.ProfileRepo, cfg *config.Config) *AuthUsecase {
	return &AuthUsecase{
		repo:    repo,
		authJWT: jwter.NewJWTManager(cfg),
	}
}
func (a *AuthUsecase) SignIn(ctx context.Context, payload *models.SignInPayload) (*models.Profile, string, time.Time, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthUsecase) SignUp(ctx context.Context, payload *models.SignUpPayload) (*models.Profile, string, time.Time, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthUsecase) CheckAuth(ctx context.Context, uuid uuid.UUID) (*models.Profile, error) {
	//TODO implement me
	panic("implement me")
}

var _ auth.AuthUsecase = &AuthUsecase{}

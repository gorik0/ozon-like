package usecase

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"ozone/internal/models"
	"ozone/internal/pkg/config"
	"ozone/internal/pkg/profile"
	"ozone/internal/pkg/profile/repo"
)

type ProfileUsecase struct {
	repo      profile.ProfileRepo
	photoPath string
}

func (p *ProfileUsecase) GetProfile(ctx context.Context, uuid uuid.UUID) (*models.Profile, error) {
	p.repo.ReadProfile(ctx, uuid)
}

func (p *ProfileUsecase) UpdatePhoto(ctx context.Context, userID uuid.UUID, filePhotoByte []byte, fileType string) (*models.Profile, error) {
	p.repo.UpdatePhoto(ctx, userID)
}

func (p *ProfileUsecase) UpdateData(ctx context.Context, uuid uuid.UUID, payload *models.UpdateProfileDataPayload) (*models.Profile, error) {
	//TODO implement me
	panic("implement me")
}

var _ profile.ProfileUsecase = &ProfileUsecase{}

func NewProfileUseCase(repo *repo.ProfileRepo, cfg *config.Config) *ProfileUsecase {

	return &ProfileUsecase{
		repo:      repo,
		photoPath: cfg.PhotosFilesPath,
	}

}

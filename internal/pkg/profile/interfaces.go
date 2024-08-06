package profile

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"ozone/internal/models"
)

type ProfileUsecase interface {
	GetProfile(context.Context, uuid.UUID) (*models.Profile, error)
	UpdatePhoto(ctx context.Context, userID uuid.UUID, filePhotoByte []byte, fileType string) (*models.Profile, error)
	UpdateData(context.Context, uuid.UUID, *models.UpdateProfileDataPayload) (*models.Profile, error)
}

type ProfileRepo interface {
	CreateProfile(context.Context, *models.Profile) error
	ReadProfile(context.Context, uuid.UUID) (*models.Profile, error)
	GetProfileIdByLogin(context.Context, string) (uuid.UUID, error)
	UpdateProfile(context.Context, *models.Profile) error
	UpdatePhoto(context.Context, uuid.UUID, string) error
}

type ProfileConfig interface {
}

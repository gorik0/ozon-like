package repo

import (
	"context"
	"github.com/jackc/pgtype/pgxtype"
	uuid "github.com/satori/go.uuid"
	"ozone/internal/models"
	"ozone/internal/pkg/profile"
)

type ProfileRepo struct {
	db pgxtype.Querier
}

const (
	createProfile       = `SELECT ... `
	readProfile         = `SELECT ... `
	getProfileIdByLogin = `SELECT ... `
	updateProfile       = `SELECT ... `
	updatePhoto         = `SELECT ... `
)

func (p ProfileRepo) CreateProfile(ctx context.Context, m *models.Profile) error {
	//TODO implement me
	panic("implement me")
}

func (p ProfileRepo) ReadProfile(ctx context.Context, uuid uuid.UUID) (*models.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileRepo) GetProfileIdByLogin(ctx context.Context, s string) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileRepo) UpdateProfile(ctx context.Context, m *models.Profile) error {
	//TODO implement me
	panic("implement me")
}

func (p ProfileRepo) UpdatePhoto(ctx context.Context, uuid uuid.UUID, s string) error {
	//TODO implement me
	panic("implement me")
}

func NewProfileRepo(db pgxtype.Querier) *ProfileRepo {
	return &ProfileRepo{db: db}
}

var _ profile.ProfileRepo = &ProfileRepo{}

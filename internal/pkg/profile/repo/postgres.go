package repo

import "github.com/jackc/pgtype/pgxtype"

type ProfileRepo struct {
	db pgxtype.Querier
}

func NewProfileRepo(db pgxtype.Querier) *ProfileRepo {
	return &ProfileRepo{db: db}
}

package repository

import (
	"anarchive/entity"
	"context"

	"github.com/uptrace/bun"
)

type ArchiveRepository struct {
	db bun.IDB
}

func NewArchiveRepository(db bun.IDB) *ArchiveRepository {
	return &ArchiveRepository{db: db}
}

func (r *ArchiveRepository) CreateArchive(ctx context.Context, a *entity.Archive) (*entity.Archive, error) {
	_, err := r.db.NewInsert().
		Model(a).
		Exec(ctx)
	return a, err
}

func (r *ArchiveRepository) UpdateArchive(ctx context.Context, a *entity.Archive) (*entity.Archive, error) {
	_, err := r.db.NewUpdate().
		Model(a).
		Where("id = ?", a.ID).
		Exec(ctx)
	return a, err
}

func (r *ArchiveRepository) DeleteArchive(ctx context.Context, a *entity.Archive) error {
	_, err := r.db.NewDelete().
		Model(a).
		Where("id = ?", a.ID).
		Exec(ctx)
	return err
}

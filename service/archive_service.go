package service

import (
	"anarchive/dto"
	"anarchive/entity"
	"anarchive/repository"
	"context"
	"time"
)

type ArchiveService struct {
	r repository.ArchiveRepository
}

func NewArchiveService(archiveRepository repository.ArchiveRepository) *ArchiveService {
	return &ArchiveService{
		r: archiveRepository,
	}
}

func (s *ArchiveService) CreateArchive(ctx *context.Context, request dto.ArchiveCreateRequest) (*entity.Archive, error) {
	archive := entity.Archive{
		Namespace: request.Namespace,
		BucketID:  request.BucketID,
		Length:    request.Length,
		FileType:  request.FileType,
		Status:    0,
	}
	return s.r.CreateArchive(*ctx, &archive)
}

func (s *ArchiveService) SetFileAsUploaded(ctx *context.Context, id entity.ID, hash string, uploader int64) (*entity.Archive, error) {
	archive := entity.Archive{
		ID:         id,
		Status:     1,
		Hash:       hash,
		UploadedAt: time.Now(),
		UploadedBy: uploader,
	}
	return s.r.UpdateArchive(*ctx, &archive)
}

func (s *ArchiveService) UpdatePartNumber(ctx *context.Context, id entity.ID, partNumber int32) (*entity.Archive, error) {
	archive := entity.Archive{
		ID:     id,
		Status: -partNumber,
	}
	return s.r.UpdateArchive(*ctx, &archive)
}

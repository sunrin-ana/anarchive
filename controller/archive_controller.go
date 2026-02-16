package controller

import (
	"anarchive/dto"
	"anarchive/pkg"
	"anarchive/service"
	"context"

	"github.com/NARUBROWN/spine/pkg/header"
	"github.com/NARUBROWN/spine/pkg/query"
	"github.com/NARUBROWN/spine/pkg/spine"
)

type ArchiveController struct {
	archiveService *service.ArchiveService
}

func NewArchiveController(archiveService *service.ArchiveService) *ArchiveController {
	return &ArchiveController{
		archiveService: archiveService,
	}
}

func (c ArchiveController) PutArchive(
	ctx context.Context,
	ns string,
	bucket string,
	key string,
	q query.Values,
	headers header.Values,
	spineCtx spine.Ctx,
) error {
	validatedHeader, err := pkg.ValidateHeaders(&headers, dto.ArchiveUploadRequestHeader{})

	if err != nil {
		return err
	}

	if !q.Has("uploadId") {
		// 단일 파일 업로드 시
		request := dto.ArchiveCreateRequest{
			Namespace: ns,
			BucketID:  bucket,
			Length:    validatedHeader.Length,
			FileType:  validatedHeader.Type,
		}
		_, err := c.archiveService.CreateArchive(&ctx, request)

		// TODO: Spine는 BodyStream을 지원하지 않음. 향후 지원 시 BodyStream으로 구현 필요

		return err
	}

	return nil
}

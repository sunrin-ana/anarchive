package dto

type ArchiveCreateRequest struct {
	Namespace string
	BucketID  string
	FileType  string
	Length    int64
}

type ArchiveUploadRequestHeader struct {
	Length        int64  `header:"Content-Length" validate:"required"`
	Type          string `header:"Content-Type" validate:"required"`
	MD5Hash       string `header:"Content-MD5" validate:"required"`
	Authorization string
}

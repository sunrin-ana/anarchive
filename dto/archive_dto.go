package dto

type ArchiveCreateRequest struct {
	Namespace string
	BucketID  int64
	FileType  string
	Length    int64
}

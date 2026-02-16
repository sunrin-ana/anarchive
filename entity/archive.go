package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Archive struct {
	Namespace  string    `bun:"namespace,type:varchar(512),notnull,index:idx_archive"`
	BucketID   int64     `bun:"bucket_id,type:bigint,notnull,index:idx_archive"`
	Bucket     *Bucket   `bun:"rel:belongs-to,join:bucket_id=id"`
	ID         ID        `bun:"id,type:uuid,default:uuid_generate_v4(),notnull,index:idx_archive"`
	FileType   string    `bun:"file_type,type:varchar(64),notnull"`
	Length     int64     `bun:"size,type:bigint,notnull"`
	Hash       string    `bun:"hash,type:varchar(32),notnull"`
	Status     int32     `bun:"status,notnull"` // 0: 업로드 되지 않음, 1: 업로드 됨, -N: Multipart로 업로드됨 (N은 Part 수)
	UploadedAt time.Time `bun:"uploaded_at,nullzero"`
	UploadedBy int64     `bun:"uploaded_by,nullzero"`

	ArchiveRefs []*ArchiveRef `bun:"m2m:archive_to_ref,join:archive_ref_id=id"`
}

type ArchiveRef struct {
	Namespace string  `bun:"type:varchar(512),notnull,index:idx_archive_ref"`
	BucketID  int64   `bun:"bucket_id,type:bigint,notnull,index:idx_archive_ref"`
	Bucket    *Bucket `bun:"rel:has-one,join:bucket_id=id"`
	ID        int64   `bun:"id,type:bigint,index:idx_archive_ref"`

	Archives []*Archive `bun:"m2m:archive_to_ref,join:archive_id=id"`
}

type ArchiveToRef struct {
	bun.BaseModel `bun:"table:archive_to_ref"`

	archiveID    ID    `bun:"archive_id,pk"`
	archiveRefID int64 `bun:"archive_ref_id,pk"`

	Archive    *Archive
	ArchiveRef *ArchiveRef
}

package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Archive struct {
	Namespace    string    `bun:"namespace,type:varchar(512),notnull,index:idx_archive"`
	BucketID     int64     `bun:"bucket_id,type:bigint,notnull,index:idx_archive"`
	Bucket       *Bucket   `bun:"rel:belongs-to,join:bucket_id=id"`
	ID           uuid.UUID `bun:"id,type:uuid,default:uuid_generate_v4(),notnull,index:idx_archive"`
	OriginalName string    `bun:"original_name,type:varchar(512),notnull"`
	Type         string    `bun:"file_type,type:varchar(64),notnull"`
	UploadedAt   time.Time `bun:"uploaded_at,notnull,default:now()"`
	UploadedBy   *int64    `bun:"uploaded_by"`

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

	archiveID    int64 `bun:"archive_id,pk"`
	archiveRefID int64 `bun:"archive_ref_id,pk"`

	Archive    *Archive
	ArchiveRef *ArchiveRef
}

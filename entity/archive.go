package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Archive struct {
	Namespace  string    `bun:"namespace,type:varchar(512),notnull,index:idx_archive"`
	BucketID   string    `bun:"bucket_id,type:varchar(512),notnull,index:idx_archive"`
	Bucket     *Bucket   `bun:"rel:belongs-to,join:bucket_id=id"`
	ID         ID        `bun:"id,type:uuid,default:uuid_generate_v4(),notnull,index:idx_archive"`
	FileType   string    `bun:"file_type,type:varchar(64),notnull"`
	Length     int64     `bun:"size,type:bigint,notnull"`
	Hash       string    `bun:"hash,type:varchar(32),notnull"`
	Status     int32     `bun:"status,notnull"` // 0: 업로드 되지 않음, 1: 업로드 됨, -N: Multipart로 업로드됨 (N은 Part 수)
	UploadedAt time.Time `bun:"uploaded_at,nullzero"`
	UploadedBy int64     `bun:"uploaded_by,nullzero"`

	ArchiveRefs []*ArchiveRef `bun:"m2m:archive_to_ref,join:namespace=namespace,join:bucket_id=bucket_id,join:id=archive_id"`
}

type ArchiveRef struct {
	Namespace string  `bun:"type:varchar(512),notnull,index:idx_archive_ref"`
	BucketID  string  `bun:"bucket_id,type:varchar(512),notnull,index:idx_archive_ref"`
	Bucket    *Bucket `bun:"rel:has-one,join:bucket_id=id"`
	ID        int64   `bun:"id,type:bigint,index:idx_archive_ref"`

	Archives []*Archive `bun:"m2m:archive_to_ref,join:namespace=namespace,join:bucket_id=bucket_id,join:id=ref_id"`
}

type ArchiveToRef struct {
	bun.BaseModel `bun:"table:archive_to_ref"`

	Namespace string `bun:"namespace,type:varchar(512),notnull,index:idx_archive,index:idx_ref"`
	BucketID  string `bun:"bucket_id,type:varchar(512),notnull,index:idx_archive,index:idx_ref"`
	archiveID ID     `bun:"archive_id,pk,index:idx_archive"`
	RefID     int64  `bun:"ref_id,index:idx_ref"`

	Archive    *Archive
	ArchiveRef *ArchiveRef
}

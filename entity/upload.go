package entity

import "time"

type UploadSession struct {
	ID        string    `bun:"type:uuid,default:gen_random_uuid()"`
	Namespace string    `bun:"type:varchar(512),notnull"`
	BucketID  int64     `bun:"bucket_id,notnull"`
	Bucket    *Bucket   `bun:"rel:belongs-to,join:bucket_id=id"`
	Target    string    `bun:"target"`
	State     int       `bun:"state,notnull"`
	RefID     int64     `bun:"ref_id,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull"`
	ExpireAt  time.Time `bun:"expire_at,notnull"`
}

type FilePart struct {
	Parent   *UploadSession `bun:"rel:belongs-to,join:id=parent_id"`
	ParentID string         `bun:"parent_id,index:idx_file_part,type:uuid,notnull"`
	Part     int            `bun:"part,index:idx_file_part"`
	ETag     string         `bun:"etag,notnull"`
}

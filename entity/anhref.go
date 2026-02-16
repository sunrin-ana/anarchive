package entity

import "time"

type Anhref struct {
	ID        string    `bun:"type:uuid,notnull"`
	Name      string    `bun:"name,notnull"`
	Namespace string    `bun:"namespace,notnull"`
	BucketID  string    `bun:"bucket_id,notnull"`
	Bucket    *Bucket   `bun:"rel:belongs-to,join:bucket_id=id"`
	ArchiveID ID        `bun:"archive_id,notnull"`
	Archive   *Archive  `bun:"rel:belongs-to,join:archive_id=id"`
	ExpireAt  time.Time `bun:"expire_at,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull"`
	CreatedBy *int64    `bun:"created_by"`
}

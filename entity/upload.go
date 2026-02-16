package entity

import "time"

type UploadSession struct {
	ID        string    `bun:"type:uuid,default:gen_random_uuid()"`
	ArchiveID int64     `bun:"archive_id,notnull"`
	Archive   *Archive  `bun:"rel:has-one,join:archive_id=id"`
	Target    string    `bun:"target"`
	CreatedAt time.Time `bun:"created_at,notnull"`
	ExpireAt  time.Time `bun:"expire_at,notnull"`
}

type FilePart struct {
	Parent   *UploadSession `bun:"rel:belongs-to,join:id=parent_id"`
	ParentID string         `bun:"parent_id,index:idx_file_part,type:uuid,notnull"`
	Part     int            `bun:"part,index:idx_file_part"`
	ETag     string         `bun:"etag,notnull"`
}

package entity

import "github.com/uptrace/bun"

type Bucket struct {
	Namespace string `bun:"type:varchar(512),notnull,index:idx_bucket"`
	ID        string `bun:"id,type:varchar(512),notnull,index:idx_bucket"`

	BucketRefs []*BucketRef `bun:"m2m:bucket_to_ref,join:namespace=namespace,join:id=bucket_id"`
}

type BucketRef struct {
	Namespace string `bun:"type:varchar(512),notnull"`
	ID        string `bun:"id,type:varchar(512),notnull"`

	Buckets []*Bucket `bun:"m2m:bucket_to_ref,join:namespace=namespace,join:id=ref_id"`
}

type BucketToRef struct {
	bun.BaseModel `bun:"table:bucket_to_ref"`
	Namespace     string `bun:"type:varchar(512),notnull,index:idx_bucket,index:idx_ref"`
	BucketID      string `bun:"bucket_id,type:varchar(512),index:idx_bucket"`
	RefID         int64  `bun:"ref_id,index:idx_bucket,index:idx_ref"`

	Bucket *Bucket    `bun:"rel:belongs-to,join:namespace=namespace,join:bucket_id=id"`
	Ref    *BucketRef `bun:"rel:belongs-to,join:namespace=namespace,join:ref_id=id"`
}

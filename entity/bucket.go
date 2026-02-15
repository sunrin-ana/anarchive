package entity

type Bucket struct {
	Namespace string `bun:"type:varchar(512),notnull,index:idx_bucket"`
	Name      string `bun:"type:varchar(512),notnull"`
	ID        int64  `bun:"id,type:bigint,notnull,index:idx_bucket"`

	BucketRefs []*BucketRef `bun:"m2m:bucket_to_ref,join:ref_id=id"`
}

type BucketRef struct {
	Namespace string `bun:"type:varchar(512),notnull"`
	ID        int64  `bun:"id,type:bigint,notnull"`

	Buckets []*Bucket `bun:"m2m:bucket_to_ref,join:bucket_id=id"`
}

type BucketToRef struct {
	BucketID int64 `bun:"bucket_id,pk"`
	RefID    int64 `bun:"ref_id,pk"`

	Bucket *Bucket    `bun:"rel:belongs-to,join:bucket_id=id"`
	Ref    *BucketRef `bun:"rel:belongs-to,join:ref_id=id"`
}

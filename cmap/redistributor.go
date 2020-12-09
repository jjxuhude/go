package cmap

type BucketStatus uint8

const (
	BUCKET_STATUS_NORMAL      BucketStatus = 0
	BUCKET_STATUS_UNDERWEIGHT BucketStatus = 1
	BUCKET_STATUS_OVERWEIGHT  BucketStatus = 2
)

type PairRedistributor interface {
	UpdateThreshold(pairTotal uint64, bucketNumber int)
	CheckBucketStatus(pariTotal uint64, bucketSize uint64) (bucketStatus BucketStatus)
	Redistribe(bucketStatus BucketStatus, buckets []Bucket) (newBuckets []Bucket, changed bool)
}

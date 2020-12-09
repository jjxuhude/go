package cmap

import "sync"

type Segment interface {
	Put(p Pair) (bool, error)
	Get(key string) Pair
	GetWithHash(key string, keyHash uint64) Pair
	Delete(key string) bool
	Size() uint64
}

type segment struct {
	buckets           []Bucket
	bucketsLen        int
	pairTotal         uint64
	pairRedistributor PairRedistributor
	lock              sync.Mutex
}

func newSegment(bucketNumber int, pairRedistributor PairRedistributor) Segment {
	if bucketNumber <= 0 {
		bucketNumber = DEFAULT_BUCKET_NUMBER
	}
	if pairRedistributor == nil {
		pairRedistributor = newDefaultPairRedistributor(
			DEFAULT_BUCKET_LOAD_FACTOR, bucketNumber)
	}
	bucktes := make([]Bucket, bucketNumber)
	for i := 0; i < bucketNumber; i++ {
		buckets[i] = NewBucket()
	}
	return &segment{}

}

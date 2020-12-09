package cmap

type ConcurrentMap interface {
	Concurrency() uint //并发量
	Put(key string, element interface{}) (bool, error)
	Get(key string) interface{}
	Delete(key string) bool
	Len() uint
}

type ConcurrentMapIns struct {
	concurrency int
	segments    []Segment
	total       int64
}

func NewConcurrentMap(concurrency int, pairRedistributor PairRedistributor) (ConcurrentMap, error) {
	if concurrency <= 0 {
		return nil, newIllegalParameterError("并发量必须大于0")
	}
	if concurrency > MAX_CONCURRENCY {
		return nil, newIllegalParameterError("并发量太大")
	}
	cmap := &ConcurrentMapIns{}
	cmap.concurrency = concurrency
	cmap.segments = make([]Segment, concurrency)
	for i := 0; i < concurrency; i++ {
		cmap.segments[i] = newSegment(DEFAULT_BUCKET_NUMBER, pairRedistributor)
	}
	return cmap, nil
}

func (cmap *ConcurrentMapIns) Concurrency() uint {
}

func (cmap *ConcurrentMapIns) Put(key string, element interface{}) (bool, error) {

}

func (cmap *ConcurrentMapIns) Get(key string) interface{} {

}

func (cmap *ConcurrentMapIns) Delete(key string) bool {

}

func (cmap *ConcurrentMapIns) Len() uint {

}

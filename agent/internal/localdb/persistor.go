package localdb

import (
	"Linda/baselibs/abstractions"
	"Linda/baselibs/abstractions/xlog"

	"github.com/nutsdb/nutsdb"
)

var (
	logger = xlog.NewForPackage()
)

// Persistor[K, V]
// 这是一个用来给LocalDB的每个Bucket做CRUD的结构
type Persistor[K, V abstractions.Serializable] struct {
	ldb    *LocalDB
	bucket string
}

func (p *Persistor[K, V]) Set(k K, v V) (err error) {
	return p.ldb.Set2(p.bucket, k.Serialize(), v.Serialize())
}

func (p *Persistor[K, V]) Delete(k K) (err error) {
	return p.ldb.Delete2(p.bucket, k.Serialize())
}

func (p *Persistor[K, V]) Get(k K, v V) (err error) {
	result, err := p.ldb.Get2(p.bucket, k.Serialize())
	v.Deserialize(result)
	return err
}

func GetPersistor[K, V abstractions.Serializable](bucket string) (p *Persistor[K, V], err error) {
	if !Instance().ExistBucket(bucket) {
		err = Instance().NewBucket(bucket)
		if err != nil && err != nutsdb.ErrBucketAlreadyExist {
			logger.Errorf("create bucket %s failed, err is %v", bucket, err)
			return
		}
	}
	p = &Persistor[K, V]{
		ldb:    localdbInstance,
		bucket: bucket,
	}
	return
}

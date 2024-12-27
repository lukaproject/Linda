package localdb

import (
	"Linda/agent/internal/config"

	"github.com/lukaproject/xerr"
	"github.com/nutsdb/nutsdb"
)

var localdbInstance *LocalDB = nil

func Initial() {
	localdbInstance = New()
}

func Instance() *LocalDB {
	return localdbInstance
}

type LocalDB struct {
	db *nutsdb.DB
}

func New() (ldb *LocalDB) {
	ldb = &LocalDB{}
	ldb.db = xerr.Must(
		nutsdb.Open(
			nutsdb.DefaultOptions,
			nutsdb.WithDir(config.Instance().LocalDBDir)))
	return
}

func (ldb *LocalDB) Set2(bucket string, k, v []byte) error {
	return ldb.db.Update(func(tx *nutsdb.Tx) error {
		if err := tx.Put(bucket, k, v, 0); err != nil {
			return err
		}
		return nil
	})
}

func (ldb *LocalDB) Delete2(bucket string, k []byte) error {
	return ldb.db.Update(func(tx *nutsdb.Tx) error {
		if err := tx.Delete(bucket, k); err != nil {
			if err == nutsdb.ErrKeyNotFound {
				return nil
			}
			return err
		}
		return nil
	})
}

func (ldb *LocalDB) Get2(bucket string, k []byte) (v []byte, err error) {
	err = ldb.db.View(
		func(tx *nutsdb.Tx) error {
			bytev, err := tx.Get(bucket, k)
			if err != nil {
				return err
			}
			v = bytev
			return nil
		})
	return
}

func (ldb *LocalDB) NewBucket(bucket string) error {
	return localdbInstance.db.Update(
		func(tx *nutsdb.Tx) error {
			if tx.ExistBucket(nutsdb.DataStructureBTree, bucket) {
				logger.Warnf("localdb bucket %s has created before!", bucket)
				return nil
			}
			return tx.NewBucket(nutsdb.DataStructureBTree, bucket)
		})
}

func (ldb *LocalDB) ExistBucket(bucket string) (exist bool) {
	exist = false
	_ = localdbInstance.db.View(func(tx *nutsdb.Tx) error {
		if tx.ExistBucket(nutsdb.DataStructureBTree, bucket) {
			exist = true
		}
		return nil
	})
	return
}

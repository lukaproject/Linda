package localdb

import (
	"Linda/agent/internal/config"

	"github.com/lukaproject/xerr"
	"github.com/nutsdb/nutsdb"
)

const (
	LocalDBBucket = "bucket"

	BagNameKey = "LINDA_BAGNAME"
)

var localdbInstance *LocalDB = nil

func Initial() {
	localdbInstance = New()
	xerr.Must0(localdbInstance.db.Update(
		func(tx *nutsdb.Tx) error {
			return tx.NewBucket(nutsdb.DataStructureBTree, LocalDBBucket)
		}))
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

func (ldb *LocalDB) Set(k, v string) error {
	return ldb.db.Update(func(tx *nutsdb.Tx) error {
		if err := tx.Put(LocalDBBucket, []byte(k), []byte(v), 0); err != nil {
			return err
		}
		return nil
	})
}

func (ldb *LocalDB) Delete(k string) error {
	return ldb.db.Update(func(tx *nutsdb.Tx) error {
		if err := tx.Delete(LocalDBBucket, []byte(k)); err != nil {
			if err == nutsdb.ErrKeyNotFound {
				return nil
			}
			return err
		}
		return nil
	})
}

func (ldb *LocalDB) Get(k string) (v string, err error) {
	err = ldb.db.View(
		func(tx *nutsdb.Tx) error {
			bytev, err := tx.Get(LocalDBBucket, []byte(k))
			if err != nil {
				v = ""
				return err
			}
			v = string(bytev)
			return nil
		})
	return
}

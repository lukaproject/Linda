package data

type KeyType struct {
	Key string
}

func (kt *KeyType) Serialize() []byte {
	return []byte(kt.Key)
}

func (kt *KeyType) Deserialize(b []byte) (err error) {
	kt.Key = string(b)
	return err
}

func NewKey(k string) *KeyType {
	return &KeyType{Key: k}
}

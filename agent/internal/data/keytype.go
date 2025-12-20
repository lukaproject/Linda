package data

type StringType struct {
	Key string
}

func (kt *StringType) Serialize() []byte {
	return []byte(kt.Key)
}

func (kt *StringType) Deserialize(b []byte) (err error) {
	if kt == nil {
		kt = &StringType{}
	}
	kt.Key = string(b)
	return err
}

func NewKey(k string) *StringType {
	return &StringType{Key: k}
}

func NewString(k string) *StringType {
	return &StringType{Key: k}
}

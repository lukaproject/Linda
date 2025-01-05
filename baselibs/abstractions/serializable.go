package abstractions

type Serializable interface {
	Serialize() []byte
	Deserialize([]byte) error
}

package generator

var (
	instance Generator
)

func Initial() {
	instance = new(generatorImpl)
}

func GetInstance() Generator {
	return instance
}

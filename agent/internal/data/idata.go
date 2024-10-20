package data

type IData interface {
	Load()
	Store()
}

func GetData[Data IData](data Data, reload bool) Data {
	if reload {
		data.Load()
	}
	return data
}

package data

type IData interface {
	// Load is a function load data from db to struct itself
	// pls confirm that the struct has unique key to
	// identify themself.
	Load()
	// Store is a function store data into db.
	// pls confirm that the struct has unique key to
	// identify themself.
	Store()
}

func GetData[Data IData](data Data, reload bool) Data {
	if reload {
		data.Load()
	}
	return data
}

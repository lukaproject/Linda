package errno

var (
	ErrMapErrNumber = map[error]int{
		ErrEmptyBag: 20001,
	}

	ErrMapNumberErr = map[int]error{
		20001: ErrEmptyBag,
	}
)

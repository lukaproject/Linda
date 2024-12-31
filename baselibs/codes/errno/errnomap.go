package errno

var (
	ErrMapErrNumber = map[error]int{
		ErrEmptyBag:      20001,
		ErrNodeIdExists:  20002,
		ErrAgentNotFound: 20003,
	}

	ErrMapNumberErr = map[int]error{
		20001: ErrEmptyBag,
		20002: ErrNodeIdExists,
		20003: ErrAgentNotFound,
	}
)

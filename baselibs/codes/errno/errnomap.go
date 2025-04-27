package errno

var (
	ErrMapErrNumber = map[error]int{
		ErrEmptyBag:                20001,
		ErrNodeIdExists:            20002,
		ErrAgentNotFound:           20003,
		ErrNodeBelongsToAnotherBag: 20004,
		ErrInvalidDownloadType:     20005,
		ErrTargetPathIsEmpty:       20006,
		ErrSourceURLIsEmpty:        20007,
		ErrFileNameIsEmpty:         20008,
		ErrInvalidTaskData:         20009,
	}

	ErrMapNumberErr = map[int]error{
		20001: ErrEmptyBag,
		20002: ErrNodeIdExists,
		20003: ErrAgentNotFound,
		20004: ErrNodeBelongsToAnotherBag,
		20005: ErrInvalidDownloadType,
		20006: ErrTargetPathIsEmpty,
		20007: ErrSourceURLIsEmpty,
		20008: ErrFileNameIsEmpty,
		20009: ErrInvalidTaskData,
	}
)

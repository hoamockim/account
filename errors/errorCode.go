package errors

type (
	ServiceName string
	ErrorCode   string
)

const (
	order        = ServiceName("ord")
	notification = ServiceName("ntf")
)

const (
	Panicked   ErrorCode = "PANICKED"
	OffService ErrorCode = "TURNOFF"
)

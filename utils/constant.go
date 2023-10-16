package utils

type Key int

const (
	Key1 Key = iota
	Key2
	Key3
)

const (
	RequestInfoMsg  string = "Request Information"
	ResponseInfoMsg string = "Response Information"
	SummaryInfoMsg  string = "Summary Information"
)

const (
	ValidateFieldError string = "Invalid Parameters"
)

const (
	ContentType     string = "Content-Type"
	ApplicationJSON string = "application/json"
	XRequestID      string = "X-Request-ID"
)

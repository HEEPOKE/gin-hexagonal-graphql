package response

type ResponseError struct {
	Code    string
	Message string
	Err     error
}

func (re *ResponseError) Error() string {
	return re.Message
}

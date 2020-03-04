package rest

type MessageError string

func (e MessageError) Error() string {
	return string(e)
}

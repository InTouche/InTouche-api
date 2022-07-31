package internal

type key int

const (
	KeyLogger key = iota
	KeyAuthToken
	KeyRequestID
)

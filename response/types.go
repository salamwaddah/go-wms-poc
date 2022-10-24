package response

type Response interface {
	ToResponse() interface{}
}

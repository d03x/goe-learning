package dto

type Response[T any] struct {
	Status  bool
	Message string
	Data    T
}

func ResponseWithData[T any](data T) Response[T] {
	return Response[T]{
		Status:  true,
		Message: "Success",
		Data:    data,
	}
}

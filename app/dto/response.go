package dto

type Response[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func ResponseWithData[T any](data T) Response[T] {
	return Response[T]{
		Status:  true,
		Message: "Success",
		Data:    data,
	}
}

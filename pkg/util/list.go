package util

func ToList[T any](values ...T) []T {
	list := make([]T, 0)

	list = append(list, values...)
	return list
}

func Add[T any](list *[]T, values ...T) {
	if list != nil && values != nil {
		*list = append(*list, values...)
	}
}

func AddAll[T any](list *[]T, values []T) {
	*list = append(*list, values...)
}

func NewList[T any]() []T {
	return make([]T, 0)
}

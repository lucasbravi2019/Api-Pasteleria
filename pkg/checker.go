package pkg

func IsNull[T any](obj *T) bool {
	return obj == nil
}

func IsNotNull[T any](obj *T) bool {
	return !IsNull(obj)
}

func HasError(err error) bool {
	return err != nil
}

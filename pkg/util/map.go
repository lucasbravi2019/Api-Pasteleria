package util

func NewMap[K comparable, V any]() map[K]V {
	newMap := make(map[K]V)

	return newMap
}

func GetValue[K comparable, V any](mapa map[K]V, key K) *V {
	value, exists := mapa[key]

	if !exists {
		return nil
	}

	return &value
}

func PutValue[K comparable, V any](mapa *map[K]V, key K, value V) {
	(*mapa)[key] = value
}

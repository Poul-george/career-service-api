package convert

func ToP[T any](v T) *T {
	return &v
}

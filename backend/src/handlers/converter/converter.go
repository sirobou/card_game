package converter

type Converter[T any, U any] interface {
	Convert(*T) *U
}

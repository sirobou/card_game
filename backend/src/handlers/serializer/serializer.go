package serializer

type Serializer[T any] interface {
	Serialize(metadata T) ([]byte, error)
	Print(metadata T)
}

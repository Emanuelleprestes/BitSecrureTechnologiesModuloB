package repositorios

type Crudinterface[T any, y any] interface {
	Save() T
	Update(T) T
	Get(y) T
}

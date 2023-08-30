package implementable

type Pager[P interface{}, Q interface{}] interface {
	NextPage(i P) (Q, error)
	More() bool
}

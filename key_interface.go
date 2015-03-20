package key

type Key interface {
	ToString() string
}
type Sum interface {
	ToString() string
	ToSum() int
}
type Int interface {
	ToString() string
	ToInt() int
}

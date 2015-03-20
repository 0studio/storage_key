package key

type Key interface {
	ToString() string
}
type Int interface {
	ToString() string
	ToInt() int
}

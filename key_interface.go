package key

type Key interface {
	ToString() string
}
type Sum interface {
	ToString() string
	ToSum() int
	SumLen() int
	GetSumByIdx(idx int) Sum
}
type Int interface {
	ToString() string
	ToInt() int
}

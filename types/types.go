package mytypes

// to learn about methods

type MyInt int
type MyStringLen string

func (i MyInt) Twice() MyInt {
	return i * 2
}

func (m MyStringLen) Len() int {
	return len(m)
}

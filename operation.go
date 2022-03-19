package math

type ByteAdder interface {
	Add(x, y byte) byte
}

type ByteSubtractor interface {
	Sub(x, y byte) byte
}

type ByteMultiplier interface {
	Mul(x, y byte) byte
}

type ByteDivider interface {
	Div(x, y byte) byte
}

type ByteAlgebra interface {
	ByteAdder
	ByteSubtractor
	ByteMultiplier
	ByteDivider
}

type ByteAlgebraImpl struct{}

func (alg *ByteAlgebraImpl) Add(x, y byte) byte {
	return x + y
}

func (alg *ByteAlgebraImpl) Sub(x, y byte) byte {
	return x - y
}

func (alg *ByteAlgebraImpl) Mul(x, y byte) byte {
	return x * y
}

func (alg *ByteAlgebraImpl) Div(x, y byte) byte {
	return x / y
}

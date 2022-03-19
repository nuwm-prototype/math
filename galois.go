package math

const (
	// PrimePoly4 is prime polynom for GF (2^4)
	PrimePoly4 = 023
	// PrimePoly8 is prime polynom for GF (2^8)
	PrimePoly8 = 0435
)

// ByteGaloisAlgebraImpl is an implementation of ByteAlgebraImpl interface
// uses gflog and gfilog for defining multiplication and division operations
type ByteGaloisAlgebraImpl struct {
	gflog  []byte
	gfilog []byte
	size   int16
}

func New(w byte) *ByteGaloisAlgebraImpl {
	var primePoly int16
	switch w {
	case 4:
		primePoly = PrimePoly4
		break
	case 8:
		primePoly = PrimePoly8
		break
	default:
		return nil
	}
	var size int16 = 1 << w
	gflog := make([]byte, size)
	gfilog := make([]byte, size)

	var log, b int16

	for log, b = 0, 1; log < size; log++ {
		gflog[b] = byte(log)
		gfilog[log] = byte(b)
		b = b << 1
		if b&int16(size) != 0 {
			b = b ^ primePoly
		}
	}
	return &ByteGaloisAlgebraImpl{gflog, gfilog, size}
}

func (alg *ByteGaloisAlgebraImpl) Add(x, y byte) byte {
	return x ^ y
}

func (alg *ByteGaloisAlgebraImpl) Sub(x, y byte) byte {
	return x ^ y
}

func (alg *ByteGaloisAlgebraImpl) Mul(x, y byte) byte {
	if x == 0 || y == 0 {
		return 0
	}
	return alg.gfilog[(int16(alg.gflog[x])+int16(alg.gflog[y]))%(alg.size-1)]
}

func (alg *ByteGaloisAlgebraImpl) Div(x, y byte) byte {
	if x == 0 {
		return 0
	}
	return alg.gfilog[(int16(alg.gflog[x])-int16(alg.gflog[y])+alg.size-1)%(alg.size-1)]
}

package base

type BitFactory struct{}

func (BitFactory) Left(v, r uint64) uint64 {
	return v << r
}

func (BitFactory) Right(v, r uint64) uint64 {
	return v >> r
}

func (BitFactory) Xor(v, r uint64) uint64 {
	return v ^ r
}

func (BitFactory) And(v, r uint64) uint64 {
	return v & r
}

func (BitFactory) Or(v, r uint64) uint64 {
	return v | r
}

func (BitFactory) Not(v uint64) uint64 {
	return ^v
}

func (BitFactory) Test(v, r uint64) bool {
	return 0 != (v>>r)&1
}

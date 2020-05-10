package bitmap

type Bitmap struct {
	buf []uint64
	cap  int
}

func NewBitmap(N int) *Bitmap {
	n := N / 64
	if N % 64 > 0 {
		n++
	}
	return &Bitmap {  make([]uint64, n), n }
}

func (t *Bitmap) IsSet(i int) bool { i--; return t.buf[i/64] & ( 1<< uint(63 - i % 64)) != 0 }
func (t *Bitmap) Set(i int)        { i--; t.buf[i/64] |= 1 << uint(63 - i % 64) }
func (t *Bitmap) Clear(i int)      { i--; t.buf[i/64] &^= 1 << uint(63 - i % 64) }



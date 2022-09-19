package geecache

type ByteView struct {
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func cloneBytes(b []byte) []byte {
	ret := make([]byte, len(b))
	copy(ret, b)
	return ret
}

func (v ByteView) String() string {
	return string(v.b)
}



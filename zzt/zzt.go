package zzt

const (
	BoardCountLimit = 101
	BoardSizeLimit = 32767
	StatusElementLimit = 151
)

const (
	True byte = 0x01
	False byte = 0x00
)

type String20 struct {
	Length byte
	Value [20]byte
}

func (s String20) String() string {
	return string(s.Value[:s.Length])
}

func NewString20(s string) String20 {
	var v String20
	v.Length = byte(copy(v.Value[:], s))
	return v
}

type String50 struct {
	Length byte
	Value [50]byte
}

func (s String50) String() string {
	return string(s.Value[:s.Length])
}

func NewString50(s string) String50 {
	var v String50
	v.Length = byte(copy(v.Value[:], s))
	return v
}

type String58 struct {
	Length byte
	Value [58]byte
}

func (s String58) String() string {
	return string(s.Value[:s.Length])
}

func NewString58(s string) String58 {
	var v String58
	v.Length = byte(copy(v.Value[:], s))
	return v
}
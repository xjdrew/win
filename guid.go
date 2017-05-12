package win

import (
	"fmt"
)

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

func (g GUID) String() string {
	return fmt.Sprintf("%08x-%04x-%04x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		g.Data1,
		g.Data2,
		g.Data3,
		g.Data4[0], g.Data4[1],
		g.Data4[2], g.Data4[3], g.Data4[4], g.Data4[5], g.Data4[6], g.Data4[7],
	)
}

func DEFINE_GUID(l uint32, w1, w2 uint16, b1, b2, b3, b4, b5, b6, b7, b8 uint8) GUID {
	return GUID{
		l, w1, w2, [8]byte{b1, b2, b3, b4, b5, b6, b7, b8},
	}
}

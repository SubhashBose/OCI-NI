package waste

import "math/rand"

var Buffers []*GiBObject

const (
	KiB = 1024
	MiB = 1024 * KiB
	GiB = 1024 * MiB
)

type GiBObject struct {
	B [MiB]byte
}

func Memory(gib float64) {
	Buffers = make([]*GiBObject, 0, int(gib*1024))
	for gib > 0 {
		o := new(GiBObject)
		rand.Read(o.B[:])
		Buffers = append(Buffers, o)
		gib -= 1
	}
}

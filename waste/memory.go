package waste

import "math/rand"

var Buffers []*1024*MiBObject

const (
	KiB = 1024
	MiB = 1024 * KiB
	GiB = 1024 * MiB
)

type GiBObject struct {
	B [GiB]byte
}

type MiBObject struct {
	B [MiB]byte
}

func Memory(gib float64) {
	mib:=int(gib*1024)
	Buffers = make([]*MiBObject, 0, mib)
	for mib > 0 {
		o := new(MiBObject)
		rand.Read(o.B[:])
		Buffers = append(Buffers, o)
		mib -= 1
	}
}

package waste

import "math/rand"

var Buffers []*GiBObject
var BuffersM []*MiBObject

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

func Memory(gibF float64) {
	gib:=int(gibF)
	mib:=int((gibF-float64(gib))*1024)
	Buffers = make([]*GiBObject, 0, gib)
	BuffersM = make([]*MiBObject, 0, mib)
	for gib > 0 {
		o := new(GiBObject)
		rand.Read(o.B[:])
		Buffers = append(Buffers, o)
		gib -= 1
	}
	for mib > 0 {
		o := new(MiBObject)
		rand.Read(o.B[:])
		BuffersM = append(BuffersM, o)
		mib -= 1
	}
}

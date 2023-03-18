package waste

import (
	"fmt"
	"math/rand"
	"time"
	"runtime"

	"golang.org/x/crypto/chacha20"
)

func CPU(interval time.Duration, duration time.Duration, percent float64, CPUcount int) {
	var buffer []byte
	if len(Buffers) > 0 {
		buffer = Buffers[0].B[:4*MiB]
	} else {
		buffer = make([]byte, 4*MiB)
	}
	rand.Read(buffer)

	// construct XChaCha20 stream cipher
	cipher, err := chacha20.NewUnauthenticatedCipher(buffer[:32], buffer[:24])
	if err != nil {
		panic(cipher)
	}

	runtime.GOMAXPROCS(CPUcount)
	for {
		for i := 0; i < CPUcount; i++ {
			go func() {
				runtime.LockOSThread()
				tend := time.Now().Add(duration)
				for ok := true; ok; ok = tend.After(time.Now()) {
					loop_st:= time.Now()
					for i := 0; i < 1; i++ {
						cipher.XORKeyStream(buffer, buffer)
					}
					loop_dur:= time.Since(loop_st)
					time.Sleep(loop_dur*time.Duration((100-percent)/percent) )
				}
			}()
		}

		fmt.Println("[CPU] Successfully wasted on", time.Now())

		// try to construct a new cipher
		newCipher, err := chacha20.NewUnauthenticatedCipher(buffer[:32], buffer[:24])
		if err == nil {
			cipher = newCipher
		}

		time.Sleep(interval)
	}
}

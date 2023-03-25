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
		buffer = Buffers[0].B[:6*MiB]
	} else {
		buffer = make([]byte, 6*MiB)
	}
	rand.Read(buffer)

	// construct XChaCha20 stream cipher
	cipher, err := chacha20.NewUnauthenticatedCipher(buffer[:32], buffer[:24])
	if err != nil {
		panic(cipher)
	}

	runtime.GOMAXPROCS(CPUcount)
	for {
		fmt.Println("[CPU] Starting stress on", time.Now())

		for i := 0; i < CPUcount; i++ {
			go func() {
				runtime.LockOSThread()
				XOR_cnt:=0
				tend := time.Now().Add(duration)
				for ok := true; ok; ok = tend.After(time.Now()) {
					loop_st:= time.Now()
					for i := 0; i < 1; i++ {
						cipher.XORKeyStream(buffer, buffer)
					}
					XOR_cnt+=1;
					loop_dur:= time.Since(loop_st)
					if XOR_cnt>=4*MiB/32/CPUcount {
						newCipher, err := chacha20.NewUnauthenticatedCipher(buffer[:32], buffer[:24])
						if err == nil {
							cipher = newCipher
							XOR_cnt=0
						}
					}
					time.Sleep(loop_dur*time.Duration((100-percent)/percent*1000)/time.Microsecond ) // percent part is rounded down to 1ns, so mult by 1000 then div by 1us
				}
			}()
		}

		// try to construct a new cipher
		newCipher, err := chacha20.NewUnauthenticatedCipher(buffer[:32], buffer[:24])
		if err == nil {
			cipher = newCipher
		}

		time.Sleep(interval)
	}
}

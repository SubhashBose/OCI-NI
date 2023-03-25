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
		fmt.Println("[CPU] Starting stress on", time.Now())

		for i := 0; i < CPUcount; i++ {
			go func() {
				runtime.LockOSThread()
				tend := time.Now().Add(duration)
				lp_cnt:=0
				for ok := true; ok; ok = tend.After(time.Now()) {
					loop_st:= time.Now()
					for i := 0; i < 1; i++ {
						cipher.XORKeyStream(buffer, buffer)
						lp_cnt++;
					}
					if lp_cnt>131000/CPUcount { // 4MiB/32B/Ncores
						newCipher, err := chacha20.NewUnauthenticatedCipher(buffer[:32], buffer[:24])
						if err == nil {
							cipher = newCipher
							lp_cnt=0
						}
					}
					loop_dur:= time.Since(loop_st)
					time.Sleep(loop_dur*time.Duration((100-percent)/percent*1000)/time.Microsecond )
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

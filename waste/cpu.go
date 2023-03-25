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

	runtime.GOMAXPROCS(CPUcount)
	for {
		fmt.Println("[CPU] Starting stress on", time.Now())

		// construct XChaCha20 stream cipher
		cipher, err := chacha20.NewUnauthenticatedCipher(buffer[:32], buffer[:24])
		if err != nil {
			panic(cipher)
		}
		XOR_cnt:=0

		for i := 0; i < CPUcount; i++ {
			go func() {
				runtime.LockOSThread()
				tend := time.Now().Add(duration)
				for ok := true; ok; ok = tend.After(time.Now()) {
					loop_st:= time.Now()
					for i := 0; i < 1; i++ {
						cipher.XORKeyStream(buffer, buffer)
						XOR_cnt++;
					}
					loop_dur:= time.Since(loop_st)
					if XOR_cnt>=4*MiB/32 {
						newCipher, err := chacha20.NewUnauthenticatedCipher(buffer[:32], buffer[:24])
						fmt.Println("[CPU] Counter reached", time.Now())
						if err == nil {
							cipher = newCipher
							XOR_cnt=0
							fmt.Println("[CPU] Replacing new", time.Now())
						}
					}
					time.Sleep(loop_dur*time.Duration((100-percent)/percent*1000)/time.Microsecond ) // percent part is rounded down to 1ns, so mult by 1000 then div by 1us
				}
			}()
		}

		time.Sleep(interval)

		// try to construct a new cipher
		//newCipher, err := chacha20.NewUnauthenticatedCipher(buffer[:32], buffer[:24])
		//if err == nil {
		//	cipher = newCipher
		//}
	}
}

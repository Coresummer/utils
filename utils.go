package utils

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"

	term "github.com/nsf/termbox-go"
)

func GetInput(CHmode chan<- int, termination chan int) error {
	//for keyboard input
	err := term.Init()
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer term.Close()

	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyArrowUp:
				CHmode <- 0
				fmt.Println("mode: Normal    ʕ◔ϖ◔ʔ")
			case term.KeyArrowDown:
				CHmode <- 1
				fmt.Println("mode: Adversary ʕ>ϖ<ʔ")
			case term.KeyCtrlC:
				fmt.Println("Process terminated.")
				termination <- 1
				term.Close()
			default:
				term.Sync()
			}
		case term.EventError:
			panic(ev.Err)
		}
		time.Sleep(time.Millisecond * 133)
	}
}

func UnixTimeRecordNano() []byte { //8byte
	// now := []byte(time.Now().Format(time.RFC3339Nano))
	now := make([]byte, 8)
	binary.LittleEndian.PutUint64(now, uint64(time.Now().UnixNano()))
	return now
}

func UnixTimeRecord() []byte { //4 byte
	// now := []byte(time.Now().Format(time.RFC3339Nano))
	now := make([]byte, 8)
	binary.LittleEndian.PutUint64(now, uint64(time.Now().Unix()))
	return now[:4]
}

func GenerateRandByteArray(bytes int) []byte {
	res := make([]byte, bytes)

	n, err := rand.Read(res) //Int(rand.Reader, big.NewInt(100))
	if n != bytes || err != nil {
		fmt.Println("Error: In utils.GenerateRandByteArray(),", err)
	}
	return res
}

func CreateConstLengthHeader(len int, space int) []byte {
	res := make([]byte, 8)
	binary.LittleEndian.PutUint64(res, uint64(len))
	return res[:space]
}

func FormattedTimeRecord() ([]byte, uint32) {
	// now := []byte(time.Now().Format(time.RFC3339Nano))
	now := make([]byte, 8)
	binary.LittleEndian.PutUint64(now, uint64(time.Now().UnixNano()))
	var timestamp []byte = now
	nowlen := len(now)

	if nowlen > 32 {
		timestamp = now[0:32]
	} else if nowlen < 32 {
		for i := 0; i < 32-nowlen; i++ {
			now = append(now, 0x00)
		}
		timestamp = now[0:32]
	}

	var ts [32]byte
	copy(ts[:], timestamp[0:32])
	return ts[:], uint32(32)
}

func FormattedTimeRecordConst() []byte {
	// now := []byte(time.Now().Format(time.RFC3339Nano))
	now := make([]byte, 8)
	binary.LittleEndian.PutUint64(now, uint64(time.Now().UnixNano()))
	var timestamp []byte = now
	nowlen := len(now)

	if nowlen > 32 {
		timestamp = now[0:32]
	} else if nowlen < 32 {
		for i := 0; i < 32-nowlen; i++ {
			now = append(now, 0x00)
		}
		timestamp = now[0:32]
	}

	var ts [32]byte
	copy(ts[:], timestamp[0:32])
	return ts[:]
}

func GetUnixNanoDiff(before, after []byte) time.Duration {
	bf := time.Unix(0, int64(binary.LittleEndian.Uint64(before)))
	af := time.Unix(0, int64(binary.LittleEndian.Uint64(after)))

	return af.Sub(bf)
}

func NowUnixNanoLittleEndian() []byte {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(time.Now().UnixNano()))
	fmt.Println(buf)
	return buf
}

func GenAIDSIDLittleEndian(Aid, Sid uint32) []byte {
	bufA := make([]byte, 4)
	bufS := make([]byte, 4)
	binary.LittleEndian.PutUint32(bufA, Aid)
	binary.LittleEndian.PutUint32(bufS, Sid)
	return append(bufA, bufS...)
}

package utils

import (
	"bytes"
	"crypto/rand"

	"encoding/binary"
	"encoding/gob"
	"fmt"
	"strconv"
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
	return buf
}

func GenAIDSIDLittleEndian(Aid, Sid uint32) []byte {
	bufA := make([]byte, 4)
	bufS := make([]byte, 4)
	binary.LittleEndian.PutUint32(bufA, Aid)
	binary.LittleEndian.PutUint32(bufS, Sid)
	return append(bufA, bufS...)
}

func GetUint8FromString2Map(field string, m map[string]string) uint8 { //PuiPuiMapCmd
	tmp, err := strconv.ParseUint(m[field], 10, 8)
	if err != nil {
		panic(err)
	}
	return uint8(tmp)
}

func GetUint32FromString2Map(field string, m map[string]string) uint32 { //PuiPuiMapCmd
	tmp, err := strconv.ParseUint(m[field], 10, 32)
	if err != nil {
		panic(err)
	}
	return uint32(tmp)
}

func GetUint64FromString2Map(field string, m map[string]string) uint64 { //PuiPuiMapCmd
	tmp, err := strconv.ParseUint(m[field], 10, 64)
	if err != nil {
		panic(err)
	}
	return uint64(tmp)
}

func GobEncoderOutString(v interface{}) string {
	buf := new(bytes.Buffer)
	err := gob.NewEncoder(buf).Encode(&v)
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(buf)
}

func GobDecoderUint32Array(val string) []uint32 {
	var resv interface{}
	// var res []uint32
	buf := bytes.NewBuffer(nil)
	buf.WriteString(val)
	err := gob.NewDecoder(buf).Decode(&resv)
	if err != nil {
		panic(err)
	}
	buf.Reset()
	return resv.([]uint32)
}

func GobDecoderUint64Array(val string) []uint64 {
	var resv interface{}
	buf := bytes.NewBuffer(nil)
	buf.WriteString(val)
	err := gob.NewDecoder(buf).Decode(&resv)
	if err != nil {
		panic(err)
	}
	buf.Reset()
	return resv.([]uint64)
}

func GobDecoderBoolArray(val string) []bool {
	var res []bool
	buf := bytes.NewBuffer(nil)
	buf.WriteString(val)
	err := gob.NewDecoder(buf).Decode(&res)
	if err != nil {
		panic(err)
	}
	buf.Reset()
	return res
}

func GobDecoderByteByteArray(val string) [][]byte {
	var res [][]byte
	buf := bytes.NewBuffer(nil)
	buf.WriteString(val)
	err := gob.NewDecoder(buf).Decode(&res)
	if err != nil {
		panic(err)
	}
	buf.Reset()
	return res
}

func GobDecoderStringArray(val string) []string {
	var res []string
	buf := bytes.NewBuffer(nil)
	buf.WriteString(val)
	err := gob.NewDecoder(buf).Decode(&res)
	if err != nil {
		panic(err)
	}
	buf.Reset()
	return res
}

func StringConvertToBool(val string) bool {

	return val == "true"

}

func ByteArrayConvertToUint(val []byte) (res int) {

	switch leng := len(val); {
	case leng <= 0:
		fmt.Println("In ByteArrayConvertToUint, Error: Input length <= 0")
		res = 0

	case leng <= 2:
		if leng < 2 {
			val = append(val, []byte{0}...)
		}
		res = int(binary.LittleEndian.Uint16(val))

	case leng <= 4:
		if leng < 4 {
			for i := 0; i < 4-leng; i++ {
				val = append(val, []byte{0}...)
			}
		}
		res = int(binary.LittleEndian.Uint32(val))

	case leng <= 8:
		if leng < 8 {
			for i := 0; i < 8-leng; i++ {
				val = append(val, []byte{0}...)
			}
		}
		res = int(binary.LittleEndian.Uint64(val))

	default:
		fmt.Println("In ByteArrayConvertToUint, Error: Input length > 8\ndonno what to do.")
		res = 0
	}
	return res
}

func UintConvertToLittleEndianByteArray(v interface{}) (res []byte) {
	switch i := v.(type) {
	case uint16:
		res = make([]byte, 2)
		binary.LittleEndian.PutUint16(res, i)
		return res
	case uint32:
		res = make([]byte, 4)
		binary.LittleEndian.PutUint32(res, i)
		return res
	case uint64:
		res = make([]byte, 8)
		binary.LittleEndian.PutUint64(res, i)
		return res
	default:
		println("Error in UintConvertToByteArray(), unknown input type!")
		return nil
	}
}

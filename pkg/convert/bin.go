package convert

import (
	"strconv"
)

// 十进制转二进制
func Bin(i int) []string {
	var sList []string
	for {
		binv := i % 2
		
		s := []string{strconv.Itoa(binv)}
		
		if i > 1 {
			sList = append(s, sList...)
			i = i / 2
		} else {
			sList = append(s, sList...)
			break
		}
	}
	return sList
}

// 1 Byte = 8 Bit, 传入 Byte, 转成带0的8个bit 的 string 格式.
func ByteToBit(i int64) string {
	binStr := strconv.FormatInt(i, 2)
	binStrLength := len(binStr)

	binStrLengthDiff := 8 - binStrLength
	
	for i := 1; i <= binStrLengthDiff; i++ {
		
		binStr = "0" + binStr
	}

	return binStr
}

// 1 Byte = 8 Bit, 传入 Byte, 转成带0的8个bit 的 string 格式.
func PayloadToHex(payload []byte) []string {
	var newPayload []string
	for _, v := range payload {
		hexStr := strconv.FormatInt(int64(v), 16)
		newPayload = append(newPayload, hexStr)
	}
	return newPayload
}

// []byte 转换 []string
func PayloadToBin(payload []byte) []string {
	var newPayload []string
	for _, v := range payload {
		vBin := ByteToBit(int64(v))
		newPayload = append(newPayload, vBin)
	}
	
	return newPayload
}

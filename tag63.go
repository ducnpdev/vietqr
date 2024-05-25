package vietqr

import "fmt"

// CRC (Cyclic Redundancy Check)
func genCRC(fullContent string) string {
	prefix := "6304"
	table := MakeTable(CRC16_CCITT_FALSE)
	h := New(table)
	h.Write([]byte(fullContent + prefix))
	sum := h.Sum16()
	crc := fmt.Sprintf("%X", sum)
	return fmt.Sprintf("%s%s", prefix, crc)
}

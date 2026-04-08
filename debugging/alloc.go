package debugging

import (
	"fmt"
	"runtime"
)

func getMessageType(msg []byte) []byte {
	return msg[:5:5] // full slice expression to be sure to get a new slice header with the same underlying array
}

func PrintAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}

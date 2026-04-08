package debugging

import (
	"fmt"
	"runtime"
)

func debugMap() {
	n := 1_000_000
	m := make(map[int][128]byte, n)
	PrintAlloc()

	for i := 0; i < n; i++ {
		m[i] = [128]byte([]byte{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		})
	}

	PrintAlloc()

	for i := 0; i < n; i++ {
		delete(m, i)
	}

	runtime.GC() // force garbage collection
	PrintAlloc()
	runtime.KeepAlive(m)
}

func addToMap() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}
	m2 := copyMap(m)

	for k, v := range m {
		m2[k] = v
		if v {
			m2[10+k] = true
		}
	}
	fmt.Println(m2)
}

func copyMap(m map[int]bool) map[int]bool {
	m2 := make(map[int]bool, len(m))
	for k, v := range m {
		m2[k] = v
	}
	return m2
}

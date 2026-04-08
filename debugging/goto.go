package debugging

import (
	"context"
	"fmt"
)

func main() {
loop:
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		switch i {
		default:
		case 2:
			break loop // break out of the loop
		}
	}

	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

loop2: // Определяется метка loop
	for {
		select {
		case <-ch:
			// Какие-то действия
		case <-ctx.Done():
			break loop2 // Прерывается выполнение привязанного к loop цикла, а не select
		}
	}

}

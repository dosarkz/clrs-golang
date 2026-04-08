package debugging

import "fmt"

func main() {
	defer func() { // Вызовы восстанавливаются внутри отложенного замыкания
		if r := recover(); r != nil {
			fmt.Println("recover", r)
		}
	}()
	f()
}
func f() {
	fmt.Println("a")
	panic("foo")
	fmt.Println("b")
	//Вызов f, которая запускает панику. Паника
	//отлавливается предыдущим восстановлением
}

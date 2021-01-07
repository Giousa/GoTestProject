package main

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for{
			select {
			case num1 := <-ch1:
				ch2 <- num1
			}
		}
	}()

	go func() {
		for{
			select {
			case num2 := <-ch2:
				ch1 <- num2
			}
		}
	}()
}

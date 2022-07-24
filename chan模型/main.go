package main

import "fmt"

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		data := i * i
		fmt.Println("生产者生产数据:", data)
		out <- data
	}
	close(out)
}

func consumer(in <-chan int) {

	for data := range in {
		fmt.Println("消费者得到数据：", data)
	}

}

func main() {

	ch := make(chan int, 5)

	go producer(ch)
	consumer(ch)
}

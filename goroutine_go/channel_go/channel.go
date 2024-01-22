package channel_go

import (
	"fmt"
	"time"
)

func BlockChannel() {
	//定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine1 结束")
		fmt.Println("goroutine1 正在运行...")
		time.Sleep(3 * time.Second)
		c <- 777 //将777发送给c
	}()

	go func() {
		defer fmt.Println("goroutine2 结束")
		fmt.Println("goroutine2 正在运行...")
		num := <-c //从c中接收数据，并赋值给num
		fmt.Println("goroutine2 接收到数据, num = ", num)
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}

func NonBlockChannel() {
	c := make(chan int, 3)
	fmt.Println("len =", len(c), "cap =", cap(c))

	go func() {
		defer fmt.Println("子goroutine结束")

		for i := 0; i < 10; i++ {
			if len(c) == cap(c) {
				fmt.Println("子goroutine发送阻塞")
			}
			c <- i
			fmt.Println("子goroutine正在运行，发送的元素:", i, "len =", len(c), "cap =", cap(c))
			time.Sleep(2 * time.Second)
		}
	}()

	for i := 0; i < 10; i++ {
		if len(c) == 0 {
			fmt.Println("主goroutine接收阻塞")
		}
		num := <-c
		fmt.Println("主接收的元素:", num, "len =", len(c), "cap =", cap(c))
		time.Sleep(1 * time.Second)
	}
}

func CloseChannel() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		//close可以关闭一个channel
		close(c)
	}()

	// for {
	// 	//ok返回的是channel是否关闭
	// 	if data, ok := <-c; ok {
	// 		fmt.Println("recrive data :", data)
	// 	} else {
	// 		break
	// 	}
	// }

	//range会阻塞等待c中的数据，可以使用range来迭代不断操作channel
	for data := range c {
		fmt.Println("recrive data :", data)
	}

	fmt.Println("main finished")
}

func fibonacii(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			//如果c可写，则该case就会进来，并且x会写入c
			pre_x := x
			x = y
			y = pre_x + y
		case <-quit:
			//如果quit可读，就退出
			fmt.Println("quit")
			return
		}
	}
}

func SelectAndChannel() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	//main
	fibonacii(c, quit)
}

// chan<-  只写
func counter(out chan<- int) {
	defer close(out)
	for i := 0; i < 5; i++ {
		out <- i //如果对方不读，会阻塞
	}
}

// <-chan  只读
func printer(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}

func SingleChannel() {
	c := make(chan int) //读写

	go counter(c) //生产者
	printer(c)    //消费者

	fmt.Println("finished")
}

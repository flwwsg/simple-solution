package interview

//启动两个线程, 一个输出 1,3,5,7…99, 另一个输出 2,4,6,8…100 最后 STDOUT 中按序输出 1,2,3,4,5…100

func PrintIntInTwoThread() {
	num := make(chan int)
	exit := make(chan int)
	go func() {
		for i := 1; i < 101; i++ {
			//打印奇数
			println("g1:", <-num)
			i++
			num <- i
		}
	}()
	go func() {
		defer func() {
			close(num)
			close(exit)
		}()
		for i := 0; i < 100; i++ {
			i++
			num <- i
			println("g2:", <-num)
		}

	}()
	<-exit
}

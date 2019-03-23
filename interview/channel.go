package interview

//启动两个线程, 一个输出 1,3,5,7…99, 另一个输出 2,4,6,8…100 最后 STDOUT 中按序输出 1,2,3,4,5…100

func PrintIntInTwoChannel() {
	maxInt := 101
	//打印偶数
	isEven := make(chan bool)
	//打印奇数
	isOdd := make(chan bool)
	res := make(chan int)
	go func() {
		i := 0
		for {
			select {
			case <-isEven:
				i += 2
				println(i)
				res <- i
			}
		}

	}()
	go func() {
		i := 1
		for {
			select {
			case <-isOdd:
				println(i)
				i += 2
				res <- i
			}
		}

	}()
	for i := 1; i < maxInt; i++ {
		if i%2 == 0 {
			isEven <- true
		} else {
			isOdd <- true
		}
		//等待结果
		select {
		case <-res:

		}
	}
}

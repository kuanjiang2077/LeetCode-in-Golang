type ZeroEvenOdd struct {
	n int
	evenCh chan int
	oddCh chan int
	zeroCh chan int
	endCh chan int
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	c1, c2, c3, c4 := make(chan int), make(chan int),make(chan int),make(chan int)
	return &ZeroEvenOdd{
		n:n,
		evenCh: c1,
		oddCh: c2,
		zeroCh: c3,
		endCh: c4,
	}
}

func printNumber(x int)  {
	fmt.Print(x)
}

// You may call global function `printNumber(int x)`
// to output "x", where x is an integer.

func (z *ZeroEvenOdd) Zero() {
	for {
		nn := <- z.zeroCh

		printNumber(0)
		nn++
		if nn % 2 == 1 {
			z.oddCh <- nn
		} else {
			z.evenCh <- nn
		}
		if nn >= z.n {
			return
		}
	}

}

func (z *ZeroEvenOdd) Even() {
	for {
		nn := <- z.evenCh
		printNumber(nn)
		if nn >= z.n {
			z.endCh <- 1
			return
		}
		z.zeroCh <- nn
	}

}

func (z *ZeroEvenOdd) Odd() {
	for {
		nn := <- z.oddCh
		printNumber(nn)
		if nn >= z.n {
			z.endCh <- 1
			return
		}
		z.zeroCh <- nn
	}

}

func (z *ZeroEvenOdd) Start() {
	z.zeroCh <- 0
}

func (z *ZeroEvenOdd) Done()  {
	<- z.endCh
}
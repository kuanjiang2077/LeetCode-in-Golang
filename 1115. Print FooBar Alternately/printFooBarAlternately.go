package fooBar

type FooBar struct {
	n        int
	fooChan  chan int
	barChan  chan int
	taskChan chan int
}

type PrintFoo interface {
	Run()
}

type PrintBar interface {
	Run()
}

func NewFooBar(n int) *FooBar {
	c1, c2, c3 := make(chan int), make(chan int), make(chan int)
	return &FooBar{
		n:        n,
		fooChan:  c1,
		barChan:  c2,
		taskChan: c3,
	}
}

func (fb *FooBar) Foo(foo PrintFoo) {
	for i := 0; i < fb.n; i++ {
		<-fb.fooChan
		foo.Run()
		fb.barChan <- 1
	}
}
func (fb *FooBar) Start() {
	fb.fooChan <- 1
}
func (fb *FooBar) Done() {
	<-fb.taskChan
}

func (fb *FooBar) Bar(bar PrintBar) {
	for i := 0; i < fb.n; i++ {
		<-fb.barChan
		bar.Run()
		if i < fb.n-1 {
			fb.fooChan <- 1
		} else {
			fb.taskChan <- 1
		}
	}
}

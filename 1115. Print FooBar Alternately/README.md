Usage:
``` go
func main() {
	fb := fooBar.NewFooBar(3)
	go bar(fb)
	go foo(fb)
	fb.Start()
	fb.Done()
}

type myFoo struct {
}

type myBar struct {
}

func (myFoo) Run() {
	fmt.Print("foo")
}

func (myBar) Run() {
	fmt.Print("bar")
}

func foo(fb *fooBar.FooBar) {
	myFoo := myFoo{}
	fb.Foo(myFoo)
}

func bar(fb *fooBar.FooBar) {
	myBar := myBar{}
	fb.Bar(myBar)
}
```
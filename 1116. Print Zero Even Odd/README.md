Usage:
```go
func main() {
	z := NewZeroEvenOdd(19)
	go z.Odd()
	go z.Even()
	go z.Zero()
	z.Start()

	z.Done()
}
```
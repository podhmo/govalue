# govalue
write go value as go code

```go
func ExmapleToCode() {
	xs := []int{1, 2, 3}
	fmt.Printf("xs := %s\n", govalue.ToCode(xs))
	ys := []string{"foo", "bar", "boo"}
	fmt.Printf("ys := %s\n", govalue.ToCode(ys))

	// Output:
	// xs := []int{1, 2, 3}
	// ys := []string{"foo", "bar", "boo"}
}
```
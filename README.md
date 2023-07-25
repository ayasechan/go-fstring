
# fsting

Simple, loose string interpolation.

# install

`go get -u github.com/ayasechan/go-fstring`

# example


```go
func TestFstring(t *testing.T) {
	got := FString("{ a } {b}", M{"a": "a"})
	assert.Equal(t, "a {b}", got)
}
```
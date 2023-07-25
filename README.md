
# go-fstring

Simple, loose string interpolation.

# Install

`go get -u github.com/ayasechan/go-fstring`

# Example


```go
func TestFstring(t *testing.T) {
	got := FString("{{ a } {b} {{}} }{", M{"a": "a"})
	assert.Equal(t, "{a {b} {{}} }{", got)
}
```

# Benchmark

```txt
goos: windows
goarch: amd64
pkg: github.com/ayasechan/go-fstring
cpu: 12th Gen Intel(R) Core(TM) i3-12100
BenchmarkFstring-8        112248             11012 ns/op            5984 B/op         36 allocs/op
PASS
ok      github.com/ayasechan/go-fstring 1.385s
```
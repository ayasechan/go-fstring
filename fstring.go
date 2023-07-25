package fstring

import (
	"fmt"
	"io"
	"strings"
)

type M map[string]string

func WithMap(reader io.Reader, m M) (io.Reader, error) {
	t := new(Template)
	err := t.Compile(reader)
	if err != nil {
		return nil, err
	}
	return t.WithFunc(func(raws, keys []string, w io.Writer) {
		for i, s := range raws {
			fmt.Fprint(w, s)
			if i < len(keys) {
				key := keys[i]
				v, ok := m[strings.Trim(key, " ")]
				if !ok {
					v = fmt.Sprintf("{%s}", key)
				}
				fmt.Fprint(w, v)
			}
		}
	}), nil
}

func FString(template string, m M) string {
	r, err := WithMap(strings.NewReader(template), m)
	if err != nil {
		panic(err)
	}
	buf, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return string(buf)
}

var F = FString

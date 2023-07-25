package fstring

import (
	"bufio"
	"bytes"
	"io"
	"sync"
)

type Template struct {
	raws []string
	keys []string
}

func NewTemplate() *Template {
	return new(Template)
}
func (t *Template) WithFunc(f func(raws, keys []string, w io.Writer)) io.ReadCloser {
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		f(t.Raws(), t.Keys(), pw)
	}()
	return pr
}

func (t *Template) Raws() []string {
	return append([]string{}, t.raws...)
}

func (t *Template) Keys() []string {
	return append([]string{}, t.keys...)
}

var bufPool = sync.Pool{New: func() any { return new(bytes.Buffer) }}

func (t *Template) Compile(reader io.Reader) error {
	rawBuf := bufPool.Get().(*bytes.Buffer)
	keyBuf := bufPool.Get().(*bytes.Buffer)
	defer func() {
		rawBuf.Reset()
		keyBuf.Reset()
		bufPool.Put(rawBuf)
		bufPool.Put(keyBuf)
	}()
	br := bufio.NewReader(reader)
	for {

		char, _, err := br.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		switch char {
		case '{':
			if keyBuf.Len() > 0 {
				io.Copy(rawBuf, keyBuf)
				keyBuf.Reset()
			}
			keyBuf.WriteRune(char)
		case '}':
			if keyBuf.Len() > 0 {
				keyBuf.Next(1)
				t.raws = append(t.raws, rawBuf.String())
				t.keys = append(t.keys, keyBuf.String())
				rawBuf.Reset()
				keyBuf.Reset()
			} else {
				rawBuf.WriteRune(char)
			}
		default:
			if keyBuf.Len() > 0 {
				keyBuf.WriteRune(char)
			} else {
				rawBuf.WriteRune(char)
			}
		}
	}
	if keyBuf.Len() > 0 {
		io.Copy(rawBuf, keyBuf)
	}
	// always len(t.raws) > len(t.keys)
	t.raws = append(t.raws, rawBuf.String())
	return nil
}

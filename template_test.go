package fstring

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompile(t *testing.T) {
	s := "{%0} {{%1} {%a} {%3} {{  }} }{"
	tpl := NewTemplate()
	err := tpl.Compile(strings.NewReader(s))
	assert.Nil(t, err)
	assert.Greater(t, len(tpl.raws), len(tpl.Keys()))
	assert.Equal(t, []string{"", " {", " ", " ", " {", "} }{"}, tpl.raws)
	assert.Equal(t, []string{"%0", "%1", "%a", "%3", "  "}, tpl.keys)
}

package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParser_Parse_single(t *testing.T) {
	p := Parser{AllStructs: false}

	err := p.Parse("../example/models/user.go", false)
	require.NoError(t, err)

	fmt.Println(p)
}
func TestParser_Parse_all(t *testing.T) {
	p := Parser{AllStructs: true}

	err := p.Parse("../example/models", true)
	require.NoError(t, err)

	fmt.Println(p)
}

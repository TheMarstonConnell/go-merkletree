package merkletree

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/wealdtech/go-merkletree/v2/blake3"
	"testing"
)

func TestSanity(t *testing.T) {
	r := require.New(t)

	h := "Hello World"
	data := make([][]byte, 0)
	data = append(data, []byte(h))

	hash := blake3.New256()
	tree, err := NewTree(
		WithData(data),
		WithHashType(hash),
	)

	r.NoError(err)

	root := tree.Root()
	fmt.Printf("%x\n", root)
}

package skiplist_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	skiplist "github.com/warrenb95/skip-list"
)

func TestSkipList(t *testing.T) {
	sl := skiplist.New()

	sl.Insert([]byte("hello"), []byte("world"))
	sl.Insert([]byte("foo"), []byte("bar"))
	sl.Insert([]byte("bye"), []byte("mars"))
	sl.Insert([]byte("blue"), []byte("planat"))

	node, err := sl.Find([]byte("hello"))
	require.NoError(t, err)
	require.NotNil(t, node)
	fmt.Println(string(node.Value))

	node, err = sl.Find([]byte("blue"))
	require.NoError(t, err)
	require.NotNil(t, node)
	fmt.Println(string(node.Value))
}

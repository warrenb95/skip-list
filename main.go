package skiplist

import (
	"bytes"
	"errors"
	"math"
	"math/rand"
)

// p = 50% = 1/2
// N = 65536 = 216
// MaxHeight = log1/p(N) = log1/(1/2)(216) = log2(216) = 16
// p is the fraction of nodes with level i pointers that also have level i+1 pointers;
// 	for example, when p = 1/2 we are saying that we want 50% (1/2) of the nodes occupying
// 	a specific level to also occupy the next level, etc. (more on this later).
// N is the predicted maximum number of nodes that you expect to store in the skip list.

const maxHeight = 16

var heightBuckets [maxHeight]uint32

type node struct {
	Key, Value []byte
	tower      [maxHeight]*node
}

type SkipList struct {
	head   *node
	height int
}

func New() *SkipList {
	n := &node{}
	sk := &SkipList{head: n, height: 1}
	return sk
}

func (s *SkipList) search(key []byte) (*node, [maxHeight]*node) {
	var journey [maxHeight]*node
	var target *node

	currentNode := s.head
	for level := s.height - 1; level >= 0; level-- {
		for nextNode := currentNode.tower[level]; nextNode != nil; nextNode = currentNode.tower[level] {
			if bytes.Equal(key, nextNode.Key) {
				target = nextNode
				break
			} else if bytes.Compare(key, nextNode.Key) <= 0 {
				break
			}

			currentNode = nextNode
		}

		journey[level] = currentNode
	}

	return target, journey
}

func init() {
	var val uint32 = math.MaxUint32
	heightBuckets[0] = val

	for i := 1; i < maxHeight; i++ {
		val = uint32(math.Round(float64(val) * 0.5))
		heightBuckets[i] = val
	}
}

func randomHeight() int {
	height := 1
	randNum := rand.Uint32()

	for height < maxHeight-1 && randNum >= heightBuckets[height] {
		height++
	}

	return height
}

func (s *SkipList) Insert(key, value []byte) {
	found, journey := s.search(key)

	if found != nil {
		found.Value = value
		return
	}

	height := randomHeight()
	newNode := &node{Key: key, Value: value, tower: [maxHeight]*node{}}

	for level := 0; level <= height; level++ {
		prevNode := journey[level]
		if prevNode == nil {
			prevNode = s.head
		}
		newNode.tower[level] = prevNode.tower[level]
		prevNode.tower[level] = newNode
	}

	if height > s.height {
		s.height = height
	}
}

func (s *SkipList) Find(key []byte) (*node, error) {
	n, _ := s.search(key)
	if n == nil {
		return nil, errors.New("not found")
	}

	return n, nil
}

package skiplists

import (
	"math/rand"
)

type node struct {
	next     []*node
	key, val int
}

type Skiplist1 struct {
	head *node
}

func getLevel() int {
	level := 0
	for rand.Intn(2) > 0 {
		level++
	}
	return level
}

func NewSkiplist() *Skiplist1 {
	return &Skiplist1{
		head: &node{
			next: make([]*node, 0),
			key:  -1,
			val:  -1,
		},
	}
}

func (s *Skiplist1) serach(target int) *node {
	c := s.head
	for level := len(s.head.next) - 1; level >= 0; level-- {
		for c.next[level] != nil && c.next[level].key < target {
			c = c.next[level]
		}

		if c.next[level] != nil && c.next[level].key == target {
			return c.next[level]
		}
	}
	return nil
}

func (s *Skiplist1) Get(key int) (int, bool) {
	if node := s.serach(key); node != nil {
		return node.val, true
	}
	return -1, false
}

func (s *Skiplist1) Put(key, val int) {
	if node := s.serach(key); node != nil {
		node.val = val
		return
	}

	level := getLevel()

	for level > len(s.head.next)-1 {
		s.head.next = append(s.head.next, nil)
	}

	newNode := node{
		key:  key,
		val:  val,
		next: make([]*node, level+1),
	}
	move := s.head

	for ; level >= 0; level-- {
		for move.next[level] != nil && move.next[level].key < key {
			move = move.next[level]
		}
		newNode.next[level] = move.next[level]
		move.next[level] = &newNode
	}

	return
}

func (s *Skiplist1) Remove(key, val int) {
	node := s.serach(key)
	if node == nil || (node != nil && node.val != val) {
		return
	}

	level := len(node.next)
	move := s.head
	for ; level >= 0; level-- {
		for move.next[level] != nil && move.next[level].key < key {
			move = move.next[level]
		}

		if move.next[level] != nil && move.next[level].key == key {
			move.next[level], move.next[level].next[level] = move.next[level].next[level], nil
		}
	}

	move = s.head
	for move.next[len(move.next)-1].next == nil {
		move.next = move.next[:len(move.next)-1]
	}
	return
}

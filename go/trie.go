package main

import "fmt"

const chars int = 26

type Node struct {
	children [chars]*Node
	isEnd    bool
}
type Trie struct {
	root *Node
}

func initTrie() *Trie {
	result := &Trie{&Node{}}
	return result
}

func (t *Trie) insert(w string) {
	l := len(w)
	curr := t.root
	for i := 0; i < l; i++ {
		ci := w[i] - 'a'
		if curr.children[ci] == nil {
			curr.children[ci] = &Node{}
		}
		curr = curr.children[ci]
	}
	curr.isEnd = true

}

func (t *Trie) search(w string) bool {
	l := len(w)
	curr := t.root
	for i := 0; i < l; i++ {
		ci := w[i] - 'a'
		if curr.children[ci] == nil {
			return false
		}
		curr = curr.children[ci]
	}
	return curr.isEnd

}

func main() {
	t := initTrie()
	t.insert("cross")
	t.insert("crossed")
	t.insert("core")
	t.insert("care")

	fmt.Println(t.search("cross"))
	fmt.Println(t.search("crossed"))
	fmt.Println(t.search("core"))
	fmt.Println(t.search("care"))

	fmt.Println(t.search("croses"))
	fmt.Println(t.search("cwossed"))
	fmt.Println(t.search("cor"))
	fmt.Println(t.search("ca"))

}

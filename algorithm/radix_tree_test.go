// Package algorithm
package algorithm

import (
	"fmt"
	"testing"
)

func TestRadixTree(t *testing.T) {
	tTire := newTire()

	tTire.insert("user")
	tTire.insert("upload")
	tTire.insert("username")

	fmt.Println(tTire.search("u"))
	fmt.Println(tTire.search("user"))
	fmt.Println(tTire.search("pload"))

	tTire.printTree()
}

// Package algorithm 基数树
package algorithm

import "fmt"

// 文档
// https://en.wikipedia.org/wiki/Radix_tree

type tireNode struct {
	children map[rune]*tireNode
	isEnd    bool
}

type tire struct {
	root *tireNode
}

func newTire() *tire {
	return &tire{
		root: &tireNode{
			children: make(map[rune]*tireNode),
		},
	}
}

// insert 向 tire 中插入一个单词
func (t *tire) insert(word string) {
	node := t.root
	fmt.Println("插入单词:", word)
	for _, char := range word {
		// 判断这个字符的节点是否存在
		fmt.Printf("检查字符: %c\n", char)
		if _, exists := node.children[char]; !exists {
			// 不存在，创建节点
			fmt.Printf("字符 %c 节点不存在，创建新节点\n", char)
			node.children[char] = &tireNode{
				children: make(map[rune]*tireNode),
			}
		}
		// 下一个字符
		node = node.children[char]
	}

	node.isEnd = true
	fmt.Printf("单词 '%s' 插入完毕\n", word)
}

// search 搜索某个单词是否在 tire 中
func (t *tire) search(word string) bool {
	node := t.root
	fmt.Println("搜索单词:", word)
	for _, char := range word {
		// 判断下是否存在某个节点
		fmt.Printf("检查字符: %c\n", char)
		if _, exists := node.children[char]; !exists {
			fmt.Println("字符不存在，搜索失败")
			return false
		}
		// 下一个字符
		node = node.children[char]
	}

	if node.isEnd {
		fmt.Println("单词存在")
		return true
	} else {
		fmt.Println("单词不存在")
		return false
	}
}

// printTree 打印 Trie 树的结构
func (t *tire) printTree() {
	printNode(t.root, "")
}

// printNode 递归打印每个节点
func printNode(node *tireNode, prefix string) {
	if node == nil {
		return
	}
	if node.isEnd {
		fmt.Printf("单词: %s\n", prefix)
	}
	for char, child := range node.children {
		printNode(child, prefix+string(char))
	}
}

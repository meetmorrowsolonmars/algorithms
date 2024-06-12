package main

import "fmt"

/* Найти наименьшего общего предка (LCA) для двух узлов в бинарном дереве (не BST).
     3
    / \
-> 5   1
  / \ / \
 6  2 0  8
   / \
  7   4 <-
*/

func main() {
	// Дерево
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5, Parent: root}
	root.Right = &TreeNode{Val: 1, Parent: root}

	root.Left.Left = &TreeNode{Val: 6, Parent: root.Left}
	root.Left.Right = &TreeNode{Val: 2, Parent: root.Left}

	root.Left.Right.Left = &TreeNode{Val: 7, Parent: root.Left.Right}
	root.Left.Right.Right = &TreeNode{Val: 4, Parent: root.Left.Right}

	root.Right.Left = &TreeNode{Val: 0, Parent: root.Right}
	root.Right.Right = &TreeNode{Val: 8, Parent: root.Right}

	// Заданные вершины
	p := root.Left             // Вершина со значением 5
	q := root.Left.Right.Right // Вершина со значением 4

	ancestor := lowestCommonAncestor(p, q)
	if ancestor != nil {
		fmt.Printf("The lowest common ancestor of %d and %d is %d\n", p.Val, q.Val, ancestor.Val)
	} else {
		fmt.Println("No common ancestor found")
	}
}

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

// Функция для нахождения высоты узла
func findHeight(node *TreeNode) int {
	height := 0
	for node != nil {
		height++
		node = node.Parent
	}
	return height
}

// Функция для подъема узла на заданное количество уровней вверх
func moveUp(node *TreeNode, levels int) *TreeNode {
	for levels > 0 && node != nil {
		node = node.Parent
		levels--
	}
	return node
}

// Функция для нахождения наименьшего общего предка
func lowestCommonAncestor(p, q *TreeNode) *TreeNode {
	heightP := findHeight(p)
	heightQ := findHeight(q)

	// Выравниваем узлы по высоте
	if heightP > heightQ {
		p = moveUp(p, heightP-heightQ)
	} else if heightQ > heightP {
		q = moveUp(q, heightQ-heightP)
	}

	// Поднимаем оба узла вверх, пока не найдем общего предка
	for p != q {
		p = p.Parent
		q = q.Parent
	}

	return p
}

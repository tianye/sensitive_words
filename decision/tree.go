package decision

import (
	"fmt"
	"strings"
)

type Node struct {
	Word     string
	Node     [] *Node
	Location int

	IsSensitive bool
}

type Tree struct {
	TreeNode [] *Node
}

//创建一个叶子节点
func CreateNode(word string, location int, isSensitive bool) *Node {
	return &Node{Word: word, Location: location, IsSensitive: isSensitive}
}

//创建一个树
func CreateTree() *Tree {
	tree := &Tree{}

	return tree
}

//查找一个Node节点
func SearchNode(str string, nodeList [] *Node) *Node {
	//查找当前层级的所有node
	for _, v := range nodeList {
		//存在则直接返回当前node
		if v.Word == str {
			return v
		}
	}

	return nil
}

//插入一个子节点
func AppendNode(nowNode, newNode *Node) (*Node) {
	nowNode.Node = append(nowNode.Node, newNode)
	return newNode
}

//Build一个树
func BuildTrue(str string, tree *Tree) *Tree {
	end := strings.Count(str, "") - 1
	var nowNode = &Node{}

	var i = 0
	for _, val := range str {
		i++

		isSensitive := false
		if i == end {
			isSensitive = true
		}

		newNode := CreateNode(string(val), i, isSensitive)

		if i == 1 {
			nowNode = SearchNode(newNode.Word, tree.TreeNode)
			if nowNode != nil {
				continue
			}

			tree.TreeNode = append(tree.TreeNode, newNode)
			nowNode = newNode

			continue
		}

		if nowNode.Node != nil {
			nowNode := SearchNode(newNode.Word, nowNode.Node)
			if nowNode != nil {
				continue
			}
		}

		nowNode = AppendNode(nowNode, newNode)
	}

	return tree
}

//匹配的敏感词汇
func MatchingSensitiveWords(tree *Tree, str string) (isSensitive bool, allLocationStr []int) {
	node := tree.TreeNode
	isSensitive = false
	locationStr := make([]int, 0)
	allLocationStr = make([]int, 0)

	var i = 0
	for _, v := range str {
		i++
		node, isSensitive = SearchLeavesNode(string(v), node)

		//没有下一个节点了 并且当前不是敏感词
		if node == nil && isSensitive == false {
			//节点回到最初
			node = tree.TreeNode
			//当前子重新匹配
			node, isSensitive = SearchLeavesNode(string(v), node)
			//记录新的本次匹配地址
			locationStr = []int{i}
		} else {
			//追加记录位置
			locationStr = append(locationStr, i)
		}

		//如果是敏感词则记录到位置中
		if isSensitive == true {
			node = tree.TreeNode

			allLocationStr = append(allLocationStr, locationStr...)
			locationStr = []int{}
		}
	}

	//匹配到了关键词
	if len(allLocationStr) > 0 {
		isSensitive = true
	}

	return isSensitive, allLocationStr
}

//搜索层级关键字
func SearchLeavesNode(str string, params []*Node) (node []*Node, isSensitive bool) {
	for _, node := range params {
		if node.Word == str {
			return node.Node, node.IsSensitive
		}
	}

	return nil, false
}

//观察树结构
func WatchPrint(params []*Node) {
	for _, watch := range params {
		fmt.Print("watch.Word:", " ", watch.Word, " ", watch.Location, watch.IsSensitive, "-----", watch, "\n")
		if watch.Node != nil {
			WatchPrint(watch.Node)
		}
	}
}

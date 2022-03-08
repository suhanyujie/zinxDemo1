package redBlackTree

// 红黑树的实现
// 参考：https://www.cnblogs.com/qxcheng/p/15505415.html
/*

## 性质
* （1）每个节点或者是黑色，或者是红色。
* （2）根节点是黑色。
* （3）每个叶子节点（NIL）是黑色。 [注意：这里叶子节点，是指为空(NIL或NULL)的叶子节点！]
* （4）如果一个节点是红色的，则它的子节点必须是黑色的。
* （5）从一个节点到该节点的子孙节点的所有路径上包含相同数目的黑节点。
	* 作用是：确保没有一条路径会比其他路径长出俩倍。因而，红黑树是相对是接近平衡的二叉树。
> 参考 https://www.cnblogs.com/skywang12345/p/3245399.html

## ref
* 红黑树(一)之 原理和算法详细介绍  https://www.cnblogs.com/skywang12345/p/3245399.html
* go 实现红黑树 https://www.jianshu.com/p/0319d7781814


## 关于旋转
* 左旋中的“左”，意味着“被旋转的节点将变成一个左节点”。同理，右旋中的“右”，意味着“被旋转的节点将变成一个右节点”
*/

const (
	ColorRed = iota + 1
	ColorBlack
)

type RBNode struct {
	color               uint8
	key, data           interface{}
	left, right, parent *RBNode
}

type RBTree struct {
	rootNode *RBNode
	len      uint32
	cmp      func(a, b interface{}) bool
}

// NewNode 实例化一个节点
func NewNode(key, val interface{}) *RBNode {
	return &RBNode{
		color:  ColorBlack,
		key:    key,
		data:   val,
		left:   nil,
		right:  nil,
		parent: nil,
	}
}

// NewRBTree 实例化一颗红黑树
func NewRBTree(cmpFunc func(a, b interface{}) bool) *RBTree {
	return &RBTree{
		cmp: cmpFunc,
	}
}

// Insert 节点的插入 todo
func (tree *RBTree) Insert(key, val interface{}) {

}

func isRed(node *RBNode) bool {
	return node != nil && node.color == ColorRed
}

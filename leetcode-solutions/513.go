/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    answer = root.Val
    maxDepth = 0
    if root.Left != nil {
        f(root.Left, 1)
    } 
    if root.Right != nil {
        f(root.Right, 1)
    }
    return answer
}

var maxDepth = 0
var answer = 0

func f(node *TreeNode, depth int) {
    fmt.Println(depth, node.Val)
    if maxDepth < depth {
        maxDepth = depth
        answer = node.Val
    }
    if node.Left != nil {
        f(node.Left, depth+1)
    }
    if node.Right != nil {
        f(node.Right, depth+1)
    }
}

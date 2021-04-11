/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode, sum int, path []int, ans *[][]int) {
    if root.Left == nil && root.Right == nil {
        if sum == root.Val {
            *ans = append(*ans, append(path, root.Val))    
        }
        return
    }
    
    if root.Left != nil {
        dfs(root.Left, sum - root.Val, append(append([]int{}, path...), root.Val), ans)    
    }
    
    if root.Right != nil {
        dfs(root.Right, sum - root.Val, append(append([]int{}, path...), root.Val), ans)    
    }
    
}

func pathSum(root *TreeNode, sum int) [][]int {
    if root == nil {
        return nil
    }
    ans := [][]int{}
    dfs(root, sum, nil, &ans)
    return ans
}

// @Title: Lowest Common Ancestor of a Binary Search Tree
// @Author: colinjxc
// @Date: 2022-02-01T04:55:48+08:00
// @URL: https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root != nil {
        if root.Val > p.Val && root.Val > q.Val {
            return lowestCommonAncestor(root.Left, p, q)
        } else if root.Val < p.Val && root.Val < q.Val {
            return lowestCommonAncestor(root.Right, p, q)
        }
    }
    return root
}

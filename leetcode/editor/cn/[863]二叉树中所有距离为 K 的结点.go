//给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 K 。 
//
// 返回到目标结点 target 距离为 K 的所有结点的值的列表。 答案可以以任何顺序返回。 
//
// 
//
// 
// 
//
// 示例 1： 
//
// 输入：root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2
//输出：[7,4,1]
//解释：
//所求结点为与目标结点（值为 5）距离为 2 的结点，
//值分别为 7，4，以及 1
//
//
//
//注意，输入的 "root" 和 "target" 实际上是树上的结点。
//上面的输入仅仅是对这些对象进行了序列化描述。
// 
//
// 
//
// 提示： 
//
// 
// 给定的树是非空的。 
// 树上的每个结点都具有唯一的值 0 <= node.val <= 500 。 
// 目标结点 target 是树上的结点。 
// 0 <= K <= 1000. 
// 
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 471 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    parent := map[*TreeNode]*TreeNode{}
    var setParent func(*TreeNode)
    setParent = func(root *TreeNode) {
        if root == nil {
            return
        }
        if root.Left != nil {
            parent[root.Left] = root
        }
        if root.Right != nil {
            parent[root.Right] = root
        }
        setParent(root.Right)
        setParent(root.Left)
    }
    setParent(root)
    var getRes func(*TreeNode, int, *TreeNode) []int
    getRes = func(target *TreeNode, depth int, from *TreeNode) []int {
        res := make([]int, 0)
        if depth == k {
            res = append(res, target.Val)
        }
        if target.Left != nil && from != target.Left {
            res = append(getRes(target.Left, depth + 1, target), res...)
        }

        if target.Right != nil && from != target.Right {
            res = append(getRes(target.Right, depth + 1, target), res...)
        }

        if parent[target] != nil && from != parent[target] {
            res = append(getRes(parent[target], depth + 1, target), res...)
        }
        return res
    }
    return getRes(target, 0, nil)
}
//leetcode submit region end(Prohibit modification and deletion)

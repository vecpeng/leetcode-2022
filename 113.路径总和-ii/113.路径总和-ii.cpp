/*
 * @lc app=leetcode.cn id=113 lang=cpp
 *
 * [113] 路径总和 II
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    vector<vector<int>> pathSum(TreeNode* root, int targetSum) {
        vector<vector<int>> result;
        if (root == nullptr) {
            return result;
        }
        if (root->left == nullptr && root->right == nullptr && root->val == targetSum) {
            result.push_back(vector<int>{root->val});
            return result;
        }
        for (auto& arr : pathSum(root->right, targetSum - root->val)) {
            arr.insert(arr.begin(), 1, root->val);
            result.push_back(arr);
        }
        for (auto& arr : pathSum(root->left, targetSum - root->val)) {
            arr.insert(arr.begin(), 1, root->val);
            result.push_back(arr);
        }
        return result;
    }
};
// @lc code=end


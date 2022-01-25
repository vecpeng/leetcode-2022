/*
 * @lc app=leetcode.cn id=23 lang=cpp
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        ListNode res;
        ListNode* temp = &res;
        while(true) {
            int i = -1;
            int min = INT_MAX;
            for (int j = 0; j < lists.size(); j++) {
                if (lists[j] != nullptr && lists[j]->val < min) {
                    i = j;
                    min = lists[j]->val;
                }
            }
            if (i != -1) {
                temp->next = lists[i];
                lists[i] = lists[i]->next;
                temp = temp->next;
            } else {
                break;
            }
        }
        return res.next;
    }
};
// @lc code=end


/**
 * Definition for singly-linked list.
 * 
 */
struct ListNode {
  int val;
  struct ListNode *next;
};

bool hasCycle(struct ListNode *head) {
  if(head == NULL || head->next == NULL) {
      return false;
  }
  struct ListNode* fast = head;
  struct ListNode* slow = head;

  while(fast->next != NULL && fast->next->next != NULL) {
    fast = fast->next->next;
    slow = slow->next;
    if(slow == fast) {
      return true;
    }
  }
  return false;
}
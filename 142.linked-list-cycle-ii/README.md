# Linked List Cycle II
## Description
Given a linked list, return the node where the cycle begins. If there is no cycle, return null.

To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is `-1`, then there is no cycle in the linked list.

## Note
`Do not modify` the linked list.

## Example
### Example 1
![](../images/142_example_1.png)
```
Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.
```

### Example 2
![](../images/142_example_2.png)
```
Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.
```

### Example 3
![](../images/142_example_3.png)
```
Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.
```

## Follow up
Can you solve it `without using extra space`?

## Tags
**[Linked List](https://leetcode.com/tag/linked-list)**

**[Two Pointers](https://leetcode.com/tag/two-pointers)**

# 662. Maximum Width of Binary Tree - Medium

```
Given a binary tree, write a function to get the maximum width of the given tree. The width of a tree is the maximum width among all levels. The binary tree has the same structure as a full binary tree, but some nodes are null.

The width of one level is defined as the length between the end-nodes (the leftmost and right most non-null nodes in the level, where the null nodes between the end-nodes are also counted into the length calculation.

Example 1:

Input: 

           1
         /   \
        3     2
       / \     \  
      5   3     9 

Output: 4
Explanation: The maximum width existing in the third level with the length 4 (5,3,null,9).
Example 2:

Input: 

          1
         /  
        3    
       / \       
      5   3     

Output: 2
Explanation: The maximum width existing in the third level with the length 2 (5,3).
Example 3:

Input: 

          1
         / \
        3   2 
       /        
      5      

Output: 2
Explanation: The maximum width existing in the second level with the length 2 (3,2).
Example 4:

Input: 

          1
         / \
        3   2
       /     \  
      5       9 
     /         \
    6           7
Output: 8
Explanation:The maximum width existing in the fourth level with the length 8 (6,null,null,null,null,null,null,7).


Note: Answer will in the range of 32-bit signed integer.
```

# Relative Topics
* Tree
* BFS


# Note
```
Time: 2020/07/10
```


# Solution
## 1. BFS

### Intuition
Tree can be repersentd by array also.
for a given node of idx, index of its left child is 2*idx+1, index of its right child is 2*idx+2
As we assign the index for each node, so for a given layer of the tree, we can get the left most node and right most node and calculate the width by (right - left + 1) 

### Complexity Analysis
*   Time Complexity: O(N) where N is the total number of the nodes in the tree.
  
*   Space Complexity: O(2^H) where H is the height of the tree. maxium space is the max node of the deepest layer.


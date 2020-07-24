# 103. Binary Tree Zigzag Level Order Traversal - Medium

```
Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its zigzag level order traversal as:
[
  [3],
  [20,9],
  [15,7]
]
```

# Relative Topics
* Stack
* Tree
* Breadth-first-Search


# Note
```
Time: 2020/07/24
```


# Solution
## 1. Stack + Breath-first Search

### Intuition
We run through the tree layer by layer. For each layer, the read from left to right, and append their child for the next layer. We have a flag to indicate the direction of the traversal, which turn oppsite when enter next layer. For each layer, we put the node value into a list and append into the result based on the traversal direction.



### Complexity Analysis
*   Time Complexity: O(N)
  
*   Space Complexity: O(2^(H-1))

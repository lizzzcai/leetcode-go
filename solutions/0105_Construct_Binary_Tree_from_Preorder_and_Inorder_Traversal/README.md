# 105. Construct Binary Tree from Preorder and Inorder Traversal - Medium

```
Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7
```

# Relative Topics
* Array
* Tree
* Depth-first Search


# Note
```
Time: 2020/09/26
```


# Solution
## 1. Breath-first Search

### Intuition
In preorder, first item is the root of the tree. [root, left-sub, right-sub]
In inorder: [left-sub, root, right-sub]

we first find the root from preorder, then find the index in inorder, then we can know the size of left sub tree, ten we can know the left-sub and right-sub in the preorder array. for each sub-tree, we resursively find the root.Left and root.Right
finally return the root.



### Complexity Analysis
*   Time Complexity: O(N)
  
*   Space Complexity: O(N)

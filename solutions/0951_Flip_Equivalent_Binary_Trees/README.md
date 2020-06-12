# 951. Flip Equivalent Binary Trees - Medium

```
For a binary tree T, we can define a flip operation as follows: choose any node, and swap the left and right child subtrees.

A binary tree X is flip equivalent to a binary tree Y if and only if we can make X equal to Y after some number of flip operations.

Write a function that determines whether two binary trees are flip equivalent.  The trees are given by root nodes root1 and root2.

 

Example 1:

Input: root1 = [1,2,3,4,5,6,null,null,null,7,8], root2 = [1,3,2,null,6,4,5,null,null,null,null,8,7]
Output: true
Explanation: We flipped at nodes with values 1, 3, and 5.
Flipped Trees Diagram
 

Note:

Each tree will have at most 100 nodes.
Each value in each tree will be a unique integer in the range [0, 99].
```

# Relative Topics
* Tree


# Note
```
Time: 2020/06/12

```


# Solution
## 1. Recursion

### Intuition
If root1 and root2 have the same root value, then we only need to check if their children are equal (up to ordering.)

### Algorithm
There are 3 cases:

If root1 or root2 is null, then they are equivalent if and only if they are both null.

Else, if root1 and root2 have different values, they aren't equivalent.

Else, let's check whether the children of root1 are equivalent to the children of root2. There are two different ways to pair these children.

### Complexity Analysis
*   Time Complexity: O(min(N1,N2)), where N1,N2 are the lengths of root1 and root2.
  
*   Space Complexity: O(min(H1, H2)), where H1, H2 are the heights of root1 and root2.
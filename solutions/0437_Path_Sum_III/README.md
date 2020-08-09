# 437. Path Sum III - Medium

```
You are given a binary tree in which each node contains an integer value.

Find the number of paths that sum to a given value.

The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).

The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.

Example:

root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

Return 3. The paths that sum to 8 are:

1.  5 -> 3
2.  5 -> 2 -> 1
3. -3 -> 11
```

# Relative Topics
* Tree
* DFS

# Note
```
Time: 2020/08/09

```


# Solution
## 1. Tree iterate + DFS

### Intuition

We iterate over each node of the tree, and start to dfs on each node.
for each dfs, we start from the tartget sum, and for each node, minus the value.
of the curr sum decrease to zero, then it is a possible path starting from the node.

### Complexity Analysis
*   Time Complexity: O(NM*H)
  
*   Space Complexity: O(MN)
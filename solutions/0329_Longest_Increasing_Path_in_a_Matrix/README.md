# 329. Longest Increasing Path in a Matrix - Hard

```
Given an integer matrix, find the length of the longest increasing path.

From each cell, you can either move to four directions: left, right, up or down. You may NOT move diagonally or move outside of the boundary (i.e. wrap-around is not allowed).

Example 1:

Input: nums = 
[
  [9,9,4],
  [6,6,8],
  [2,1,1]
] 
Output: 4 
Explanation: The longest increasing path is [1, 2, 6, 9].
Example 2:

Input: nums = 
[
  [3,4,5],
  [3,2,6],
  [2,2,1]
] 
Output: 4 
Explanation: The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.
```

# Relative Topics
* Depth-first Search
* Topological Sort
* Memoization


# Note
```
Time: 2020/06/27
```


# Solution
## 1. Depth-first Search

### Intuition
    We iterate each cell, and start from this cell, using dfs to find the next value larger than current, and find the max path and assign to the current cell. We use memoization to redue the duplication.

    for each DFS,
    curr = max_path start from four directions + 1
        

### Complexity Analysis
*   Time Complexity: O(MN)
  
*   Space Complexity: O(MN)

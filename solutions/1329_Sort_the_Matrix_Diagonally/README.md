# 1329. Sort the Matrix Diagonally - Medium

```
Given a m * n matrix mat of integers, sort it diagonally in ascending order from the top-left to the bottom-right then return the sorted array.

 

Example 1:


Input: mat = [[3,3,1,1],[2,2,1,2],[1,1,1,2]]
Output: [[1,1,1,1],[1,2,2,2],[1,2,3,3]]
 

Constraints:

m == mat.length
n == mat[i].length
1 <= m, n <= 100
1 <= mat[i][j] <= 100

```

# Relative Topics
* Array
* Sort


# Note
```
Time: 2020/07/12

```


# Solution
## 1. Sort

### Intuition
for given i, j in mat, element in the same diagonal has the same i-j
so we can collect the elements in the same diagonal into a map,
then sort them seperately and assign it back into the mat 

### Complexity Analysis
*   Time Complexity: O(mnlogD) where D is min(m,n)
  
*   Space Complexity: O(mn)
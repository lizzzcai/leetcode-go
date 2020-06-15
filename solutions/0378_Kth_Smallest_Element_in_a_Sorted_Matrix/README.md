# 378. Kth Smallest Element in a Sorted Matrix - Medium

```
Given a n x n matrix where each of the rows and columns are sorted in ascending order, find the kth smallest element in the matrix.

Note that it is the kth smallest element in the sorted order, not the kth distinct element.

Example:

matrix = [
   [ 1,  5,  9],
   [10, 11, 13],
   [12, 13, 15]
],
k = 8,

return 13.
Note:
You may assume k is always valid, 1 ≤ k ≤ n2.
```

# Relative Topics
* Binary Search
* Heap


# Note
```
Time: 2020/06/14

```


# Solution
## 1. Binary Search

### Intuition
Use binary search to search the range from smallest number to the biggest number, and count number smaller than mid


### Complexity Analysis
*   Time Complexity: O(NlogN)
  
*   Space Complexity: O(1)
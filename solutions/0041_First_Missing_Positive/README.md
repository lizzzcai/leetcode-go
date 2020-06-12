# 41. First Missing Positive - Hard

```
Given an unsorted integer array, find the smallest missing positive integer.

Example 1:

Input: [1,2,0]
Output: 3
Example 2:

Input: [3,4,-1,1]
Output: 2
Example 3:

Input: [7,8,9,11,12]
Output: 1
Note:

Your algorithm should run in O(n) time and uses constant extra space.
```

# Relative Topics
* Array


# Note
```
Time: 2020/06/13
```


# Solution

## 1. In place

### Intuition
The first posivite must be with [1, n+1]. we make all the element < 1 or > n become n+1.
we iterate the num, and make the value at idx = abs(num)-1 become negative.
we go through the array and find the first number which is not negative, and return the index + 1 as the first missing positive.

### Algorithm




### Complexity Analysis
*   Time Complexity: O(N).
  
*   Space Complexity: O(1).
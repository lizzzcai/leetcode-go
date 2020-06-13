# 137. Single Number II - Medium

```
Given a non-empty array of integers, every element appears three times except for one, which appears exactly once. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:

Input: [2,2,3,2]
Output: 3
Example 2:

Input: [0,1,0,1,0,1,99]
Output: 99
```

# Relative Topics
* Bit Manipulation


# Note
```
Time: 2020/06/13
```


# Solution
## 1. Bit Manipulation

### Intuition
iterate each bit, and count the 1 in all nums, if nums % 3 != 0, then this bit is 1 in res

### Complexity Analysis
*   Time Complexity: O(32N)
  
*   Space Complexity: O(1)

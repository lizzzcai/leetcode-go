# 171. Excel Sheet Column Number - Easy

```
Given a column title as appear in an Excel sheet, return its corresponding column number.

For example:

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28 
    ...
Example 1:

Input: "A"
Output: 1
Example 2:

Input: "AB"
Output: 28
Example 3:

Input: "ZY"
Output: 701
 

Constraints:

1 <= s.length <= 7
s consists only of uppercase English letters.
s is between "A" and "FXSHRXW".
```

# Relative Topics
* Math


# Note
```
Time: 2020/08/11
```


# Solution
## 1. Math

### Intuition
it is a base 26 operation. iterate over the string, compare it with int('A) to get the value starting from 1. convert it to base 10

base 10 = 26^0*a0 + 26^1*a1 + .. + 26^n*an

### Complexity Analysis
*   Time Complexity: O(N)
  
*   Space Complexity: O(1)

# 949. Largest Time for Given Digits - Easy

```
Given an array of 4 digits, return the largest 24 hour time that can be made.

The smallest 24 hour time is 00:00, and the largest is 23:59.  Starting from 00:00, a time is larger if more time has elapsed since midnight.

Return the answer as a string of length 5.  If no valid time can be made, return an empty string.

 

Example 1:

Input: [1,2,3,4]
Output: "23:41"
Example 2:

Input: [5,5,5,5]
Output: ""
 

Note:

A.length == 4
0 <= A[i] <= 9
```

# Relative Topics
* Math


# Note
```
Time: 2020/09/01
```


# Solution

## 1. Permutation via Backtracking

### Intuition
As we discussed before, the hard part of the problem is not enumerating over the permutations, but actually constructing the permutations itself.

In this approach, we present a solution to reinvent the wheel, i.e. generating permutations, which itself is a fun problem to solve. For practice, one can implement the permutation algorithms on these two problem: permutations and next permutation.

### Algorithm
Now we can put together all the ideas that we presented before, and implement the permutation algorithm.

Here we implement the permutation algorithm as the function permutate(array, start) which generates the permutations for the postfix subarray of array[start:len(array)]. Once we implement the function, we invocate it as permutate(array, 0) to generate all the permutations from the array.

### Complexity Analysis
*   Time Complexity: O(1)
*   Space Complexity: O(1).
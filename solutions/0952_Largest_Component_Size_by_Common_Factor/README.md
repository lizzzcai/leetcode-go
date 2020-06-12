# 952. Largest Component Size by Common Factor - Hard

```
Given a non-empty array of unique positive integers A, consider the following graph:

There are A.length nodes, labelled A[0] to A[A.length - 1];
There is an edge between A[i] and A[j] if and only if A[i] and A[j] share a common factor greater than 1.
Return the size of the largest connected component in the graph.

 

Example 1:

Input: [4,6,15,35]
Output: 4

Example 2:

Input: [20,50,9,63]
Output: 2

Example 3:

Input: [2,3,6,7,4,12,21,39]
Output: 8

Note:

1 <= A.length <= 20000
1 <= A[i] <= 100000
```

# Relative Topics
* Math
* Union Find


# Note
```
Time: 2020/06/12
```


# Solution

## 1. Union-Find 

### Intuition
Let W = \max(A[i])W=max(A[i]), and R = \sqrt{W}R= 
W
​	
 . For each value A[i]A[i], there is at most one prime factor p \geq Rp≥R dividing A[i]A[i]. Let's call A[i]A[i]'s "big prime" this pp, if it exists.

This means that there are at most R + A\text{.length}R+A.length unique prime divisors of elements in AA: the big primes correspond to a maximum of A\text{.length}A.length values, and the small primes are all less than RR, so there's at most RR of them too.

### Algorithm

Factor each A[i]A[i] into prime factors, and index every occurrence of these primes. (To save time, we can use a sieve. Please see this article's comments for more details.)

Then, use a union-find structure to union together any prime factors that came from the same A[i]A[i].

Finally, we can count the size of each component, by inspecting and counting the id of the component each A[i]A[i] belongs to.



### Complexity Analysis
*   Time Complexity: O(N\sqrt{W})O(Nsqr(W)) where NN is the length of A, and W = \max(A[i])W=max(A[i]).
  
*   Space Complexity: O(N).
# 309. Best Time to Buy and Sell Stock with Cooldown - Medium

```
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times) with the following restrictions:

You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)
Example:

Input: [1,2,3,0,2]
Output: 3 
Explanation: transactions = [buy, sell, cooldown, buy, sell]
```

# Relative Topics
* Dynamic Programming


# Note
```
Time: 2020/06/13
```


# Solution
## 1. Dynamic Programming

### Intuition
 dp[i] = max(dp[i-1], prices[i] - prices[j] + dp[j-2]) j = [0, i-1]
 If we sell the shares on i-th day bought on j-th day, 
 we couldn't trade on (j-1)-th day because of cooldown. So the last one is dp[j-2].
        

### Complexity Analysis
*   Time Complexity: O(N)
  
*   Space Complexity: O(N)

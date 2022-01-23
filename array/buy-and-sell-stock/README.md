# Best Time to Buy and Sell Stock

[Leet #121](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) - Easy

## Status

Started:        November 16, 2021     
Completed:      January 2, 2022

## Topics

• Array         
• Comparison    
• High and low values   
• Optimize time complexity      

## Description

You are given an array of stock prices where `prices[i]` is the price of a stock on the _ith_ day.

Maximize your profit by choosing a single day to buy the stock and a different day in the future to sell the stock.

Return the maximum profit possible.  If no profit is possible, return 0.

## Examples

```text
prices := []int{7, 1, 5, 3, 6, 4}
maxProfit = 5
```

```text
prices := []int{7, 6, 4, 3, 1}
maxProfit = 0
```

## Algorithm

1. set initial buyPrice as prices[0]
2. set maxProfit = 0
3. loop through prices[] starting at prices[1]  
        a. calculate profit = prices[i] - buyPrice     
        b. if prices[i] < buyPrice, set buyPrice = prices[i]    
        c. else if profit > maxProfit, set maxProfit = profit   

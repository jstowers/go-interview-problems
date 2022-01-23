# House Robber

[Leet #198](https://leetcode.com/problems/house-robber/) - Medium

## Topics

• Dynamic Programming  
• Array  
• Amazon    
• Google        

## Description

You are a professional cat burglar who burgles houses.  Each house contains stashed money, an integer value between $0 and $400.  Your only constraint is that adjacent houses have connected security systems that will automatically call the police.

Given an integer array `nums` representing the amount of money in each house, return the maximum amount of money you can rob tonight without alerting the police.

## Examples

Input: nums = [1,2,3,1]
Output: 4
Rob house 1 ($1) and house 3 ($3).

Input: nums = [2,7,9,3,1]
Output: 12
Rob house 1 ($2), house 3 ($9), and house 5 ($1)

## Constraints

•   0 <= nums[i] <= 400

## Algorithm

At any house, _i_, we have only two options:

1. _Rob house_ - we get the cash, `nums[i]`, in the current house, plus the total cash stolen up to two houses earlier dp(i-2).

2. _Don't rob house_ - we don't get any cash from the current house, but we keep the total cash stolen up to the last house dp(i-1).

In deciding whether to rob or not rob the current house, _i_, we pick the max of dp(i-2) + nums[i] vs. dp(i-1):

```text
    dp(i) = max((dp(i-2) + nums[i], dp(i-1)))
```

## Base Cases

These base cases will ensure the recursive calls don't error out:

1. For the first house, dp(0) = nums[0]

2. For the first two houses, dp(1) = max(nums[0], nums[1])

## Status

Started:    January 9, 2022    
Completed:  January 12, 2022
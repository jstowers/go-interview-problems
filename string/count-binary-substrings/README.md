# Count Binary Substrings

[Leet #696](https://leetcode.com/problems/count-binary-substrings/)

Easy

## Topics

• Substrings    
• Pointers  
• Loops

## Problem Statement

Given a binary string made of 0's and 1's, return the number of non-empty substrings that have the same number of 0's and 1's.  All the 0's and 1's in these substrings must be grouped consecutively.

## Example

```text
Input: s = "10101"

Output: 4

Why?    "10", "01", "10", "01"
```

```text
Input: s = "00110011"

Output: 6

Why?    "0011", "01", "1100", "10", "01", "0011"
```

## Algorithm

1. Check for an empty string.  If empty, return 0.

2. Initialize a groups[] with a value of 1.

3. Loop through the string, starting at index = 1.  Compare the current value (i) to the previous value (i-1).  If they are different, append 1 to groups[].  If they are equal, increment the current groups element by 1.

4. Loop through groups[]. Compare consecutive elements and choose the smallest number. This number represents the number of substrings that can be created between the two elements.

5. Keep a running sum and return at the end of the loop.
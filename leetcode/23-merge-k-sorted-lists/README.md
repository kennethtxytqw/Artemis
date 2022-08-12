# 23. Merge k sorted lists

## Solution 1 - Brute Force

Space = O(1)
Time = O((k^2)n) or O(kN)
where N is the total number of items

## Solution 2 - In Place Divide and Conquer

Space = O(1)
Time = O(Nlogk)

Every phase we do N merges and there are log k phases.

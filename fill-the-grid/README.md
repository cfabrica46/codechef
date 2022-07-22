# Fill The Grid
## Problem
You have a grid with N rows and M columns. You have two types of tiles — one of dimensions 2 \times 2×2 and the other of dimensions 1 \times 1×1. You want to cover the grid using these two types of tiles in such a way that:

Each cell of the grid is covered by exactly one tile; and
The number of 1 \times 1×1 tiles used is minimized.
Find the minimum number of 1 \times 11×1 tiles you have to use to fill the grid.

### Input Format
The first line of input will contain a single integer T, denoting the number of test cases.
Each test case consists of a single line containing two space-separated integers N, MN,M.
### Output Format
For each test case, print on a new line the minimum number of 1 \times 1×1 tiles needed to fill the grid.

### Constraints
1 <= N, M <= 1000
 
### Sample 1:
#### Input
4
1 1
4 5
6 8
3 2

#### Output
1
4
0
2

### Explanation:
Test case 11: There is only one square in the grid, and it must be filled with a single 1 \times 11×1 tile.

Test case 22: One way of tiling the grid using 1\times 11×1 tiles exactly 44 times is as follows:

Test case 33: One way of tiling the grid using no 1\times 11×1 tiles is as follows:

Test case 44: One way of tiling the grid using 1\times 11×1 tiles exactly twice is:


# Cricket Ranking
## Problem
In a season, each player has three statistics: runs, wickets, and catches. Given the season stats of two players AA and BB, denoted by RR, WW, and CC respectively, the person who is better than the other in the most statistics is regarded as the better overall player. Tell who is better amongst AA and BB. It is known that in each statistic, the players have different values.

### Input Format
The first line contains an integer TT, the number of test cases. Then the test cases follow.
Each test case contains two lines of input.
The first line contains three integers R1, W1, C1, the stats for player AA.
The second line contains three integers R2, W2, C2 , the stats for player BB.

### Output Format
For each test case, output in a single line "A" (without quotes) if player AA is better than player BB and "B" (without quotes) otherwise.

### Constraints
0 <= R1, R2 <= 500
0 <= W1, W2 <= 20
0 <= C1, C2 <= 20
R1 != R2
W1 != W2
C1 != C2
 
### Sample 1:
#### Input
3
0 1 2
2 3 4
10 10 10
8 8 8
10 0 10
0 10 0

#### Output
B
A
A

### Explanation:
Test Case 11: Player BB is better than AA in all 3 fields.

Test Case 22: Player AA is better than BB in all 3 fields.

Test Case 33: Player AA is better than BB in runs scored and number of catches.

# Greedy
# O(R + C), where R is the len(rowCosts) and C is the len(colCosts)
# O(1) memory
class Solution:

    def minCost(
        self,
        startPos: List[int],
        homePos: List[int],
        rowCosts: List[int],
        colCosts: List[int],
    ) -> int:
        startRow, startCol = startPos[0], startPos[1]
        homeRow, homeCol = homePos[0], homePos[1]

        cost = 0
        if startRow <= homeRow:
            for row in range(startRow + 1, homeRow + 1):
                cost += rowCosts[row]
        else:
            for row in range(startRow - 1, homeRow - 1, -1):
                cost += rowCosts[row]

        if startCol <= homeCol:
            for col in range(startCol + 1, homeCol + 1):
                cost += colCosts[col]
        else:
            for col in range(startCol - 1, homeCol - 1, -1):
                cost += colCosts[col]

        return cost


# DP
# Memo is wrong
class Solution:
    def minCost(
        self,
        startPos: List[int],
        homePos: List[int],
        rowCosts: List[int],
        colCosts: List[int],
    ) -> int:
        ROWS, COLS = len(rowCosts), len(colCosts)
        forbidden = float("infinity")
        visited = set()  # (row, col)
        memo = {}

        def dfs(row, col):
            # print("({}, {}) ====================== ".format(row, col))
            if row == homePos[0] and col == homePos[1]:
                # print("({}, {}) ".format(row, col))
                return 0

            if (row, col) in visited:
                # print("({}, {}) already visited".format(row, col))
                return forbidden

            if (row, col) in memo:
                # print("({}, {}) IN memo => memo[(row, col)]: {}".format(row, col, memo[(row, col)]))
                return memo[(row, col)]

            visited.add((row, col))

            top = forbidden
            if row - 1 >= 0:
                top = rowCosts[row - 1] + dfs(row - 1, col)

            down = forbidden
            if row + 1 < ROWS:
                down = rowCosts[row + 1] + dfs(row + 1, col)

            left = forbidden
            if col - 1 >= 0:
                left = colCosts[col - 1] + dfs(row, col - 1)

            right = forbidden
            if col + 1 < COLS:
                right = colCosts[col + 1] + dfs(row, col + 1)

            visited.remove((row, col))
            minValue = min(top, down, left, right)
            # print("({}, {}) minValue: {}".format(row, col, minValue))

            # if minValue < forbidden:
            memo[(row, col)] = minValue
            # memo[(row, col)] = min(memo.get((row, col), forbidden), minValue)
            # memo[(row, col)] = minValue

            return minValue

        result = dfs(startPos[0], startPos[1])
        print(memo)
        return result


# [1,0]
# [2,3]
# [5,4,3]
# [8,2,6,7]
# [0,0]
# [0,0]
# [5]
# [26]
# [3,0]
# [4,1]
# [10,5,6,7,11]
# [4,19,13,16,29,28,22,13]
# [2,0]
# [6,0]
# [4,3,9,0,10,0,15,14,12]
# [9,14]

# [45,70]
# [32,35]
# [4,10,5,8,0,10,5,9,10,2,7,7,7,6,1,1,5,1,5,9,7,1,0,3,1,2,6,4,6,4,2,4,1,1,5,2,3,9,3,9,8,4,1,4,6,6728,8650,6586,9762,9034,9246,5033,9632,6907,7237,6422,5603,6062,5033,5109,8118,7210,9672,8268,5157,5854,7723,8101]
# [7,8,9,8,10,3,10,4,8,10,7,5,5,9,7,5,7,10,8,6,2,2,4,10,7,3,6,2,1,8,1,6,5,5,1,9,10,6,3,2,8,6,3,0,10,4,5,4,6,2,1,6,8,9,3,0,5,6,8,3,8,6,10,4,6,4,3,3,6,3,6893,9916,9592,7854,8030,6396,8904,5191,9514,9417,9701,9840,9194,6515,5643,7804,9768,8898,6735,8548,6859,6024,9551,9520,5537,5135]

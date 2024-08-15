# https://leetcode.com/problems/word-search


# O(m * n) time, O(1) memory
class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        ROWS, COLS = len(board), len(board[0])

        def backtracking(row, col, letterIndex):
            if letterIndex >= len(word):
                return True

            # out of bound
            if row < 0 or row >= ROWS or col < 0 or col >= COLS:
                return False

            # already used
            if board[row][col] == ".":
                return False

            # current letter doesn't match
            letter = board[row][col]
            if letter != word[letterIndex]:
                return False

            # mark as used
            board[row][col] = "."

            # look to neighboors
            top = backtracking(row - 1, col, letterIndex + 1)
            if top:
                return True

            bottom = backtracking(row + 1, col, letterIndex + 1)
            if bottom:
                return True

            left = backtracking(row, col - 1, letterIndex + 1)
            if left:
                return True

            right = backtracking(row, col + 1, letterIndex + 1)
            if right:
                return True

            board[row][col] = letter
            return False

        for row in range(ROWS):
            for col in range(COLS):
                if backtracking(row, col, 0):
                    return True

        return False

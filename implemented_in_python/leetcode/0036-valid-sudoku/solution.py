# https://leetcode.com/problems/valid-sudoku/description/

import collections


def solution(board):
    rows = collections.defaultdict(set)
    columns = collections.defaultdict(set)
    squares = collections.defaultdict(set)

    for row in range(9):
        for column in range(9):
            value = board[row][column]
            if value == ".":
                continue

            if (
                value in rows[row]
                or value in columns[column]
                or value in squares[(row // 3, column // 3)]
            ):
                return False

            rows[row].add(value)
            columns[column].add(value)
            squares[(row // 3, column // 3)].add(value)

    return True


def worst_solution(board):
    return check_lines(board) and check_columns(board) and check_sub_boards(board)


def check_lines(board):
    num_set = set({})
    for line_index in range(9):
        num_set.clear()
        for col_index in range(9):
            cell = board[line_index][col_index]
            if cell == ".":
                continue

            if cell in num_set:
                return False

            num_set.add(cell)

    return True


def check_columns(board):
    num_set = set({})
    for line_index in range(9):
        num_set.clear()

        for col_index in range(9):
            cell = board[col_index][line_index]

            if cell == ".":
                continue

            if cell in num_set:
                return False

            num_set.add(cell)

    return True


def check_sub_boards(board):
    result = True

    for i in range(0, 7, 3):
        for j in range(0, 7, 3):
            result = result and check_sub_board(board, i, j)

    return result


def check_sub_board(board, line_start, col_start):
    num_set = set({})

    for line in range(line_start, line_start + 3):
        for col in range(col_start, col_start + 3):
            cell = board[line][col]

            if cell == ".":
                continue

            if cell in num_set:
                return False

            num_set.add(cell)

    return True


board = [
    ["5", "3", ".", ".", "7", ".", ".", ".", "."],
    ["6", ".", ".", "1", "9", "5", ".", ".", "."],
    [".", "9", "8", ".", ".", ".", ".", "6", "."],
    ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
    ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
    [".", "6", ".", ".", ".", ".", "2", "8", "."],
    [".", ".", ".", "4", "1", "9", ".", ".", "5"],
    [".", ".", ".", ".", "8", ".", ".", "7", "9"],
]

print(solution(board))

board = [
    ["8", "3", ".", ".", "7", ".", ".", ".", "."],
    ["6", ".", ".", "1", "9", "5", ".", ".", "."],
    [".", "9", "8", ".", ".", ".", ".", "6", "."],
    ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
    ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
    [".", "6", ".", ".", ".", ".", "2", "8", "."],
    [".", ".", ".", "4", "1", "9", ".", ".", "5"],
    [".", ".", ".", ".", "8", ".", ".", "7", "9"],
]

print(solution(board))

board = [
    ["7", ".", ".", ".", "4", ".", ".", ".", "."],
    [".", ".", ".", "8", "6", "5", ".", ".", "."],
    [".", "1", ".", "2", ".", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", "9", ".", ".", "."],
    [".", ".", ".", ".", "5", ".", "5", ".", "."],
    [".", ".", ".", ".", ".", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", ".", "2", ".", "."],
    [".", ".", ".", ".", ".", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", ".", ".", ".", "."],
]

print(solution(board))

board = [
    [".", ".", ".", ".", "5", ".", ".", "1", "."],
    [".", "4", ".", "3", ".", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", "3", ".", ".", "1"],
    ["8", ".", ".", ".", ".", ".", ".", "2", "."],
    [".", ".", "2", ".", "7", ".", ".", ".", "."],
    [".", "1", "5", ".", ".", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", "2", ".", ".", "."],
    [".", "2", ".", "9", ".", ".", ".", ".", "."],
    [".", ".", "4", ".", ".", ".", ".", ".", "."],
]

print(solution(board))

board = [
    [".", ".", ".", ".", ".", ".", "5", ".", "."],
    [".", ".", ".", ".", ".", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", ".", ".", ".", "."],
    ["9", "3", ".", ".", "2", ".", "4", ".", "."],
    [".", ".", "7", ".", ".", ".", "3", ".", "."],
    [".", ".", ".", ".", ".", ".", ".", ".", "."],
    [".", ".", ".", "3", "4", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", "3", ".", ".", "."],
    [".", ".", ".", ".", ".", "5", "2", ".", "."],
]
print(solution(board))

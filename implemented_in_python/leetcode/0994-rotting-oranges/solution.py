# https://leetcode.com/problems/rotting-oranges/
#
# Time: O(m * n)
# Memory: O(m * n)
class Solution:
    def orangesRotting(self, grid: List[List[int]]) -> int:
        queue = deque([])

        # O(m * n)
        for row in range(len(grid)):
            for col in range(len(grid[0])):
                if grid[row][col] == 2:
                    queue.append((row, col))

        minutes = 0
        neighboors = [(-1, 0), (1, 0), (0, -1), (0, 1)]
        while queue:
            length = len(queue)

            for i in range(length):
                row, col = queue.popleft()

                for r, c in neighboors:  # 4 -> O(1)
                    if (
                        row + r < 0
                        or row + r >= len(grid)
                        or col + c < 0
                        or col + c >= len(grid[0])
                    ):
                        continue

                    if grid[row + r][col + c] in [0, 2]:
                        continue

                    grid[row + r][col + c] = 2
                    queue.append((row + r, col + c))

            if len(queue) > 0:
                minutes += 1

        # O(m * n)
        for row in range(len(grid)):
            for col in range(len(grid[0])):
                if grid[row][col] == 1:
                    return -1

        return minutes

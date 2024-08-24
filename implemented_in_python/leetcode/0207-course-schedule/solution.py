# https://leetcode.com/problems/course-schedule/


# O(c + p) => c: course, p: pre-requisites
class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        preRequisitesMap = {}
        for course, pre in prerequisites:
            array = preRequisitesMap.get(course, [])
            array.append(pre)
            preRequisitesMap[course] = array

        visited = set()

        def dfs(course):
            if course in visited:
                return False

            coursePreRequisites = preRequisitesMap.get(course, [])
            if len(coursePreRequisites) == 0:
                return True

            visited.add(course)
            for pre in coursePreRequisites:
                if not dfs(pre):
                    return False

            visited.remove(course)
            preRequisitesMap[course] = []
            return True

        for course in range(numCourses):
            if not dfs(course):
                return False

        return True

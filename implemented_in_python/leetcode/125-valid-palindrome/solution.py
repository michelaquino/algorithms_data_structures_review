class Solution(object):
    def isPalindrome(self, s):
        """
        :type s: str
        :rtype: bool
        """

        s = s.lower()
        i = 0
        j = len(s) - 1

        while i <= j:
            if not self.isAlfaNumeric(s[i]):
                i += 1
                continue

            if not self.isAlfaNumeric(s[j]):
                j -= 1
                continue

            if s[i] != s[j]:
                return False

            i += 1
            j -= 1

        return True

    def isAlfaNumeric(self, c):
        return (
            ord("A") <= ord(c) <= ord("Z")
            or ord("a") <= ord(c) <= ord("z")
            or ord("0") <= ord(c) <= ord("9")
        )


solution = Solution()

s = "A man, a plan, a canal: Panama"
print(solution.isPalindrome(s))

s = "race a car"
print(solution.isPalindrome(s))


s = " "
print(solution.isPalindrome(s))

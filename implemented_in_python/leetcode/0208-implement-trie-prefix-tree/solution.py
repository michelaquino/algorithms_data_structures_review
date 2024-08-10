# https://leetcode.com/problems/implement-trie-prefix-tree/


class Node:
    def __init__(self):
        self.children = {}
        self.terminal = False


class Trie:
    def __init__(self):
        self.root = None

    def insert(self, word):
        if not self.root:
            self.root = Node()

        curr = self.root

        lastLetter = ""
        for i, char in enumerate(word):
            lastLetter = char
            if char not in curr.children:
                node = Node()
                curr.children[char] = node

            curr = curr.children[char]

        curr.terminal = True

    def search(self, word):
        def dfs(root, word, wordIndex):
            if not root:
                return False

            if wordIndex == len(word) and root.terminal:
                return True

            letter = word[wordIndex] if wordIndex < len(word) else ""
            if letter not in root.children:
                return False

            child = root.children[letter]
            return dfs(child, word, wordIndex + 1)

        return dfs(self.root, word, 0)

    def startsWith(self, prefix):
        def dfs(root, word, wordIndex):
            if not root:
                return False

            if wordIndex >= len(word):
                return True

            letter = word[wordIndex]

            if letter not in root.children:
                return False

            child = root.children[letter]
            return dfs(child, word, wordIndex + 1)

        return dfs(self.root, prefix, 0)


# Your Trie object will be instantiated and called as such:
# obj = Trie()
# obj.insert(word)
# param_2 = obj.search(word)
# param_3 = obj.startsWith(prefix)

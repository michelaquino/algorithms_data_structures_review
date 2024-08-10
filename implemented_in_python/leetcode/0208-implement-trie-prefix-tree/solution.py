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


class TrieNode:
    def __init__(self):
        self.children = [None] * 26
        self.end = False


class SimplerTrie:
    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.root = TrieNode()

    def insert(self, word: str) -> None:
        """
        Inserts a word into the trie.
        """
        curr = self.root
        for c in word:
            i = ord(c) - ord("a")
            if curr.children[i] == None:
                curr.children[i] = TrieNode()
            curr = curr.children[i]
        curr.end = True

    def search(self, word: str) -> bool:
        """
        Returns if the word is in the trie.
        """
        curr = self.root
        for c in word:
            i = ord(c) - ord("a")
            if curr.children[i] == None:
                return False
            curr = curr.children[i]
        return curr.end

    def startsWith(self, prefix: str) -> bool:
        """
        Returns if there is any word in the trie that starts with the given prefix.
        """
        curr = self.root
        for c in prefix:
            i = ord(c) - ord("a")
            if curr.children[i] == None:
                return False
            curr = curr.children[i]
        return True


# Your Trie object will be instantiated and called as such:
# obj = Trie()
# obj.insert(word)
# param_2 = obj.search(word)
# param_3 = obj.startsWith(prefix)

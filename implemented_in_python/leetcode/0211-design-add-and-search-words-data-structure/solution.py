# https://leetcode.com/problems/design-add-and-search-words-data-structure/


class TrieNode:
    def __init__(self):
        self.children = {}
        self.endOfWord = False


class WordDictionary:
    def __init__(self):
        self.root = TrieNode()

    def addWord(self, word: str) -> None:
        curr = self.root

        for letter in word:
            node = curr.children.get(letter, TrieNode())
            curr.children[letter] = node
            curr = node

        curr.endOfWord = True

    def search(self, word: str) -> bool:
        def dfs(startIndex, curr):
            for i in range(startIndex, len(word)):
                letter = word[i]
                if letter == ".":
                    for child in curr.children.values():
                        if dfs(i + 1, child):
                            return True
                    return False
                else:
                    if letter not in curr.children:
                        return False

                    curr = curr.children[letter]

            return curr.endOfWord

        return dfs(0, self.root)

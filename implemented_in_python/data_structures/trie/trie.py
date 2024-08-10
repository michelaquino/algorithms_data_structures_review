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

    def print(self):
        if not self.root:
            print("Empty trie")
            return

        def dfs(root, prefix):
            if not root:
                return

            for letter, child in root.children.items():
                if child.terminal:
                    print("--> ", prefix + letter)

                dfs(child, prefix + letter)

        dfs(self.root, "")


t = Trie()
print(t.search("word"))

t.insert("where")
t.insert("was")
t.insert("cat")
t.insert("cattle")
t.insert("happy")
t.print()

# print(t.root.children)
print(t.search("was"))
print(t.search("wass"))

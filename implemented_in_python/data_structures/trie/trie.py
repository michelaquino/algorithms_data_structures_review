class TrieNode:
    def __init__(self):
        self.children = [None] * 26
        self.end = False


class SimplerTrie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word: str) -> None:
        curr = self.root
        for c in word:
            i = ord(c) - ord("a")
            if curr.children[i] == None:
                curr.children[i] = TrieNode()
            curr = curr.children[i]
        curr.end = True

    def search(self, word: str) -> bool:
        curr = self.root
        for c in word:
            i = ord(c) - ord("a")
            if curr.children[i] == None:
                return False
            curr = curr.children[i]
        return curr.end

    def startsWith(self, prefix: str) -> bool:
        curr = self.root
        for c in prefix:
            i = ord(c) - ord("a")
            if curr.children[i] == None:
                return False
            curr = curr.children[i]
        return True

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

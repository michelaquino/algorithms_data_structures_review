class Node:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None


class Tree:
    def __init__(self, root):
        self.root = root

    def insert(self, node):
        return None

    def delete(self):
        return None

    def search(self, value):
        return None

    def printInOrder(self):
        return None

    def printPreOrder(self):
        return None

    def printPostOrder(self):
        return None

    def minimum(self):
        return None

    def maximum(self):
        return self._iterative_maximum(self.root)
        # return self._recursive_maximum(self.root)

    def _iterative_maximum(self, node):
        curr = node
        while curr and curr.right:
            curr = curr.right

        return curr

    def _recursive_maximum(self, node):
        if not node or not node.right:
            return node

        return self._recursive_maximum(node.right)

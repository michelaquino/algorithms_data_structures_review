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
        return self._iterative_minimum(self.root)
        # return self._recursive_minimum(self.root)

    def _recursive_minimum(self, node):
        if not node or not node.left:
            return node

        return self._recursive_minimum(node.left)

    def _iterative_minimum(self, node):
        curr = node
        while curr and curr.left:
            curr = curr.left

        return curr

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


node1 = Node(1)
node2 = Node(2)
node3 = Node(3)
node4 = Node(4)

node6 = Node(6)
node7 = Node(7)
node9 = Node(9)


node2.left = node1
node2.right = node3

node7.left = node6
node7.right = node9

node4.left = node2
node4.right = node7

tree = Tree(node4)

minNode = tree.minimum()
print("Min node: ", minNode.value if minNode else None)

maxNode = tree.maximum()
print("Max node: ", maxNode.value if maxNode else None)

class Node:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None


class Tree:
    def __init__(self, root):
        self.root = root

    def insert(self, node):
        # self.root = self._insert_recursive(self.root, node)
        self.root = self._insert_iterative(self.root, node)

    def _insert_iterative(self, root, node):
        parent = root
        curr = root
        while curr:
            parent = curr
            if curr.value < node.value:
                curr = curr.right
            else:
                curr = curr.left

        if not parent:
            return node

        if node.value > parent.value:
            parent.right = node
        else:
            parent.left = node

        return root

    def _insert_recursive(self, root, node):
        if not root:
            return node

        if root.value < node.value:
            right = self._insert_recursive(root.right, node)
            root.right = right
        else:
            left = self._insert_recursive(root.left, node)
            root.left = left

        return root

    def delete(self):
        return None

    def search(self, value):
        # return self._search_iterative(self.root, value)
        return self._search_recursively(self.root, value)

    def _search_iterative(self, root, value):
        curr = root
        while curr:
            if curr.value < value:
                return self._search_recursively(curr.right, value)
            elif curr.value > value:
                return self._search_recursively(curr.left, value)
            else:
                return root

        return None

    def _search_recursively(self, root, value):
        if not root:
            return None

        if root.value < value:
            return self._search_recursively(root.right, value)
        elif root.value > value:
            return self._search_recursively(root.left, value)
        else:
            return root

    def printInOrder(self):
        print("==== PRINT IN ORDER ====")
        self._inOrderTraversal(self.root)

    def _inOrderTraversal(self, root):
        if not root:
            return

        self._inOrderTraversal(root.left)
        print(root.value)
        self._inOrderTraversal(root.right)

    def printPreOrder(self):
        print("==== PRINT PRE ORDER ====")
        self._preOrderTraversal(self.root)

    def _preOrderTraversal(self, root):
        if not root:
            return

        print(root.value)
        self._preOrderTraversal(root.left)
        self._preOrderTraversal(root.right)

    def printPostOrder(self):
        print("==== PRINT POST ORDER ====")
        self._postOrderTraversal(self.root)

    def _postOrderTraversal(self, root):
        if not root:
            return

        self._postOrderTraversal(root.left)
        self._postOrderTraversal(root.right)
        print(root.value)

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

tree.printInOrder()
tree.printPreOrder()
tree.printPostOrder()

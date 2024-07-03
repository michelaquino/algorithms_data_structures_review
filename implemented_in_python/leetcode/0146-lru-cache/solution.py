class Node:
    def __init__(self, key, value):
        self.key = key
        self.value = value
        self.prev = None
        self.next = None


class LRUCache:
    def __init__(self, capacity: int):
        self.size = capacity
        self.head = None
        self.tail = None
        self.index = {}

    def get(self, key):
        node = self.index.get(key, None)
        if not node:
            return -1

        self._updatePriority(key)
        return node.value

    def put(self, key, value):
        node = self.index.get(key, None)
        if node:
            self._update(key, value)
            return

        self._add(key, value)

    def _add(self, key, value):
        newNode = Node(key, value)
        newNode.prev = self.tail

        if self.tail:
            self.tail.next = newNode

        if not self.head:
            self.head = newNode

        self.tail = newNode

        self.index[key] = newNode
        self.size -= 1
        self._evict()

    def _update(self, key, value):
        node = self.index.get(key, None)
        if not node:
            return

        node.value = value
        self._updatePriority(key)

    def _evict(self):
        if self.size >= 0 or not self.head:
            return

        del self.index[self.head.key]
        self.head = self.head.next

    def _updatePriority(self, key):
        node = self.index.get(key, None)
        if not node or node == self.tail:
            return

        if node == self.head:
            node.next.prev = None
            self.head = node.next
        else:
            prev = node.prev
            prev.next = node.next
            node.next.prev = prev

        node.next = None
        self.tail.next = node
        node.prev = self.tail
        self.tail = node


# Simpler solution


# class Node:
#     def __init__(self, key, val):
#         self.key, self.val = key, val
#         self.prev = self.next = None


# class LRUCache:
#     def __init__(self, capacity: int):
#         self.cap = capacity
#         self.cache = {}  # map key to node

#         self.left, self.right = Node(0, 0), Node(0, 0)
#         self.left.next, self.right.prev = self.right, self.left

#     # remove node from list
#     def remove(self, node):
#         prev, nxt = node.prev, node.next
#         prev.next, nxt.prev = nxt, prev

#     # insert node at right
#     def insert(self, node):
#         prev, nxt = self.right.prev, self.right
#         prev.next = nxt.prev = node
#         node.next, node.prev = nxt, prev

#     def get(self, key: int) -> int:
#         if key in self.cache:
#             self.remove(self.cache[key])
#             self.insert(self.cache[key])
#             return self.cache[key].val
#         return -1

#     def put(self, key: int, value: int) -> None:
#         if key in self.cache:
#             self.remove(self.cache[key])
#         self.cache[key] = Node(key, value)
#         self.insert(self.cache[key])

#         if len(self.cache) > self.cap:
#             # remove from the list and delete the LRU from hashmap
#             lru = self.left.next
#             self.remove(lru)
#             del self.cache[lru.key]

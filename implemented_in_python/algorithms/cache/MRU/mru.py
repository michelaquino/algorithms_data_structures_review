class Node:
    def __init__(self, key, value):
        self.key = key
        self.value = value
        self.prev = None
        self.next = None


class Cache:
    def __init__(self, size):
        self.size = size
        self.head = None
        self.index = {}

    def put(self, key, value):
        node = self.index.get(key, None)
        if node:
            self._update(key, value)
            return

        self._add(key, value)

    def get(self, key):
        node = self.index.get(key, None)
        if not node:
            return -1

        self._updatePriority(key)
        return node.value

    def _add(self, key, value):
        self.size -= 1
        self._evict()
        newNode = Node(key, value)

        if self.head:
            self.head.prev = newNode
            newNode.next = self.head

        self.head = newNode
        self.index[key] = newNode

    def _update(self, key, value):
        node = self.index.get(key, None)
        if not node:
            return

        node.value = value
        self._updatePriority(key)

    def _evict(self):
        if self.size >= 0:
            return

        del self.index[self.head.key]
        self.head = self.head.next
        self.head.prev = None

    def _updatePriority(self, key):
        node = self.index.get(key, None)
        if not node or node == self.head:
            return

        prev = node.prev
        prev.next = node.next

        node.prev = None
        self.head.prev = node
        node.next = self.head
        self.head = node

    def display(self):
        print("----")
        # print("Head: ", self.head)
        # print("head: ", self.head)
        # print()

        curr = self.head
        while curr:
            print(curr.value, " --> ", end=" ")
            curr = curr.next

        print()
        print("----")


cache = Cache(4)


def display():
    print("*****")
    print("Size: ", cache.size)
    cache.display()


print("get z: ", cache.get("z"))

print("=======================")
cache.put("a", 1)
# display()
# print("get a: ", cache.get("a"))

# print("******")
cache.put("b", 2)
# display()
# print("get a: ", cache.get("a"))
# print("get b: ", cache.get("b"))

# print("******")
cache.put("c", 3)
# display()
# print("get a: ", cache.get("a"))
# print("get b: ", cache.get("b"))
# print("get c: ", cache.get("c"))

# print("******")
cache.put("d", 4)
# display()
# print("get a: ", cache.get("a"))
# print("get b: ", cache.get("b"))
# print("get c: ", cache.get("c"))
# print("get d: ", cache.get("d"))

# print("******")
# print("get a: ", cache.get("a"))
# display()
# print("get c: ", cache.get("c"))

print("=======================")

# print("*")
cache.put("e", 5)  # bug
display()
print("get b: ", cache.get("b"))
print("get d: ", cache.get("d"))
display()

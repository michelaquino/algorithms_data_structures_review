import collections

# Implementing with array


class StackArray:
    def __init__(self):
        self.data = []
        self.items_quantity = 0

    def push(self, value):
        self.data.append(value)
        self.items_quantity += 1

    def top(self):
        if self.items_quantity == 0:
            return None

        return self.data[self.items_quantity - 1]

    def pop(self):
        if self.items_quantity == 0:
            return None

        value = self.data[self.items_quantity - 1]
        del self.data[self.items_quantity - 1]
        self.items_quantity -= 1

        return value

    def print(self):
        print("Data => ", self.data)


class Node:
    def __init__(self, value):
        self.next = None
        self.value = value


class StackLinkedList:
    def __init__(self):
        self.head = None

    def push(self, value):
        node = Node(value)

        if self.head is None:
            self.head = node
            return

        node.next = self.head
        self.head = node

    def top(self):
        if self.head is None:
            return None

        return self.head.value

    def pop(self):
        if self.head is None:
            return None

        value = self.head.value
        self.head = self.head.next
        return value

    def print(self):
        print("Data -> ")
        node = self.head
        while node is not None:
            print("-> ", node.value)
            node = node.next


# stack = StackArray()
stack = StackLinkedList()
stack.print()
stack.push(1)
stack.push(2)
stack.push(3)
stack.push(4)

stack.print()
print("top => ", stack.top())

print("pop => ", stack.pop())

stack.print()
print("top => ", stack.top())

stack.pop()
stack.pop()
stack.pop()
stack.pop()

stack.print()



class Node:
    def __init__(self, value):
        self.value = value
        self.next = None


class LinkedList:
    def __init__(self):
        self.head = None
        self.node_quantity = 0

    # insert(index, value) - insert value at index, so the current item at that index is pointed to by the new item at the index
    def insert(self, index, value):
        if index < 0:
            return

        newNode = Node(value)

        if index == 0:
            newNode.next = self.head
            self.head = newNode
            self.node_quantity += 1
            return

        i = 0
        prev_noe = self.head
        curr_node = self.head

        while True:
            if i == index:
                prev.next = newNode
                newNode.next = curr_node
                break
            else:
                if curr_node is None:
                    return

                prev = curr_node
                curr_node = curr_node.next
                i += 1

        self.node_quantity += 1
        return

    # erase(index) - removes node at given index
    def erase(self, index):
        if self.head is None:
            return

        if index == 0:
            self.head = self.head.next
            self.node_quantity -= 1
            return

        prev = self.head
        curr = self.head.next
        i = 1
        while curr is not None:
            if i == index:
                prev.next = curr.next
                self.node_quantity -= 1
                return

            prev = curr
            curr = curr.next
            i += 1

    # remove_value(value) - removes the first item in the list with this value
    def remove_value(self, value):
        prev = self.head
        curr = self.head

        if self.head is not None and self.head.value == value:
            self.head = self.head.next
            self.node_quantity -= 1
            return

        while curr is not None:
            if curr.value == value:
                prev.next = curr.next
                self.node_quantity -= 1
                return

            prev = curr
            curr = curr.next

    # size() - returns the number of data elements in the list
    def size(self):
        return self.node_quantity

    # empty() - bool returns true if empty
    def empty(self):
        return self.node_quantity == 0

    # value_at(index) - returns the value of the nth item (starting at 0 for first)
    def value_at(self, index):
        if index < 0:
            return None

        i = 0
        curr_node = self.head
        while curr_node is not None:
            if i == index:
                return curr_node.value

            curr_node = curr_node.next
            i += 1

        return None

    # push_front(value) - adds an item to the front of the list
    def push_front(self, value):
        self.insert(0, value)

    # pop_front() - remove the front item and return its value
    def pop_front(self):
        self.erase(0)

    # push_back(value) - adds an item at the end
    def push_back(self, value):
        self.insert(self.size(), value)

    # pop_back() - removes end item and returns its value
    def pop_back(self):
        self.erase(self.size()-1)

    # front() - get the value of the front item
    def front(self):
        return self.value_at(0)

    # back() - get the value of the end item
    def back(self):
        return self.value_at(self.size()-1)

    # reverse() - reverses the list
    def reverse(self):
        if self.head is None:
            return

        prev = None
        curr = self.head

        while curr is not None:
            aux = curr.next
            curr.next = prev
            prev = curr
            curr = aux

        self.head = prev

    # print_all - print all node values
    def print_all(self):
        print("<--------")
        curr_node = self.head
        index = 0
        while curr_node is not None:
            print("[%d] %d"%(index, curr_node.value))
            index += 1
            curr_node = curr_node.next

        print("-------->")



l = LinkedList()

l.insert(0, 1)
# l.print_all()

l.insert(1, 2)
# l.print_all()

l.insert(10, 2)
# l.print_all()

l.insert(1, 3)
# l.print_all()

l.insert(0, 4)
# l.print_all()

# print ("size: ", l.size())
# print("value[0]: ", l.value_at(0))

l.push_front(111)
l.push_back(99)
# l.print_all()

l.remove_value(99)
# l.print_all()

# print ("size: ", l.size())
# print("front: ", l.front())
# print("back: ", l.back())


# l.print_all()
# l.insert(2, 33)
# l.pop_front()
# l.pop_back()
l.print_all()

l.reverse()
l.print_all()

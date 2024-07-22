class Graph:
    def __init__(self):
        self.nodes = []

    def insert_edge(self, node):
        if not node:
            return

        self.nodes.append(node)

    def search(self, value):
        print("-----------------------")
        if len(self.nodes) == 0:
            return None

        visited = set()

        def dfs(node, value):
            if node.value == value:
                return node

            visited.add(node)
            for neighboor in node.neighboors:
                if neighboor in visited:
                    continue

                found = dfs(neighboor, value)
                if found:
                    return found

            return None

        return dfs(self.nodes[0], value)

    def print(self):
        print(self.nodes)


class Node:
    def __init__(self, value):
        self.value = value
        self.neighboors = []

    def connectTo(self, node):
        if not node:
            return

        self.neighboors.append(node)


n0 = Node(0)
n1 = Node(1)
n2 = Node(2)
n3 = Node(3)
n4 = Node(4)
n5 = Node(5)

n0.connectTo(n2)
n1.connectTo(n2)
n1.connectTo(n3)
n1.connectTo(n4)
n2.connectTo(n0)
n2.connectTo(n1)
n2.connectTo(n5)
n3.connectTo(n1)
n4.connectTo(n1)
n5.connectTo(n2)


g = Graph()
g.insert_edge(n0)
g.insert_edge(n1)
g.insert_edge(n2)
g.insert_edge(n3)
g.insert_edge(n4)
g.insert_edge(n5)

print("Found: ", g.search(1))
print("Found: ", g.search(2))
print("Found: ", g.search(3))
print("Found: ", g.search(4))
print("Found: ", g.search(5))
print("Found: ", g.search(6))

# n6 = Node(0)
# n7 = Node(0)
# n8 = Node(0)
# n8 = Node(0)
# n8 = Node(0)

class Graph:
    def __init__(self):
        self.nodes = {}

    def insert_edge(self, node):
        if not node or node.value in self.nodes:
            return

        self.nodes[node.value] = node

    def insert_vertice(self, node1Value, node2Value):
        if node1Value not in self.nodes or node2Value not in self.nodes:
            return

        node1 = self.nodes[node1Value]
        node2 = self.nodes[node2Value]

        node1.connectTo(node2)
        node2.connectTo(node1)

    def print(self):
        print(self.nodes)


class Node:
    def __init__(self, value):
        self.value = value
        self.neighboors = {}

    def connectTo(self, node):
        if not node:
            return

        self.neighboors[node.value] = node


g = Graph()
# g.print()

g.insert_edge(Node(0))
g.insert_edge(Node(3))

g.insert_vertice(0, 3)
# g.print()

print(g.nodes[0].neighboors)
print(g.nodes[3].neighboors)

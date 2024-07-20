class Graph:
    def __init__(self, edges):
        self.adjacency_matrix = []
        for r in range(edges):
            row = [0] * edges
            self.adjacency_matrix.append(row)

    def insert_vertice(self, v1, v2):
        self.adjacency_matrix[v1][v2] = 1
        self.adjacency_matrix[v2][v1] = 1

    def print(self):
        print(self.adjacency_matrix)


g = Graph(4)
g.print()
g.insert_vertice(0, 3)
g.print()

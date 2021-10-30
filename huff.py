from display import create_diagram
from collections import defaultdict


def compress(data):
    freqs = defaultdict(int)
    for l in data:
        freqs[l] += 1

    nodes = [Node(letter=l, freq=f) for l, f in freqs.items()]

    while len(nodes) > 1:
        nodes = sorted(nodes, key=lambda n: n.freq, reverse=True)
        r, l = nodes.pop(), nodes.pop()
        parent = Node(freq=r.freq+l.freq, l=l, r=r)
        nodes += [parent]

    assert len(nodes) == 1, "Should be left with only one node"
    create_diagram(nodes[0])

class Node():
    def __init__(self, freq=0, letter=None, l=None, r=None):
        assert isinstance(freq, int), "freq should be int"
        self.freq=freq
        self.letter=letter
        self.l=l
        self.r=r

    def __repr__(self):
        ntype = "Leaf" if self.letter else "Node"
        return f"{ntype}(freq={self.freq}, letter={self.letter})"

compress("aaaaaaabbbbbbccddaphvpjoireqihqpyreqbvkxncjvnzxcvpirehgpiuqehnkv")

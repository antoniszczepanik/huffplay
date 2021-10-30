#!/usr/bin/env python3

import os
import sys
from collections import defaultdict

from display import draw_tree


def compress(data):
    freqs = defaultdict(int)
    for l in data:
        freqs[l] += 1

    nodes = [Node(letter=l, freq=f) for l, f in freqs.items()]
    root = construct_tree(nodes)
    draw_tree(root)

    translation_table = get_translation_table(root, "")

    return encode(data, translation_table), translation_table


def decompress(binary_data, translation_table):
    decoded = ""
    acc = ""
    for bit in binary_data:
        acc += bit
        if translation_table.get(acc):
            decoded += translation_table[acc]
            acc = ""
    return decoded


def construct_tree(nodes):
    while len(nodes) > 1:
        nodes = sorted(nodes, key=lambda n: n.freq, reverse=True)
        r, l = nodes.pop(), nodes.pop()
        parent = Node(freq=r.freq+l.freq, l=l, r=r)
        nodes += [parent]
    return nodes[0]

# This method that this is correct huffman tree which means it is full binary
# tree.
def get_translation_table(root, prefix):
    # it has a letter => it is a leaf
    if root.letter:
        return {root.letter: prefix}
    if root.l:
        lt = get_translation_table(root.l, prefix+"0")
    if root.r:
        rt = get_translation_table(root.r, prefix+"1")
    return {**lt, **rt}

def encode(data, translation_table):
    return "".join([translation_table[l] for l in data])

class Node():
    def __init__(self, freq=0, letter=None, l=None, r=None):
        assert isinstance(freq, int), "freq should be int"
        self.freq=freq
        self.letter=letter
        self.l=l
        self.r=r

    def __repr__(self):
        node_type = "Leaf" if self.letter else "Node"
        return f"{node_type}(freq={self.freq}, letter={self.letter})"

if __name__ == "__main__":
    raw_data = "".join(sys.stdin.readlines())
    binary, translation_table = compress(raw_data)
    print(f"Binary: {len(binary)}")
    original_data = decompress(binary, {v: k for k, v in translation_table.items()})
    print(f"Original data: {len(original_data.encode('utf-8'))*8}")

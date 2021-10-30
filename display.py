import graphviz

counter = 0

def draw_tree(root):
    graph = graphviz.Graph(comment='Huffman Coding Binary Tree', format="png")
    construct_tree(graph, root, None, "")
    graph.render('graph.gv', view=True)

def construct_tree(graph, root, parent, parent_uid):
    global counter
    uid = str(counter)
    counter += 1
    graph.node(uid, label=str(root))
    if parent:
        graph.edge(parent_uid, uid)
    if root.l:
        construct_tree(graph, root.l, root, uid)
    if root.r:
        construct_tree(graph, root.r, root, uid)

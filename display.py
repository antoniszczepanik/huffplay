import graphviz


counter = 0

def create_diagram(root):
    graph = graphviz.Graph(comment='Huffman Coding Binary Tree', format="png")
    create_diagram_internal(graph, root, None, "")
    graph.render('graph.png', view=True)

def create_diagram_internal(graph, root, parent, parent_uid):
    global counter
    uid = str(counter)
    counter += 1
    graph.node(uid, label=str(root))
    if parent:
        graph.edge(parent_uid, uid)
    if root.l:
        create_diagram_internal(graph, root.l, root, uid)
    if root.r:
        create_diagram_internal(graph, root.r, root, uid)

import networkx as nx
import matplotlib.pyplot as plt
import numpy as np

G = nx.MultiGraph()

graph_data_nodes = np.loadtxt('nodes.csv', dtype='str', delimiter=',', encoding="utf-8-sig")
graph_data_edges = np.loadtxt('edges.csv', dtype='str', delimiter=',', encoding="utf-8-sig")

for node in graph_data_nodes:
    G.add_node(node[5],pos=(float(node[0]), float(node[1])))

G.add_edges_from(graph_data_edges)

pos=nx.get_node_attributes(G,'pos')

nx.draw(G, pos, with_labels=False, node_size=50)

plt.show()
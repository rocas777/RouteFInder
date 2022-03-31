from xml.dom.minicompat import NodeList
import networkx as nx
import matplotlib.pyplot as plt
import numpy as np
import re

G = nx.MultiGraph()

graph_data_nodes = np.loadtxt('nodes.csv', dtype='str', delimiter=',', encoding="utf-8-sig")
graph_data_edges = np.loadtxt('edges.csv', dtype='str', delimiter=',', encoding="utf-8-sig")

for node in graph_data_nodes:
        G.add_node(node[5],pos=(float(node[0]), float(node[1])))

G.add_edges_from(graph_data_edges)

pos=nx.get_node_attributes(G,'pos')

busNodes = []
metroNodes = []
busEdges = []
metroEdges = []

for node in G.nodes:
    if 'bus' in node:
        busNodes.append(node)
    if 'metro' in node:
        metroNodes.append(node)

for node in G.edges:
    if 'bus' in node[0]:
        busEdges.append(node)
    if 'metro' in node[1]:
        metroEdges.append(node)

# Draw Bus Nodes
nx.draw_networkx_nodes(G, pos, nodelist=busNodes, node_size=25, node_color='tab:blue', alpha=0.9)

# Draw Metro Nodes
nx.draw_networkx_nodes(G, pos, nodelist=metroNodes, node_size=25, node_color='tab:red', alpha=0.9)

# Draw Bus Edges
nx.draw_networkx_edges(G, pos, edgelist=busEdges, edge_color='tab:blue')

# Draw Metro Edges
nx.draw_networkx_edges(G, pos, edgelist=metroEdges, edge_color='tab:red')

# Draw Labels
#nx.draw_networkx_labels(G, pos, font_size=10)

plt.show()
import networkx as nx
import matplotlib.pyplot as plt
import numpy as np
import math


def latlon_to_xy(lat, lon):
    """Convert angluar to cartesian coordiantes

    latitude is the 90deg - zenith angle in range [-90;90]
    lonitude is the azimuthal angle in range [-180;180]
    """
    r = 6371  # https://en.wikipedia.org/wiki/Earth_radius
    theta = math.pi / 2 - math.radians(float(lat))
    phi = math.radians(float(lon))
    x = r * math.sin(theta) * math.cos(phi)  # bronstein (3.381a)
    y = r * math.sin(theta) * math.sin(phi)
    return [x, y]


G = nx.MultiGraph()

graph_data_nodes = np.loadtxt('data/reuse/nodes.csv', dtype='str', delimiter=',', encoding="utf-8-sig", skiprows=1)
graph_data_edges = np.loadtxt('data/reuse/edges.csv', dtype='str', delimiter=',', encoding="utf-8-sig", skiprows=1)
path_edges = []
try:
    path_edges = np.loadtxt('data/reuse/path_edges.csv', dtype='str', delimiter=',', encoding="utf-8-sig", skiprows=1)
except:
    pass

for node in graph_data_nodes:
    x, y = latlon_to_xy((node[0]), (node[1]))
    G.add_node(node[4], pos=(y, -x))

G.add_edges_from(graph_data_edges)

pos = nx.get_node_attributes(G, 'pos')

busNodes = []
metroNodes = []
busEdges = []
metroEdges = []
walkNodes = []
walkEdges = []
pathEdges = []

pathEdgeExist = set()

for edge in path_edges:
    pathEdgeExist.add(edge[0] + "::" + edge[1])

for node in G.nodes:
    if 'bus' in node:
        busNodes.append(node)
    elif 'metro' in node:
        metroNodes.append(node)
    elif 'walk' in node:
        walkNodes.append(node)

for node in G.edges:
    if (node[0] + "::" + node[1] in pathEdgeExist or node[1] + "::" + node[0] in pathEdgeExist):
        pathEdges.append(node)
    elif ('bus' in node[0] and 'bus' in node[1]):
        busEdges.append(node)
    elif ('metro' in node[0] and 'metro' in node[1]):
        metroEdges.append(node)
    elif ('walk' in node[0] and 'walk' in node[1]):
        walkEdges.append(node)

# Draw Bus Nodes
try:
    nx.draw_networkx_nodes(G, pos, nodelist=busNodes, node_size=15, node_color='tab:blue', alpha=0.5)
except nx.NetworkXError:
    print('Error drawing bus nodes')

# Draw Metro Nodes
try:
    nx.draw_networkx_nodes(G, pos, nodelist=metroNodes, node_size=25, node_color='tab:red', alpha=0.5)
except nx.NetworkXError:
    print('Error drawing metro nodes')

# Draw Walk Nodes
# try: nx.draw_networkx_nodes(G, pos, nodelist=walkNodes, node_size=5, node_color='tab:green', alpha=0.9)
# except nx.NetworkXError:
#    print('Error drawing walk nodes')

# Draw Bus Edges
try:
    nx.draw_networkx_edges(G, pos, edgelist=busEdges, edge_color='tab:blue', alpha=0.2)
except nx.NetworkXError:
    print('Error drawing bus edges')

# Draw Metro Edges
try:
    nx.draw_networkx_edges(G, pos, edgelist=metroEdges, edge_color='tab:red', alpha=0.2)
except nx.NetworkXError:
    print('Error drawing metro edges')

# Draw Walk Edges
try:
    nx.draw_networkx_edges(G, pos, edgelist=walkEdges, edge_color='tab:green', alpha=0.2)
except nx.NetworkXError:
    print('Error drawing walk edges')

# Draw Path Edges
try:
    nx.draw_networkx_edges(G, pos, edgelist=pathEdges, edge_color='black', width=5)
except nx.NetworkXError:
    print('Error drawing path edges')


# Draw Labels	
# nx.draw_networkx_labels(G, pos, font_size=10)

def maximize():
    plot_backend = plt.get_backend()
    mng = plt.get_current_fig_manager()
    if plot_backend == 'TkAgg':
        mng.resize(*mng.window.maxsize())
    elif plot_backend == 'wxAgg':
        mng.frame.Maximize(True)
    elif plot_backend == 'Qt4Agg':
        mng.window.showMaximized()


plt.axis('equal')
maximize()
plt.show()

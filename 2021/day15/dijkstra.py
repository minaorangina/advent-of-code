import math

graph = {
  'a':{'b':3,'c':4,'d':7},
  'b':{'c':1,'f':5},
  'c':{'f':6, 'd':2},
  'd':{'e':3, 'g':6},
  'e':{'g':3,'h':4},
  'f':{'e':1,'h':8},
  'g':{'h':2},
  'h':{'g':2},
}

def dijkstra(graph, start, end):
    best_route_costs = {} # best route from start to <ref node>
    preceding_nodes = {}
    unseen_nodes = graph
    path = []

    for node in unseen_nodes:
        best_route_costs[node] = math.inf
    
    best_route_costs[start] = 0
 
    while unseen_nodes:
        current_node = None

        # find the nearest/cheapest node relative to the start
        # just a lookup of unseen nodes in the best_route_costs
        for node in unseen_nodes:
            if current_node is None:
                current_node = node 
            elif best_route_costs[node] < best_route_costs[current_node]: 
                current_node = node
        
        neighbours = graph[current_node].items()

        # before we go, opportunity to see if we've incidentally found better routes from start to each neighbour
        # "is it cheaper to go via me to <neighbour> than via some other previously-discovered route to <neighbour>?"
        # 'start -> me -> neighbour' vs 'start -> unknown -> neighbour'
        for neighbour, weight in neighbours:
            if weight + best_route_costs[current_node] < best_route_costs[neighbour]: # best route out of here
                best_route_costs[neighbour] = weight + best_route_costs[current_node]
                # update the directions from that neighbour node towards the start
                preceding_nodes[neighbour] = current_node

        unseen_nodes.pop(current_node)


dijkstra(graph, 'a', 'h')
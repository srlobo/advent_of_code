# -*- coding: utf-8 -*-
from pprint import pprint

def find_next_hop(graph, start_node, visited_nodes, visited_once, print_offset):
    local_print_offset = print_offset + "    "
    print(f"{local_print_offset}start_node: {start_node}, visited_nodes: {visited_nodes}"
          f" visited_once: {visited_once}")

    local_visited_nodes = visited_nodes.copy()
    local_visited_once = visited_once.copy()
    if start_node.islower() and start_node != 'start':
        if start_node in local_visited_once or len(local_visited_nodes) >= 1:
            local_visited_nodes.add(start_node)
            local_visited_nodes.update(local_visited_once)
        else:
            local_visited_once.add(start_node)

    for edge in graph:
        # print(f"{local_print_offset}Probando {edge}")
        if edge[0] != start_node:
            # print(f"Edge {edge} no empieza en {start_node}")
            pass
        elif edge[1] == 'start':
            print(f"{local_print_offset}Edge {edge} vuelve a start -> nada")
        elif edge[1] == 'end':
            print(f"{local_print_offset}Edge {edge} termina en {edge[1]}")
            yield([edge[1], start_node])
        elif edge[1] not in local_visited_nodes:
            print(f"{local_print_offset}Edge {edge} si empieza en {start_node} y no aparece en {local_visited_nodes} -> Drill down en {edge[1]}")
            for g in find_next_hop(graph, edge[1], local_visited_nodes, local_visited_once, local_print_offset):
                print(f"{local_print_offset}res find_next_hop: {g}")
                if g:
                    g.append(start_node)
                    print(f"{local_print_offset}g is not none -> {g}")
                    yield(g)
        else:
            print(f"{local_print_offset}Edge {edge} empieza en {start_node} y está en {local_visited_nodes}")

graph = []
with open("input.txt") as f:
#with open("mini_input.txt") as f:
#with open("mini_input2.txt") as f:
#with open("mini_input3.txt") as f:
    for line in f:
        line = line.strip()
        graph.append(line.split('-'))
        graph.append(line.split('-')[::-1])

for edge in graph:
    print(edge)

start = 'start'
visited_nodes = set()
visited_once = set()
res = []
for g in find_next_hop(graph, start, visited_nodes, visited_once, ""):
    if g:
        res.append(','.join(g[::-1]))

res.sort()
for el in res:
    print(el)
print(len(res))

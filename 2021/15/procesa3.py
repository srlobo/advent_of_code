# -*- coding: utf-8 -*-
from pprint import pprint
import heapq
from sty import fg

def dijkstra(matrix, start, end):
    unvisited_nodes = set()
    for y in range(len(matrix)):
        for x in range(len(matrix)):
            unvisited_nodes.add((x, y))

    tentative_distance = {}

    current_node = start
    current_val = 0
    tentative_distance[current_node] = current_val

    while end in unvisited_nodes:
        print(f"Entramos en {current_node}: {current_val}")
        for new_node in get_4_adjacent_pos(matrix, current_node):
            print(f"Probando {new_node}")
            if new_node in unvisited_nodes:
                print("Esta en unvisited_nodes")
                new_node_val = matrix[new_node[1]][new_node[0]] + current_val
                if new_node in tentative_distance:
                    if new_node_val < tentative_distance[new_node]:
                        tentative_distance[new_node] = new_node_val
                else:
                    tentative_distance[new_node] = new_node_val

        print(f"Borramos {current_node} de unvisited_nodes")
        unvisited_nodes.remove(current_node)
        del(tentative_distance[current_node])

        if end not in unvisited_nodes:
            return current_val

        current_node = min(tentative_distance, key=tentative_distance.get)
        current_val = tentative_distance[current_node]


    return current_val



def paint_matrix_with_glowing_numbers(matrix, glowing_numbers):
    for y in range(len(matrix)):
        row = []
        for x in range(len(matrix[y])):
            if (x, y) in glowing_numbers:
                row.append(fg.blue)
                row.append(matrix[y][x])
                row.append(fg.rs)
            else:
                row.append(matrix[y][x])
        print(''.join([str(c) for c in row]))

def get_4_adjacent_pos(matrix, pos):
    """ Get at most 4 adjacent positions (if they are inside the matrix)"""
    matrix_size = (len(matrix[0]), len(matrix))

    ret_pos = (pos[0] + 1, pos[1])
    if ret_pos[0] >= 0 and ret_pos[0] < matrix_size[0] \
            and ret_pos[1] >= 0 and ret_pos[1] < matrix_size[1]:
        yield ret_pos

    ret_pos = (pos[0], pos[1] + 1)
    if ret_pos[0] >= 0 and ret_pos[0] < matrix_size[0] \
            and ret_pos[1] >= 0 and ret_pos[1] < matrix_size[1]:
        yield ret_pos

    ret_pos = (pos[0] - 1, pos[1])
    if ret_pos[0] >= 0 and ret_pos[0] < matrix_size[0] \
            and ret_pos[1] >= 0 and ret_pos[1] < matrix_size[1]:
        yield ret_pos

    ret_pos = (pos[0], pos[1] - 1)
    if ret_pos[0] >= 0 and ret_pos[0] < matrix_size[0] \
            and ret_pos[1] >= 0 and ret_pos[1] < matrix_size[1]:
        yield ret_pos


matrix = []
#with open("mini_input.txt") as f:
with open("input.txt") as f:
    for line in f:
        line = line.strip()
        matrix.append([int(c) for c in line])

derivated_matrix = []
for step in range(5):
    for row in matrix:
        new_row = []
        for step_x in range(5):
            for el in row:
                el = el + step + step_x
                el = (el % 10) + (el // 10)
                new_row.append(el)
        derivated_matrix.append(new_row)

matrix = derivated_matrix

print(len(matrix), len(matrix[0]))
start_pos = (0, 0)
end_pos = (len(matrix) - 1, len(matrix[0]) - 1)
print(start_pos)
print(end_pos)

paint_matrix_with_glowing_numbers(matrix, [])



print(dijkstra(matrix, start_pos, end_pos))

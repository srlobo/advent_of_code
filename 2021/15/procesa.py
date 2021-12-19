# -*- coding: utf-8 -*-
from pprint import pprint
from sty import fg

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

def calculate_shortest_cost_path(matrix, start_pos, end_pos, cost, path):
    global MAX_COST
    # print(f"Entramos en calculate_shortest_cost_path start: {start_pos} "
    #      f"end_pos:{end_pos}; cost: {cost}; path: {path}; MAX_COST: {MAX_COST}")
    matrix_size = (len(matrix[0]), len(matrix))
    total_cost = cost + matrix[start_pos[1]][start_pos[0]]

    pos = start_pos
    new_path = path.copy()
    new_path.append(pos)

    # paint_matrix_with_glowing_numbers(matrix, new_path)

    estimated_cost_to_end = end_pos[0] - pos[0] + end_pos[1] - pos[1]
    if total_cost + estimated_cost_to_end > MAX_COST:
        # print(f"Nos hemos pasado -> {total_cost + estimated_cost_to_end}")
        return (MAX_COST, new_path)

    if pos == end_pos:
        MAX_COST = total_cost
        print(f"Encontrado!!!!! {MAX_COST}")
        return (total_cost, new_path)


    costs = {}
    partial_paths = {}
    for new_pos in get_4_adjacent_pos(matrix, pos):
        if new_pos in new_path:
            # print("No volvemos atras")
            continue

        new_cost, partial_path = calculate_shortest_cost_path(matrix, new_pos, end_pos,
                                                total_cost, new_path)

        costs[new_pos] = new_cost
        partial_paths[new_pos] = partial_path
    if len(costs) == 0:
        # Nos hemos metido en un callejon sin salida
        return (MAX_COST, new_path)

    min_cost = min(costs, key=costs.get)
    new_path.extend(partial_paths[min_cost])
    # print(f"min_cost: {min_cost}")

    return (total_cost + costs[min_cost],
            new_path)

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

pos = (0, 0)
end_pos = (len(matrix) - 1, len(matrix[0]) - 1)
global MAX_COST
# MAX_COST = 9 * len(matrix) * len(matrix[0])
MAX_COST = 1242

path = []
cost = 0

cost, path = calculate_shortest_cost_path(matrix, pos, end_pos, cost, path)

print(MAX_COST - matrix[0][0])
print(path)

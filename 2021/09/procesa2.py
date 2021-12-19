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

matrix = []
v_points = set()
h_points = set()
with open("input.txt") as f:
#with open("mini_input.txt") as f:
    for line in f:
        matrix_row = [int(c) for c in line.strip()]
        matrix.append(matrix_row)

for y in range(len(matrix)):
    row = matrix[y]
    for x in range(len(row)):
        if x == 0:
            if row[x] < row[x + 1]:
                h_points.add((x,y))
        elif x == len(row) - 1:
            if row[x] < row[x - 1]:
                h_points.add((x,y))
        else:
            if (row[x] < row[x + 1]) and (row[x] < row[x - 1]):
                h_points.add((x,y))

for y in range(len(matrix)):
    for x in range(len(matrix[y])):
        if y == 0:
            if matrix[y][x] < matrix[y + 1][x]:
                v_points.add((x,y))
        elif y == len(matrix) - 1:
            if matrix[y][x] < matrix[y - 1][x]:
                v_points.add((x,y))
        else:
            if (matrix[y][x] < matrix[y + 1][x]) \
                    and (matrix[y][x] < matrix[y - 1][x]):
                v_points.add((x,y))


print("Numero de puntos a tratar: %s" % len(h_points.intersection(v_points)))
s = 0
lowest_points = h_points.intersection(v_points)

def get_8_around_points(x, y, matrix):
    val = matrix[y][x]
    # print(f"Calculando alrededores de ({x},{y}), valor {val}")
    for r_x in range(x - 1, x + 2):
        for r_y in range(y - 1, y + 2):
            # print(f"Viendo si ({r_x},{r_y}) >= (0, 0)")
            if r_x >= 0 and r_y >= 0 \
                    and r_y < len(matrix) and r_x < len(matrix[r_y]):
                # print(f"Comparando ({r_x},{r_y}), valor {matrix[r_y][r_x]}")
                if matrix[r_y][r_x] == 9:
                    continue
                if matrix[r_y][r_x] == val + 1:
                    # print(f"Encontrado ({r_x},{r_y}), valor {matrix[r_y][r_x]} == {val} + 1")
                    yield (r_x, r_y)

basins = []
for lp_x, lp_y in lowest_points:
    candidates = set()
    processed = set()
    processed.add((lp_x, lp_y))

    candidates.update(get_8_around_points(lp_x, lp_y, matrix))
    candidates = candidates - processed
    while candidates != set():
        new_candidates = set()
        for x, y in candidates:
            new_candidates.update(get_8_around_points(x, y, matrix))
            processed.add((x, y))

        candidates = new_candidates - processed
    print(processed)
    print(len(processed))
    paint_matrix_with_glowing_numbers(matrix, processed)
    basins.append(len(processed))
basins.sort(reverse=True)
print(basins)
print(basins[0] * basins[1] * basins[2])

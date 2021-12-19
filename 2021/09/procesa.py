# -*- coding: utf-8 -*-
from pprint import pprint
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


print(len(h_points.intersection(v_points)))
s = 0
for x, y in h_points.intersection(v_points):
    s += matrix[y][x] + 1

print(s)



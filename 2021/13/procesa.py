# -*- coding: utf-8 -*-
from pprint import pprint

points = set()
with open("input.txt") as f:
#with open("mini_input.txt") as f:
    instruction_mode = False
    instructions = []
    for line in f:
        line = line.strip()
        if line == '':
            instruction_mode = True
            continue
        if not instruction_mode:
            points.add(tuple(map(int, line.split(','))))
        else:
            instruction, number = line.split('=')
            if 'y' in instruction:
                instructions.append(('y', number))
            else:
                instructions.append(('x', number))

for axis, n in instructions:
    axis = instructions[0][0]
    n = int(n)

    new_points = set()
    if axis == 'x':
        for point in points:
            x, y = point
            if x > n:
                new_x = n - (x - n)
                new_points.add((new_x, y))
            else:
                new_points.add(point)
    elif axis == 'y':
        for point in points:
            x, y = point
            if y > n:
                new_y = n - (y - n)
                new_points.add((x, new_y))
            else:
                new_points.add(point)
    else:
        print(f"Problemas con el axis: {axis}")

    points = new_points
    break

def print_matrix(points):
    max_x = max([point[0] for point in points])
    max_y = max([point[1] for point in points])

    for y in range(max_y + 1):
        row = []
        for x in range(max_x + 1):
            if (x, y) in points:
                row.append('#')
            else:
                row.append('.')
        print(''.join(row))

print(points)
print_matrix(points)
print(len(points))

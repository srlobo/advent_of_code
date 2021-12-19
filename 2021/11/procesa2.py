# -*- coding: utf-8 -*-
from pprint import pprint
from sty import fg

flash_count = 0

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

def get_8_around_points(x, y, matrix):
    for r_x in range(x - 1, x + 2):
        for r_y in range(y - 1, y + 2):
            if r_x >= 0 and r_y >= 0 \
                    and r_y < len(matrix) and r_x < len(matrix[r_y]) and \
                    (r_x, r_y) != (x, y):
                yield(r_x, r_y)

def print_matrix(matrix):
    for row in matrix:
        print(''.join(map(str, row)))

matrix = []
with open("input.txt") as f:
#with open("mini_input.txt") as f:
    for line in f:
        line = line.strip()
        print(line)
        matrix.append([int(c) for c in line])

step = 0
must_exit = False
while not must_exit:
    step += 1
    print(f"Step {step}")

    global_flashed = set()
    local_flashed = set()
    # We increase every position by one
    for y in range(len(matrix)):
        for x in range(len(matrix[y])):
            matrix[y][x] += 1
            # And capture every position greater than 9
            if matrix[y][x] > 9:
                local_flashed.add((x, y))

    while len(local_flashed) > 0:
        #print("flash:")
        #paint_matrix_with_glowing_numbers(matrix, local_flashed)

        new_local_flashed = set()
        #print(local_flashed)
        for x, y in local_flashed:
            for point in get_8_around_points(x, y, matrix):
                if point not in global_flashed:
                    rx, ry = point
                    if matrix[ry][rx] > 9:
                        continue
                    matrix[ry][rx] += 1
                    if matrix[ry][rx] > 9:
                        new_local_flashed.add((rx,ry))
            global_flashed.add((x, y))

        # print("new_local_flashed", new_local_flashed)
        local_flashed = new_local_flashed

    #print(global_flashed)
    #print(len(global_flashed))
    for x, y in global_flashed:
        matrix[y][x] = 0
    if len(global_flashed) == len(matrix) * len(matrix[1]):
        must_exit = True

    print_matrix(matrix)
    print("------------")
    flash_count += len(global_flashed)
print(f"Total: {flash_count}")

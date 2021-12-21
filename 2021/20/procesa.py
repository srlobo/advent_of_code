# -*- coding: utf-8 -*-
from pprint import pprint
import logging

def print_matrix(matrix):
    for row in matrix:
        print(''.join(row))

def iterate_9_by_9(matrix, translate):
    new_matrix = []
    for y in range(1, len(matrix) - 1):
        row = []
        for x in range(1, len(matrix[y]) - 1):
            el0 = pixel2bin(matrix[y - 1][x - 1] +
                            matrix[y - 1][x] +
                            matrix[y - 1][x + 1])
            el1 = pixel2bin(matrix[y][x -1] + matrix[y][x] + matrix[y][x + 1])
            el2 = pixel2bin(matrix[y + 1][x - 1] +
                            matrix[y + 1][x] +
                            matrix[y + 1][x + 1])

            #print(el0 + el1 + el2)
            res = translate[int(el0 + el1 + el2, 2)]
            #print(res)
            row.append(res)
        new_matrix.append(''.join(row))
    return new_matrix


def pixel2bin(pixels):
    s = []
    for p in pixels:
        if p == '.':
            s.append('0')
        else:
            s.append('1')
    return ''.join(s)

logger = logging.getLogger(__name__)
logging.basicConfig(level=logging.CRITICAL)
def main():
    #with open("mini_input.txt") as f:
    with open("input.txt") as f:
        c = 0
        matrix = []
        for line in f:
            line = line.strip()
            if c == 0:
                translate = line
            elif c == 1:
                pass
            else:
                matrix.append([c for c in line])
            c += 1

    steps = 2
    print(matrix)
    matrix = reformat_matrix(matrix, steps)
    print("reformat")
    print_matrix(matrix)
    for s in range(steps):
        print("iterate")
        matrix = iterate_9_by_9(matrix, translate)
        print_matrix(matrix)
    print(count_pixels(matrix))

def count_pixels(matrix):
    c = 0
    for row in matrix:
        for p in row:
            if p == '#':
                c += 1

    return c

def reformat_matrix(matrix, steps):
    """Add space on the border to be able to process matrix again"""
    res_matrix = []
    row_length = len(matrix[0]) + (steps * 6)
    # print(f"{len(matrix[0])} vs {row_length}")
    for c in range(steps * 3):
        res_matrix.append(['.'] * (row_length))
    for row in matrix:
        new_row = ['.'] * (3 * steps)
        new_row.extend(row)
        new_row.extend(['.'] * (3 * steps))
        res_matrix.append(new_row)
    for c in range(steps * 3):
        res_matrix.append(['.'] * (row_length))

    return res_matrix

if __name__ == '__main__':
    main()

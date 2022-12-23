# -*- coding: utf-8 -*-
from pprint import pprint
import logging

logger = logging.getLogger(__name__)
logging.basicConfig(level=logging.DEBUG)
def main():
    bottom = []
    #with open("mini_input.txt") as f:
    with open("input.txt") as f:
        for line in f:
            line = line.strip()
            row = [c for c in line]
            bottom.append(row)
    res = bottom
    print_bottom(res)
    for s in range(10000):
        new_res = step(res)
        print(f'--------- step {s + 1} -----------')
        # print_bottom(res)
        if compare_bottoms(new_res, res):
            print("Encontrado")
            break
        res = new_res

def print_bottom(bottom):
    for row in bottom:
        print(''.join(row))

def compare_bottoms(bottom1, bottom2):
    for y in range(len(bottom1)):
        for x in range(len(bottom1[y])):
            if bottom1[y][x] != bottom2[y][x]:
                return False
    return True


def step(bottom):
    res = []
    for row in bottom:
        res_row = []
        i = 0
        while i < len(row):
            if row[i] == '>':
                # Comprobamos si hay hueco
                next_index = (i + 1) % len(row)
                if row[next_index] == '.':
                    res_row.append('.')
                    if next_index == 0:
                        res_row[0] = '>'
                    else:
                        res_row.append('>')
                        i += 1
                else:
                    res_row.append(row[i])
            else:
                if len(res_row) <= i:
                    res_row.append(row[i])
            i += 1
        res.append(res_row)

    bottom = res
    res = []
    for _ in range(len(bottom)):
        res.append(['-'] * len(bottom[0]))
    for x in range(len(res[0])):
        y = 0
        while y < len(bottom):
            if bottom[y][x] == 'v':
                # Comprobamos si hay hueco
                next_index = (y + 1) % len(bottom)
                if bottom[next_index][x] == '.':
                    res[y][x] = '.'
                    if next_index == 0:
                        res[0][x] = 'v'
                    else:
                        y += 1
                        res[y][x] = 'v'
                else:
                    res[y][x] = bottom[y][x]
            else:
                res[y][x] = bottom[y][x]
            y += 1

    return res


if __name__ == '__main__':
    main()

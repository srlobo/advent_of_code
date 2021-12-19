# -*- coding: utf-8 -*-
#with open("mini_input.txt") as f:
with open("input.txt") as f:
    pos = 0
    depth = 0
    aim = 0
    for l in f:
        order, n = l.strip().split(' ')
        n = int(n)
        if order == 'forward':
            pos += n
            depth += n * aim
        elif order == 'down':
            aim += n
        elif order == 'up':
            aim -= n

print(pos * depth)



# -*- coding: utf-8 -*-
from collections import Counter

#with open("mini_input.txt") as f:
with open("input.txt") as f:
    points = []
    for l in f:
        print('----------')
        src, dst = l.strip().split(' -> ')
        src_x, src_y = src.split(',')
        dst_x, dst_y = dst.split(',')
        print(f'{src_x} {src_y}')
        print(f'{dst_x} {dst_y}')
        x_coordinate = None
        y_coordinate = None
        if src_x == dst_x:
            # Trabajamos en el eje y
            s = int(src_y)
            d = int(dst_y)
            x_coordinate = int(src_x)
        elif src_y == dst_y:
            # Trabajamos en el eje x
            s = int(src_x)
            d = int(dst_x)
            y_coordinate = int(src_y)
        else:
            print("No son iguales:")
            print(f'{src_x} {src_y}')
            print(f'{dst_x} {dst_y}')
            continue

        if s > d:
            s, d = d, s
        print(f'de {s} a {d}')
        for point in range(s, d + 1):
            print(f"{point} ", )
            if x_coordinate:
                points.append((x_coordinate, point))
            else:
                points.append((point, y_coordinate))
        print()

    total = 0
    for value in Counter(points).values():
        if value > 1:
            total += 1
    print(total)

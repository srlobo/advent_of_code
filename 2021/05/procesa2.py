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
        src_x = int(src_x)
        dst_x = int(dst_x)
        src_y = int(src_y)
        dst_y = int(dst_y)
        print(f'{src_x} {src_y}')
        print(f'{dst_x} {dst_y}')
        x_coordinate = None
        y_coordinate = None
        if src_x == dst_x:
            # Trabajamos en el eje y
            s = int(src_y)
            d = int(dst_y)
            x_coordinate = int(src_x)
            if s > d:
                s, d = d, s
            print(f'de {s} a {d}')
            for point in range(s, d + 1):
                print(f"{point} ", )
                points.append((x_coordinate, point))
            print('------')
        elif src_y == dst_y:
            # Trabajamos en el eje x
            s = int(src_x)
            d = int(dst_x)
            y_coordinate = int(src_y)
            if s > d:
                s, d = d, s
            print(f'de {s} a {d}')
            for point in range(s, d + 1):
                print(f"{point} ", )
                points.append((point, y_coordinate))
            print('------')
        else:
            # Trabajamos en ambos ejes
            if src_x < dst_x:
                step_x = 1
            else:
                step_x = -1

            if src_y < dst_y:
                step_y = 1
            else:
                step_y = -1

            point = (src_x, src_y)
            while point != (dst_x, dst_y):
                points.append(point)
                point = (point[0] + step_x, point[1] + step_y)
            # Aniadimos tambien el ultimo punto
            points.append(point)

    total = 0
    for value in Counter(points).values():
        if value > 1:
            total += 1
    print(total)

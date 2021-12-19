# -*- coding: utf-8 -*-
from pprint import pprint
import logging

logger = logging.getLogger(__name__)
logging.basicConfig(level=logging.CRITICAL)
#with open("mini_input.txt") as f:
with open("input.txt") as f:
    for line in f:
        line = line.strip()
        logger.debug("Processing %s", line)
        target_area_arr = line.split(': ')[1].split(', ')
        for target_area_str in target_area_arr:
            if 'x' in target_area_str:
                payload = target_area_str.split('=')[1]
                x_range = list(map(int, payload.split('..')))
            if 'y' in target_area_str:
                payload = target_area_str.split('=')[1]
                y_range = list(map(int, payload.split('..')))


def point_inside_square(point, x_range, y_range):
    x, y = point
    if x >= x_range[0] and x <= x_range[1]:
        if y >= y_range[0] and y <= y_range[1]:
            return True

    return False

def simulate_shoot(vector, x_range, y_range):
    pos = (0,0)
    steps = 0
    positions = []
    while pos[0] <= x_range[1] and pos[1] >= y_range[1]:
        new_pos_x = pos[0] + vector[0]
        new_pos_y = pos[1] + vector[1]
        new_vector_x = vector[0] - 1
        if new_vector_x < 0:
            new_vector_x = 0
        new_vector_y = vector[1] - 1
        pos = (new_pos_x, new_pos_y)
        positions.append(pos)
        vector = (new_vector_x, new_vector_y)
    return positions

def simulate_shoot_and_check_target(vector, x_range, y_range):
    seq = simulate_shoot(vector, x_range, y_range)
    #print(f"Simulating {vector} -> {seq}")
    for pos in seq:
        if point_inside_square(pos, x_range, y_range):
            return True, seq
    return False, seq

points_and_heights = {}
for x in range(1, x_range[1] + 1):
    for y in range(200, y_range[1] - 1, -1):
        success, seq = simulate_shoot_and_check_target(
            (x, y), x_range, y_range)
        if success:
            try:
                max_height = max(seq, key=lambda x: x[1])[1]
            except Exception as e:
                print(f"Problemas, {(x,y)}; {seq}")
                raise e
            # print(f"{(x,y)} -> {max_height}")
            points_and_heights[(x,y)] = max_height

print(f"{x_range} {y_range}")

pprint(points_and_heights)
print(max(points_and_heights.values()))

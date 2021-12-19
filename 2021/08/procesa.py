# -*- coding: utf-8 -*-
from pprint import pprint
import statistics
#with open("mini_input.txt") as f:
with open("input.txt") as f:
    c = 0
    for line in f:
        # De momento nos quedamos con la segunda parte despues del |
        car, cdr = line.split('|')
        for el in cdr.strip().split(' '):
            size_el = len(el)
            if size_el == 2: # Es un 1
                c += 1
            elif size_el == 4: # Es un 4
                c += 1
            elif size_el == 3: # Es un 7
                c += 1
            elif size_el == 7: # Es un 8
                c += 1
print(c)

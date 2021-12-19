# -*- coding: utf-8 -*-
stack = []
open_signs = ["{", "[", "<", "("]
close_signs = ["}", "]", ">", ")"]
scores = [1197, 57, 25137, 3]
s = 0
with open("input.txt") as f:
#with open("mini_input.txt") as f:
    for line in f:
        line = line.strip()
        print(line)
        for c in line:
            if c in open_signs:
                stack.append(c)
            elif c in close_signs:
                c_i = close_signs.index(c)

                o = stack.pop()
                c_o = open_signs.index(o)

                if c_i != c_o:
                    print(f"Error!!!! Buscabamos {close_signs[c_o]} y encontramos {close_signs[c_i]}")
                    s += scores[c_i]
print(s)

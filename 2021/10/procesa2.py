# -*- coding: utf-8 -*-
open_signs = ["{", "[", "<", "("]
close_signs = ["}", "]", ">", ")"]
scores = [3, 2, 4, 1]
final_scores = []
with open("input.txt") as f:
#with open("mini_input.txt") as f:
    for line in f:
        s = 0
        stack = []
        line = line.strip()
        errors = False
        for c in line:
            if c in open_signs:
                stack.append(c)
            elif c in close_signs:
                c_i = close_signs.index(c)

                o = stack.pop()
                c_o = open_signs.index(o)

                if c_i != c_o:
                    # print(line)
                    # print(f"Error!!!! Buscabamos {close_signs[c_o]} y encontramos {close_signs[c_i]}")
                    errors = True
                    break
        if errors:
            continue

        completion = []
        for el in stack[::-1]:
            c_o = open_signs.index(el)
            completion.append(close_signs[c_o])
            s = s*5 + scores[c_o]
        print(line)
        print("Completion: %s" % ''.join(completion))
        print(s)
        final_scores.append(s)
final_scores.sort()
print(final_scores)
print(final_scores[len(final_scores)//2])


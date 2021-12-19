# -*- coding: utf-8 -*-
with open("kkfu") as f:
    for l in f:
        row = []
        c = ""
        print(l.strip())
        for c_pos in range(len(l.strip())):
            print(f"row {row}")
            if l[c_pos] != ' ':
                c += l[c_pos]
            else:
                if c != '':
                    row.append(c)
                    c = ''

        if c != '':
            row.append(c)
        print(f"row (final) {row}")


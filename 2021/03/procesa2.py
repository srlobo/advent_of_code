# -*- coding: utf-8 -*-
#with open("mini_input.txt") as f:
with open("input.txt") as f:
    t = []
    for l in f:
        t.append(l.strip())

    l = len(t[0])
    todo = t
    for pos in range(l):
        group_1 = []
        group_0 = []
        for el in todo:
            print(f"el {el} -> {el[pos]}")
            if el[pos] == '1':
                print("1")
                group_1.append(el)
            else:
                print("0")
                group_0.append(el)

        if len(group_1) >= len(group_0):
            todo = group_1
        else:
            todo = group_0

        if len(todo) == 1:
            break

    print(f"todo {todo[0]}")
    num_1 = int(todo[0], 2)

    todo = t
    for pos in range(l):
        print(f"_________{pos}__________")
        group_1 = []
        group_0 = []
        for el in todo:
            print(f"el {el} -> {el[pos]}")
            if el[pos] == '1':
                print("1")
                group_1.append(el)
            else:
                print("0")
                group_0.append(el)

        print(f"g1 {len(group_1)} g2 {len(group_0)}")
        if len(group_1) < len(group_0):
            todo = group_1
        else:
            todo = group_0

        if len(todo) == 1:
            break

    print(f"todo {todo}")

    num_2 = int(todo[0], 2)

    print(f'{num_1}, {num_2}, {num_1*num_2}')

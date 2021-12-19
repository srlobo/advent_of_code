# -*- coding: utf-8 -*-
#with open("mini_input.txt") as f:
with open("input.txt") as f:
    cuenta_total = []
    num_filas = 0
    for l in f:
        if num_filas == 0:
            cuenta_total = [0] * len(l.strip())
        num_filas += 1
        # Contamos los 1s
        count = 0
        for c in l.strip():
            if c == '1':
                cuenta_total[count] += 1

            count +=1

res = ''
not_res = ''
for a in cuenta_total:
    if a > num_filas / 2:
        res += '1'
        not_res += '0'
    else:
        res += '0'
        not_res += '1'

res_binary = int(res, 2)
not_res_binary = int(not_res, 2)
print(res_binary)
print(not_res_binary)
print(res_binary * not_res_binary)

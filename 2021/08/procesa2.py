# -*- coding: utf-8 -*-
from pprint import pprint
import sys

bigsum = 0
#with open("mini_input.txt") as f:
with open("input.txt") as f:
    for line in f:
        car, cdr = line.split('|')
        translation = {}
        data = {}
        while len(data) != 10:
            for el in car.strip().split(' '):
                set_el = set([l for l in el])
                size_el = len(set_el)
                print(el)
                localizado = False
                for key, value in data.items():
                    if set_el == value:
                        print(f"Localizado: Es un {key}")
                        localizado = True
                        break
                if localizado:
                    continue

                if size_el == 2: # Es un 1
                    print("Es un 1")
                    data[1] = set_el
                elif size_el == 4: # Es un 4
                    print("Es un 4")
                    data[4] = set_el
                elif size_el == 3: # Es un 7
                    print("Es un 7")
                    data[7] = set_el
                elif size_el == 7: # Es un 8
                    print("Es un 8")
                    data[8] = set_el
                elif size_el == 6: # Es un 0, 6 o 9
                    if 1 in data:
                        if set_el.intersection(data[1]) == data[1]: # La barra de la derecha
                            if 4 in data:
                                if data[4].intersection(set_el) == data[4]:
                                    print("Es un 9")
                                    data[9] = set_el
                                else:
                                    print("Es un 0")
                                    data[0] = set_el
                            else:
                                print("No sabemos si es un 9 o un 0, nos falta el 4")
                        else:
                            print("es un 6")
                            data[6] = set_el
                    else:
                        print("Es un 0, 6 o 9, nos falta un 1")
                elif size_el == 5: # Es un 2, 3 o 5
                    if 1 in data:
                        if set_el.intersection(data[1]) == data[1]:
                            print("Es un 3")
                            data[3] = set_el
                        else:
                            if 6 in data:
                                if set_el - data[6] == set(): # 5 - 6 -> vacio
                                    print("Es un 5")
                                    data[5] = set_el
                                else:
                                    print("Es un 2")
                                    data[2] = set_el
                            else:
                                print("Es un 2 o 5, necesitamos un 6")
                    else:
                        print("No sabemos si es un 2, 3 o 5, nos falta un 1")
                else:
                    print(f"Size: {size_el}")
                    sys.exit(1)
        num = []
        for el in cdr.strip().split(' '):
            set_el = set([l for l in el])
            for key, value in data.items():
                if set_el == value:
                    num.append(key)
                    break
        print("Resultado: %s" % int(''.join(map(str, num))))
        bigsum += int(''.join(map(str, num)))
print(f"Final res: {bigsum}")

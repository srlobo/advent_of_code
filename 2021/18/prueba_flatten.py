# -*- coding: utf-8 -*-
import procesa

expr = eval('[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]')
expr = eval('[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]')

tree = procesa.binary_tree_constructor(expr)

print(tree)
print(procesa.flatten_tree(tree))

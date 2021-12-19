# -*- coding: utf-8 -*-
import procesa
import logging

expr = eval('[[[[7,7],[7,8]],[[9,5],[8,0]]],[[[9,10],20],[8,[9,0]]]]')

tree = procesa.binary_tree_constructor(expr)


logger = logging.getLogger(__name__)
logging.basicConfig(level=logging.DEBUG)
logging.getLogger(procesa.__name__).setLevel(logging.DEBUG)
print(tree)
procesa.snail_reduce((tree))
print(tree)

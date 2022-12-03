# -*- coding: utf-8 -*-
import pytest

@pytest.mark.parametrize("test_input,expected", [
    (("[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
     "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]"),
     "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]"),

    (("[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
      "[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]"),
     "[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]"),

    (("[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]",
      "[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]"),
     "[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]"),

    (("[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]",
      "[7,[5,[[3,8],[1,4]]]]"),
     "[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]"),

    (("[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]",
      "[[2,[2,2]],[8,[8,1]]]"),
      "[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]"),

    (("[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]",
      "[2,9]"),
     "[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]"),

    (("[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]",
      "[1,[[[9,3],9],[[9,0],[0,7]]]]"),
     "[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]"),

    (("[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]",
      "[[[5,[7,4]],7],1]"),
     "[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]"),

    (("[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]",
      "[[[[4,2],2],6],[8,7]]"),
     "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"),
])
def test_1(test_input, expected):
    import procesa

    expr = map(eval, test_input)
    res = procesa.binary_tree_constructor(eval(expected))

    assert repr(procesa.calculate_sum(expr)) == repr(res)




@pytest.mark.parametrize("test_input,expected", [
    ('[[1,2],[[3,4],5]]', 143),
    ('[[[[0,7],4],[[7,8],[6,0]]],[8,1]]', 1384),
    ('[[[[1,1],[2,2]],[3,3]],[4,4]]', 445),
    ('[[[[3,0],[5,3]],[4,4]],[5,5]]', 791),
    ('[[[[5,0],[7,4]],[5,5]],[6,6]]', 1137),
    ('[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]', 3488),
])
def test_magnitudes(test_input, expected):
    import procesa

    expr = eval(test_input)
    tree = procesa.binary_tree_constructor(expr)

    assert procesa.calculate_magnitude(tree) == expected


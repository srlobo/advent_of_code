# -*- coding: utf-8 -*-
from pprint import pprint
import logging
import sys

class binary_tree_node(object):
    def __init__(self, parent=None):
        self.left = None
        self.right = None
        self.parent = parent
    def __repr__(self):
        return f"[{repr(self.left)},{repr(self.right)}]"

    def to_list(self):
        if type(self.left) is int:
            left = self.left
        else:
            left = self.left.to_list()
        if type(self.right) is int:
            right = self.right
        else:
            right = self.right.to_list()

        return [left, right]


def binary_tree_constructor(expression, parent=None):
    # logger.debug("Entrando en binary_tree_constructor: %s", expression)
    if type(expression) == int:
        # logger.debug("Entero")
        return expression
    # logger.debug("Array")

    ret = binary_tree_node(parent)

    left = binary_tree_constructor(expression[0], ret)
    # logger.debug("left: %s", repr(left))

    right = binary_tree_constructor(expression[1], ret)
    # logger.debug("right: %s", repr(right))

    ret.left = left
    ret.right = right

    return ret


def snail_sum(a, b):
    ret = binary_tree_node(parent=None)
    a.parent = ret
    b.parent = ret
    ret.left = a
    ret.right = b

    while snail_reduce(ret):
        pass
    return ret

def snail_reduce(tree):
    if snail_reduce1(tree, tree) is False:
        logger.critical("Explode pass: %s", tree)
        return True

    if snail_reduce2(tree, tree) is False:
        logger.critical("Split pass: %s", tree)
        return True

    return False

def snail_reduce2(tree, actual_node):
    logger.debug("snail_reduce2, node: %s", actual_node)
    if type(actual_node.left) is int and actual_node.left >= 10:
        val = actual_node.left
        actual_node.left = binary_tree_node(parent=actual_node)
        actual_node.left.left = val // 2
        actual_node.left.right = val // 2 if val % 2 == 0 else val // 2 + 1
        logger.info("split left -> %s", tree)
        return False
    elif type(actual_node.left) is not int:
        if not snail_reduce2(tree, actual_node.left):
            logger.info("snail_reduce2 left -> %s", tree)
            return False

    if type(actual_node.right) is int and actual_node.right >= 10:
        val = actual_node.right
        actual_node.right = binary_tree_node(parent=actual_node)
        actual_node.right.left = int(val / 2)
        actual_node.right.right = val // 2 if val % 2 == 0 else val // 2 + 1
        logger.info("split right -> %s", tree)
        return False
    elif type(actual_node.right) is not int:
        if not snail_reduce2(tree, actual_node.right):
            logger.info("snail_reduce2 right -> %s", tree)
            return False

    return True


def snail_reduce1(tree, actual_node, depth=0):
    if depth == 4:
        snail_explode(tree, actual_node)
        logger.info("snail_explode -> %s", tree)
        return False

    d = depth + 1

    if type(actual_node.left) is not int:
        if not snail_reduce1(tree, actual_node.left, d):
            #logger.info("snail_reduce1 left -> %s", tree)
            return False

    if type(actual_node.right) is not int:
        if not snail_reduce1(tree, actual_node.right, d):
            #logger.info("snail_reduce1 right -> %s", tree)
            return False

    return True

def snail_explode(tree, exploding_node):
    tree_flattened = flatten_tree(tree)
    logger.debug("snail_explode tree: %s; exploding_node: %s"
                 ", tree_flattened: %s",
                 tree, exploding_node, tree_flattened)

    for i in range(len(tree_flattened)):
        if tree_flattened[i] == exploding_node:
            break
    if i > 0:
        previous_node = tree_flattened[i - 1]
        if type(previous_node.right) is int:
            previous_node.right += exploding_node.left
        elif type(previous_node.left) is int:
            previous_node.left += exploding_node.left

    if i < len(tree_flattened) - 1:
        next_node = tree_flattened[i + 1]
        if type(next_node.left) is int:
            next_node.left += exploding_node.right
        elif type(next_node.right) is int:
            next_node.right += exploding_node.right


#    last_node_with_number, found = search_node_and_previous_number(
#        tree, exploding_node)
#    logger.debug("Res last_node_with_number: %s", last_node_with_number)
#    if found and last_node_with_number is not None:
#        if type(last_node_with_number.right) is int:
#            last_node_with_number.right += exploding_node.left
#        else:
#            last_node_with_number.left += exploding_node.left
#
#    logger.debug("Partial explode, %s", tree)
#    next_node_with_number, found = search_node_and_next_number(tree, exploding_node)
#    if found and next_node_with_number is not None:
#        if type(next_node_with_number.left) is int:
#            next_node_with_number.left += exploding_node.right
#        else:
#            next_node_with_number.right += exploding_node.right

    # And now we change the exploding_node with 0
    if exploding_node.parent.right == exploding_node:
        exploding_node.parent.right = 0
    else:
        exploding_node.parent.left = 0

    logger.debug("snail_explode despues tree: %s; exploding_node: %s"
                 ", tree_flattened: %s",
                 tree, exploding_node, tree_flattened)




def search_node_and_previous_number(node, node_to_search,
                                    last_node_with_number=None):
    logger.debug("search_node_and_previous_number "
                 "current_node: %s look for node: %s node_with_number: %s", node, node_to_search,
                 last_node_with_number)
    if node == node_to_search:
        logger.debug("Node found, returning %s", last_node_with_number)
        return (last_node_with_number, True)

    if type(node.left) is int:
        last_node_with_number = node
    else:
        ret, found = search_node_and_previous_number(node.left, node_to_search,
                                           last_node_with_number)
        if found:
            return ret, True
        else:
            last_node_with_number = ret

    if type(node.right) is int:
        last_node_with_number = node
    else:
        ret, found  = search_node_and_previous_number(node.right, node_to_search,
                                                      last_node_with_number)
        if found:
            return ret, True
        else:
            last_node_with_number = ret

    return last_node_with_number, False

def search_node_and_next_number(node, node_to_search,
                                    node_to_search_found=False):
    logger.debug("search_node_and_next_number "
                 "current_node: %s look for node: %s node_to_search_found: %s",
                 node, node_to_search, node_to_search_found)
    if node == node_to_search:
        logger.debug("Found, returning")
        return None, True

    if type(node.left) is not int:
        ret, success = search_node_and_next_number(node.left, node_to_search,
                                          node_to_search_found)
        if success is True:
            node_to_search_found = True
            logger.debug("left success is true, node: %s, ret: %s", node, ret)
            if ret is not None:
                return ret, True
    elif node_to_search_found: # left is int and we have found the node_to_search
        logger.debug("Found left -> node: %s", node)
        return node, True

    logger.debug("node.right: %s", node.right)
    if type(node.right) is not int:
        ret, success = search_node_and_next_number(node.right, node_to_search,
                                                   node_to_search_found)
        if success is True:
            logger.debug("right success is true, node: %s, ret: %s", node, ret)
            return ret, True
    elif node_to_search_found:
        logger.debug("Found right -> node: %s", node)
        return node, True

    return None, False

def flatten_tree(tree):
    if type(tree.left) is int:
        left = [tree]
    else:
        left = flatten_tree(tree.left)

    if type(tree.right) is int:
        right = [tree]
    else:
        right = flatten_tree(tree.right)

    if left == right:
        return left
    else:
        left.extend(right)
        return left

def calculate_magnitude(tree):
    if type(tree.left) is int:
        left = tree.left
    else:
        left = calculate_magnitude(tree.left)

    if type(tree.right) is int:
        right = tree.right
    else:
        right = calculate_magnitude(tree.right)

    return 3 * left +  2 * right


logger = logging.getLogger(__name__)
logging.basicConfig(level=logging.CRITICAL)
def main():
    exprs = []
    with open(sys.argv[1]) as f:
    #with open("mini_input.txt") as f:
    #with open("input.txt") as f:
        for line in f:
            line = line.strip()

            exprs.append(eval(line))

    s = calculate_sum(exprs)
    print(s)
    print(calculate_magnitude(s))

def calculate_sum(exprs):
    acc = None
    for expr in exprs:
        #logger.setLevel(logging.DEBUG)
        # logger.setLevel(51)
        tree = binary_tree_constructor(expr)
        if acc is None:
            acc = tree
            continue
        tree_str = repr(tree)
        acc_str = repr(acc)

        acc = snail_sum(acc, tree)
        print(f"{acc_str} + {tree_str} -> {acc}")
    return acc

if __name__ == '__main__':
    main()

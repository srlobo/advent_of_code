# -*- coding: utf-8 -*-
from pprint import pprint
import logging

logger = logging.getLogger(__name__)
logging.basicConfig(level=logging.DEBUG)
def main():
    players = {}
    #with open("mini_input.txt") as f:
    with open("input.txt") as f:
        for line in f:
            line = line.strip()

if __name__ == '__main__':
    main()

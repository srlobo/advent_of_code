# -*- coding: utf-8 -*-
from pprint import pprint
import logging
import sys

class ALU(object):
    def __init__(self, w=0, x=0, y=0, z=0):
        self.registers = {
            'w': w,
            'x': x,
            'y': y,
            'z': z,
        }

    def inp(self, variable, value_from):
        self.registers[variable] = next(value_from)

    def add(self, op1, op2):
        if op2 in "wxyz":
            self.registers[op1] += self.registers[op2]
        else:
            self.registers[op1] += int(op2)

    def mul(self, op1, op2):
        if op2 in "wxyz":
            self.registers[op1] *= self.registers[op2]
        else:
            self.registers[op1] *= int(op2)

    def div(self, op1, op2):
        if op2 in "wxyz":
            self.registers[op1] = self.registers[op1] // self.registers[op2]
        else:
            self.registers[op1] = self.registers[op1] // int(op2)

    def mod(self, op1, op2):
        if op2 in "wxyz":
            self.registers[op1] = self.registers[op1] % self.registers[op2]
        else:
            self.registers[op1] = self.registers[op1] % int(op2)

    def eql(self, op1, op2):
        if op2 in "wxyz":
            second_operand = self.registers[op2]
        else:
            second_operand = int(op2)

        if self.registers[op1] == second_operand:
            self.registers[op1] = 1
        else:
            self.registers[op1] = 0

    def print_state(self):
        print(self.registers)

def eval_program(program, input_, w, x, y, z):
    state = ALU(w, x, y, z)
    for statement in program:
        instruction = statement.split(' ')
        if instruction[0] == 'inp':
            state.inp(instruction[1], input_)
        elif instruction[0] == 'add':
            state.add(instruction[1], instruction[2])
        elif instruction[0] == 'mul':
            state.mul(instruction[1], instruction[2])
        elif instruction[0] == 'div':
            state.div(instruction[1], instruction[2])
        elif instruction[0] == 'mod':
            state.mod(instruction[1], instruction[2])
        elif instruction[0] == 'eql':
            state.eql(instruction[1], instruction[2])
    return state

def gen_input(input_text):
    for c in input_text:
        yield int(c)

logger = logging.getLogger(__name__)
logging.basicConfig(level=logging.DEBUG)
def main():
    players = {}
    #with open("mini_input.txt") as f:
    with open(sys.argv[1]) as f:
        instructions = []
        for line in f:
            line = line.strip()
            instructions.append(line)
    x, y, w = 0, 0, 0
    for input_ in range(1, 10):
        for z in range(15, 25):
            res = eval_program(instructions, gen_input(
                str(input_)), w, x, y, z)
            print(f"input: {input_} {w} {x} {y} {z}, "
                  f"output: {res.registers['w']}"
                  f" {res.registers['x']}"
                  f" {res.registers['y']}"
                  f" {res.registers['z']}")
            if res.registers['z'] == 0:
                print("Este!!!")

if __name__ == '__main__':
    main()

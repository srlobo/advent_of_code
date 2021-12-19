# -*- coding: utf-8 -*-
from pprint import pprint
def sum_board(board):
    s = 0
    for x in range(len(board[0])):
        for y in range(len(board)):
            if board[y][x] != 'x':
                s += int(board[y][x])

    return s

def check_if_board_wins(board):
    # Count number of 'x' vertically and horizontally. If number is 5, then the
    # board is a winner
    for l in board:
        for y in range(len(board)):
            count = 0
            for x in range(len(board[0])):
                if board[y][x] == 'x':
                    count += 1
            if count == 5:
                print(f"filas {y}")
                return True

        for x in range(len(board[0])):
            count = 0
            for y in range(len(board)):
                if board[y][x] == 'x':
                    count += 1
            if count == 5:
                print(f"columnas {x}")
                return True
    return False

#with open("mini_input.txt") as f:
with open("input.txt") as f:
    line_count = 0
    boards =[]
    current_board = None
    for l in f:
        if line_count == 0:
            print(l)
            num_sequence = l.strip().split(',')
            line_count += 1
            continue
        if l.strip() == '':
            # Separator
            if current_board:
                boards.append(current_board)
            current_board = []
        else:
            row = []
            c = ""
            l = l.strip()
            print(l)
            for c_pos in range(len(l)):
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
            current_board.append(row)
    boards.append(current_board)

for n in num_sequence:
    print(f'--------------- {n} ----------------')
    for board in boards:
        # pprint(board)
        try:
            for y in range(len(board)):
                for x in range(len(board[0])):
                    # print(f"{x}, {y} -> {board[y][x]}")
                    if board[y][x] == n:
                        board[y][x] = 'x'
        except:
            pprint(board)
            raise Exception("chungo")

    for board in boards:
        if check_if_board_wins(board):
            print("Winner!!!")
            pprint(board)
            print(sum_board(board))
            print(int(n) * sum_board(board))
            break

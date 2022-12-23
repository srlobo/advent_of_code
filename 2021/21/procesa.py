# -*- coding: utf-8 -*-
from pprint import pprint
import logging

def inc_dice(d):
    res = d + 1
    if res > 100:
        res = 1

    return res

def dice():
    c = 0
    while True:
        c = inc_dice(c)
        res = [c]
        c = inc_dice(c)
        res.append(c)
        c = inc_dice(c)
        res.append(c)

        yield res

logger = logging.getLogger(__name__)
logging.basicConfig(level=logging.CRITICAL)
def main():
    players = {}
    #with open("mini_input.txt") as f:
    with open("input.txt") as f:
        for line in f:
            line = line.strip()
            player, position = get_player_and_position(line)
            players[player] = position

    print(players)

    scores = [0, 0, 0]
    roll_number = 0
    current_player = 1
    for dice_roll in dice():
        roll_number += 3
        dice_roll = sum(dice_roll)
        players[current_player] = (players[current_player] + dice_roll) % 10
        if players[current_player] == 0:
            scores[current_player] += 10
        else:
            scores[current_player] += players[current_player]

        if current_player == 1:
            current_player = 2
        else:
            current_player = 1

        print(scores)
        print(players)

        print(f"roll_number: {roll_number}; players: {players}; scores: {scores}")
        if max(scores) > 1000:
            break
    print(roll_number * min(scores[1:]))

def get_player_and_position(line):
    player, position = line.split(': ')
    if '1' in player:
        player = 1
    else:
        player = 2

    position = int(position)

    return player, position

if __name__ == '__main__':
    main()

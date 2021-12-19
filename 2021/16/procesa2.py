# -*- coding: utf-8 -*-
from pprint import pprint
import logging

logger = logging.getLogger(__name__)

translation_dict = {
    '0': '0000',
    '1': '0001',
    '2': '0010',
    '3': '0011',
    '4': '0100',
    '5': '0101',
    '6': '0110',
    '7': '0111',
    '8': '1000',
    '9': '1001',
    'A': '1010',
    'B': '1011',
    'C': '1100',
    'D': '1101',
    'E': '1110',
    'F': '1111',
}

def decode_subpackets(bin_str):
    logger.debug(f"decode_subpackets: {bin_str}")
    length_type_id = bin_str[0]
    if length_type_id == '0': # 15bits -> length subpackets
        length_subpackets = int(bin_str[1:16], 2)
        logger.debug(f"subpackets by length, {length_subpackets}")
        subpackets, rest = decode_packets(bin_str[16:16 + length_subpackets])
        rest += bin_str[16 + length_subpackets:]
    else:
        number_packets = int(bin_str[1:12], 2)
        logger.debug(f"subpackets by number, {number_packets}")
        subpackets, rest = decode_packets(bin_str[12:], number_packets)

    logger.debug(f"subpackets[{bin_str}]: {subpackets}; rest: {rest}")

    return subpackets, rest

def decode_literal_value(bin_str):
    logger.debug(f"decode_literal_value: {bin_str}")
    number_arr = []
    for i in range(len(bin_str) // 5):
        piece = bin_str[i * 5:i * 5 + 5]
        logger.debug(f"piece: {piece}; {piece[0]}")
        if piece[0] == '1': # continuamos
            number_arr.append(piece[1:])
        else:
            number_arr.append(piece[1:])
            break
    rest = bin_str[i * 5 + 5:]

    number_str = ''.join(number_arr)
    number_int = int(number_str, 2)
    logger.debug(f"number_arr: {number_arr}; Number_str: {number_str}; number_int: {number_int}")
    return number_int, rest

def decode_packets(bin_str, number=0):
    """Decodes a series of packets, until number is reached (or there's no more
    packets left)"""
    logger.debug(f"decode_packets: {bin_str}; {number}")
    rest = bin_str
    packets = []

    c = 0
    while len(rest) != 0:
        logger.debug(f"decode_packets: while {bin_str}; limit:{number} c:{c} len(rest):{len(rest)}")
        version, type_, payload, new_rest = decode_packet(rest)
        packets.append((version, type_, payload))
        c += 1
        rest = new_rest
        if '1' not in rest:
            logger.debug('Only 0 on rest -> reset it %s', rest)
            rest = ''
        logger.debug(f"decode_packets: while after {bin_str}; len(rest):{len(rest)}, packet: {(version, type_, payload)}")
        logger.debug(f"Iteracion en decode_packets, {rest}, {packets}")

        if number != 0 and c == number:
            break

    logger.debug(f"decode_packets[{bin_str}]: {packets}; rest: {rest}")
    return packets, rest

def decode_packet(bin_str):
    logger.debug(f"decode_packet: {bin_str}")
    version_str = bin_str[0:3]
    version_int = int(version_str, 2)
    type_str = bin_str[3:6]
    type_int = int(type_str, 2)

    rest = bin_str[6:]
    if type_int == 4: # Literal value
        payload, new_rest = decode_literal_value(rest)
        if '1'  in new_rest:
            rest = new_rest
        else:
            rest = ''
    else:
        payload, new_rest = decode_subpackets(rest)
        rest = new_rest

    return (version_int, type_int, payload, rest)

def line2bin(line):
    ret = []
    for c in line:
        ret.append(translation_dict[c])

    ret_str = ''.join(ret)
    packets, rest = decode_packets(ret_str)
    return packets, rest

def process_line(line):
    packets, rest = line2bin(line)
    return packets, rest

def sum_version(packets):
    sum_versions = 0
    if type(packets) is tuple: # Only one package
        sum_versions += packets[0]
        if type(packets[2]) is list:
            sum_versions += sum_version(packets[2])
    else:
        for p in packets: # array of packets
            sum_versions += sum_version(p)
    return sum_versions

def process_expression(exp):
    packet_version = exp[0]
    packet_type = exp[1]
    packet_payload = exp[2]

    if packet_type == 4:
        res = packet_payload
    elif packet_type == 0: # sum
        res = 0
        for p in packet_payload:
            res += process_expression(p)
    elif packet_type == 1: # Product
        res = 1
        for p in packet_payload:
            res *= process_expression(p)
    elif packet_type == 2: # Minimum
        res_arr = []
        for p in packet_payload:
            res_arr.append(process_expression(p))
        res = min(res_arr)
    elif packet_type == 3: # Maximum
        res_arr = []
        for p in packet_payload:
            res_arr.append(process_expression(p))
        res = max(res_arr)
    elif packet_type == 5: # greater
        if process_expression(packet_payload[0]) > process_expression(packet_payload[1]):
            res = 1
        else:
            res = 0
    elif packet_type == 6: # less than
        if process_expression(packet_payload[0]) < process_expression(packet_payload[1]):
            res = 1
        else:
            res = 0
    elif packet_type == 7: # equal
        if process_expression(packet_payload[1]) == process_expression(packet_payload[0]):
            res = 1
        else:
            res = 0
    else:
        res = exp

    return res

logging.basicConfig(level=logging.CRITICAL)
#with open("micro_input.txt") as f:
#with open("mini_input2.txt") as f:
with open("input.txt") as f:
    for line in f:
        logger.setLevel(logging.CRITICAL)
        line = line.strip()
        logger.debug("Processing %s", line)
        packets, res = process_line(line)
        # s = sum_version(packets)
        # print(f"{line} -> {packets}; len(res): {len(res)}; sum_Version: {s}")
        logger.setLevel(logging.DEBUG)
        res = process_expression(packets[0])
        print(f"{line}:{packets[0]} -> {res}")



#!/usr/bin/env python3.7

import json

with open('opcodes.json') as f:
    opcodes_json = json.load(f)

for row in opcodes_json:
    row_prefix = ''
    for td in row['children']:
        if td['tag'] == 'th':
            continue
        first_child = td['children'][0]
        if first_child['tag'] == 'strong':
            row_prefix = first_child['text'][0]
            index = 0

        else:
            opcode = '0x' + row_prefix + hex(index)[-1].upper()
            disassembly = first_child['text']
            description = first_child['title']
            print('{' + opcode + ', "' + disassembly + '", "' + description + '"},')
            index += 1

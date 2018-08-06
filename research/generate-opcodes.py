#!/usr/bin/env python3.7

import json
import re

def printOpcodeMap(name, opcodes):
    print()
    print(f'var {name} = map[uint8]opcode{{')
    for o, oid in opcodes.items():
        print(f'    {o}: {oid},')
    print('}')

with open('opcodes.json') as f:
    opcodes_json = json.load(f)

prefix = 'Opcode'
print('var (')
opcodeToId = {}
for row in opcodes_json:
    row_prefix = ''
    for td in row['children']:
        if td['tag'] == 'th':
            if prefix !=  'OpcodeExt':
                prefix = 'OpcodeExt'
                print(')')
                print()
                printOpcodeMap('opcodeMap', opcodeToId)
                print()
                print('var (')
            continue
        first_child = td['children'][0]
        if first_child['tag'] == 'strong':
            row_prefix = first_child['text'][0]
            index = 0

        else:
            opcode = '0x' + row_prefix + hex(index)[-1].upper()
            disassembly = first_child['text']
            opcodeId = prefix + ''.join([re.sub(r'\W', '', c).lower().capitalize() for c in disassembly.split(' ')])
            opcodeToId[opcode] = opcodeId
            description = first_child['title']
            print(f'    {opcodeId} = opcode{{{opcode}, "{disassembly}", "{description}", 0, 1}}')
            index += 1
print(')')

printOpcodeMap('opcodeMapExt', opcodeToId)

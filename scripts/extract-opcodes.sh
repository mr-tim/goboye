#!/bin/bash

curl -sv http://imrannazar.com/Gameboy-Z80-Opcode-Map | ~/go/bin/pup --color 'tr json{}'
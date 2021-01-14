package cpu

var (
	OpcodeExtRlcB   = opcode{0x00, "RLC B", "Rotate B left with carry", 0, 8, rotateRegLeftWithCarry(RegisterB)}
	OpcodeExtRlcC   = opcode{0x01, "RLC C", "Rotate C left with carry", 0, 8, rotateRegLeftWithCarry(RegisterC)}
	OpcodeExtRlcD   = opcode{0x02, "RLC D", "Rotate D left with carry", 0, 8, rotateRegLeftWithCarry(RegisterD)}
	OpcodeExtRlcE   = opcode{0x03, "RLC E", "Rotate E left with carry", 0, 8, rotateRegLeftWithCarry(RegisterE)}
	OpcodeExtRlcH   = opcode{0x04, "RLC H", "Rotate H left with carry", 0, 8, rotateRegLeftWithCarry(RegisterH)}
	OpcodeExtRlcL   = opcode{0x05, "RLC L", "Rotate L left with carry", 0, 8, rotateRegLeftWithCarry(RegisterL)}
	OpcodeExtRlcHl  = opcode{0x06, "RLC (HL)", "Rotate value pointed by HL left with carry", 0, 16, rotateHLAddrLeftWithCarry}
	OpcodeExtRlcA   = opcode{0x07, "RLC A", "Rotate A left with carry", 0, 8, rotateRegLeftWithCarry(RegisterA)}
	OpcodeExtRrcB   = opcode{0x08, "RRC B", "Rotate B right with carry", 0, 8, rotateRegRightWithCarry(RegisterB)}
	OpcodeExtRrcC   = opcode{0x09, "RRC C", "Rotate C right with carry", 0, 8, rotateRegRightWithCarry(RegisterC)}
	OpcodeExtRrcD   = opcode{0x0A, "RRC D", "Rotate D right with carry", 0, 8, rotateRegRightWithCarry(RegisterD)}
	OpcodeExtRrcE   = opcode{0x0B, "RRC E", "Rotate E right with carry", 0, 8, rotateRegRightWithCarry(RegisterE)}
	OpcodeExtRrcH   = opcode{0x0C, "RRC H", "Rotate H right with carry", 0, 8, rotateRegRightWithCarry(RegisterH)}
	OpcodeExtRrcL   = opcode{0x0D, "RRC L", "Rotate L right with carry", 0, 8, rotateRegRightWithCarry(RegisterL)}
	OpcodeExtRrcHl  = opcode{0x0E, "RRC (HL)", "Rotate value pointed by HL right with carry", 0, 16, rotateHLAddrRightWithCarry}
	OpcodeExtRrcA   = opcode{0x0F, "RRC A", "Rotate A right with carry", 0, 8, rotateRegRightWithCarry(RegisterA)}
	OpcodeExtRlB    = opcode{0x10, "RL B", "Rotate B left", 0, 8, rotateRegLeft(RegisterB)}
	OpcodeExtRlC    = opcode{0x11, "RL C", "Rotate C left", 0, 8, rotateRegLeft(RegisterC)}
	OpcodeExtRlD    = opcode{0x12, "RL D", "Rotate D left", 0, 8, rotateRegLeft(RegisterD)}
	OpcodeExtRlE    = opcode{0x13, "RL E", "Rotate E left", 0, 8, rotateRegLeft(RegisterE)}
	OpcodeExtRlH    = opcode{0x14, "RL H", "Rotate H left", 0, 8, rotateRegLeft(RegisterH)}
	OpcodeExtRlL    = opcode{0x15, "RL L", "Rotate L left", 0, 8, rotateRegLeft(RegisterL)}
	OpcodeExtRlHl   = opcode{0x16, "RL (HL)", "Rotate value pointed by HL left", 0, 16, rotateHLAddrLeft}
	OpcodeExtRlA    = opcode{0x17, "RL A", "Rotate A left", 0, 8, rotateRegLeft(RegisterA)}
	OpcodeExtRrB    = opcode{0x18, "RR B", "Rotate B right", 0, 8, rotateRegRight(RegisterB)}
	OpcodeExtRrC    = opcode{0x19, "RR C", "Rotate C right", 0, 8, rotateRegRight(RegisterC)}
	OpcodeExtRrD    = opcode{0x1A, "RR D", "Rotate D right", 0, 8, rotateRegRight(RegisterD)}
	OpcodeExtRrE    = opcode{0x1B, "RR E", "Rotate E right", 0, 8, rotateRegRight(RegisterE)}
	OpcodeExtRrH    = opcode{0x1C, "RR H", "Rotate H right", 0, 8, rotateRegRight(RegisterH)}
	OpcodeExtRrL    = opcode{0x1D, "RR L", "Rotate L right", 0, 8, rotateRegRight(RegisterL)}
	OpcodeExtRrHl   = opcode{0x1E, "RR (HL)", "Rotate value pointed by HL right", 0, 16, rotateHLAddrRight}
	OpcodeExtRrA    = opcode{0x1F, "RR A", "Rotate A right", 0, 8, rotateRegRight(RegisterA)}
	OpcodeExtSlaB   = opcode{0x20, "SLA B", "Shift B left preserving sign", 0, 8, shiftRegLeftPreservingSign(RegisterB)}
	OpcodeExtSlaC   = opcode{0x21, "SLA C", "Shift C left preserving sign", 0, 8, shiftRegLeftPreservingSign(RegisterC)}
	OpcodeExtSlaD   = opcode{0x22, "SLA D", "Shift D left preserving sign", 0, 8, shiftRegLeftPreservingSign(RegisterD)}
	OpcodeExtSlaE   = opcode{0x23, "SLA E", "Shift E left preserving sign", 0, 8, shiftRegLeftPreservingSign(RegisterE)}
	OpcodeExtSlaH   = opcode{0x24, "SLA H", "Shift H left preserving sign", 0, 8, shiftRegLeftPreservingSign(RegisterH)}
	OpcodeExtSlaL   = opcode{0x25, "SLA L", "Shift L left preserving sign", 0, 8, shiftRegLeftPreservingSign(RegisterL)}
	OpcodeExtSlaHl  = opcode{0x26, "SLA (HL)", "Shift value pointed by HL left preserving sign", 0, 16, shiftHLAddrLeftPreservingSign}
	OpcodeExtSlaA   = opcode{0x27, "SLA A", "Shift A left preserving sign", 0, 8, shiftRegLeftPreservingSign(RegisterA)}
	OpcodeExtSraB   = opcode{0x28, "SRA B", "Shift B right preserving sign", 0, 8, shiftRegRightPreservingSign(RegisterB)}
	OpcodeExtSraC   = opcode{0x29, "SRA C", "Shift C right preserving sign", 0, 8, shiftRegRightPreservingSign(RegisterC)}
	OpcodeExtSraD   = opcode{0x2A, "SRA D", "Shift D right preserving sign", 0, 8, shiftRegRightPreservingSign(RegisterD)}
	OpcodeExtSraE   = opcode{0x2B, "SRA E", "Shift E right preserving sign", 0, 8, shiftRegRightPreservingSign(RegisterE)}
	OpcodeExtSraH   = opcode{0x2C, "SRA H", "Shift H right preserving sign", 0, 8, shiftRegRightPreservingSign(RegisterH)}
	OpcodeExtSraL   = opcode{0x2D, "SRA L", "Shift L right preserving sign", 0, 8, shiftRegRightPreservingSign(RegisterL)}
	OpcodeExtSraHl  = opcode{0x2E, "SRA (HL)", "Shift value pointed by HL right preserving sign", 0, 16, shiftHLAddrRightPreservingSign}
	OpcodeExtSraA   = opcode{0x2F, "SRA A", "Shift A right preserving sign", 0, 8, shiftRegRightPreservingSign(RegisterA)}
	OpcodeExtSwapB  = opcode{0x30, "SWAP B", "Swap nybbles in B", 0, 8, swapRegNybbles(RegisterB)}
	OpcodeExtSwapC  = opcode{0x31, "SWAP C", "Swap nybbles in C", 0, 8, swapRegNybbles(RegisterC)}
	OpcodeExtSwapD  = opcode{0x32, "SWAP D", "Swap nybbles in D", 0, 8, swapRegNybbles(RegisterD)}
	OpcodeExtSwapE  = opcode{0x33, "SWAP E", "Swap nybbles in E", 0, 8, swapRegNybbles(RegisterE)}
	OpcodeExtSwapH  = opcode{0x34, "SWAP H", "Swap nybbles in H", 0, 8, swapRegNybbles(RegisterH)}
	OpcodeExtSwapL  = opcode{0x35, "SWAP L", "Swap nybbles in L", 0, 8, swapRegNybbles(RegisterL)}
	OpcodeExtSwapHl = opcode{0x36, "SWAP (HL)", "Swap nybbles in value pointed by HL", 0, 16, swapHLAddrNybbles}
	OpcodeExtSwapA  = opcode{0x37, "SWAP A", "Swap nybbles in A", 0, 8, swapRegNybbles(RegisterA)}
	OpcodeExtSrlB   = opcode{0x38, "SRL B", "Shift B right", 0, 8, shiftRegRight(RegisterB)}
	OpcodeExtSrlC   = opcode{0x39, "SRL C", "Shift C right", 0, 8, shiftRegRight(RegisterC)}
	OpcodeExtSrlD   = opcode{0x3A, "SRL D", "Shift D right", 0, 8, shiftRegRight(RegisterD)}
	OpcodeExtSrlE   = opcode{0x3B, "SRL E", "Shift E right", 0, 8, shiftRegRight(RegisterE)}
	OpcodeExtSrlH   = opcode{0x3C, "SRL H", "Shift H right", 0, 8, shiftRegRight(RegisterH)}
	OpcodeExtSrlL   = opcode{0x3D, "SRL L", "Shift L right", 0, 8, shiftRegRight(RegisterL)}
	OpcodeExtSrlHl  = opcode{0x3E, "SRL (HL)", "Shift value pointed by HL right", 0, 16, shiftHLAddrRight}
	OpcodeExtSrlA   = opcode{0x3F, "SRL A", "Shift A right", 0, 8, shiftRegRight(RegisterA)}
	OpcodeExtBit0b  = opcode{0x40, "BIT 0,B", "Test bit 0 of B", 0, 8, testBitOfReg(0, RegisterB)}
	OpcodeExtBit0c  = opcode{0x41, "BIT 0,C", "Test bit 0 of C", 0, 8, testBitOfReg(0, RegisterC)}
	OpcodeExtBit0d  = opcode{0x42, "BIT 0,D", "Test bit 0 of D", 0, 8, testBitOfReg(0, RegisterD)}
	OpcodeExtBit0e  = opcode{0x43, "BIT 0,E", "Test bit 0 of E", 0, 8, testBitOfReg(0, RegisterE)}
	OpcodeExtBit0h  = opcode{0x44, "BIT 0,H", "Test bit 0 of H", 0, 8, testBitOfReg(0, RegisterH)}
	OpcodeExtBit0l  = opcode{0x45, "BIT 0,L", "Test bit 0 of L", 0, 8, testBitOfReg(0, RegisterL)}
	OpcodeExtBit0hl = opcode{0x46, "BIT 0,(HL)", "Test bit 0 of value pointed by HL", 0, 16, testBitOfHLAddr(0)}
	OpcodeExtBit0a  = opcode{0x47, "BIT 0,A", "Test bit 0 of A", 0, 8, testBitOfReg(0, RegisterA)}
	OpcodeExtBit1b  = opcode{0x48, "BIT 1,B", "Test bit 1 of B", 0, 8, testBitOfReg(1, RegisterB)}
	OpcodeExtBit1c  = opcode{0x49, "BIT 1,C", "Test bit 1 of C", 0, 8, testBitOfReg(1, RegisterC)}
	OpcodeExtBit1d  = opcode{0x4A, "BIT 1,D", "Test bit 1 of D", 0, 8, testBitOfReg(1, RegisterD)}
	OpcodeExtBit1e  = opcode{0x4B, "BIT 1,E", "Test bit 1 of E", 0, 8, testBitOfReg(1, RegisterE)}
	OpcodeExtBit1h  = opcode{0x4C, "BIT 1,H", "Test bit 1 of H", 0, 8, testBitOfReg(1, RegisterH)}
	OpcodeExtBit1l  = opcode{0x4D, "BIT 1,L", "Test bit 1 of L", 0, 8, testBitOfReg(1, RegisterL)}
	OpcodeExtBit1hl = opcode{0x4E, "BIT 1,(HL)", "Test bit 1 of value pointed by HL", 0, 16, testBitOfHLAddr(1)}
	OpcodeExtBit1a  = opcode{0x4F, "BIT 1,A", "Test bit 1 of A", 0, 8, testBitOfReg(1, RegisterA)}
	OpcodeExtBit2b  = opcode{0x50, "BIT 2,B", "Test bit 2 of B", 0, 8, testBitOfReg(2, RegisterB)}
	OpcodeExtBit2c  = opcode{0x51, "BIT 2,C", "Test bit 2 of C", 0, 8, testBitOfReg(2, RegisterC)}
	OpcodeExtBit2d  = opcode{0x52, "BIT 2,D", "Test bit 2 of D", 0, 8, testBitOfReg(2, RegisterD)}
	OpcodeExtBit2e  = opcode{0x53, "BIT 2,E", "Test bit 2 of E", 0, 8, testBitOfReg(2, RegisterE)}
	OpcodeExtBit2h  = opcode{0x54, "BIT 2,H", "Test bit 2 of H", 0, 8, testBitOfReg(2, RegisterH)}
	OpcodeExtBit2l  = opcode{0x55, "BIT 2,L", "Test bit 2 of L", 0, 8, testBitOfReg(2, RegisterL)}
	OpcodeExtBit2hl = opcode{0x56, "BIT 2,(HL)", "Test bit 2 of value pointed by HL", 0, 16, testBitOfHLAddr(2)}
	OpcodeExtBit2a  = opcode{0x57, "BIT 2,A", "Test bit 2 of A", 0, 8, testBitOfReg(2, RegisterA)}
	OpcodeExtBit3b  = opcode{0x58, "BIT 3,B", "Test bit 3 of B", 0, 8, testBitOfReg(3, RegisterB)}
	OpcodeExtBit3c  = opcode{0x59, "BIT 3,C", "Test bit 3 of C", 0, 8, testBitOfReg(3, RegisterC)}
	OpcodeExtBit3d  = opcode{0x5A, "BIT 3,D", "Test bit 3 of D", 0, 8, testBitOfReg(3, RegisterD)}
	OpcodeExtBit3e  = opcode{0x5B, "BIT 3,E", "Test bit 3 of E", 0, 8, testBitOfReg(3, RegisterE)}
	OpcodeExtBit3h  = opcode{0x5C, "BIT 3,H", "Test bit 3 of H", 0, 8, testBitOfReg(3, RegisterH)}
	OpcodeExtBit3l  = opcode{0x5D, "BIT 3,L", "Test bit 3 of L", 0, 8, testBitOfReg(3, RegisterL)}
	OpcodeExtBit3hl = opcode{0x5E, "BIT 3,(HL)", "Test bit 3 of value pointed by HL", 0, 16, testBitOfHLAddr(3)}
	OpcodeExtBit3a  = opcode{0x5F, "BIT 3,A", "Test bit 3 of A", 0, 8, testBitOfReg(3, RegisterA)}
	OpcodeExtBit4b  = opcode{0x60, "BIT 4,B", "Test bit 4 of B", 0, 8, testBitOfReg(4, RegisterB)}
	OpcodeExtBit4c  = opcode{0x61, "BIT 4,C", "Test bit 4 of C", 0, 8, testBitOfReg(4, RegisterC)}
	OpcodeExtBit4d  = opcode{0x62, "BIT 4,D", "Test bit 4 of D", 0, 8, testBitOfReg(4, RegisterD)}
	OpcodeExtBit4e  = opcode{0x63, "BIT 4,E", "Test bit 4 of E", 0, 8, testBitOfReg(4, RegisterE)}
	OpcodeExtBit4h  = opcode{0x64, "BIT 4,H", "Test bit 4 of H", 0, 8, testBitOfReg(4, RegisterH)}
	OpcodeExtBit4l  = opcode{0x65, "BIT 4,L", "Test bit 4 of L", 0, 8, testBitOfReg(4, RegisterL)}
	OpcodeExtBit4hl = opcode{0x66, "BIT 4,(HL)", "Test bit 4 of value pointed by HL", 0, 16, testBitOfHLAddr(4)}
	OpcodeExtBit4a  = opcode{0x67, "BIT 4,A", "Test bit 4 of A", 0, 8, testBitOfReg(4, RegisterA)}
	OpcodeExtBit5b  = opcode{0x68, "BIT 5,B", "Test bit 5 of B", 0, 8, testBitOfReg(5, RegisterB)}
	OpcodeExtBit5c  = opcode{0x69, "BIT 5,C", "Test bit 5 of C", 0, 8, testBitOfReg(5, RegisterC)}
	OpcodeExtBit5d  = opcode{0x6A, "BIT 5,D", "Test bit 5 of D", 0, 8, testBitOfReg(5, RegisterD)}
	OpcodeExtBit5e  = opcode{0x6B, "BIT 5,E", "Test bit 5 of E", 0, 8, testBitOfReg(5, RegisterE)}
	OpcodeExtBit5h  = opcode{0x6C, "BIT 5,H", "Test bit 5 of H", 0, 8, testBitOfReg(5, RegisterH)}
	OpcodeExtBit5l  = opcode{0x6D, "BIT 5,L", "Test bit 5 of L", 0, 8, testBitOfReg(5, RegisterL)}
	OpcodeExtBit5hl = opcode{0x6E, "BIT 5,(HL)", "Test bit 5 of value pointed by HL", 0, 16, testBitOfHLAddr(5)}
	OpcodeExtBit5a  = opcode{0x6F, "BIT 5,A", "Test bit 5 of A", 0, 8, testBitOfReg(5, RegisterA)}
	OpcodeExtBit6b  = opcode{0x70, "BIT 6,B", "Test bit 6 of B", 0, 8, testBitOfReg(6, RegisterB)}
	OpcodeExtBit6c  = opcode{0x71, "BIT 6,C", "Test bit 6 of C", 0, 8, testBitOfReg(6, RegisterC)}
	OpcodeExtBit6d  = opcode{0x72, "BIT 6,D", "Test bit 6 of D", 0, 8, testBitOfReg(6, RegisterD)}
	OpcodeExtBit6e  = opcode{0x73, "BIT 6,E", "Test bit 6 of E", 0, 8, testBitOfReg(6, RegisterE)}
	OpcodeExtBit6h  = opcode{0x74, "BIT 6,H", "Test bit 6 of H", 0, 8, testBitOfReg(6, RegisterH)}
	OpcodeExtBit6l  = opcode{0x75, "BIT 6,L", "Test bit 6 of L", 0, 8, testBitOfReg(6, RegisterL)}
	OpcodeExtBit6hl = opcode{0x76, "BIT 6,(HL)", "Test bit 6 of value pointed by HL", 0, 16, testBitOfHLAddr(6)}
	OpcodeExtBit6a  = opcode{0x77, "BIT 6,A", "Test bit 6 of A", 0, 8, testBitOfReg(6, RegisterA)}
	OpcodeExtBit7b  = opcode{0x78, "BIT 7,B", "Test bit 7 of B", 0, 8, testBitOfReg(7, RegisterB)}
	OpcodeExtBit7c  = opcode{0x79, "BIT 7,C", "Test bit 7 of C", 0, 8, testBitOfReg(7, RegisterC)}
	OpcodeExtBit7d  = opcode{0x7A, "BIT 7,D", "Test bit 7 of D", 0, 8, testBitOfReg(7, RegisterD)}
	OpcodeExtBit7e  = opcode{0x7B, "BIT 7,E", "Test bit 7 of E", 0, 8, testBitOfReg(7, RegisterE)}
	OpcodeExtBit7h  = opcode{0x7C, "BIT 7,H", "Test bit 7 of H", 0, 8, testBitOfReg(7, RegisterH)}
	OpcodeExtBit7l  = opcode{0x7D, "BIT 7,L", "Test bit 7 of L", 0, 8, testBitOfReg(7, RegisterL)}
	OpcodeExtBit7hl = opcode{0x7E, "BIT 7,(HL)", "Test bit 7 of value pointed by HL", 0, 16, testBitOfHLAddr(7)}
	OpcodeExtBit7a  = opcode{0x7F, "BIT 7,A", "Test bit 7 of A", 0, 8, testBitOfReg(7, RegisterA)}
	OpcodeExtRes0b  = opcode{0x80, "RES 0,B", "Clear (reset) bit 0 of B", 0, 8, clearBitOfReg(0, RegisterB)}
	OpcodeExtRes0c  = opcode{0x81, "RES 0,C", "Clear (reset) bit 0 of C", 0, 8, clearBitOfReg(0, RegisterC)}
	OpcodeExtRes0d  = opcode{0x82, "RES 0,D", "Clear (reset) bit 0 of D", 0, 8, clearBitOfReg(0, RegisterD)}
	OpcodeExtRes0e  = opcode{0x83, "RES 0,E", "Clear (reset) bit 0 of E", 0, 8, clearBitOfReg(0, RegisterE)}
	OpcodeExtRes0h  = opcode{0x84, "RES 0,H", "Clear (reset) bit 0 of H", 0, 8, clearBitOfReg(0, RegisterH)}
	OpcodeExtRes0l  = opcode{0x85, "RES 0,L", "Clear (reset) bit 0 of L", 0, 8, clearBitOfReg(0, RegisterL)}
	OpcodeExtRes0hl = opcode{0x86, "RES 0,(HL)", "Clear (reset) bit 0 of value pointed by HL", 0, 16, clearBitOfHLAddr(0)}
	OpcodeExtRes0a  = opcode{0x87, "RES 0,A", "Clear (reset) bit 0 of A", 0, 8, clearBitOfReg(0, RegisterA)}
	OpcodeExtRes1b  = opcode{0x88, "RES 1,B", "Clear (reset) bit 1 of B", 0, 8, clearBitOfReg(1, RegisterB)}
	OpcodeExtRes1c  = opcode{0x89, "RES 1,C", "Clear (reset) bit 1 of C", 0, 8, clearBitOfReg(1, RegisterC)}
	OpcodeExtRes1d  = opcode{0x8A, "RES 1,D", "Clear (reset) bit 1 of D", 0, 8, clearBitOfReg(1, RegisterD)}
	OpcodeExtRes1e  = opcode{0x8B, "RES 1,E", "Clear (reset) bit 1 of E", 0, 8, clearBitOfReg(1, RegisterE)}
	OpcodeExtRes1h  = opcode{0x8C, "RES 1,H", "Clear (reset) bit 1 of H", 0, 8, clearBitOfReg(1, RegisterH)}
	OpcodeExtRes1l  = opcode{0x8D, "RES 1,L", "Clear (reset) bit 1 of L", 0, 8, clearBitOfReg(1, RegisterL)}
	OpcodeExtRes1hl = opcode{0x8E, "RES 1,(HL)", "Clear (reset) bit 1 of value pointed by HL", 0, 16, clearBitOfHLAddr(1)}
	OpcodeExtRes1a  = opcode{0x8F, "RES 1,A", "Clear (reset) bit 1 of A", 0, 8, clearBitOfReg(1, RegisterA)}
	OpcodeExtRes2b  = opcode{0x90, "RES 2,B", "Clear (reset) bit 2 of B", 0, 8, clearBitOfReg(2, RegisterB)}
	OpcodeExtRes2c  = opcode{0x91, "RES 2,C", "Clear (reset) bit 2 of C", 0, 8, clearBitOfReg(2, RegisterC)}
	OpcodeExtRes2d  = opcode{0x92, "RES 2,D", "Clear (reset) bit 2 of D", 0, 8, clearBitOfReg(2, RegisterD)}
	OpcodeExtRes2e  = opcode{0x93, "RES 2,E", "Clear (reset) bit 2 of E", 0, 8, clearBitOfReg(2, RegisterE)}
	OpcodeExtRes2h  = opcode{0x94, "RES 2,H", "Clear (reset) bit 2 of H", 0, 8, clearBitOfReg(2, RegisterH)}
	OpcodeExtRes2l  = opcode{0x95, "RES 2,L", "Clear (reset) bit 2 of L", 0, 8, clearBitOfReg(2, RegisterL)}
	OpcodeExtRes2hl = opcode{0x96, "RES 2,(HL)", "Clear (reset) bit 2 of value pointed by HL", 0, 16, clearBitOfHLAddr(2)}
	OpcodeExtRes2a  = opcode{0x97, "RES 2,A", "Clear (reset) bit 2 of A", 0, 8, clearBitOfReg(2, RegisterA)}
	OpcodeExtRes3b  = opcode{0x98, "RES 3,B", "Clear (reset) bit 3 of B", 0, 8, clearBitOfReg(3, RegisterB)}
	OpcodeExtRes3c  = opcode{0x99, "RES 3,C", "Clear (reset) bit 3 of C", 0, 8, clearBitOfReg(3, RegisterC)}
	OpcodeExtRes3d  = opcode{0x9A, "RES 3,D", "Clear (reset) bit 3 of D", 0, 8, clearBitOfReg(3, RegisterD)}
	OpcodeExtRes3e  = opcode{0x9B, "RES 3,E", "Clear (reset) bit 3 of E", 0, 8, clearBitOfReg(3, RegisterE)}
	OpcodeExtRes3h  = opcode{0x9C, "RES 3,H", "Clear (reset) bit 3 of H", 0, 8, clearBitOfReg(3, RegisterH)}
	OpcodeExtRes3l  = opcode{0x9D, "RES 3,L", "Clear (reset) bit 3 of L", 0, 8, clearBitOfReg(3, RegisterL)}
	OpcodeExtRes3hl = opcode{0x9E, "RES 3,(HL)", "Clear (reset) bit 3 of value pointed by HL", 0, 16, clearBitOfHLAddr(3)}
	OpcodeExtRes3a  = opcode{0x9F, "RES 3,A", "Clear (reset) bit 3 of A", 0, 8, clearBitOfReg(3, RegisterA)}
	OpcodeExtRes4b  = opcode{0xA0, "RES 4,B", "Clear (reset) bit 4 of B", 0, 8, clearBitOfReg(4, RegisterB)}
	OpcodeExtRes4c  = opcode{0xA1, "RES 4,C", "Clear (reset) bit 4 of C", 0, 8, clearBitOfReg(4, RegisterC)}
	OpcodeExtRes4d  = opcode{0xA2, "RES 4,D", "Clear (reset) bit 4 of D", 0, 8, clearBitOfReg(4, RegisterD)}
	OpcodeExtRes4e  = opcode{0xA3, "RES 4,E", "Clear (reset) bit 4 of E", 0, 8, clearBitOfReg(4, RegisterE)}
	OpcodeExtRes4h  = opcode{0xA4, "RES 4,H", "Clear (reset) bit 4 of H", 0, 8, clearBitOfReg(4, RegisterH)}
	OpcodeExtRes4l  = opcode{0xA5, "RES 4,L", "Clear (reset) bit 4 of L", 0, 8, clearBitOfReg(4, RegisterL)}
	OpcodeExtRes4hl = opcode{0xA6, "RES 4,(HL)", "Clear (reset) bit 4 of value pointed by HL", 0, 16, clearBitOfHLAddr(4)}
	OpcodeExtRes4a  = opcode{0xA7, "RES 4,A", "Clear (reset) bit 4 of A", 0, 8, clearBitOfReg(4, RegisterA)}
	OpcodeExtRes5b  = opcode{0xA8, "RES 5,B", "Clear (reset) bit 5 of B", 0, 8, clearBitOfReg(5, RegisterB)}
	OpcodeExtRes5c  = opcode{0xA9, "RES 5,C", "Clear (reset) bit 5 of C", 0, 8, clearBitOfReg(5, RegisterC)}
	OpcodeExtRes5d  = opcode{0xAA, "RES 5,D", "Clear (reset) bit 5 of D", 0, 8, clearBitOfReg(5, RegisterD)}
	OpcodeExtRes5e  = opcode{0xAB, "RES 5,E", "Clear (reset) bit 5 of E", 0, 8, clearBitOfReg(5, RegisterE)}
	OpcodeExtRes5h  = opcode{0xAC, "RES 5,H", "Clear (reset) bit 5 of H", 0, 8, clearBitOfReg(5, RegisterH)}
	OpcodeExtRes5l  = opcode{0xAD, "RES 5,L", "Clear (reset) bit 5 of L", 0, 8, clearBitOfReg(5, RegisterL)}
	OpcodeExtRes5hl = opcode{0xAE, "RES 5,(HL)", "Clear (reset) bit 5 of value pointed by HL", 0, 16, clearBitOfHLAddr(5)}
	OpcodeExtRes5a  = opcode{0xAF, "RES 5,A", "Clear (reset) bit 5 of A", 0, 8, clearBitOfReg(5, RegisterA)}
	OpcodeExtRes6b  = opcode{0xB0, "RES 6,B", "Clear (reset) bit 6 of B", 0, 8, clearBitOfReg(6, RegisterB)}
	OpcodeExtRes6c  = opcode{0xB1, "RES 6,C", "Clear (reset) bit 6 of C", 0, 8, clearBitOfReg(6, RegisterC)}
	OpcodeExtRes6d  = opcode{0xB2, "RES 6,D", "Clear (reset) bit 6 of D", 0, 8, clearBitOfReg(6, RegisterD)}
	OpcodeExtRes6e  = opcode{0xB3, "RES 6,E", "Clear (reset) bit 6 of E", 0, 8, clearBitOfReg(6, RegisterE)}
	OpcodeExtRes6h  = opcode{0xB4, "RES 6,H", "Clear (reset) bit 6 of H", 0, 8, clearBitOfReg(6, RegisterH)}
	OpcodeExtRes6l  = opcode{0xB5, "RES 6,L", "Clear (reset) bit 6 of L", 0, 8, clearBitOfReg(6, RegisterL)}
	OpcodeExtRes6hl = opcode{0xB6, "RES 6,(HL)", "Clear (reset) bit 6 of value pointed by HL", 0, 16, clearBitOfHLAddr(6)}
	OpcodeExtRes6a  = opcode{0xB7, "RES 6,A", "Clear (reset) bit 6 of A", 0, 8, clearBitOfReg(6, RegisterA)}
	OpcodeExtRes7b  = opcode{0xB8, "RES 7,B", "Clear (reset) bit 7 of B", 0, 8, clearBitOfReg(7, RegisterB)}
	OpcodeExtRes7c  = opcode{0xB9, "RES 7,C", "Clear (reset) bit 7 of C", 0, 8, clearBitOfReg(7, RegisterC)}
	OpcodeExtRes7d  = opcode{0xBA, "RES 7,D", "Clear (reset) bit 7 of D", 0, 8, clearBitOfReg(7, RegisterD)}
	OpcodeExtRes7e  = opcode{0xBB, "RES 7,E", "Clear (reset) bit 7 of E", 0, 8, clearBitOfReg(7, RegisterE)}
	OpcodeExtRes7h  = opcode{0xBC, "RES 7,H", "Clear (reset) bit 7 of H", 0, 8, clearBitOfReg(7, RegisterH)}
	OpcodeExtRes7l  = opcode{0xBD, "RES 7,L", "Clear (reset) bit 7 of L", 0, 8, clearBitOfReg(7, RegisterL)}
	OpcodeExtRes7hl = opcode{0xBE, "RES 7,(HL)", "Clear (reset) bit 7 of value pointed by HL", 0, 16, clearBitOfHLAddr(7)}
	OpcodeExtRes7a  = opcode{0xBF, "RES 7,A", "Clear (reset) bit 7 of A", 0, 8, clearBitOfReg(7, RegisterA)}
	OpcodeExtSet0b  = opcode{0xC0, "SET 0,B", "Set bit 0 of B", 0, 8, setBitOfReg(0, RegisterB)}
	OpcodeExtSet0c  = opcode{0xC1, "SET 0,C", "Set bit 0 of C", 0, 8, setBitOfReg(0, RegisterC)}
	OpcodeExtSet0d  = opcode{0xC2, "SET 0,D", "Set bit 0 of D", 0, 8, setBitOfReg(0, RegisterD)}
	OpcodeExtSet0e  = opcode{0xC3, "SET 0,E", "Set bit 0 of E", 0, 8, setBitOfReg(0, RegisterE)}
	OpcodeExtSet0h  = opcode{0xC4, "SET 0,H", "Set bit 0 of H", 0, 8, setBitOfReg(0, RegisterH)}
	OpcodeExtSet0l  = opcode{0xC5, "SET 0,L", "Set bit 0 of L", 0, 8, setBitOfReg(0, RegisterL)}
	OpcodeExtSet0hl = opcode{0xC6, "SET 0,(HL)", "Set bit 0 of value pointed by HL", 0, 16, setBitOfHLAddr(0)}
	OpcodeExtSet0a  = opcode{0xC7, "SET 0,A", "Set bit 0 of A", 0, 8, setBitOfReg(0, RegisterA)}
	OpcodeExtSet1b  = opcode{0xC8, "SET 1,B", "Set bit 1 of B", 0, 8, setBitOfReg(1, RegisterB)}
	OpcodeExtSet1c  = opcode{0xC9, "SET 1,C", "Set bit 1 of C", 0, 8, setBitOfReg(1, RegisterC)}
	OpcodeExtSet1d  = opcode{0xCA, "SET 1,D", "Set bit 1 of D", 0, 8, setBitOfReg(1, RegisterD)}
	OpcodeExtSet1e  = opcode{0xCB, "SET 1,E", "Set bit 1 of E", 0, 8, setBitOfReg(1, RegisterE)}
	OpcodeExtSet1h  = opcode{0xCC, "SET 1,H", "Set bit 1 of H", 0, 8, setBitOfReg(1, RegisterH)}
	OpcodeExtSet1l  = opcode{0xCD, "SET 1,L", "Set bit 1 of L", 0, 8, setBitOfReg(1, RegisterL)}
	OpcodeExtSet1hl = opcode{0xCE, "SET 1,(HL)", "Set bit 1 of value pointed by HL", 0, 16, setBitOfHLAddr(1)}
	OpcodeExtSet1a  = opcode{0xCF, "SET 1,A", "Set bit 1 of A", 0, 8, setBitOfReg(1, RegisterA)}
	OpcodeExtSet2b  = opcode{0xD0, "SET 2,B", "Set bit 2 of B", 0, 8, setBitOfReg(2, RegisterB)}
	OpcodeExtSet2c  = opcode{0xD1, "SET 2,C", "Set bit 2 of C", 0, 8, setBitOfReg(2, RegisterC)}
	OpcodeExtSet2d  = opcode{0xD2, "SET 2,D", "Set bit 2 of D", 0, 8, setBitOfReg(2, RegisterD)}
	OpcodeExtSet2e  = opcode{0xD3, "SET 2,E", "Set bit 2 of E", 0, 8, setBitOfReg(2, RegisterE)}
	OpcodeExtSet2h  = opcode{0xD4, "SET 2,H", "Set bit 2 of H", 0, 8, setBitOfReg(2, RegisterH)}
	OpcodeExtSet2l  = opcode{0xD5, "SET 2,L", "Set bit 2 of L", 0, 8, setBitOfReg(2, RegisterL)}
	OpcodeExtSet2hl = opcode{0xD6, "SET 2,(HL)", "Set bit 2 of value pointed by HL", 0, 16, setBitOfHLAddr(2)}
	OpcodeExtSet2a  = opcode{0xD7, "SET 2,A", "Set bit 2 of A", 0, 8, setBitOfReg(2, RegisterA)}
	OpcodeExtSet3b  = opcode{0xD8, "SET 3,B", "Set bit 3 of B", 0, 8, setBitOfReg(3, RegisterB)}
	OpcodeExtSet3c  = opcode{0xD9, "SET 3,C", "Set bit 3 of C", 0, 8, setBitOfReg(3, RegisterC)}
	OpcodeExtSet3d  = opcode{0xDA, "SET 3,D", "Set bit 3 of D", 0, 8, setBitOfReg(3, RegisterD)}
	OpcodeExtSet3e  = opcode{0xDB, "SET 3,E", "Set bit 3 of E", 0, 8, setBitOfReg(3, RegisterE)}
	OpcodeExtSet3h  = opcode{0xDC, "SET 3,H", "Set bit 3 of H", 0, 8, setBitOfReg(3, RegisterH)}
	OpcodeExtSet3l  = opcode{0xDD, "SET 3,L", "Set bit 3 of L", 0, 8, setBitOfReg(3, RegisterL)}
	OpcodeExtSet3hl = opcode{0xDE, "SET 3,(HL)", "Set bit 3 of value pointed by HL", 0, 16, setBitOfHLAddr(3)}
	OpcodeExtSet3a  = opcode{0xDF, "SET 3,A", "Set bit 3 of A", 0, 8, setBitOfReg(3, RegisterA)}
	OpcodeExtSet4b  = opcode{0xE0, "SET 4,B", "Set bit 4 of B", 0, 8, setBitOfReg(4, RegisterB)}
	OpcodeExtSet4c  = opcode{0xE1, "SET 4,C", "Set bit 4 of C", 0, 8, setBitOfReg(4, RegisterC)}
	OpcodeExtSet4d  = opcode{0xE2, "SET 4,D", "Set bit 4 of D", 0, 8, setBitOfReg(4, RegisterD)}
	OpcodeExtSet4e  = opcode{0xE3, "SET 4,E", "Set bit 4 of E", 0, 8, setBitOfReg(4, RegisterE)}
	OpcodeExtSet4h  = opcode{0xE4, "SET 4,H", "Set bit 4 of H", 0, 8, setBitOfReg(4, RegisterH)}
	OpcodeExtSet4l  = opcode{0xE5, "SET 4,L", "Set bit 4 of L", 0, 8, setBitOfReg(4, RegisterL)}
	OpcodeExtSet4hl = opcode{0xE6, "SET 4,(HL)", "Set bit 4 of value pointed by HL", 0, 16, setBitOfHLAddr(4)}
	OpcodeExtSet4a  = opcode{0xE7, "SET 4,A", "Set bit 4 of A", 0, 8, setBitOfReg(4, RegisterA)}
	OpcodeExtSet5b  = opcode{0xE8, "SET 5,B", "Set bit 5 of B", 0, 8, setBitOfReg(5, RegisterB)}
	OpcodeExtSet5c  = opcode{0xE9, "SET 5,C", "Set bit 5 of C", 0, 8, setBitOfReg(5, RegisterC)}
	OpcodeExtSet5d  = opcode{0xEA, "SET 5,D", "Set bit 5 of D", 0, 8, setBitOfReg(5, RegisterD)}
	OpcodeExtSet5e  = opcode{0xEB, "SET 5,E", "Set bit 5 of E", 0, 8, setBitOfReg(5, RegisterE)}
	OpcodeExtSet5h  = opcode{0xEC, "SET 5,H", "Set bit 5 of H", 0, 8, setBitOfReg(5, RegisterH)}
	OpcodeExtSet5l  = opcode{0xED, "SET 5,L", "Set bit 5 of L", 0, 8, setBitOfReg(5, RegisterL)}
	OpcodeExtSet5hl = opcode{0xEE, "SET 5,(HL)", "Set bit 5 of value pointed by HL", 0, 16, setBitOfHLAddr(5)}
	OpcodeExtSet5a  = opcode{0xEF, "SET 5,A", "Set bit 5 of A", 0, 8, setBitOfReg(5, RegisterA)}
	OpcodeExtSet6b  = opcode{0xF0, "SET 6,B", "Set bit 6 of B", 0, 8, setBitOfReg(6, RegisterB)}
	OpcodeExtSet6c  = opcode{0xF1, "SET 6,C", "Set bit 6 of C", 0, 8, setBitOfReg(6, RegisterC)}
	OpcodeExtSet6d  = opcode{0xF2, "SET 6,D", "Set bit 6 of D", 0, 8, setBitOfReg(6, RegisterD)}
	OpcodeExtSet6e  = opcode{0xF3, "SET 6,E", "Set bit 6 of E", 0, 8, setBitOfReg(6, RegisterE)}
	OpcodeExtSet6h  = opcode{0xF4, "SET 6,H", "Set bit 6 of H", 0, 8, setBitOfReg(6, RegisterH)}
	OpcodeExtSet6l  = opcode{0xF5, "SET 6,L", "Set bit 6 of L", 0, 8, setBitOfReg(6, RegisterL)}
	OpcodeExtSet6hl = opcode{0xF6, "SET 6,(HL)", "Set bit 6 of value pointed by HL", 0, 16, setBitOfHLAddr(6)}
	OpcodeExtSet6a  = opcode{0xF7, "SET 6,A", "Set bit 6 of A", 0, 8, setBitOfReg(6, RegisterA)}
	OpcodeExtSet7b  = opcode{0xF8, "SET 7,B", "Set bit 7 of B", 0, 8, setBitOfReg(7, RegisterB)}
	OpcodeExtSet7c  = opcode{0xF9, "SET 7,C", "Set bit 7 of C", 0, 8, setBitOfReg(7, RegisterC)}
	OpcodeExtSet7d  = opcode{0xFA, "SET 7,D", "Set bit 7 of D", 0, 8, setBitOfReg(7, RegisterD)}
	OpcodeExtSet7e  = opcode{0xFB, "SET 7,E", "Set bit 7 of E", 0, 8, setBitOfReg(7, RegisterE)}
	OpcodeExtSet7h  = opcode{0xFC, "SET 7,H", "Set bit 7 of H", 0, 8, setBitOfReg(7, RegisterH)}
	OpcodeExtSet7l  = opcode{0xFD, "SET 7,L", "Set bit 7 of L", 0, 8, setBitOfReg(7, RegisterL)}
	OpcodeExtSet7hl = opcode{0xFE, "SET 7,(HL)", "Set bit 7 of value pointed by HL", 0, 16, setBitOfHLAddr(7)}
	OpcodeExtSet7a  = opcode{0xFF, "SET 7,A", "Set bit 7 of A", 0, 8, setBitOfReg(7, RegisterA)}
)

func LookupExtOpcode(opcodeByte byte) opcode {
	switch opcodeByte {
		case 0x00: return OpcodeExtRlcB 
		case 0x01: return OpcodeExtRlcC 
		case 0x02: return OpcodeExtRlcD 
		case 0x03: return OpcodeExtRlcE 
		case 0x04: return OpcodeExtRlcH 
		case 0x05: return OpcodeExtRlcL 
		case 0x06: return OpcodeExtRlcHl 
		case 0x07: return OpcodeExtRlcA 
		case 0x08: return OpcodeExtRrcB 
		case 0x09: return OpcodeExtRrcC 
		case 0x0A: return OpcodeExtRrcD 
		case 0x0B: return OpcodeExtRrcE 
		case 0x0C: return OpcodeExtRrcH 
		case 0x0D: return OpcodeExtRrcL 
		case 0x0E: return OpcodeExtRrcHl 
		case 0x0F: return OpcodeExtRrcA 
		case 0x10: return OpcodeExtRlB 
		case 0x11: return OpcodeExtRlC 
		case 0x12: return OpcodeExtRlD 
		case 0x13: return OpcodeExtRlE 
		case 0x14: return OpcodeExtRlH 
		case 0x15: return OpcodeExtRlL 
		case 0x16: return OpcodeExtRlHl 
		case 0x17: return OpcodeExtRlA 
		case 0x18: return OpcodeExtRrB 
		case 0x19: return OpcodeExtRrC 
		case 0x1A: return OpcodeExtRrD 
		case 0x1B: return OpcodeExtRrE 
		case 0x1C: return OpcodeExtRrH 
		case 0x1D: return OpcodeExtRrL 
		case 0x1E: return OpcodeExtRrHl 
		case 0x1F: return OpcodeExtRrA 
		case 0x20: return OpcodeExtSlaB 
		case 0x21: return OpcodeExtSlaC 
		case 0x22: return OpcodeExtSlaD 
		case 0x23: return OpcodeExtSlaE 
		case 0x24: return OpcodeExtSlaH 
		case 0x25: return OpcodeExtSlaL 
		case 0x26: return OpcodeExtSlaHl 
		case 0x27: return OpcodeExtSlaA 
		case 0x28: return OpcodeExtSraB 
		case 0x29: return OpcodeExtSraC 
		case 0x2A: return OpcodeExtSraD 
		case 0x2B: return OpcodeExtSraE 
		case 0x2C: return OpcodeExtSraH 
		case 0x2D: return OpcodeExtSraL 
		case 0x2E: return OpcodeExtSraHl 
		case 0x2F: return OpcodeExtSraA 
		case 0x30: return OpcodeExtSwapB 
		case 0x31: return OpcodeExtSwapC 
		case 0x32: return OpcodeExtSwapD 
		case 0x33: return OpcodeExtSwapE 
		case 0x34: return OpcodeExtSwapH 
		case 0x35: return OpcodeExtSwapL 
		case 0x36: return OpcodeExtSwapHl 
		case 0x37: return OpcodeExtSwapA 
		case 0x38: return OpcodeExtSrlB 
		case 0x39: return OpcodeExtSrlC 
		case 0x3A: return OpcodeExtSrlD 
		case 0x3B: return OpcodeExtSrlE 
		case 0x3C: return OpcodeExtSrlH 
		case 0x3D: return OpcodeExtSrlL 
		case 0x3E: return OpcodeExtSrlHl 
		case 0x3F: return OpcodeExtSrlA 
		case 0x40: return OpcodeExtBit0b 
		case 0x41: return OpcodeExtBit0c 
		case 0x42: return OpcodeExtBit0d 
		case 0x43: return OpcodeExtBit0e 
		case 0x44: return OpcodeExtBit0h 
		case 0x45: return OpcodeExtBit0l 
		case 0x46: return OpcodeExtBit0hl 
		case 0x47: return OpcodeExtBit0a 
		case 0x48: return OpcodeExtBit1b 
		case 0x49: return OpcodeExtBit1c 
		case 0x4A: return OpcodeExtBit1d 
		case 0x4B: return OpcodeExtBit1e 
		case 0x4C: return OpcodeExtBit1h 
		case 0x4D: return OpcodeExtBit1l 
		case 0x4E: return OpcodeExtBit1hl 
		case 0x4F: return OpcodeExtBit1a 
		case 0x50: return OpcodeExtBit2b 
		case 0x51: return OpcodeExtBit2c 
		case 0x52: return OpcodeExtBit2d 
		case 0x53: return OpcodeExtBit2e 
		case 0x54: return OpcodeExtBit2h 
		case 0x55: return OpcodeExtBit2l 
		case 0x56: return OpcodeExtBit2hl 
		case 0x57: return OpcodeExtBit2a 
		case 0x58: return OpcodeExtBit3b 
		case 0x59: return OpcodeExtBit3c 
		case 0x5A: return OpcodeExtBit3d 
		case 0x5B: return OpcodeExtBit3e 
		case 0x5C: return OpcodeExtBit3h 
		case 0x5D: return OpcodeExtBit3l 
		case 0x5E: return OpcodeExtBit3hl 
		case 0x5F: return OpcodeExtBit3a 
		case 0x60: return OpcodeExtBit4b 
		case 0x61: return OpcodeExtBit4c 
		case 0x62: return OpcodeExtBit4d 
		case 0x63: return OpcodeExtBit4e 
		case 0x64: return OpcodeExtBit4h 
		case 0x65: return OpcodeExtBit4l 
		case 0x66: return OpcodeExtBit4hl 
		case 0x67: return OpcodeExtBit4a 
		case 0x68: return OpcodeExtBit5b 
		case 0x69: return OpcodeExtBit5c 
		case 0x6A: return OpcodeExtBit5d 
		case 0x6B: return OpcodeExtBit5e 
		case 0x6C: return OpcodeExtBit5h 
		case 0x6D: return OpcodeExtBit5l 
		case 0x6E: return OpcodeExtBit5hl 
		case 0x6F: return OpcodeExtBit5a 
		case 0x70: return OpcodeExtBit6b 
		case 0x71: return OpcodeExtBit6c 
		case 0x72: return OpcodeExtBit6d 
		case 0x73: return OpcodeExtBit6e 
		case 0x74: return OpcodeExtBit6h 
		case 0x75: return OpcodeExtBit6l 
		case 0x76: return OpcodeExtBit6hl 
		case 0x77: return OpcodeExtBit6a 
		case 0x78: return OpcodeExtBit7b 
		case 0x79: return OpcodeExtBit7c 
		case 0x7A: return OpcodeExtBit7d 
		case 0x7B: return OpcodeExtBit7e 
		case 0x7C: return OpcodeExtBit7h 
		case 0x7D: return OpcodeExtBit7l 
		case 0x7E: return OpcodeExtBit7hl 
		case 0x7F: return OpcodeExtBit7a 
		case 0x80: return OpcodeExtRes0b 
		case 0x81: return OpcodeExtRes0c 
		case 0x82: return OpcodeExtRes0d 
		case 0x83: return OpcodeExtRes0e 
		case 0x84: return OpcodeExtRes0h 
		case 0x85: return OpcodeExtRes0l 
		case 0x86: return OpcodeExtRes0hl 
		case 0x87: return OpcodeExtRes0a 
		case 0x88: return OpcodeExtRes1b 
		case 0x89: return OpcodeExtRes1c 
		case 0x8A: return OpcodeExtRes1d 
		case 0x8B: return OpcodeExtRes1e 
		case 0x8C: return OpcodeExtRes1h 
		case 0x8D: return OpcodeExtRes1l 
		case 0x8E: return OpcodeExtRes1hl 
		case 0x8F: return OpcodeExtRes1a 
		case 0x90: return OpcodeExtRes2b 
		case 0x91: return OpcodeExtRes2c 
		case 0x92: return OpcodeExtRes2d 
		case 0x93: return OpcodeExtRes2e 
		case 0x94: return OpcodeExtRes2h 
		case 0x95: return OpcodeExtRes2l 
		case 0x96: return OpcodeExtRes2hl 
		case 0x97: return OpcodeExtRes2a 
		case 0x98: return OpcodeExtRes3b 
		case 0x99: return OpcodeExtRes3c 
		case 0x9A: return OpcodeExtRes3d 
		case 0x9B: return OpcodeExtRes3e 
		case 0x9C: return OpcodeExtRes3h 
		case 0x9D: return OpcodeExtRes3l 
		case 0x9E: return OpcodeExtRes3hl 
		case 0x9F: return OpcodeExtRes3a 
		case 0xA0: return OpcodeExtRes4b 
		case 0xA1: return OpcodeExtRes4c 
		case 0xA2: return OpcodeExtRes4d 
		case 0xA3: return OpcodeExtRes4e 
		case 0xA4: return OpcodeExtRes4h 
		case 0xA5: return OpcodeExtRes4l 
		case 0xA6: return OpcodeExtRes4hl 
		case 0xA7: return OpcodeExtRes4a 
		case 0xA8: return OpcodeExtRes5b 
		case 0xA9: return OpcodeExtRes5c 
		case 0xAA: return OpcodeExtRes5d 
		case 0xAB: return OpcodeExtRes5e 
		case 0xAC: return OpcodeExtRes5h 
		case 0xAD: return OpcodeExtRes5l 
		case 0xAE: return OpcodeExtRes5hl 
		case 0xAF: return OpcodeExtRes5a 
		case 0xB0: return OpcodeExtRes6b 
		case 0xB1: return OpcodeExtRes6c 
		case 0xB2: return OpcodeExtRes6d 
		case 0xB3: return OpcodeExtRes6e 
		case 0xB4: return OpcodeExtRes6h 
		case 0xB5: return OpcodeExtRes6l 
		case 0xB6: return OpcodeExtRes6hl 
		case 0xB7: return OpcodeExtRes6a 
		case 0xB8: return OpcodeExtRes7b 
		case 0xB9: return OpcodeExtRes7c 
		case 0xBA: return OpcodeExtRes7d 
		case 0xBB: return OpcodeExtRes7e 
		case 0xBC: return OpcodeExtRes7h 
		case 0xBD: return OpcodeExtRes7l 
		case 0xBE: return OpcodeExtRes7hl 
		case 0xBF: return OpcodeExtRes7a 
		case 0xC0: return OpcodeExtSet0b 
		case 0xC1: return OpcodeExtSet0c 
		case 0xC2: return OpcodeExtSet0d 
		case 0xC3: return OpcodeExtSet0e 
		case 0xC4: return OpcodeExtSet0h 
		case 0xC5: return OpcodeExtSet0l 
		case 0xC6: return OpcodeExtSet0hl 
		case 0xC7: return OpcodeExtSet0a 
		case 0xC8: return OpcodeExtSet1b 
		case 0xC9: return OpcodeExtSet1c 
		case 0xCA: return OpcodeExtSet1d 
		case 0xCB: return OpcodeExtSet1e 
		case 0xCC: return OpcodeExtSet1h 
		case 0xCD: return OpcodeExtSet1l 
		case 0xCE: return OpcodeExtSet1hl 
		case 0xCF: return OpcodeExtSet1a 
		case 0xD0: return OpcodeExtSet2b 
		case 0xD1: return OpcodeExtSet2c 
		case 0xD2: return OpcodeExtSet2d 
		case 0xD3: return OpcodeExtSet2e 
		case 0xD4: return OpcodeExtSet2h 
		case 0xD5: return OpcodeExtSet2l 
		case 0xD6: return OpcodeExtSet2hl 
		case 0xD7: return OpcodeExtSet2a 
		case 0xD8: return OpcodeExtSet3b 
		case 0xD9: return OpcodeExtSet3c 
		case 0xDA: return OpcodeExtSet3d 
		case 0xDB: return OpcodeExtSet3e 
		case 0xDC: return OpcodeExtSet3h 
		case 0xDD: return OpcodeExtSet3l 
		case 0xDE: return OpcodeExtSet3hl 
		case 0xDF: return OpcodeExtSet3a 
		case 0xE0: return OpcodeExtSet4b 
		case 0xE1: return OpcodeExtSet4c 
		case 0xE2: return OpcodeExtSet4d 
		case 0xE3: return OpcodeExtSet4e 
		case 0xE4: return OpcodeExtSet4h 
		case 0xE5: return OpcodeExtSet4l 
		case 0xE6: return OpcodeExtSet4hl 
		case 0xE7: return OpcodeExtSet4a 
		case 0xE8: return OpcodeExtSet5b 
		case 0xE9: return OpcodeExtSet5c 
		case 0xEA: return OpcodeExtSet5d 
		case 0xEB: return OpcodeExtSet5e 
		case 0xEC: return OpcodeExtSet5h 
		case 0xED: return OpcodeExtSet5l 
		case 0xEE: return OpcodeExtSet5hl 
		case 0xEF: return OpcodeExtSet5a 
		case 0xF0: return OpcodeExtSet6b 
		case 0xF1: return OpcodeExtSet6c 
		case 0xF2: return OpcodeExtSet6d 
		case 0xF3: return OpcodeExtSet6e 
		case 0xF4: return OpcodeExtSet6h 
		case 0xF5: return OpcodeExtSet6l 
		case 0xF6: return OpcodeExtSet6hl 
		case 0xF7: return OpcodeExtSet6a 
		case 0xF8: return OpcodeExtSet7b 
		case 0xF9: return OpcodeExtSet7c 
		case 0xFA: return OpcodeExtSet7d 
		case 0xFB: return OpcodeExtSet7e 
		case 0xFC: return OpcodeExtSet7h 
		case 0xFD: return OpcodeExtSet7l 
		case 0xFE: return OpcodeExtSet7hl 
		case 0xFF: return OpcodeExtSet7a 
	}
	panic("Unreachable code!")
}

func doNothing() {

}

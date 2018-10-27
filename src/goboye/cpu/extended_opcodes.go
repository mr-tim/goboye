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

var opcodeMapExt = map[uint8]opcode{
	0x00: OpcodeExtRlcB,
	0x01: OpcodeExtRlcC,
	0x02: OpcodeExtRlcD,
	0x03: OpcodeExtRlcE,
	0x04: OpcodeExtRlcH,
	0x05: OpcodeExtRlcL,
	0x06: OpcodeExtRlcHl,
	0x07: OpcodeExtRlcA,
	0x08: OpcodeExtRrcB,
	0x09: OpcodeExtRrcC,
	0x0A: OpcodeExtRrcD,
	0x0B: OpcodeExtRrcE,
	0x0C: OpcodeExtRrcH,
	0x0D: OpcodeExtRrcL,
	0x0E: OpcodeExtRrcHl,
	0x0F: OpcodeExtRrcA,
	0x10: OpcodeExtRlB,
	0x11: OpcodeExtRlC,
	0x12: OpcodeExtRlD,
	0x13: OpcodeExtRlE,
	0x14: OpcodeExtRlH,
	0x15: OpcodeExtRlL,
	0x16: OpcodeExtRlHl,
	0x17: OpcodeExtRlA,
	0x18: OpcodeExtRrB,
	0x19: OpcodeExtRrC,
	0x1A: OpcodeExtRrD,
	0x1B: OpcodeExtRrE,
	0x1C: OpcodeExtRrH,
	0x1D: OpcodeExtRrL,
	0x1E: OpcodeExtRrHl,
	0x1F: OpcodeExtRrA,
	0x20: OpcodeExtSlaB,
	0x21: OpcodeExtSlaC,
	0x22: OpcodeExtSlaD,
	0x23: OpcodeExtSlaE,
	0x24: OpcodeExtSlaH,
	0x25: OpcodeExtSlaL,
	0x26: OpcodeExtSlaHl,
	0x27: OpcodeExtSlaA,
	0x28: OpcodeExtSraB,
	0x29: OpcodeExtSraC,
	0x2A: OpcodeExtSraD,
	0x2B: OpcodeExtSraE,
	0x2C: OpcodeExtSraH,
	0x2D: OpcodeExtSraL,
	0x2E: OpcodeExtSraHl,
	0x2F: OpcodeExtSraA,
	0x30: OpcodeExtSwapB,
	0x31: OpcodeExtSwapC,
	0x32: OpcodeExtSwapD,
	0x33: OpcodeExtSwapE,
	0x34: OpcodeExtSwapH,
	0x35: OpcodeExtSwapL,
	0x36: OpcodeExtSwapHl,
	0x37: OpcodeExtSwapA,
	0x38: OpcodeExtSrlB,
	0x39: OpcodeExtSrlC,
	0x3A: OpcodeExtSrlD,
	0x3B: OpcodeExtSrlE,
	0x3C: OpcodeExtSrlH,
	0x3D: OpcodeExtSrlL,
	0x3E: OpcodeExtSrlHl,
	0x3F: OpcodeExtSrlA,
	0x40: OpcodeExtBit0b,
	0x41: OpcodeExtBit0c,
	0x42: OpcodeExtBit0d,
	0x43: OpcodeExtBit0e,
	0x44: OpcodeExtBit0h,
	0x45: OpcodeExtBit0l,
	0x46: OpcodeExtBit0hl,
	0x47: OpcodeExtBit0a,
	0x48: OpcodeExtBit1b,
	0x49: OpcodeExtBit1c,
	0x4A: OpcodeExtBit1d,
	0x4B: OpcodeExtBit1e,
	0x4C: OpcodeExtBit1h,
	0x4D: OpcodeExtBit1l,
	0x4E: OpcodeExtBit1hl,
	0x4F: OpcodeExtBit1a,
	0x50: OpcodeExtBit2b,
	0x51: OpcodeExtBit2c,
	0x52: OpcodeExtBit2d,
	0x53: OpcodeExtBit2e,
	0x54: OpcodeExtBit2h,
	0x55: OpcodeExtBit2l,
	0x56: OpcodeExtBit2hl,
	0x57: OpcodeExtBit2a,
	0x58: OpcodeExtBit3b,
	0x59: OpcodeExtBit3c,
	0x5A: OpcodeExtBit3d,
	0x5B: OpcodeExtBit3e,
	0x5C: OpcodeExtBit3h,
	0x5D: OpcodeExtBit3l,
	0x5E: OpcodeExtBit3hl,
	0x5F: OpcodeExtBit3a,
	0x60: OpcodeExtBit4b,
	0x61: OpcodeExtBit4c,
	0x62: OpcodeExtBit4d,
	0x63: OpcodeExtBit4e,
	0x64: OpcodeExtBit4h,
	0x65: OpcodeExtBit4l,
	0x66: OpcodeExtBit4hl,
	0x67: OpcodeExtBit4a,
	0x68: OpcodeExtBit5b,
	0x69: OpcodeExtBit5c,
	0x6A: OpcodeExtBit5d,
	0x6B: OpcodeExtBit5e,
	0x6C: OpcodeExtBit5h,
	0x6D: OpcodeExtBit5l,
	0x6E: OpcodeExtBit5hl,
	0x6F: OpcodeExtBit5a,
	0x70: OpcodeExtBit6b,
	0x71: OpcodeExtBit6c,
	0x72: OpcodeExtBit6d,
	0x73: OpcodeExtBit6e,
	0x74: OpcodeExtBit6h,
	0x75: OpcodeExtBit6l,
	0x76: OpcodeExtBit6hl,
	0x77: OpcodeExtBit6a,
	0x78: OpcodeExtBit7b,
	0x79: OpcodeExtBit7c,
	0x7A: OpcodeExtBit7d,
	0x7B: OpcodeExtBit7e,
	0x7C: OpcodeExtBit7h,
	0x7D: OpcodeExtBit7l,
	0x7E: OpcodeExtBit7hl,
	0x7F: OpcodeExtBit7a,
	0x80: OpcodeExtRes0b,
	0x81: OpcodeExtRes0c,
	0x82: OpcodeExtRes0d,
	0x83: OpcodeExtRes0e,
	0x84: OpcodeExtRes0h,
	0x85: OpcodeExtRes0l,
	0x86: OpcodeExtRes0hl,
	0x87: OpcodeExtRes0a,
	0x88: OpcodeExtRes1b,
	0x89: OpcodeExtRes1c,
	0x8A: OpcodeExtRes1d,
	0x8B: OpcodeExtRes1e,
	0x8C: OpcodeExtRes1h,
	0x8D: OpcodeExtRes1l,
	0x8E: OpcodeExtRes1hl,
	0x8F: OpcodeExtRes1a,
	0x90: OpcodeExtRes2b,
	0x91: OpcodeExtRes2c,
	0x92: OpcodeExtRes2d,
	0x93: OpcodeExtRes2e,
	0x94: OpcodeExtRes2h,
	0x95: OpcodeExtRes2l,
	0x96: OpcodeExtRes2hl,
	0x97: OpcodeExtRes2a,
	0x98: OpcodeExtRes3b,
	0x99: OpcodeExtRes3c,
	0x9A: OpcodeExtRes3d,
	0x9B: OpcodeExtRes3e,
	0x9C: OpcodeExtRes3h,
	0x9D: OpcodeExtRes3l,
	0x9E: OpcodeExtRes3hl,
	0x9F: OpcodeExtRes3a,
	0xA0: OpcodeExtRes4b,
	0xA1: OpcodeExtRes4c,
	0xA2: OpcodeExtRes4d,
	0xA3: OpcodeExtRes4e,
	0xA4: OpcodeExtRes4h,
	0xA5: OpcodeExtRes4l,
	0xA6: OpcodeExtRes4hl,
	0xA7: OpcodeExtRes4a,
	0xA8: OpcodeExtRes5b,
	0xA9: OpcodeExtRes5c,
	0xAA: OpcodeExtRes5d,
	0xAB: OpcodeExtRes5e,
	0xAC: OpcodeExtRes5h,
	0xAD: OpcodeExtRes5l,
	0xAE: OpcodeExtRes5hl,
	0xAF: OpcodeExtRes5a,
	0xB0: OpcodeExtRes6b,
	0xB1: OpcodeExtRes6c,
	0xB2: OpcodeExtRes6d,
	0xB3: OpcodeExtRes6e,
	0xB4: OpcodeExtRes6h,
	0xB5: OpcodeExtRes6l,
	0xB6: OpcodeExtRes6hl,
	0xB7: OpcodeExtRes6a,
	0xB8: OpcodeExtRes7b,
	0xB9: OpcodeExtRes7c,
	0xBA: OpcodeExtRes7d,
	0xBB: OpcodeExtRes7e,
	0xBC: OpcodeExtRes7h,
	0xBD: OpcodeExtRes7l,
	0xBE: OpcodeExtRes7hl,
	0xBF: OpcodeExtRes7a,
	0xC0: OpcodeExtSet0b,
	0xC1: OpcodeExtSet0c,
	0xC2: OpcodeExtSet0d,
	0xC3: OpcodeExtSet0e,
	0xC4: OpcodeExtSet0h,
	0xC5: OpcodeExtSet0l,
	0xC6: OpcodeExtSet0hl,
	0xC7: OpcodeExtSet0a,
	0xC8: OpcodeExtSet1b,
	0xC9: OpcodeExtSet1c,
	0xCA: OpcodeExtSet1d,
	0xCB: OpcodeExtSet1e,
	0xCC: OpcodeExtSet1h,
	0xCD: OpcodeExtSet1l,
	0xCE: OpcodeExtSet1hl,
	0xCF: OpcodeExtSet1a,
	0xD0: OpcodeExtSet2b,
	0xD1: OpcodeExtSet2c,
	0xD2: OpcodeExtSet2d,
	0xD3: OpcodeExtSet2e,
	0xD4: OpcodeExtSet2h,
	0xD5: OpcodeExtSet2l,
	0xD6: OpcodeExtSet2hl,
	0xD7: OpcodeExtSet2a,
	0xD8: OpcodeExtSet3b,
	0xD9: OpcodeExtSet3c,
	0xDA: OpcodeExtSet3d,
	0xDB: OpcodeExtSet3e,
	0xDC: OpcodeExtSet3h,
	0xDD: OpcodeExtSet3l,
	0xDE: OpcodeExtSet3hl,
	0xDF: OpcodeExtSet3a,
	0xE0: OpcodeExtSet4b,
	0xE1: OpcodeExtSet4c,
	0xE2: OpcodeExtSet4d,
	0xE3: OpcodeExtSet4e,
	0xE4: OpcodeExtSet4h,
	0xE5: OpcodeExtSet4l,
	0xE6: OpcodeExtSet4hl,
	0xE7: OpcodeExtSet4a,
	0xE8: OpcodeExtSet5b,
	0xE9: OpcodeExtSet5c,
	0xEA: OpcodeExtSet5d,
	0xEB: OpcodeExtSet5e,
	0xEC: OpcodeExtSet5h,
	0xED: OpcodeExtSet5l,
	0xEE: OpcodeExtSet5hl,
	0xEF: OpcodeExtSet5a,
	0xF0: OpcodeExtSet6b,
	0xF1: OpcodeExtSet6c,
	0xF2: OpcodeExtSet6d,
	0xF3: OpcodeExtSet6e,
	0xF4: OpcodeExtSet6h,
	0xF5: OpcodeExtSet6l,
	0xF6: OpcodeExtSet6hl,
	0xF7: OpcodeExtSet6a,
	0xF8: OpcodeExtSet7b,
	0xF9: OpcodeExtSet7c,
	0xFA: OpcodeExtSet7d,
	0xFB: OpcodeExtSet7e,
	0xFC: OpcodeExtSet7h,
	0xFD: OpcodeExtSet7l,
	0xFE: OpcodeExtSet7hl,
	0xFF: OpcodeExtSet7a,
}

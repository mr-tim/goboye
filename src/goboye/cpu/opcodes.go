package cpu

type opcode struct {
	code uint8
	disassembly string
	description string
}

var opcodes = []opcode{
	{0x00, "NOP", "No Operation"},
	{0x01, "LD BC,nn", "Load 16-bit immediate into BC"},
	{0x02, "LD (BC),A", "Save A to address pointed by BC"},
	{0x03, "INC BC", "Increment 16-bit BC"},
	{0x04, "INC B", "Increment B"},
	{0x05, "DEC B", "Decrement B"},
	{0x06, "LD B,n", "Load 8-bit immediate into B"},
	{0x07, "RLC A", "Rotate A left with carry"},
	{0x08, "LD (nn),SP", "Save SP to given address"},
	{0x09, "ADD HL,BC", "Add 16-bit BC to HL"},
	{0x0A, "LD A,(BC)", "Load A from address pointed to by BC"},
	{0x0B, "DEC BC", "Decrement 16-bit BC"},
	{0x0C, "INC C", "Increment C"},
	{0x0D, "DEC C", "Decrement C"},
	{0x0E, "LD C,n", "Load 8-bit immediate into C"},
	{0x0F, "RRC A", "Rotate A right with carry"},
	{0x10, "STOP", "Stop processor"},
	{0x11, "LD DE,nn", "Load 16-bit immediate into DE"},
	{0x12, "LD (DE),A", "Save A to address pointed by DE"},
	{0x13, "INC DE", "Increment 16-bit DE"},
	{0x14, "INC D", "Increment D"},
	{0x15, "DEC D", "Decrement D"},
	{0x16, "LD D,n", "Load 8-bit immediate into D"},
	{0x17, "RL A", "Rotate A left"},
	{0x18, "JR n", "Relative jump by signed immediate"},
	{0x19, "ADD HL,DE", "Add 16-bit DE to HL"},
	{0x1A, "LD A,(DE)", "Load A from address pointed to by DE"},
	{0x1B, "DEC DE", "Decrement 16-bit DE"},
	{0x1C, "INC E", "Increment E"},
	{0x1D, "DEC E", "Decrement E"},
	{0x1E, "LD E,n", "Load 8-bit immediate into E"},
	{0x1F, "RR A", "Rotate A right"},
	{0x20, "JR NZ,n", "Relative jump by signed immediate if last result was not zero"},
	{0x21, "LD HL,nn", "Load 16-bit immediate into HL"},
	{0x22, "LDI (HL),A", "Save A to address pointed by HL, and increment HL"},
	{0x23, "INC HL", "Increment 16-bit HL"},
	{0x24, "INC H", "Increment H"},
	{0x25, "DEC H", "Decrement H"},
	{0x26, "LD H,n", "Load 8-bit immediate into H"},
	{0x27, "DAA", "Adjust A for BCD addition"},
	{0x28, "JR Z,n", "Relative jump by signed immediate if last result was zero"},
	{0x29, "ADD HL,HL", "Add 16-bit HL to HL"},
	{0x2A, "LDI A,(HL)", "Load A from address pointed to by HL, and increment HL"},
	{0x2B, "DEC HL", "Decrement 16-bit HL"},
	{0x2C, "INC L", "Increment L"},
	{0x2D, "DEC L", "Decrement L"},
	{0x2E, "LD L,n", "Load 8-bit immediate into L"},
	{0x2F, "CPL", "Complement (logical NOT) on A"},
	{0x30, "JR NC,n", "Relative jump by signed immediate if last result caused no carry"},
	{0x31, "LD SP,nn", "Load 16-bit immediate into SP"},
	{0x32, "LDD (HL),A", "Save A to address pointed by HL, and decrement HL"},
	{0x33, "INC SP", "Increment 16-bit HL"},
	{0x34, "INC (HL)", "Increment value pointed by HL"},
	{0x35, "DEC (HL)", "Decrement value pointed by HL"},
	{0x36, "LD (HL),n", "Load 8-bit immediate into address pointed by HL"},
	{0x37, "SCF", "Set carry flag"},
	{0x38, "JR C,n", "Relative jump by signed immediate if last result caused carry"},
	{0x39, "ADD HL,SP", "Add 16-bit SP to HL"},
	{0x3A, "LDD A,(HL)", "Load A from address pointed to by HL, and decrement HL"},
	{0x3B, "DEC SP", "Decrement 16-bit SP"},
	{0x3C, "INC A", "Increment A"},
	{0x3D, "DEC A", "Decrement A"},
	{0x3E, "LD A,n", "Load 8-bit immediate into A"},
	{0x3F, "CCF", "Clear carry flag"},
	{0x40, "LD B,B", "Copy B to B"},
	{0x41, "LD B,C", "Copy C to B"},
	{0x42, "LD B,D", "Copy D to B"},
	{0x43, "LD B,E", "Copy E to B"},
	{0x44, "LD B,H", "Copy H to B"},
	{0x45, "LD B,L", "Copy L to B"},
	{0x46, "LD B,(HL)", "Copy value pointed by HL to B"},
	{0x47, "LD B,A", "Copy A to B"},
	{0x48, "LD C,B", "Copy B to C"},
	{0x49, "LD C,C", "Copy C to C"},
	{0x4A, "LD C,D", "Copy D to C"},
	{0x4B, "LD C,E", "Copy E to C"},
	{0x4C, "LD C,H", "Copy H to C"},
	{0x4D, "LD C,L", "Copy L to C"},
	{0x4E, "LD C,(HL)", "Copy value pointed by HL to C"},
	{0x4F, "LD C,A", "Copy A to C"},
	{0x50, "LD D,B", "Copy B to D"},
	{0x51, "LD D,C", "Copy C to D"},
	{0x52, "LD D,D", "Copy D to D"},
	{0x53, "LD D,E", "Copy E to D"},
	{0x54, "LD D,H", "Copy H to D"},
	{0x55, "LD D,L", "Copy L to D"},
	{0x56, "LD D,(HL)", "Copy value pointed by HL to D"},
	{0x57, "LD D,A", "Copy A to D"},
	{0x58, "LD E,B", "Copy B to E"},
	{0x59, "LD E,C", "Copy C to E"},
	{0x5A, "LD E,D", "Copy D to E"},
	{0x5B, "LD E,E", "Copy E to E"},
	{0x5C, "LD E,H", "Copy H to E"},
	{0x5D, "LD E,L", "Copy L to E"},
	{0x5E, "LD E,(HL)", "Copy value pointed by HL to E"},
	{0x5F, "LD E,A", "Copy A to E"},
	{0x60, "LD H,B", "Copy B to H"},
	{0x61, "LD H,C", "Copy C to H"},
	{0x62, "LD H,D", "Copy D to H"},
	{0x63, "LD H,E", "Copy E to H"},
	{0x64, "LD H,H", "Copy H to H"},
	{0x65, "LD H,L", "Copy L to H"},
	{0x66, "LD H,(HL)", "Copy value pointed by HL to H"},
	{0x67, "LD H,A", "Copy A to H"},
	{0x68, "LD L,B", "Copy B to L"},
	{0x69, "LD L,C", "Copy C to L"},
	{0x6A, "LD L,D", "Copy D to L"},
	{0x6B, "LD L,E", "Copy E to L"},
	{0x6C, "LD L,H", "Copy H to L"},
	{0x6D, "LD L,L", "Copy L to L"},
	{0x6E, "LD L,(HL)", "Copy value pointed by HL to L"},
	{0x6F, "LD L,A", "Copy A to L"},
	{0x70, "LD (HL),B", "Copy B to address pointed by HL"},
	{0x71, "LD (HL),C", "Copy C to address pointed by HL"},
	{0x72, "LD (HL),D", "Copy D to address pointed by HL"},
	{0x73, "LD (HL),E", "Copy E to address pointed by HL"},
	{0x74, "LD (HL),H", "Copy H to address pointed by HL"},
	{0x75, "LD (HL),L", "Copy L to address pointed by HL"},
	{0x76, "HALT", "Halt processor"},
	{0x77, "LD (HL),A", "Copy A to address pointed by HL"},
	{0x78, "LD A,B", "Copy B to A"},
	{0x79, "LD A,C", "Copy C to A"},
	{0x7A, "LD A,D", "Copy D to A"},
	{0x7B, "LD A,E", "Copy E to A"},
	{0x7C, "LD A,H", "Copy H to A"},
	{0x7D, "LD A,L", "Copy L to A"},
	{0x7E, "LD A,(HL)", "Copy value pointed by HL to A"},
	{0x7F, "LD A,A", "Copy A to A"},
	{0x80, "ADD A,B", "Add B to A"},
	{0x81, "ADD A,C", "Add C to A"},
	{0x82, "ADD A,D", "Add D to A"},
	{0x83, "ADD A,E", "Add E to A"},
	{0x84, "ADD A,H", "Add H to A"},
	{0x85, "ADD A,L", "Add L to A"},
	{0x86, "ADD A,(HL)", "Add value pointed by HL to A"},
	{0x87, "ADD A,A", "Add A to A"},
	{0x88, "ADC A,B", "Add B and carry flag to A"},
	{0x89, "ADC A,C", "Add C and carry flag to A"},
	{0x8A, "ADC A,D", "Add D and carry flag to A"},
	{0x8B, "ADC A,E", "Add E and carry flag to A"},
	{0x8C, "ADC A,H", "Add H and carry flag to A"},
	{0x8D, "ADC A,L", "Add and carry flag L to A"},
	{0x8E, "ADC A,(HL)", "Add value pointed by HL and carry flag to A"},
	{0x8F, "ADC A,A", "Add A and carry flag to A"},
	{0x90, "SUB A,B", "Subtract B from A"},
	{0x91, "SUB A,C", "Subtract C from A"},
	{0x92, "SUB A,D", "Subtract D from A"},
	{0x93, "SUB A,E", "Subtract E from A"},
	{0x94, "SUB A,H", "Subtract H from A"},
	{0x95, "SUB A,L", "Subtract L from A"},
	{0x96, "SUB A,(HL)", "Subtract value pointed by HL from A"},
	{0x97, "SUB A,A", "Subtract A from A"},
	{0x98, "SBC A,B", "Subtract B and carry flag from A"},
	{0x99, "SBC A,C", "Subtract C and carry flag from A"},
	{0x9A, "SBC A,D", "Subtract D and carry flag from A"},
	{0x9B, "SBC A,E", "Subtract E and carry flag from A"},
	{0x9C, "SBC A,H", "Subtract H and carry flag from A"},
	{0x9D, "SBC A,L", "Subtract and carry flag L from A"},
	{0x9E, "SBC A,(HL)", "Subtract value pointed by HL and carry flag from A"},
	{0x9F, "SBC A,A", "Subtract A and carry flag from A"},
	{0xA0, "AND B", "Logical AND B against A"},
	{0xA1, "AND C", "Logical AND C against A"},
	{0xA2, "AND D", "Logical AND D against A"},
	{0xA3, "AND E", "Logical AND E against A"},
	{0xA4, "AND H", "Logical AND H against A"},
	{0xA5, "AND L", "Logical AND L against A"},
	{0xA6, "AND (HL)", "Logical AND value pointed by HL against A"},
	{0xA7, "AND A", "Logical AND A against A"},
	{0xA8, "XOR B", "Logical XOR B against A"},
	{0xA9, "XOR C", "Logical XOR C against A"},
	{0xAA, "XOR D", "Logical XOR D against A"},
	{0xAB, "XOR E", "Logical XOR E against A"},
	{0xAC, "XOR H", "Logical XOR H against A"},
	{0xAD, "XOR L", "Logical XOR L against A"},
	{0xAE, "XOR (HL)", "Logical XOR value pointed by HL against A"},
	{0xAF, "XOR A", "Logical XOR A against A"},
	{0xB0, "OR B", "Logical OR B against A"},
	{0xB1, "OR C", "Logical OR C against A"},
	{0xB2, "OR D", "Logical OR D against A"},
	{0xB3, "OR E", "Logical OR E against A"},
	{0xB4, "OR H", "Logical OR H against A"},
	{0xB5, "OR L", "Logical OR L against A"},
	{0xB6, "OR (HL)", "Logical OR value pointed by HL against A"},
	{0xB7, "OR A", "Logical OR A against A"},
	{0xB8, "CP B", "Compare B against A"},
	{0xB9, "CP C", "Compare C against A"},
	{0xBA, "CP D", "Compare D against A"},
	{0xBB, "CP E", "Compare E against A"},
	{0xBC, "CP H", "Compare H against A"},
	{0xBD, "CP L", "Compare L against A"},
	{0xBE, "CP (HL)", "Compare value pointed by HL against A"},
	{0xBF, "CP A", "Compare A against A"},
	{0xC0, "RET NZ", "Return if last result was not zero"},
	{0xC1, "POP BC", "Pop 16-bit value from stack into BC"},
	{0xC2, "JP NZ,nn", "Absolute jump to 16-bit location if last result was not zero"},
	{0xC3, "JP nn", "Absolute jump to 16-bit location"},
	{0xC4, "CALL NZ,nn", "Call routine at 16-bit location if last result was not zero"},
	{0xC5, "PUSH BC", "Push 16-bit BC onto stack"},
	{0xC6, "ADD A,n", "Add 8-bit immediate to A"},
	{0xC7, "RST 0", "Call routine at address 0000h"},
	{0xC8, "RET Z", "Return if last result was zero"},
	{0xC9, "RET", "Return to calling routine"},
	{0xCA, "JP Z,nn", "Absolute jump to 16-bit location if last result was zero"},
	{0xCB, "Ext ops", "Extended operations (two-byte instruction code)"},
	{0xCC, "CALL Z,nn", "Call routine at 16-bit location if last result was zero"},
	{0xCD, "CALL nn", "Call routine at 16-bit location"},
	{0xCE, "ADC A,n", "Add 8-bit immediate and carry to A"},
	{0xCF, "RST 8", "Call routine at address 0008h"},
	{0xD0, "RET NC", "Return if last result caused no carry"},
	{0xD1, "POP DE", "Pop 16-bit value from stack into DE"},
	{0xD2, "JP NC,nn", "Absolute jump to 16-bit location if last result caused no carry"},
	//{0xD3, "XX", "Operation removed in this CPU"},
	{0xD4, "CALL NC,nn", "Call routine at 16-bit location if last result caused no carry"},
	{0xD5, "PUSH DE", "Push 16-bit DE onto stack"},
	{0xD6, "SUB A,n", "Subtract 8-bit immediate from A"},
	{0xD7, "RST 10", "Call routine at address 0010h"},
	{0xD8, "RET C", "Return if last result caused carry"},
	{0xD9, "RETI", "Enable interrupts and return to calling routine"},
	{0xDA, "JP C,nn", "Absolute jump to 16-bit location if last result caused carry"},
	{0xDB, "XX", "Operation removed in this CPU"},
	{0xDC, "CALL C,nn", "Call routine at 16-bit location if last result caused carry"},
	{0xDD, "XX", "Operation removed in this CPU"},
	{0xDE, "SBC A,n", "Subtract 8-bit immediate and carry from A"},
	{0xDF, "RST 18", "Call routine at address 0018h"},
	{0xE0, "LDH (n),A", "Save A at address pointed to by (FF00h + 8-bit immediate)"},
	{0xE1, "POP HL", "Pop 16-bit value from stack into HL"},
	{0xE2, "LDH (C),A", "Save A at address pointed to by (FF00h + C)"},
	//{0xE3, "XX", "Operation removed in this CPU"},
	//{0xE4, "XX", "Operation removed in this CPU"},
	{0xE5, "PUSH HL", "Push 16-bit HL onto stack"},
	{0xE6, "AND n", "Logical AND 8-bit immediate against A"},
	{0xE7, "RST 20", "Call routine at address 0020h"},
	{0xE8, "ADD SP,d", "Add signed 8-bit immediate to SP"},
	{0xE9, "JP (HL)", "Jump to 16-bit value pointed by HL"},
	{0xEA, "LD (nn),A", "Save A at given 16-bit address"},
	//{0xEB, "XX", "Operation removed in this CPU"},
	//{0xEC, "XX", "Operation removed in this CPU"},
	//{0xED, "XX", "Operation removed in this CPU"},
	{0xEE, "XOR n", "Logical XOR 8-bit immediate against A"},
	{0xEF, "RST 28", "Call routine at address 0028h"},
	{0xF0, "LDH A,(n)", "Load A from address pointed to by (FF00h + 8-bit immediate)"},
	{0xF1, "POP AF", "Pop 16-bit value from stack into AF"},
	//{0xF2, "XX", "Operation removed in this CPU"},
	{0xF3, "DI", "DIsable interrupts"},
	//{0xF4, "XX", "Operation removed in this CPU"},
	{0xF5, "PUSH AF", "Push 16-bit AF onto stack"},
	{0xF6, "OR n", "Logical OR 8-bit immediate against A"},
	{0xF7, "RST 30", "Call routine at address 0030h"},
	{0xF8, "LDHL SP,d", "Add signed 8-bit immediate to SP and save result in HL"},
	{0xF9, "LD SP,HL", "Copy HL to SP"},
	{0xFA, "LD A,(nn)", "Load A from given 16-bit address"},
	{0xFB, "EI", "Enable interrupts"},
	//{0xFC, "XX", "Operation removed in this CPU"},
	//{0xFD, "XX", "Operation removed in this CPU"},
	{0xFE, "CP n", "Compare 8-bit immediate against A"},
	{0xFF, "RST 38", "Call routine at address 0038h"},
}

var extendedOpcodes = []opcode{
	{0x00, "RLC B", "Rotate B left with carry"},
	{0x01, "RLC C", "Rotate C left with carry"},
	{0x02, "RLC D", "Rotate D left with carry"},
	{0x03, "RLC E", "Rotate E left with carry"},
	{0x04, "RLC H", "Rotate H left with carry"},
	{0x05, "RLC L", "Rotate L left with carry"},
	{0x06, "RLC (HL)", "Rotate value pointed by HL left with carry"},
	{0x07, "RLC A", "Rotate A left with carry"},
	{0x08, "RRC B", "Rotate B right with carry"},
	{0x09, "RRC C", "Rotate C right with carry"},
	{0x0A, "RRC D", "Rotate D right with carry"},
	{0x0B, "RRC E", "Rotate E right with carry"},
	{0x0C, "RRC H", "Rotate H right with carry"},
	{0x0D, "RRC L", "Rotate L right with carry"},
	{0x0E, "RRC (HL)", "Rotate value pointed by HL right with carry"},
	{0x0F, "RRC A", "Rotate A right with carry"},
	{0x10, "RL B", "Rotate B left"},
	{0x11, "RL C", "Rotate C left"},
	{0x12, "RL D", "Rotate D left"},
	{0x13, "RL E", "Rotate E left"},
	{0x14, "RL H", "Rotate H left"},
	{0x15, "RL L", "Rotate L left"},
	{0x16, "RL (HL)", "Rotate value pointed by HL left"},
	{0x17, "RL A", "Rotate A left"},
	{0x18, "RR B", "Rotate B right"},
	{0x19, "RR C", "Rotate C right"},
	{0x1A, "RR D", "Rotate D right"},
	{0x1B, "RR E", "Rotate E right"},
	{0x1C, "RR H", "Rotate H right"},
	{0x1D, "RR L", "Rotate L right"},
	{0x1E, "RR (HL)", "Rotate value pointed by HL right"},
	{0x1F, "RR A", "Rotate A right"},
	{0x20, "SLA B", "Shift B left preserving sign"},
	{0x21, "SLA C", "Shift C left preserving sign"},
	{0x22, "SLA D", "Shift D left preserving sign"},
	{0x23, "SLA E", "Shift E left preserving sign"},
	{0x24, "SLA H", "Shift H left preserving sign"},
	{0x25, "SLA L", "Shift L left preserving sign"},
	{0x26, "SLA (HL)", "Shift value pointed by HL left preserving sign"},
	{0x27, "SLA A", "Shift A left preserving sign"},
	{0x28, "SRA B", "Shift B right preserving sign"},
	{0x29, "SRA C", "Shift C right preserving sign"},
	{0x2A, "SRA D", "Shift D right preserving sign"},
	{0x2B, "SRA E", "Shift E right preserving sign"},
	{0x2C, "SRA H", "Shift H right preserving sign"},
	{0x2D, "SRA L", "Shift L right preserving sign"},
	{0x2E, "SRA (HL)", "Shift value pointed by HL right preserving sign"},
	{0x2F, "SRA A", "Shift A right preserving sign"},
	{0x30, "SWAP B", "Swap nybbles in B"},
	{0x31, "SWAP C", "Swap nybbles in C"},
	{0x32, "SWAP D", "Swap nybbles in D"},
	{0x33, "SWAP E", "Swap nybbles in E"},
	{0x34, "SWAP H", "Swap nybbles in H"},
	{0x35, "SWAP L", "Swap nybbles in L"},
	{0x36, "SWAP (HL)", "Swap nybbles in value pointed by HL"},
	{0x37, "SWAP A", "Swap nybbles in A"},
	{0x38, "SRL B", "Shift B right"},
	{0x39, "SRL C", "Shift C right"},
	{0x3A, "SRL D", "Shift D right"},
	{0x3B, "SRL E", "Shift E right"},
	{0x3C, "SRL H", "Shift H right"},
	{0x3D, "SRL L", "Shift L right"},
	{0x3E, "SRL (HL)", "Shift value pointed by HL right"},
	{0x3F, "SRL A", "Shift A right"},
	{0x40, "BIT 0,B", "Test bit 0 of B"},
	{0x41, "BIT 0,C", "Test bit 0 of C"},
	{0x42, "BIT 0,D", "Test bit 0 of D"},
	{0x43, "BIT 0,E", "Test bit 0 of E"},
	{0x44, "BIT 0,H", "Test bit 0 of H"},
	{0x45, "BIT 0,L", "Test bit 0 of L"},
	{0x46, "BIT 0,(HL)", "Test bit 0 of value pointed by HL"},
	{0x47, "BIT 0,A", "Test bit 0 of A"},
	{0x48, "BIT 1,B", "Test bit 1 of B"},
	{0x49, "BIT 1,C", "Test bit 1 of C"},
	{0x4A, "BIT 1,D", "Test bit 1 of D"},
	{0x4B, "BIT 1,E", "Test bit 1 of E"},
	{0x4C, "BIT 1,H", "Test bit 1 of H"},
	{0x4D, "BIT 1,L", "Test bit 1 of L"},
	{0x4E, "BIT 1,(HL)", "Test bit 1 of value pointed by HL"},
	{0x4F, "BIT 1,A", "Test bit 1 of A"},
	{0x50, "BIT 2,B", "Test bit 2 of B"},
	{0x51, "BIT 2,C", "Test bit 2 of C"},
	{0x52, "BIT 2,D", "Test bit 2 of D"},
	{0x53, "BIT 2,E", "Test bit 2 of E"},
	{0x54, "BIT 2,H", "Test bit 2 of H"},
	{0x55, "BIT 2,L", "Test bit 2 of L"},
	{0x56, "BIT 2,(HL)", "Test bit 2 of value pointed by HL"},
	{0x57, "BIT 2,A", "Test bit 2 of A"},
	{0x58, "BIT 3,B", "Test bit 3 of B"},
	{0x59, "BIT 3,C", "Test bit 3 of C"},
	{0x5A, "BIT 3,D", "Test bit 3 of D"},
	{0x5B, "BIT 3,E", "Test bit 3 of E"},
	{0x5C, "BIT 3,H", "Test bit 3 of H"},
	{0x5D, "BIT 3,L", "Test bit 3 of L"},
	{0x5E, "BIT 3,(HL)", "Test bit 3 of value pointed by HL"},
	{0x5F, "BIT 3,A", "Test bit 3 of A"},
	{0x60, "BIT 4,B", "Test bit 4 of B"},
	{0x61, "BIT 4,C", "Test bit 4 of C"},
	{0x62, "BIT 4,D", "Test bit 4 of D"},
	{0x63, "BIT 4,E", "Test bit 4 of E"},
	{0x64, "BIT 4,H", "Test bit 4 of H"},
	{0x65, "BIT 4,L", "Test bit 4 of L"},
	{0x66, "BIT 4,(HL)", "Test bit 4 of value pointed by HL"},
	{0x67, "BIT 4,A", "Test bit 4 of A"},
	{0x68, "BIT 5,B", "Test bit 5 of B"},
	{0x69, "BIT 5,C", "Test bit 5 of C"},
	{0x6A, "BIT 5,D", "Test bit 5 of D"},
	{0x6B, "BIT 5,E", "Test bit 5 of E"},
	{0x6C, "BIT 5,H", "Test bit 5 of H"},
	{0x6D, "BIT 5,L", "Test bit 5 of L"},
	{0x6E, "BIT 5,(HL)", "Test bit 5 of value pointed by HL"},
	{0x6F, "BIT 5,A", "Test bit 5 of A"},
	{0x70, "BIT 6,B", "Test bit 6 of B"},
	{0x71, "BIT 6,C", "Test bit 6 of C"},
	{0x72, "BIT 6,D", "Test bit 6 of D"},
	{0x73, "BIT 6,E", "Test bit 6 of E"},
	{0x74, "BIT 6,H", "Test bit 6 of H"},
	{0x75, "BIT 6,L", "Test bit 6 of L"},
	{0x76, "BIT 6,(HL)", "Test bit 6 of value pointed by HL"},
	{0x77, "BIT 6,A", "Test bit 6 of A"},
	{0x78, "BIT 7,B", "Test bit 7 of B"},
	{0x79, "BIT 7,C", "Test bit 7 of C"},
	{0x7A, "BIT 7,D", "Test bit 7 of D"},
	{0x7B, "BIT 7,E", "Test bit 7 of E"},
	{0x7C, "BIT 7,H", "Test bit 7 of H"},
	{0x7D, "BIT 7,L", "Test bit 7 of L"},
	{0x7E, "BIT 7,(HL)", "Test bit 7 of value pointed by HL"},
	{0x7F, "BIT 7,A", "Test bit 7 of A"},
	{0x80, "RES 0,B", "Clear (reset) bit 0 of B"},
	{0x81, "RES 0,C", "Clear (reset) bit 0 of C"},
	{0x82, "RES 0,D", "Clear (reset) bit 0 of D"},
	{0x83, "RES 0,E", "Clear (reset) bit 0 of E"},
	{0x84, "RES 0,H", "Clear (reset) bit 0 of H"},
	{0x85, "RES 0,L", "Clear (reset) bit 0 of L"},
	{0x86, "RES 0,(HL)", "Clear (reset) bit 0 of value pointed by HL"},
	{0x87, "RES 0,A", "Clear (reset) bit 0 of A"},
	{0x88, "RES 1,B", "Clear (reset) bit 1 of B"},
	{0x89, "RES 1,C", "Clear (reset) bit 1 of C"},
	{0x8A, "RES 1,D", "Clear (reset) bit 1 of D"},
	{0x8B, "RES 1,E", "Clear (reset) bit 1 of E"},
	{0x8C, "RES 1,H", "Clear (reset) bit 1 of H"},
	{0x8D, "RES 1,L", "Clear (reset) bit 1 of L"},
	{0x8E, "RES 1,(HL)", "Clear (reset) bit 1 of value pointed by HL"},
	{0x8F, "RES 1,A", "Clear (reset) bit 1 of A"},
	{0x90, "RES 2,B", "Clear (reset) bit 2 of B"},
	{0x91, "RES 2,C", "Clear (reset) bit 2 of C"},
	{0x92, "RES 2,D", "Clear (reset) bit 2 of D"},
	{0x93, "RES 2,E", "Clear (reset) bit 2 of E"},
	{0x94, "RES 2,H", "Clear (reset) bit 2 of H"},
	{0x95, "RES 2,L", "Clear (reset) bit 2 of L"},
	{0x96, "RES 2,(HL)", "Clear (reset) bit 2 of value pointed by HL"},
	{0x97, "RES 2,A", "Clear (reset) bit 2 of A"},
	{0x98, "RES 3,B", "Clear (reset) bit 3 of B"},
	{0x99, "RES 3,C", "Clear (reset) bit 3 of C"},
	{0x9A, "RES 3,D", "Clear (reset) bit 3 of D"},
	{0x9B, "RES 3,E", "Clear (reset) bit 3 of E"},
	{0x9C, "RES 3,H", "Clear (reset) bit 3 of H"},
	{0x9D, "RES 3,L", "Clear (reset) bit 3 of L"},
	{0x9E, "RES 3,(HL)", "Clear (reset) bit 3 of value pointed by HL"},
	{0x9F, "RES 3,A", "Clear (reset) bit 3 of A"},
	{0xA0, "RES 4,B", "Clear (reset) bit 4 of B"},
	{0xA1, "RES 4,C", "Clear (reset) bit 4 of C"},
	{0xA2, "RES 4,D", "Clear (reset) bit 4 of D"},
	{0xA3, "RES 4,E", "Clear (reset) bit 4 of E"},
	{0xA4, "RES 4,H", "Clear (reset) bit 4 of H"},
	{0xA5, "RES 4,L", "Clear (reset) bit 4 of L"},
	{0xA6, "RES 4,(HL)", "Clear (reset) bit 4 of value pointed by HL"},
	{0xA7, "RES 4,A", "Clear (reset) bit 4 of A"},
	{0xA8, "RES 5,B", "Clear (reset) bit 5 of B"},
	{0xA9, "RES 5,C", "Clear (reset) bit 5 of C"},
	{0xAA, "RES 5,D", "Clear (reset) bit 5 of D"},
	{0xAB, "RES 5,E", "Clear (reset) bit 5 of E"},
	{0xAC, "RES 5,H", "Clear (reset) bit 5 of H"},
	{0xAD, "RES 5,L", "Clear (reset) bit 5 of L"},
	{0xAE, "RES 5,(HL)", "Clear (reset) bit 5 of value pointed by HL"},
	{0xAF, "RES 5,A", "Clear (reset) bit 5 of A"},
	{0xB0, "RES 6,B", "Clear (reset) bit 6 of B"},
	{0xB1, "RES 6,C", "Clear (reset) bit 6 of C"},
	{0xB2, "RES 6,D", "Clear (reset) bit 6 of D"},
	{0xB3, "RES 6,E", "Clear (reset) bit 6 of E"},
	{0xB4, "RES 6,H", "Clear (reset) bit 6 of H"},
	{0xB5, "RES 6,L", "Clear (reset) bit 6 of L"},
	{0xB6, "RES 6,(HL)", "Clear (reset) bit 6 of value pointed by HL"},
	{0xB7, "RES 6,A", "Clear (reset) bit 6 of A"},
	{0xB8, "RES 7,B", "Clear (reset) bit 7 of B"},
	{0xB9, "RES 7,C", "Clear (reset) bit 7 of C"},
	{0xBA, "RES 7,D", "Clear (reset) bit 7 of D"},
	{0xBB, "RES 7,E", "Clear (reset) bit 7 of E"},
	{0xBC, "RES 7,H", "Clear (reset) bit 7 of H"},
	{0xBD, "RES 7,L", "Clear (reset) bit 7 of L"},
	{0xBE, "RES 7,(HL)", "Clear (reset) bit 7 of value pointed by HL"},
	{0xBF, "RES 7,A", "Clear (reset) bit 7 of A"},
	{0xC0, "SET 0,B", "Set bit 0 of B"},
	{0xC1, "SET 0,C", "Set bit 0 of C"},
	{0xC2, "SET 0,D", "Set bit 0 of D"},
	{0xC3, "SET 0,E", "Set bit 0 of E"},
	{0xC4, "SET 0,H", "Set bit 0 of H"},
	{0xC5, "SET 0,L", "Set bit 0 of L"},
	{0xC6, "SET 0,(HL)", "Set bit 0 of value pointed by HL"},
	{0xC7, "SET 0,A", "Set bit 0 of A"},
	{0xC8, "SET 1,B", "Set bit 1 of B"},
	{0xC9, "SET 1,C", "Set bit 1 of C"},
	{0xCA, "SET 1,D", "Set bit 1 of D"},
	{0xCB, "SET 1,E", "Set bit 1 of E"},
	{0xCC, "SET 1,H", "Set bit 1 of H"},
	{0xCD, "SET 1,L", "Set bit 1 of L"},
	{0xCE, "SET 1,(HL)", "Set bit 1 of value pointed by HL"},
	{0xCF, "SET 1,A", "Set bit 1 of A"},
	{0xD0, "SET 2,B", "Set bit 2 of B"},
	{0xD1, "SET 2,C", "Set bit 2 of C"},
	{0xD2, "SET 2,D", "Set bit 2 of D"},
	{0xD3, "SET 2,E", "Set bit 2 of E"},
	{0xD4, "SET 2,H", "Set bit 2 of H"},
	{0xD5, "SET 2,L", "Set bit 2 of L"},
	{0xD6, "SET 2,(HL)", "Set bit 2 of value pointed by HL"},
	{0xD7, "SET 2,A", "Set bit 2 of A"},
	{0xD8, "SET 3,B", "Set bit 3 of B"},
	{0xD9, "SET 3,C", "Set bit 3 of C"},
	{0xDA, "SET 3,D", "Set bit 3 of D"},
	{0xDB, "SET 3,E", "Set bit 3 of E"},
	{0xDC, "SET 3,H", "Set bit 3 of H"},
	{0xDD, "SET 3,L", "Set bit 3 of L"},
	{0xDE, "SET 3,(HL)", "Set bit 3 of value pointed by HL"},
	{0xDF, "SET 3,A", "Set bit 3 of A"},
	{0xE0, "SET 4,B", "Set bit 4 of B"},
	{0xE1, "SET 4,C", "Set bit 4 of C"},
	{0xE2, "SET 4,D", "Set bit 4 of D"},
	{0xE3, "SET 4,E", "Set bit 4 of E"},
	{0xE4, "SET 4,H", "Set bit 4 of H"},
	{0xE5, "SET 4,L", "Set bit 4 of L"},
	{0xE6, "SET 4,(HL)", "Set bit 4 of value pointed by HL"},
	{0xE7, "SET 4,A", "Set bit 4 of A"},
	{0xE8, "SET 5,B", "Set bit 5 of B"},
	{0xE9, "SET 5,C", "Set bit 5 of C"},
	{0xEA, "SET 5,D", "Set bit 5 of D"},
	{0xEB, "SET 5,E", "Set bit 5 of E"},
	{0xEC, "SET 5,H", "Set bit 5 of H"},
	{0xED, "SET 5,L", "Set bit 5 of L"},
	{0xEE, "SET 5,(HL)", "Set bit 5 of value pointed by HL"},
	{0xEF, "SET 5,A", "Set bit 5 of A"},
	{0xF0, "SET 6,B", "Set bit 6 of B"},
	{0xF1, "SET 6,C", "Set bit 6 of C"},
	{0xF2, "SET 6,D", "Set bit 6 of D"},
	{0xF3, "SET 6,E", "Set bit 6 of E"},
	{0xF4, "SET 6,H", "Set bit 6 of H"},
	{0xF5, "SET 6,L", "Set bit 6 of L"},
	{0xF6, "SET 6,(HL)", "Set bit 6 of value pointed by HL"},
	{0xF7, "SET 6,A", "Set bit 6 of A"},
	{0xF8, "SET 7,B", "Set bit 7 of B"},
	{0xF9, "SET 7,C", "Set bit 7 of C"},
	{0xFA, "SET 7,D", "Set bit 7 of D"},
	{0xFB, "SET 7,E", "Set bit 7 of E"},
	{0xFC, "SET 7,H", "Set bit 7 of H"},
	{0xFD, "SET 7,L", "Set bit 7 of L"},
	{0xFE, "SET 7,(HL)", "Set bit 7 of value pointed by HL"},
	{0xFF, "SET 7,A", "Set bit 7 of A"},
}
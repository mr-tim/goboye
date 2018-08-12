package cpu

type opcodeHandler func(opcode, *processor)

type opcode struct {
	code          uint8
	disassembly   string
	description   string
	payloadLength uint8
	cycles        uint8
	handler       opcodeHandler
}

var (
	OpcodeNop       = opcode{0x00, "NOP", "No Operation", 0, 1, nopHandler}
	OpcodeLdBcnn    = opcode{0x01, "LD BC,nn", "Load 16-bit immediate into BC", 0, 1, load16BitToBC}
	OpcodeLdBca     = opcode{0x02, "LD (BC),A", "Save A to address pointed by BC", 0, 1, saveAToBCAddr}
	OpcodeIncBc     = opcode{0x03, "INC BC", "Increment 16-bit BC", 0, 1, incrementBC}
	OpcodeIncB      = opcode{0x04, "INC B", "Increment B", 0, 1, incrementB}
	OpcodeDecB      = opcode{0x05, "DEC B", "Decrement B", 0, 1, decrementB}
	OpcodeLdBn      = opcode{0x06, "LD B,n", "Load 8-bit immediate into B", 0, 1, load8BitToB}
	OpcodeRlcA      = opcode{0x07, "RLC A", "Rotate A left with carry", 0, 1, unimplementedHandler}
	OpcodeLdNnsp    = opcode{0x08, "LD (nn),SP", "Save SP to given address", 0, 1, unimplementedHandler}
	OpcodeAddHlbc   = opcode{0x09, "ADD HL,BC", "Add 16-bit BC to HL", 0, 1, unimplementedHandler}
	OpcodeLdAbc     = opcode{0x0A, "LD A,(BC)", "Load A from address pointed to by BC", 0, 1, unimplementedHandler}
	OpcodeDecBc     = opcode{0x0B, "DEC BC", "Decrement 16-bit BC", 0, 1, decrementBC}
	OpcodeIncC      = opcode{0x0C, "INC C", "Increment C", 0, 1, incrementC}
	OpcodeDecC      = opcode{0x0D, "DEC C", "Decrement C", 0, 1, decrementC}
	OpcodeLdCn      = opcode{0x0E, "LD C,n", "Load 8-bit immediate into C", 0, 1, load8BitToC}
	OpcodeRrcA      = opcode{0x0F, "RRC A", "Rotate A right with carry", 0, 1, unimplementedHandler}
	OpcodeStop      = opcode{0x10, "STOP", "Stop processor", 0, 1, unimplementedHandler}
	OpcodeLdDenn    = opcode{0x11, "LD DE,nn", "Load 16-bit immediate into DE", 0, 1, load16BitToDE}
	OpcodeLdDea     = opcode{0x12, "LD (DE),A", "Save A to address pointed by DE", 0, 1, saveAToDEAddr}
	OpcodeIncDe     = opcode{0x13, "INC DE", "Increment 16-bit DE", 0, 1, incrementDE}
	OpcodeIncD      = opcode{0x14, "INC D", "Increment D", 0, 1, incrementD}
	OpcodeDecD      = opcode{0x15, "DEC D", "Decrement D", 0, 1, decrementD}
	OpcodeLdDn      = opcode{0x16, "LD D,n", "Load 8-bit immediate into D", 0, 1, load8BitToD}
	OpcodeRlA       = opcode{0x17, "RL A", "Rotate A left", 0, 1, unimplementedHandler}
	OpcodeJrN       = opcode{0x18, "JR n", "Relative jump by signed immediate", 0, 1, unimplementedHandler}
	OpcodeAddHlde   = opcode{0x19, "ADD HL,DE", "Add 16-bit DE to HL", 0, 1, unimplementedHandler}
	OpcodeLdAde     = opcode{0x1A, "LD A,(DE)", "Load A from address pointed to by DE", 0, 1, unimplementedHandler}
	OpcodeDecDe     = opcode{0x1B, "DEC DE", "Decrement 16-bit DE", 0, 1, decrementDE}
	OpcodeIncE      = opcode{0x1C, "INC E", "Increment E", 0, 1, incrementE}
	OpcodeDecE      = opcode{0x1D, "DEC E", "Decrement E", 0, 1, decrementE}
	OpcodeLdEn      = opcode{0x1E, "LD E,n", "Load 8-bit immediate into E", 0, 1, load8BitToE}
	OpcodeRrA       = opcode{0x1F, "RR A", "Rotate A right", 0, 1, unimplementedHandler}
	OpcodeJrNzn     = opcode{0x20, "JR NZ,n", "Relative jump by signed immediate if last result was not zero", 0, 1, unimplementedHandler}
	OpcodeLdHlnn    = opcode{0x21, "LD HL,nn", "Load 16-bit immediate into HL", 0, 1, load16BitToHL}
	OpcodeLdiHla    = opcode{0x22, "LDI (HL),A", "Save A to address pointed by HL, and increment HL", 0, 1, saveAToHLAddrInc}
	OpcodeIncHl     = opcode{0x23, "INC HL", "Increment 16-bit HL", 0, 1, incrementHL}
	OpcodeIncH      = opcode{0x24, "INC H", "Increment H", 0, 1, incrementH}
	OpcodeDecH      = opcode{0x25, "DEC H", "Decrement H", 0, 1, decrementH}
	OpcodeLdHn      = opcode{0x26, "LD H,n", "Load 8-bit immediate into H", 0, 1, load8BitToH}
	OpcodeDaa       = opcode{0x27, "DAA", "Adjust A for BCD addition", 0, 1, unimplementedHandler}
	OpcodeJrZn      = opcode{0x28, "JR Z,n", "Relative jump by signed immediate if last result was zero", 0, 1, unimplementedHandler}
	OpcodeAddHlhl   = opcode{0x29, "ADD HL,HL", "Add 16-bit HL to HL", 0, 1, unimplementedHandler}
	OpcodeLdiAhl    = opcode{0x2A, "LDI A,(HL)", "Load A from address pointed to by HL, and increment HL", 0, 1, unimplementedHandler}
	OpcodeDecHl     = opcode{0x2B, "DEC HL", "Decrement 16-bit HL", 0, 1, decrementHL}
	OpcodeIncL      = opcode{0x2C, "INC L", "Increment L", 0, 1, incrementL}
	OpcodeDecL      = opcode{0x2D, "DEC L", "Decrement L", 0, 1, decrementL}
	OpcodeLdLn      = opcode{0x2E, "LD L,n", "Load 8-bit immediate into L", 0, 1, load8BitToL}
	OpcodeCpl       = opcode{0x2F, "CPL", "Complement (logical NOT) on A", 0, 1, unimplementedHandler}
	OpcodeJrNcn     = opcode{0x30, "JR NC,n", "Relative jump by signed immediate if last result caused no carry", 0, 1, unimplementedHandler}
	OpcodeLdSpnn    = opcode{0x31, "LD SP,nn", "Load 16-bit immediate into SP", 0, 1, load16BitToSP}
	OpcodeLddHla    = opcode{0x32, "LDD (HL),A", "Save A to address pointed by HL, and decrement HL", 0, 1, saveAToHLAddrDec}
	OpcodeIncSp     = opcode{0x33, "INC SP", "Increment 16-bit SP", 0, 1, incrementSP}
	OpcodeIncHlAddr = opcode{0x34, "INC (HL)", "Increment value pointed by HL", 0, 1, incrementHLAddr}
	OpcodeDecHlAddr = opcode{0x35, "DEC (HL)", "Decrement value pointed by HL", 0, 1, decrementHLAddr}
	OpcodeLdHln     = opcode{0x36, "LD (HL),n", "Load 8-bit immediate into address pointed by HL", 0, 1, load8BitToHLAddr}
	OpcodeScf       = opcode{0x37, "SCF", "Set carry flag", 0, 1, unimplementedHandler}
	OpcodeJrCn      = opcode{0x38, "JR C,n", "Relative jump by signed immediate if last result caused carry", 0, 1, unimplementedHandler}
	OpcodeAddHlsp   = opcode{0x39, "ADD HL,SP", "Add 16-bit SP to HL", 0, 1, unimplementedHandler}
	OpcodeLddAhl    = opcode{0x3A, "LDD A,(HL)", "Load A from address pointed to by HL, and decrement HL", 0, 1, unimplementedHandler}
	OpcodeDecSp     = opcode{0x3B, "DEC SP", "Decrement 16-bit SP", 0, 1, decrementSP}
	OpcodeIncA      = opcode{0x3C, "INC A", "Increment A", 0, 1, incrementA}
	OpcodeDecA      = opcode{0x3D, "DEC A", "Decrement A", 0, 1, decrementA}
	OpcodeLdAn      = opcode{0x3E, "LD A,n", "Load 8-bit immediate into A", 0, 1, load8BitToA}
	OpcodeCcf       = opcode{0x3F, "CCF", "Clear carry flag", 0, 1, unimplementedHandler}
	OpcodeLdBb      = opcode{0x40, "LD B,B", "Copy B to B", 0, 1, unimplementedHandler}
	OpcodeLdBc      = opcode{0x41, "LD B,C", "Copy C to B", 0, 1, unimplementedHandler}
	OpcodeLdBd      = opcode{0x42, "LD B,D", "Copy D to B", 0, 1, unimplementedHandler}
	OpcodeLdBe      = opcode{0x43, "LD B,E", "Copy E to B", 0, 1, unimplementedHandler}
	OpcodeLdBh      = opcode{0x44, "LD B,H", "Copy H to B", 0, 1, unimplementedHandler}
	OpcodeLdBl      = opcode{0x45, "LD B,L", "Copy L to B", 0, 1, unimplementedHandler}
	OpcodeLdBhl     = opcode{0x46, "LD B,(HL)", "Copy value pointed by HL to B", 0, 1, unimplementedHandler}
	OpcodeLdBa      = opcode{0x47, "LD B,A", "Copy A to B", 0, 1, unimplementedHandler}
	OpcodeLdCb      = opcode{0x48, "LD C,B", "Copy B to C", 0, 1, unimplementedHandler}
	OpcodeLdCc      = opcode{0x49, "LD C,C", "Copy C to C", 0, 1, unimplementedHandler}
	OpcodeLdCd      = opcode{0x4A, "LD C,D", "Copy D to C", 0, 1, unimplementedHandler}
	OpcodeLdCe      = opcode{0x4B, "LD C,E", "Copy E to C", 0, 1, unimplementedHandler}
	OpcodeLdCh      = opcode{0x4C, "LD C,H", "Copy H to C", 0, 1, unimplementedHandler}
	OpcodeLdCl      = opcode{0x4D, "LD C,L", "Copy L to C", 0, 1, unimplementedHandler}
	OpcodeLdChl     = opcode{0x4E, "LD C,(HL)", "Copy value pointed by HL to C", 0, 1, unimplementedHandler}
	OpcodeLdCa      = opcode{0x4F, "LD C,A", "Copy A to C", 0, 1, unimplementedHandler}
	OpcodeLdDb      = opcode{0x50, "LD D,B", "Copy B to D", 0, 1, unimplementedHandler}
	OpcodeLdDc      = opcode{0x51, "LD D,C", "Copy C to D", 0, 1, unimplementedHandler}
	OpcodeLdDd      = opcode{0x52, "LD D,D", "Copy D to D", 0, 1, unimplementedHandler}
	OpcodeLdDe      = opcode{0x53, "LD D,E", "Copy E to D", 0, 1, unimplementedHandler}
	OpcodeLdDh      = opcode{0x54, "LD D,H", "Copy H to D", 0, 1, unimplementedHandler}
	OpcodeLdDl      = opcode{0x55, "LD D,L", "Copy L to D", 0, 1, unimplementedHandler}
	OpcodeLdDhl     = opcode{0x56, "LD D,(HL)", "Copy value pointed by HL to D", 0, 1, unimplementedHandler}
	OpcodeLdDa      = opcode{0x57, "LD D,A", "Copy A to D", 0, 1, unimplementedHandler}
	OpcodeLdEb      = opcode{0x58, "LD E,B", "Copy B to E", 0, 1, unimplementedHandler}
	OpcodeLdEc      = opcode{0x59, "LD E,C", "Copy C to E", 0, 1, unimplementedHandler}
	OpcodeLdEd      = opcode{0x5A, "LD E,D", "Copy D to E", 0, 1, unimplementedHandler}
	OpcodeLdEe      = opcode{0x5B, "LD E,E", "Copy E to E", 0, 1, unimplementedHandler}
	OpcodeLdEh      = opcode{0x5C, "LD E,H", "Copy H to E", 0, 1, unimplementedHandler}
	OpcodeLdEl      = opcode{0x5D, "LD E,L", "Copy L to E", 0, 1, unimplementedHandler}
	OpcodeLdEhl     = opcode{0x5E, "LD E,(HL)", "Copy value pointed by HL to E", 0, 1, unimplementedHandler}
	OpcodeLdEa      = opcode{0x5F, "LD E,A", "Copy A to E", 0, 1, unimplementedHandler}
	OpcodeLdHb      = opcode{0x60, "LD H,B", "Copy B to H", 0, 1, unimplementedHandler}
	OpcodeLdHc      = opcode{0x61, "LD H,C", "Copy C to H", 0, 1, unimplementedHandler}
	OpcodeLdHd      = opcode{0x62, "LD H,D", "Copy D to H", 0, 1, unimplementedHandler}
	OpcodeLdHe      = opcode{0x63, "LD H,E", "Copy E to H", 0, 1, unimplementedHandler}
	OpcodeLdHh      = opcode{0x64, "LD H,H", "Copy H to H", 0, 1, unimplementedHandler}
	OpcodeLdHl      = opcode{0x65, "LD H,L", "Copy L to H", 0, 1, unimplementedHandler}
	OpcodeLdHhl     = opcode{0x66, "LD H,(HL)", "Copy value pointed by HL to H", 0, 1, unimplementedHandler}
	OpcodeLdHa      = opcode{0x67, "LD H,A", "Copy A to H", 0, 1, unimplementedHandler}
	OpcodeLdLb      = opcode{0x68, "LD L,B", "Copy B to L", 0, 1, unimplementedHandler}
	OpcodeLdLc      = opcode{0x69, "LD L,C", "Copy C to L", 0, 1, unimplementedHandler}
	OpcodeLdLd      = opcode{0x6A, "LD L,D", "Copy D to L", 0, 1, unimplementedHandler}
	OpcodeLdLe      = opcode{0x6B, "LD L,E", "Copy E to L", 0, 1, unimplementedHandler}
	OpcodeLdLh      = opcode{0x6C, "LD L,H", "Copy H to L", 0, 1, unimplementedHandler}
	OpcodeLdLl      = opcode{0x6D, "LD L,L", "Copy L to L", 0, 1, unimplementedHandler}
	OpcodeLdLhl     = opcode{0x6E, "LD L,(HL)", "Copy value pointed by HL to L", 0, 1, unimplementedHandler}
	OpcodeLdLa      = opcode{0x6F, "LD L,A", "Copy A to L", 0, 1, unimplementedHandler}
	OpcodeLdHlb     = opcode{0x70, "LD (HL),B", "Copy B to address pointed by HL", 0, 1, unimplementedHandler}
	OpcodeLdHlc     = opcode{0x71, "LD (HL),C", "Copy C to address pointed by HL", 0, 1, unimplementedHandler}
	OpcodeLdHld     = opcode{0x72, "LD (HL),D", "Copy D to address pointed by HL", 0, 1, unimplementedHandler}
	OpcodeLdHle     = opcode{0x73, "LD (HL),E", "Copy E to address pointed by HL", 0, 1, unimplementedHandler}
	OpcodeLdHlh     = opcode{0x74, "LD (HL),H", "Copy H to address pointed by HL", 0, 1, unimplementedHandler}
	OpcodeLdHll     = opcode{0x75, "LD (HL),L", "Copy L to address pointed by HL", 0, 1, unimplementedHandler}
	OpcodeHalt      = opcode{0x76, "HALT", "Halt processor", 0, 1, unimplementedHandler}
	OpcodeLdHla     = opcode{0x77, "LD (HL),A", "Copy A to address pointed by HL", 0, 1, unimplementedHandler}
	OpcodeLdAb      = opcode{0x78, "LD A,B", "Copy B to A", 0, 1, unimplementedHandler}
	OpcodeLdAc      = opcode{0x79, "LD A,C", "Copy C to A", 0, 1, unimplementedHandler}
	OpcodeLdAd      = opcode{0x7A, "LD A,D", "Copy D to A", 0, 1, unimplementedHandler}
	OpcodeLdAe      = opcode{0x7B, "LD A,E", "Copy E to A", 0, 1, unimplementedHandler}
	OpcodeLdAh      = opcode{0x7C, "LD A,H", "Copy H to A", 0, 1, unimplementedHandler}
	OpcodeLdAl      = opcode{0x7D, "LD A,L", "Copy L to A", 0, 1, unimplementedHandler}
	OpcodeLdAhl     = opcode{0x7E, "LD A,(HL)", "Copy value pointed by HL to A", 0, 1, unimplementedHandler}
	OpcodeLdAa      = opcode{0x7F, "LD A,A", "Copy A to A", 0, 1, unimplementedHandler}
	OpcodeAddAb     = opcode{0x80, "ADD A,B", "Add B to A", 0, 1, unimplementedHandler}
	OpcodeAddAc     = opcode{0x81, "ADD A,C", "Add C to A", 0, 1, unimplementedHandler}
	OpcodeAddAd     = opcode{0x82, "ADD A,D", "Add D to A", 0, 1, unimplementedHandler}
	OpcodeAddAe     = opcode{0x83, "ADD A,E", "Add E to A", 0, 1, unimplementedHandler}
	OpcodeAddAh     = opcode{0x84, "ADD A,H", "Add H to A", 0, 1, unimplementedHandler}
	OpcodeAddAl     = opcode{0x85, "ADD A,L", "Add L to A", 0, 1, unimplementedHandler}
	OpcodeAddAhl    = opcode{0x86, "ADD A,(HL)", "Add value pointed by HL to A", 0, 1, unimplementedHandler}
	OpcodeAddAa     = opcode{0x87, "ADD A,A", "Add A to A", 0, 1, unimplementedHandler}
	OpcodeAdcAb     = opcode{0x88, "ADC A,B", "Add B and carry flag to A", 0, 1, unimplementedHandler}
	OpcodeAdcAc     = opcode{0x89, "ADC A,C", "Add C and carry flag to A", 0, 1, unimplementedHandler}
	OpcodeAdcAd     = opcode{0x8A, "ADC A,D", "Add D and carry flag to A", 0, 1, unimplementedHandler}
	OpcodeAdcAe     = opcode{0x8B, "ADC A,E", "Add E and carry flag to A", 0, 1, unimplementedHandler}
	OpcodeAdcAh     = opcode{0x8C, "ADC A,H", "Add H and carry flag to A", 0, 1, unimplementedHandler}
	OpcodeAdcAl     = opcode{0x8D, "ADC A,L", "Add and carry flag L to A", 0, 1, unimplementedHandler}
	OpcodeAdcAhl    = opcode{0x8E, "ADC A,(HL)", "Add value pointed by HL and carry flag to A", 0, 1, unimplementedHandler}
	OpcodeAdcAa     = opcode{0x8F, "ADC A,A", "Add A and carry flag to A", 0, 1, unimplementedHandler}
	OpcodeSubAb     = opcode{0x90, "SUB A,B", "Subtract B from A", 0, 1, unimplementedHandler}
	OpcodeSubAc     = opcode{0x91, "SUB A,C", "Subtract C from A", 0, 1, unimplementedHandler}
	OpcodeSubAd     = opcode{0x92, "SUB A,D", "Subtract D from A", 0, 1, unimplementedHandler}
	OpcodeSubAe     = opcode{0x93, "SUB A,E", "Subtract E from A", 0, 1, unimplementedHandler}
	OpcodeSubAh     = opcode{0x94, "SUB A,H", "Subtract H from A", 0, 1, unimplementedHandler}
	OpcodeSubAl     = opcode{0x95, "SUB A,L", "Subtract L from A", 0, 1, unimplementedHandler}
	OpcodeSubAhl    = opcode{0x96, "SUB A,(HL)", "Subtract value pointed by HL from A", 0, 1, unimplementedHandler}
	OpcodeSubAa     = opcode{0x97, "SUB A,A", "Subtract A from A", 0, 1, unimplementedHandler}
	OpcodeSbcAb     = opcode{0x98, "SBC A,B", "Subtract B and carry flag from A", 0, 1, unimplementedHandler}
	OpcodeSbcAc     = opcode{0x99, "SBC A,C", "Subtract C and carry flag from A", 0, 1, unimplementedHandler}
	OpcodeSbcAd     = opcode{0x9A, "SBC A,D", "Subtract D and carry flag from A", 0, 1, unimplementedHandler}
	OpcodeSbcAe     = opcode{0x9B, "SBC A,E", "Subtract E and carry flag from A", 0, 1, unimplementedHandler}
	OpcodeSbcAh     = opcode{0x9C, "SBC A,H", "Subtract H and carry flag from A", 0, 1, unimplementedHandler}
	OpcodeSbcAl     = opcode{0x9D, "SBC A,L", "Subtract and carry flag L from A", 0, 1, unimplementedHandler}
	OpcodeSbcAhl    = opcode{0x9E, "SBC A,(HL)", "Subtract value pointed by HL and carry flag from A", 0, 1, unimplementedHandler}
	OpcodeSbcAa     = opcode{0x9F, "SBC A,A", "Subtract A and carry flag from A", 0, 1, unimplementedHandler}
	OpcodeAndB      = opcode{0xA0, "AND B", "Logical AND B against A", 0, 1, unimplementedHandler}
	OpcodeAndC      = opcode{0xA1, "AND C", "Logical AND C against A", 0, 1, unimplementedHandler}
	OpcodeAndD      = opcode{0xA2, "AND D", "Logical AND D against A", 0, 1, unimplementedHandler}
	OpcodeAndE      = opcode{0xA3, "AND E", "Logical AND E against A", 0, 1, unimplementedHandler}
	OpcodeAndH      = opcode{0xA4, "AND H", "Logical AND H against A", 0, 1, unimplementedHandler}
	OpcodeAndL      = opcode{0xA5, "AND L", "Logical AND L against A", 0, 1, unimplementedHandler}
	OpcodeAndHl     = opcode{0xA6, "AND (HL)", "Logical AND value pointed by HL against A", 0, 1, unimplementedHandler}
	OpcodeAndA      = opcode{0xA7, "AND A", "Logical AND A against A", 0, 1, unimplementedHandler}
	OpcodeXorB      = opcode{0xA8, "XOR B", "Logical XOR B against A", 0, 1, unimplementedHandler}
	OpcodeXorC      = opcode{0xA9, "XOR C", "Logical XOR C against A", 0, 1, unimplementedHandler}
	OpcodeXorD      = opcode{0xAA, "XOR D", "Logical XOR D against A", 0, 1, unimplementedHandler}
	OpcodeXorE      = opcode{0xAB, "XOR E", "Logical XOR E against A", 0, 1, unimplementedHandler}
	OpcodeXorH      = opcode{0xAC, "XOR H", "Logical XOR H against A", 0, 1, unimplementedHandler}
	OpcodeXorL      = opcode{0xAD, "XOR L", "Logical XOR L against A", 0, 1, unimplementedHandler}
	OpcodeXorHl     = opcode{0xAE, "XOR (HL)", "Logical XOR value pointed by HL against A", 0, 1, unimplementedHandler}
	OpcodeXorA      = opcode{0xAF, "XOR A", "Logical XOR A against A", 0, 1, unimplementedHandler}
	OpcodeOrB       = opcode{0xB0, "OR B", "Logical OR B against A", 0, 1, unimplementedHandler}
	OpcodeOrC       = opcode{0xB1, "OR C", "Logical OR C against A", 0, 1, unimplementedHandler}
	OpcodeOrD       = opcode{0xB2, "OR D", "Logical OR D against A", 0, 1, unimplementedHandler}
	OpcodeOrE       = opcode{0xB3, "OR E", "Logical OR E against A", 0, 1, unimplementedHandler}
	OpcodeOrH       = opcode{0xB4, "OR H", "Logical OR H against A", 0, 1, unimplementedHandler}
	OpcodeOrL       = opcode{0xB5, "OR L", "Logical OR L against A", 0, 1, unimplementedHandler}
	OpcodeOrHl      = opcode{0xB6, "OR (HL)", "Logical OR value pointed by HL against A", 0, 1, unimplementedHandler}
	OpcodeOrA       = opcode{0xB7, "OR A", "Logical OR A against A", 0, 1, unimplementedHandler}
	OpcodeCpB       = opcode{0xB8, "CP B", "Compare B against A", 0, 1, unimplementedHandler}
	OpcodeCpC       = opcode{0xB9, "CP C", "Compare C against A", 0, 1, unimplementedHandler}
	OpcodeCpD       = opcode{0xBA, "CP D", "Compare D against A", 0, 1, unimplementedHandler}
	OpcodeCpE       = opcode{0xBB, "CP E", "Compare E against A", 0, 1, unimplementedHandler}
	OpcodeCpH       = opcode{0xBC, "CP H", "Compare H against A", 0, 1, unimplementedHandler}
	OpcodeCpL       = opcode{0xBD, "CP L", "Compare L against A", 0, 1, unimplementedHandler}
	OpcodeCpHl      = opcode{0xBE, "CP (HL)", "Compare value pointed by HL against A", 0, 1, unimplementedHandler}
	OpcodeCpA       = opcode{0xBF, "CP A", "Compare A against A", 0, 1, unimplementedHandler}
	OpcodeRetNz     = opcode{0xC0, "RET NZ", "Return if last result was not zero", 0, 1, unimplementedHandler}
	OpcodePopBc     = opcode{0xC1, "POP BC", "Pop 16-bit value from stack into BC", 0, 1, unimplementedHandler}
	OpcodeJpNznn    = opcode{0xC2, "JP NZ,nn", "Absolute jump to 16-bit location if last result was not zero", 0, 1, unimplementedHandler}
	OpcodeJpNn      = opcode{0xC3, "JP nn", "Absolute jump to 16-bit location", 0, 1, unimplementedHandler}
	OpcodeCallNznn  = opcode{0xC4, "CALL NZ,nn", "Call routine at 16-bit location if last result was not zero", 0, 1, unimplementedHandler}
	OpcodePushBc    = opcode{0xC5, "PUSH BC", "Push 16-bit BC onto stack", 0, 1, unimplementedHandler}
	OpcodeAddAn     = opcode{0xC6, "ADD A,n", "Add 8-bit immediate to A", 0, 1, unimplementedHandler}
	OpcodeRst0      = opcode{0xC7, "RST 0", "Call routine at address 0000h", 0, 1, unimplementedHandler}
	OpcodeRetZ      = opcode{0xC8, "RET Z", "Return if last result was zero", 0, 1, unimplementedHandler}
	OpcodeRet       = opcode{0xC9, "RET", "Return to calling routine", 0, 1, unimplementedHandler}
	OpcodeJpZnn     = opcode{0xCA, "JP Z,nn", "Absolute jump to 16-bit location if last result was zero", 0, 1, unimplementedHandler}
	OpcodeExtOps    = opcode{0xCB, "Ext ops", "Extended operations (two-byte instruction code)", 0, 1, unimplementedHandler}
	OpcodeCallZnn   = opcode{0xCC, "CALL Z,nn", "Call routine at 16-bit location if last result was zero", 0, 1, unimplementedHandler}
	OpcodeCallNn    = opcode{0xCD, "CALL nn", "Call routine at 16-bit location", 0, 1, unimplementedHandler}
	OpcodeAdcAn     = opcode{0xCE, "ADC A,n", "Add 8-bit immediate and carry to A", 0, 1, unimplementedHandler}
	OpcodeRst8      = opcode{0xCF, "RST 8", "Call routine at address 0008h", 0, 1, unimplementedHandler}
	OpcodeRetNc     = opcode{0xD0, "RET NC", "Return if last result caused no carry", 0, 1, unimplementedHandler}
	OpcodePopDe     = opcode{0xD1, "POP DE", "Pop 16-bit value from stack into DE", 0, 1, unimplementedHandler}
	OpcodeJpNcnn    = opcode{0xD2, "JP NC,nn", "Absolute jump to 16-bit location if last result caused no carry", 0, 1, unimplementedHandler}
	OpcodeXxD3      = opcode{0xD3, "XX", "Operation removed in this CPU", 0, 1, unimplementedHandler}
	OpcodeCallNcnn  = opcode{0xD4, "CALL NC,nn", "Call routine at 16-bit location if last result caused no carry", 0, 1, unimplementedHandler}
	OpcodePushDe    = opcode{0xD5, "PUSH DE", "Push 16-bit DE onto stack", 0, 1, unimplementedHandler}
	OpcodeSubAn     = opcode{0xD6, "SUB A,n", "Subtract 8-bit immediate from A", 0, 1, unimplementedHandler}
	OpcodeRst10     = opcode{0xD7, "RST 10", "Call routine at address 0010h", 0, 1, unimplementedHandler}
	OpcodeRetC      = opcode{0xD8, "RET C", "Return if last result caused carry", 0, 1, unimplementedHandler}
	OpcodeReti      = opcode{0xD9, "RETI", "Enable interrupts and return to calling routine", 0, 1, unimplementedHandler}
	OpcodeJpCnn     = opcode{0xDA, "JP C,nn", "Absolute jump to 16-bit location if last result caused carry", 0, 1, unimplementedHandler}
	OpcodeXxDB      = opcode{0xDB, "XX", "Operation removed in this CPU", 0, 1, unimplementedHandler}
	OpcodeCallCnn   = opcode{0xDC, "CALL C,nn", "Call routine at 16-bit location if last result caused carry", 0, 1, unimplementedHandler}
	OpcodeXxDD      = opcode{0xDD, "XX", "Operation removed in this CPU", 0, 1, unimplementedHandler}
	OpcodeSbcAn     = opcode{0xDE, "SBC A,n", "Subtract 8-bit immediate and carry from A", 0, 1, unimplementedHandler}
	OpcodeRst18     = opcode{0xDF, "RST 18", "Call routine at address 0018h", 0, 1, unimplementedHandler}
	OpcodeLdhNa     = opcode{0xE0, "LDH (n),A", "Save A at address pointed to by (FF00h + 8-bit immediate)", 0, 1, unimplementedHandler}
	OpcodePopHl     = opcode{0xE1, "POP HL", "Pop 16-bit value from stack into HL", 0, 1, unimplementedHandler}
	OpcodeLdhCa     = opcode{0xE2, "LDH (C),A", "Save A at address pointed to by (FF00h + C)", 0, 1, unimplementedHandler}
	OpcodeXxE3      = opcode{0xE3, "XX", "Operation removed in this CPU", 0, 1, unimplementedHandler}
	OpcodeXxE4      = opcode{0xE4, "XX", "Operation removed in this CPU", 0, 1, unimplementedHandler}
	OpcodePushHl    = opcode{0xE5, "PUSH HL", "Push 16-bit HL onto stack", 0, 1, unimplementedHandler}
	OpcodeAndN      = opcode{0xE6, "AND n", "Logical AND 8-bit immediate against A", 0, 1, unimplementedHandler}
	OpcodeRst20     = opcode{0xE7, "RST 20", "Call routine at address 0020h", 0, 1, unimplementedHandler}
	OpcodeAddSpd    = opcode{0xE8, "ADD SP,d", "Add signed 8-bit immediate to SP", 0, 1, unimplementedHandler}
	OpcodeJpHl      = opcode{0xE9, "JP (HL)", "Jump to 16-bit value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeLdNna     = opcode{0xEA, "LD (nn),A", "Save A at given 16-bit address", 0, 1, unimplementedHandler}
	OpcodeXxEB      = opcode{0xEB, "XX", "Operation removed in this CPU", 0, 1, unimplementedHandler}
	OpcodeXxEC      = opcode{0xEC, "XX", "Operation removed in this CPU", 0, 1, unimplementedHandler}
	OpcodeXxED      = opcode{0xED, "XX", "Operation removed in this CPU", 0, 1, unimplementedHandler}
	OpcodeXorN      = opcode{0xEE, "XOR n", "Logical XOR 8-bit immediate against A", 0, 1, unimplementedHandler}
	OpcodeRst28     = opcode{0xEF, "RST 28", "Call routine at address 0028h", 0, 1, unimplementedHandler}
	OpcodeLdhAn     = opcode{0xF0, "LDH A,(n)", "Load A from address pointed to by (FF00h + 8-bit immediate)", 0, 1, unimplementedHandler}
	OpcodePopAf     = opcode{0xF1, "POP AF", "Pop 16-bit value from stack into AF", 0, 1, unimplementedHandler}
	OpcodeXxF2      = opcode{0xF2, "XX", "Operation removed in this CPU", 0, 1, unimplementedHandler}
	OpcodeDi        = opcode{0xF3, "DI", "DIsable interrupts", 0, 1, unimplementedHandler}
	OpcodeXxF4      = opcode{0xF4, "XX", "Operation removed in this CPU", 0, 1, unimplementedHandler}
	OpcodePushAf    = opcode{0xF5, "PUSH AF", "Push 16-bit AF onto stack", 0, 1, unimplementedHandler}
	OpcodeOrN       = opcode{0xF6, "OR n", "Logical OR 8-bit immediate against A", 0, 1, unimplementedHandler}
	OpcodeRst30     = opcode{0xF7, "RST 30", "Call routine at address 0030h", 0, 1, unimplementedHandler}
	OpcodeLdhlSpd   = opcode{0xF8, "LDHL SP,d", "Add signed 8-bit immediate to SP and save result in HL", 0, 1, unimplementedHandler}
	OpcodeLdSphl    = opcode{0xF9, "LD SP,HL", "Copy HL to SP", 0, 1, unimplementedHandler}
	OpcodeLdAnn     = opcode{0xFA, "LD A,(nn)", "Load A from given 16-bit address", 0, 1, unimplementedHandler}
	OpcodeEi        = opcode{0xFB, "EI", "Enable interrupts", 0, 1, unimplementedHandler}
	OpcodeXxFC      = opcode{0xFC, "XX", "Operation removed in this CPU", 0, 1, unimplementedHandler}
	OpcodeXxFD      = opcode{0xFD, "XX", "Operation removed in this CPU", 0, 1, unimplementedHandler}
	OpcodeCpN       = opcode{0xFE, "CP n", "Compare 8-bit immediate against A", 0, 1, unimplementedHandler}
	OpcodeRst38     = opcode{0xFF, "RST 38", "Call routine at address 0038h", 0, 1, unimplementedHandler}
)

var opcodeMap = map[uint8]opcode{
	0x00: OpcodeNop,
	0x01: OpcodeLdBcnn,
	0x02: OpcodeLdBca,
	0x03: OpcodeIncBc,
	0x04: OpcodeIncB,
	0x05: OpcodeDecB,
	0x06: OpcodeLdBn,
	0x07: OpcodeRlcA,
	0x08: OpcodeLdNnsp,
	0x09: OpcodeAddHlbc,
	0x0A: OpcodeLdAbc,
	0x0B: OpcodeDecBc,
	0x0C: OpcodeIncC,
	0x0D: OpcodeDecC,
	0x0E: OpcodeLdCn,
	0x0F: OpcodeRrcA,
	0x10: OpcodeStop,
	0x11: OpcodeLdDenn,
	0x12: OpcodeLdDea,
	0x13: OpcodeIncDe,
	0x14: OpcodeIncD,
	0x15: OpcodeDecD,
	0x16: OpcodeLdDn,
	0x17: OpcodeRlA,
	0x18: OpcodeJrN,
	0x19: OpcodeAddHlde,
	0x1A: OpcodeLdAde,
	0x1B: OpcodeDecDe,
	0x1C: OpcodeIncE,
	0x1D: OpcodeDecE,
	0x1E: OpcodeLdEn,
	0x1F: OpcodeRrA,
	0x20: OpcodeJrNzn,
	0x21: OpcodeLdHlnn,
	0x22: OpcodeLdiHla,
	0x23: OpcodeIncHl,
	0x24: OpcodeIncH,
	0x25: OpcodeDecH,
	0x26: OpcodeLdHn,
	0x27: OpcodeDaa,
	0x28: OpcodeJrZn,
	0x29: OpcodeAddHlhl,
	0x2A: OpcodeLdiAhl,
	0x2B: OpcodeDecHl,
	0x2C: OpcodeIncL,
	0x2D: OpcodeDecL,
	0x2E: OpcodeLdLn,
	0x2F: OpcodeCpl,
	0x30: OpcodeJrNcn,
	0x31: OpcodeLdSpnn,
	0x32: OpcodeLddHla,
	0x33: OpcodeIncSp,
	0x34: OpcodeIncHlAddr,
	0x35: OpcodeDecHlAddr,
	0x36: OpcodeLdHln,
	0x37: OpcodeScf,
	0x38: OpcodeJrCn,
	0x39: OpcodeAddHlsp,
	0x3A: OpcodeLddAhl,
	0x3B: OpcodeDecSp,
	0x3C: OpcodeIncA,
	0x3D: OpcodeDecA,
	0x3E: OpcodeLdAn,
	0x3F: OpcodeCcf,
	0x40: OpcodeLdBb,
	0x41: OpcodeLdBc,
	0x42: OpcodeLdBd,
	0x43: OpcodeLdBe,
	0x44: OpcodeLdBh,
	0x45: OpcodeLdBl,
	0x46: OpcodeLdBhl,
	0x47: OpcodeLdBa,
	0x48: OpcodeLdCb,
	0x49: OpcodeLdCc,
	0x4A: OpcodeLdCd,
	0x4B: OpcodeLdCe,
	0x4C: OpcodeLdCh,
	0x4D: OpcodeLdCl,
	0x4E: OpcodeLdChl,
	0x4F: OpcodeLdCa,
	0x50: OpcodeLdDb,
	0x51: OpcodeLdDc,
	0x52: OpcodeLdDd,
	0x53: OpcodeLdDe,
	0x54: OpcodeLdDh,
	0x55: OpcodeLdDl,
	0x56: OpcodeLdDhl,
	0x57: OpcodeLdDa,
	0x58: OpcodeLdEb,
	0x59: OpcodeLdEc,
	0x5A: OpcodeLdEd,
	0x5B: OpcodeLdEe,
	0x5C: OpcodeLdEh,
	0x5D: OpcodeLdEl,
	0x5E: OpcodeLdEhl,
	0x5F: OpcodeLdEa,
	0x60: OpcodeLdHb,
	0x61: OpcodeLdHc,
	0x62: OpcodeLdHd,
	0x63: OpcodeLdHe,
	0x64: OpcodeLdHh,
	0x65: OpcodeLdHl,
	0x66: OpcodeLdHhl,
	0x67: OpcodeLdHa,
	0x68: OpcodeLdLb,
	0x69: OpcodeLdLc,
	0x6A: OpcodeLdLd,
	0x6B: OpcodeLdLe,
	0x6C: OpcodeLdLh,
	0x6D: OpcodeLdLl,
	0x6E: OpcodeLdLhl,
	0x6F: OpcodeLdLa,
	0x70: OpcodeLdHlb,
	0x71: OpcodeLdHlc,
	0x72: OpcodeLdHld,
	0x73: OpcodeLdHle,
	0x74: OpcodeLdHlh,
	0x75: OpcodeLdHll,
	0x76: OpcodeHalt,
	0x77: OpcodeLdHla,
	0x78: OpcodeLdAb,
	0x79: OpcodeLdAc,
	0x7A: OpcodeLdAd,
	0x7B: OpcodeLdAe,
	0x7C: OpcodeLdAh,
	0x7D: OpcodeLdAl,
	0x7E: OpcodeLdAhl,
	0x7F: OpcodeLdAa,
	0x80: OpcodeAddAb,
	0x81: OpcodeAddAc,
	0x82: OpcodeAddAd,
	0x83: OpcodeAddAe,
	0x84: OpcodeAddAh,
	0x85: OpcodeAddAl,
	0x86: OpcodeAddAhl,
	0x87: OpcodeAddAa,
	0x88: OpcodeAdcAb,
	0x89: OpcodeAdcAc,
	0x8A: OpcodeAdcAd,
	0x8B: OpcodeAdcAe,
	0x8C: OpcodeAdcAh,
	0x8D: OpcodeAdcAl,
	0x8E: OpcodeAdcAhl,
	0x8F: OpcodeAdcAa,
	0x90: OpcodeSubAb,
	0x91: OpcodeSubAc,
	0x92: OpcodeSubAd,
	0x93: OpcodeSubAe,
	0x94: OpcodeSubAh,
	0x95: OpcodeSubAl,
	0x96: OpcodeSubAhl,
	0x97: OpcodeSubAa,
	0x98: OpcodeSbcAb,
	0x99: OpcodeSbcAc,
	0x9A: OpcodeSbcAd,
	0x9B: OpcodeSbcAe,
	0x9C: OpcodeSbcAh,
	0x9D: OpcodeSbcAl,
	0x9E: OpcodeSbcAhl,
	0x9F: OpcodeSbcAa,
	0xA0: OpcodeAndB,
	0xA1: OpcodeAndC,
	0xA2: OpcodeAndD,
	0xA3: OpcodeAndE,
	0xA4: OpcodeAndH,
	0xA5: OpcodeAndL,
	0xA6: OpcodeAndHl,
	0xA7: OpcodeAndA,
	0xA8: OpcodeXorB,
	0xA9: OpcodeXorC,
	0xAA: OpcodeXorD,
	0xAB: OpcodeXorE,
	0xAC: OpcodeXorH,
	0xAD: OpcodeXorL,
	0xAE: OpcodeXorHl,
	0xAF: OpcodeXorA,
	0xB0: OpcodeOrB,
	0xB1: OpcodeOrC,
	0xB2: OpcodeOrD,
	0xB3: OpcodeOrE,
	0xB4: OpcodeOrH,
	0xB5: OpcodeOrL,
	0xB6: OpcodeOrHl,
	0xB7: OpcodeOrA,
	0xB8: OpcodeCpB,
	0xB9: OpcodeCpC,
	0xBA: OpcodeCpD,
	0xBB: OpcodeCpE,
	0xBC: OpcodeCpH,
	0xBD: OpcodeCpL,
	0xBE: OpcodeCpHl,
	0xBF: OpcodeCpA,
	0xC0: OpcodeRetNz,
	0xC1: OpcodePopBc,
	0xC2: OpcodeJpNznn,
	0xC3: OpcodeJpNn,
	0xC4: OpcodeCallNznn,
	0xC5: OpcodePushBc,
	0xC6: OpcodeAddAn,
	0xC7: OpcodeRst0,
	0xC8: OpcodeRetZ,
	0xC9: OpcodeRet,
	0xCA: OpcodeJpZnn,
	0xCB: OpcodeExtOps,
	0xCC: OpcodeCallZnn,
	0xCD: OpcodeCallNn,
	0xCE: OpcodeAdcAn,
	0xCF: OpcodeRst8,
	0xD0: OpcodeRetNc,
	0xD1: OpcodePopDe,
	0xD2: OpcodeJpNcnn,
	0xD3: OpcodeXxD3,
	0xD4: OpcodeCallNcnn,
	0xD5: OpcodePushDe,
	0xD6: OpcodeSubAn,
	0xD7: OpcodeRst10,
	0xD8: OpcodeRetC,
	0xD9: OpcodeReti,
	0xDA: OpcodeJpCnn,
	0xDB: OpcodeXxDB,
	0xDC: OpcodeCallCnn,
	0xDD: OpcodeXxDD,
	0xDE: OpcodeSbcAn,
	0xDF: OpcodeRst18,
	0xE0: OpcodeLdhNa,
	0xE1: OpcodePopHl,
	0xE2: OpcodeLdhCa,
	0xE3: OpcodeXxE3,
	0xE4: OpcodeXxE4,
	0xE5: OpcodePushHl,
	0xE6: OpcodeAndN,
	0xE7: OpcodeRst20,
	0xE8: OpcodeAddSpd,
	0xE9: OpcodeJpHl,
	0xEA: OpcodeLdNna,
	0xEB: OpcodeXxEB,
	0xEC: OpcodeXxEC,
	0xED: OpcodeXxED,
	0xEE: OpcodeXorN,
	0xEF: OpcodeRst28,
	0xF0: OpcodeLdhAn,
	0xF1: OpcodePopAf,
	0xF2: OpcodeXxF2,
	0xF3: OpcodeDi,
	0xF4: OpcodeXxF4,
	0xF5: OpcodePushAf,
	0xF6: OpcodeOrN,
	0xF7: OpcodeRst30,
	0xF8: OpcodeLdhlSpd,
	0xF9: OpcodeLdSphl,
	0xFA: OpcodeLdAnn,
	0xFB: OpcodeEi,
	0xFC: OpcodeXxFC,
	0xFD: OpcodeXxFD,
	0xFE: OpcodeCpN,
	0xFF: OpcodeRst38,
}

var (
	OpcodeExtRlcB   = opcode{0x00, "RLC B", "Rotate B left with carry", 0, 1, unimplementedHandler}
	OpcodeExtRlcC   = opcode{0x01, "RLC C", "Rotate C left with carry", 0, 1, unimplementedHandler}
	OpcodeExtRlcD   = opcode{0x02, "RLC D", "Rotate D left with carry", 0, 1, unimplementedHandler}
	OpcodeExtRlcE   = opcode{0x03, "RLC E", "Rotate E left with carry", 0, 1, unimplementedHandler}
	OpcodeExtRlcH   = opcode{0x04, "RLC H", "Rotate H left with carry", 0, 1, unimplementedHandler}
	OpcodeExtRlcL   = opcode{0x05, "RLC L", "Rotate L left with carry", 0, 1, unimplementedHandler}
	OpcodeExtRlcHl  = opcode{0x06, "RLC (HL)", "Rotate value pointed by HL left with carry", 0, 1, unimplementedHandler}
	OpcodeExtRlcA   = opcode{0x07, "RLC A", "Rotate A left with carry", 0, 1, unimplementedHandler}
	OpcodeExtRrcB   = opcode{0x08, "RRC B", "Rotate B right with carry", 0, 1, unimplementedHandler}
	OpcodeExtRrcC   = opcode{0x09, "RRC C", "Rotate C right with carry", 0, 1, unimplementedHandler}
	OpcodeExtRrcD   = opcode{0x0A, "RRC D", "Rotate D right with carry", 0, 1, unimplementedHandler}
	OpcodeExtRrcE   = opcode{0x0B, "RRC E", "Rotate E right with carry", 0, 1, unimplementedHandler}
	OpcodeExtRrcH   = opcode{0x0C, "RRC H", "Rotate H right with carry", 0, 1, unimplementedHandler}
	OpcodeExtRrcL   = opcode{0x0D, "RRC L", "Rotate L right with carry", 0, 1, unimplementedHandler}
	OpcodeExtRrcHl  = opcode{0x0E, "RRC (HL)", "Rotate value pointed by HL right with carry", 0, 1, unimplementedHandler}
	OpcodeExtRrcA   = opcode{0x0F, "RRC A", "Rotate A right with carry", 0, 1, unimplementedHandler}
	OpcodeExtRlB    = opcode{0x10, "RL B", "Rotate B left", 0, 1, unimplementedHandler}
	OpcodeExtRlC    = opcode{0x11, "RL C", "Rotate C left", 0, 1, unimplementedHandler}
	OpcodeExtRlD    = opcode{0x12, "RL D", "Rotate D left", 0, 1, unimplementedHandler}
	OpcodeExtRlE    = opcode{0x13, "RL E", "Rotate E left", 0, 1, unimplementedHandler}
	OpcodeExtRlH    = opcode{0x14, "RL H", "Rotate H left", 0, 1, unimplementedHandler}
	OpcodeExtRlL    = opcode{0x15, "RL L", "Rotate L left", 0, 1, unimplementedHandler}
	OpcodeExtRlHl   = opcode{0x16, "RL (HL)", "Rotate value pointed by HL left", 0, 1, unimplementedHandler}
	OpcodeExtRlA    = opcode{0x17, "RL A", "Rotate A left", 0, 1, unimplementedHandler}
	OpcodeExtRrB    = opcode{0x18, "RR B", "Rotate B right", 0, 1, unimplementedHandler}
	OpcodeExtRrC    = opcode{0x19, "RR C", "Rotate C right", 0, 1, unimplementedHandler}
	OpcodeExtRrD    = opcode{0x1A, "RR D", "Rotate D right", 0, 1, unimplementedHandler}
	OpcodeExtRrE    = opcode{0x1B, "RR E", "Rotate E right", 0, 1, unimplementedHandler}
	OpcodeExtRrH    = opcode{0x1C, "RR H", "Rotate H right", 0, 1, unimplementedHandler}
	OpcodeExtRrL    = opcode{0x1D, "RR L", "Rotate L right", 0, 1, unimplementedHandler}
	OpcodeExtRrHl   = opcode{0x1E, "RR (HL)", "Rotate value pointed by HL right", 0, 1, unimplementedHandler}
	OpcodeExtRrA    = opcode{0x1F, "RR A", "Rotate A right", 0, 1, unimplementedHandler}
	OpcodeExtSlaB   = opcode{0x20, "SLA B", "Shift B left preserving sign", 0, 1, unimplementedHandler}
	OpcodeExtSlaC   = opcode{0x21, "SLA C", "Shift C left preserving sign", 0, 1, unimplementedHandler}
	OpcodeExtSlaD   = opcode{0x22, "SLA D", "Shift D left preserving sign", 0, 1, unimplementedHandler}
	OpcodeExtSlaE   = opcode{0x23, "SLA E", "Shift E left preserving sign", 0, 1, unimplementedHandler}
	OpcodeExtSlaH   = opcode{0x24, "SLA H", "Shift H left preserving sign", 0, 1, unimplementedHandler}
	OpcodeExtSlaL   = opcode{0x25, "SLA L", "Shift L left preserving sign", 0, 1, unimplementedHandler}
	OpcodeExtSlaHl  = opcode{0x26, "SLA (HL)", "Shift value pointed by HL left preserving sign", 0, 1, unimplementedHandler}
	OpcodeExtSlaA   = opcode{0x27, "SLA A", "Shift A left preserving sign", 0, 1, unimplementedHandler}
	OpcodeExtSraB   = opcode{0x28, "SRA B", "Shift B right preserving sign", 0, 1, unimplementedHandler}
	OpcodeExtSraC   = opcode{0x29, "SRA C", "Shift C right preserving sign", 0, 1, unimplementedHandler}
	OpcodeExtSraD   = opcode{0x2A, "SRA D", "Shift D right preserving sign", 0, 1, unimplementedHandler}
	OpcodeExtSraE   = opcode{0x2B, "SRA E", "Shift E right preserving sign", 0, 1, unimplementedHandler}
	OpcodeExtSraH   = opcode{0x2C, "SRA H", "Shift H right preserving sign", 0, 1, unimplementedHandler}
	OpcodeExtSraL   = opcode{0x2D, "SRA L", "Shift L right preserving sign", 0, 1, unimplementedHandler}
	OpcodeExtSraHl  = opcode{0x2E, "SRA (HL)", "Shift value pointed by HL right preserving sign", 0, 1, unimplementedHandler}
	OpcodeExtSraA   = opcode{0x2F, "SRA A", "Shift A right preserving sign", 0, 1, unimplementedHandler}
	OpcodeExtSwapB  = opcode{0x30, "SWAP B", "Swap nybbles in B", 0, 1, unimplementedHandler}
	OpcodeExtSwapC  = opcode{0x31, "SWAP C", "Swap nybbles in C", 0, 1, unimplementedHandler}
	OpcodeExtSwapD  = opcode{0x32, "SWAP D", "Swap nybbles in D", 0, 1, unimplementedHandler}
	OpcodeExtSwapE  = opcode{0x33, "SWAP E", "Swap nybbles in E", 0, 1, unimplementedHandler}
	OpcodeExtSwapH  = opcode{0x34, "SWAP H", "Swap nybbles in H", 0, 1, unimplementedHandler}
	OpcodeExtSwapL  = opcode{0x35, "SWAP L", "Swap nybbles in L", 0, 1, unimplementedHandler}
	OpcodeExtSwapHl = opcode{0x36, "SWAP (HL)", "Swap nybbles in value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtSwapA  = opcode{0x37, "SWAP A", "Swap nybbles in A", 0, 1, unimplementedHandler}
	OpcodeExtSrlB   = opcode{0x38, "SRL B", "Shift B right", 0, 1, unimplementedHandler}
	OpcodeExtSrlC   = opcode{0x39, "SRL C", "Shift C right", 0, 1, unimplementedHandler}
	OpcodeExtSrlD   = opcode{0x3A, "SRL D", "Shift D right", 0, 1, unimplementedHandler}
	OpcodeExtSrlE   = opcode{0x3B, "SRL E", "Shift E right", 0, 1, unimplementedHandler}
	OpcodeExtSrlH   = opcode{0x3C, "SRL H", "Shift H right", 0, 1, unimplementedHandler}
	OpcodeExtSrlL   = opcode{0x3D, "SRL L", "Shift L right", 0, 1, unimplementedHandler}
	OpcodeExtSrlHl  = opcode{0x3E, "SRL (HL)", "Shift value pointed by HL right", 0, 1, unimplementedHandler}
	OpcodeExtSrlA   = opcode{0x3F, "SRL A", "Shift A right", 0, 1, unimplementedHandler}
	OpcodeExtBit0b  = opcode{0x40, "BIT 0,B", "Test bit 0 of B", 0, 1, unimplementedHandler}
	OpcodeExtBit0c  = opcode{0x41, "BIT 0,C", "Test bit 0 of C", 0, 1, unimplementedHandler}
	OpcodeExtBit0d  = opcode{0x42, "BIT 0,D", "Test bit 0 of D", 0, 1, unimplementedHandler}
	OpcodeExtBit0e  = opcode{0x43, "BIT 0,E", "Test bit 0 of E", 0, 1, unimplementedHandler}
	OpcodeExtBit0h  = opcode{0x44, "BIT 0,H", "Test bit 0 of H", 0, 1, unimplementedHandler}
	OpcodeExtBit0l  = opcode{0x45, "BIT 0,L", "Test bit 0 of L", 0, 1, unimplementedHandler}
	OpcodeExtBit0hl = opcode{0x46, "BIT 0,(HL)", "Test bit 0 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtBit0a  = opcode{0x47, "BIT 0,A", "Test bit 0 of A", 0, 1, unimplementedHandler}
	OpcodeExtBit1b  = opcode{0x48, "BIT 1,B", "Test bit 1 of B", 0, 1, unimplementedHandler}
	OpcodeExtBit1c  = opcode{0x49, "BIT 1,C", "Test bit 1 of C", 0, 1, unimplementedHandler}
	OpcodeExtBit1d  = opcode{0x4A, "BIT 1,D", "Test bit 1 of D", 0, 1, unimplementedHandler}
	OpcodeExtBit1e  = opcode{0x4B, "BIT 1,E", "Test bit 1 of E", 0, 1, unimplementedHandler}
	OpcodeExtBit1h  = opcode{0x4C, "BIT 1,H", "Test bit 1 of H", 0, 1, unimplementedHandler}
	OpcodeExtBit1l  = opcode{0x4D, "BIT 1,L", "Test bit 1 of L", 0, 1, unimplementedHandler}
	OpcodeExtBit1hl = opcode{0x4E, "BIT 1,(HL)", "Test bit 1 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtBit1a  = opcode{0x4F, "BIT 1,A", "Test bit 1 of A", 0, 1, unimplementedHandler}
	OpcodeExtBit2b  = opcode{0x50, "BIT 2,B", "Test bit 2 of B", 0, 1, unimplementedHandler}
	OpcodeExtBit2c  = opcode{0x51, "BIT 2,C", "Test bit 2 of C", 0, 1, unimplementedHandler}
	OpcodeExtBit2d  = opcode{0x52, "BIT 2,D", "Test bit 2 of D", 0, 1, unimplementedHandler}
	OpcodeExtBit2e  = opcode{0x53, "BIT 2,E", "Test bit 2 of E", 0, 1, unimplementedHandler}
	OpcodeExtBit2h  = opcode{0x54, "BIT 2,H", "Test bit 2 of H", 0, 1, unimplementedHandler}
	OpcodeExtBit2l  = opcode{0x55, "BIT 2,L", "Test bit 2 of L", 0, 1, unimplementedHandler}
	OpcodeExtBit2hl = opcode{0x56, "BIT 2,(HL)", "Test bit 2 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtBit2a  = opcode{0x57, "BIT 2,A", "Test bit 2 of A", 0, 1, unimplementedHandler}
	OpcodeExtBit3b  = opcode{0x58, "BIT 3,B", "Test bit 3 of B", 0, 1, unimplementedHandler}
	OpcodeExtBit3c  = opcode{0x59, "BIT 3,C", "Test bit 3 of C", 0, 1, unimplementedHandler}
	OpcodeExtBit3d  = opcode{0x5A, "BIT 3,D", "Test bit 3 of D", 0, 1, unimplementedHandler}
	OpcodeExtBit3e  = opcode{0x5B, "BIT 3,E", "Test bit 3 of E", 0, 1, unimplementedHandler}
	OpcodeExtBit3h  = opcode{0x5C, "BIT 3,H", "Test bit 3 of H", 0, 1, unimplementedHandler}
	OpcodeExtBit3l  = opcode{0x5D, "BIT 3,L", "Test bit 3 of L", 0, 1, unimplementedHandler}
	OpcodeExtBit3hl = opcode{0x5E, "BIT 3,(HL)", "Test bit 3 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtBit3a  = opcode{0x5F, "BIT 3,A", "Test bit 3 of A", 0, 1, unimplementedHandler}
	OpcodeExtBit4b  = opcode{0x60, "BIT 4,B", "Test bit 4 of B", 0, 1, unimplementedHandler}
	OpcodeExtBit4c  = opcode{0x61, "BIT 4,C", "Test bit 4 of C", 0, 1, unimplementedHandler}
	OpcodeExtBit4d  = opcode{0x62, "BIT 4,D", "Test bit 4 of D", 0, 1, unimplementedHandler}
	OpcodeExtBit4e  = opcode{0x63, "BIT 4,E", "Test bit 4 of E", 0, 1, unimplementedHandler}
	OpcodeExtBit4h  = opcode{0x64, "BIT 4,H", "Test bit 4 of H", 0, 1, unimplementedHandler}
	OpcodeExtBit4l  = opcode{0x65, "BIT 4,L", "Test bit 4 of L", 0, 1, unimplementedHandler}
	OpcodeExtBit4hl = opcode{0x66, "BIT 4,(HL)", "Test bit 4 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtBit4a  = opcode{0x67, "BIT 4,A", "Test bit 4 of A", 0, 1, unimplementedHandler}
	OpcodeExtBit5b  = opcode{0x68, "BIT 5,B", "Test bit 5 of B", 0, 1, unimplementedHandler}
	OpcodeExtBit5c  = opcode{0x69, "BIT 5,C", "Test bit 5 of C", 0, 1, unimplementedHandler}
	OpcodeExtBit5d  = opcode{0x6A, "BIT 5,D", "Test bit 5 of D", 0, 1, unimplementedHandler}
	OpcodeExtBit5e  = opcode{0x6B, "BIT 5,E", "Test bit 5 of E", 0, 1, unimplementedHandler}
	OpcodeExtBit5h  = opcode{0x6C, "BIT 5,H", "Test bit 5 of H", 0, 1, unimplementedHandler}
	OpcodeExtBit5l  = opcode{0x6D, "BIT 5,L", "Test bit 5 of L", 0, 1, unimplementedHandler}
	OpcodeExtBit5hl = opcode{0x6E, "BIT 5,(HL)", "Test bit 5 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtBit5a  = opcode{0x6F, "BIT 5,A", "Test bit 5 of A", 0, 1, unimplementedHandler}
	OpcodeExtBit6b  = opcode{0x70, "BIT 6,B", "Test bit 6 of B", 0, 1, unimplementedHandler}
	OpcodeExtBit6c  = opcode{0x71, "BIT 6,C", "Test bit 6 of C", 0, 1, unimplementedHandler}
	OpcodeExtBit6d  = opcode{0x72, "BIT 6,D", "Test bit 6 of D", 0, 1, unimplementedHandler}
	OpcodeExtBit6e  = opcode{0x73, "BIT 6,E", "Test bit 6 of E", 0, 1, unimplementedHandler}
	OpcodeExtBit6h  = opcode{0x74, "BIT 6,H", "Test bit 6 of H", 0, 1, unimplementedHandler}
	OpcodeExtBit6l  = opcode{0x75, "BIT 6,L", "Test bit 6 of L", 0, 1, unimplementedHandler}
	OpcodeExtBit6hl = opcode{0x76, "BIT 6,(HL)", "Test bit 6 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtBit6a  = opcode{0x77, "BIT 6,A", "Test bit 6 of A", 0, 1, unimplementedHandler}
	OpcodeExtBit7b  = opcode{0x78, "BIT 7,B", "Test bit 7 of B", 0, 1, unimplementedHandler}
	OpcodeExtBit7c  = opcode{0x79, "BIT 7,C", "Test bit 7 of C", 0, 1, unimplementedHandler}
	OpcodeExtBit7d  = opcode{0x7A, "BIT 7,D", "Test bit 7 of D", 0, 1, unimplementedHandler}
	OpcodeExtBit7e  = opcode{0x7B, "BIT 7,E", "Test bit 7 of E", 0, 1, unimplementedHandler}
	OpcodeExtBit7h  = opcode{0x7C, "BIT 7,H", "Test bit 7 of H", 0, 1, unimplementedHandler}
	OpcodeExtBit7l  = opcode{0x7D, "BIT 7,L", "Test bit 7 of L", 0, 1, unimplementedHandler}
	OpcodeExtBit7hl = opcode{0x7E, "BIT 7,(HL)", "Test bit 7 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtBit7a  = opcode{0x7F, "BIT 7,A", "Test bit 7 of A", 0, 1, unimplementedHandler}
	OpcodeExtRes0b  = opcode{0x80, "RES 0,B", "Clear (reset) bit 0 of B", 0, 1, unimplementedHandler}
	OpcodeExtRes0c  = opcode{0x81, "RES 0,C", "Clear (reset) bit 0 of C", 0, 1, unimplementedHandler}
	OpcodeExtRes0d  = opcode{0x82, "RES 0,D", "Clear (reset) bit 0 of D", 0, 1, unimplementedHandler}
	OpcodeExtRes0e  = opcode{0x83, "RES 0,E", "Clear (reset) bit 0 of E", 0, 1, unimplementedHandler}
	OpcodeExtRes0h  = opcode{0x84, "RES 0,H", "Clear (reset) bit 0 of H", 0, 1, unimplementedHandler}
	OpcodeExtRes0l  = opcode{0x85, "RES 0,L", "Clear (reset) bit 0 of L", 0, 1, unimplementedHandler}
	OpcodeExtRes0hl = opcode{0x86, "RES 0,(HL)", "Clear (reset) bit 0 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtRes0a  = opcode{0x87, "RES 0,A", "Clear (reset) bit 0 of A", 0, 1, unimplementedHandler}
	OpcodeExtRes1b  = opcode{0x88, "RES 1,B", "Clear (reset) bit 1 of B", 0, 1, unimplementedHandler}
	OpcodeExtRes1c  = opcode{0x89, "RES 1,C", "Clear (reset) bit 1 of C", 0, 1, unimplementedHandler}
	OpcodeExtRes1d  = opcode{0x8A, "RES 1,D", "Clear (reset) bit 1 of D", 0, 1, unimplementedHandler}
	OpcodeExtRes1e  = opcode{0x8B, "RES 1,E", "Clear (reset) bit 1 of E", 0, 1, unimplementedHandler}
	OpcodeExtRes1h  = opcode{0x8C, "RES 1,H", "Clear (reset) bit 1 of H", 0, 1, unimplementedHandler}
	OpcodeExtRes1l  = opcode{0x8D, "RES 1,L", "Clear (reset) bit 1 of L", 0, 1, unimplementedHandler}
	OpcodeExtRes1hl = opcode{0x8E, "RES 1,(HL)", "Clear (reset) bit 1 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtRes1a  = opcode{0x8F, "RES 1,A", "Clear (reset) bit 1 of A", 0, 1, unimplementedHandler}
	OpcodeExtRes2b  = opcode{0x90, "RES 2,B", "Clear (reset) bit 2 of B", 0, 1, unimplementedHandler}
	OpcodeExtRes2c  = opcode{0x91, "RES 2,C", "Clear (reset) bit 2 of C", 0, 1, unimplementedHandler}
	OpcodeExtRes2d  = opcode{0x92, "RES 2,D", "Clear (reset) bit 2 of D", 0, 1, unimplementedHandler}
	OpcodeExtRes2e  = opcode{0x93, "RES 2,E", "Clear (reset) bit 2 of E", 0, 1, unimplementedHandler}
	OpcodeExtRes2h  = opcode{0x94, "RES 2,H", "Clear (reset) bit 2 of H", 0, 1, unimplementedHandler}
	OpcodeExtRes2l  = opcode{0x95, "RES 2,L", "Clear (reset) bit 2 of L", 0, 1, unimplementedHandler}
	OpcodeExtRes2hl = opcode{0x96, "RES 2,(HL)", "Clear (reset) bit 2 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtRes2a  = opcode{0x97, "RES 2,A", "Clear (reset) bit 2 of A", 0, 1, unimplementedHandler}
	OpcodeExtRes3b  = opcode{0x98, "RES 3,B", "Clear (reset) bit 3 of B", 0, 1, unimplementedHandler}
	OpcodeExtRes3c  = opcode{0x99, "RES 3,C", "Clear (reset) bit 3 of C", 0, 1, unimplementedHandler}
	OpcodeExtRes3d  = opcode{0x9A, "RES 3,D", "Clear (reset) bit 3 of D", 0, 1, unimplementedHandler}
	OpcodeExtRes3e  = opcode{0x9B, "RES 3,E", "Clear (reset) bit 3 of E", 0, 1, unimplementedHandler}
	OpcodeExtRes3h  = opcode{0x9C, "RES 3,H", "Clear (reset) bit 3 of H", 0, 1, unimplementedHandler}
	OpcodeExtRes3l  = opcode{0x9D, "RES 3,L", "Clear (reset) bit 3 of L", 0, 1, unimplementedHandler}
	OpcodeExtRes3hl = opcode{0x9E, "RES 3,(HL)", "Clear (reset) bit 3 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtRes3a  = opcode{0x9F, "RES 3,A", "Clear (reset) bit 3 of A", 0, 1, unimplementedHandler}
	OpcodeExtRes4b  = opcode{0xA0, "RES 4,B", "Clear (reset) bit 4 of B", 0, 1, unimplementedHandler}
	OpcodeExtRes4c  = opcode{0xA1, "RES 4,C", "Clear (reset) bit 4 of C", 0, 1, unimplementedHandler}
	OpcodeExtRes4d  = opcode{0xA2, "RES 4,D", "Clear (reset) bit 4 of D", 0, 1, unimplementedHandler}
	OpcodeExtRes4e  = opcode{0xA3, "RES 4,E", "Clear (reset) bit 4 of E", 0, 1, unimplementedHandler}
	OpcodeExtRes4h  = opcode{0xA4, "RES 4,H", "Clear (reset) bit 4 of H", 0, 1, unimplementedHandler}
	OpcodeExtRes4l  = opcode{0xA5, "RES 4,L", "Clear (reset) bit 4 of L", 0, 1, unimplementedHandler}
	OpcodeExtRes4hl = opcode{0xA6, "RES 4,(HL)", "Clear (reset) bit 4 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtRes4a  = opcode{0xA7, "RES 4,A", "Clear (reset) bit 4 of A", 0, 1, unimplementedHandler}
	OpcodeExtRes5b  = opcode{0xA8, "RES 5,B", "Clear (reset) bit 5 of B", 0, 1, unimplementedHandler}
	OpcodeExtRes5c  = opcode{0xA9, "RES 5,C", "Clear (reset) bit 5 of C", 0, 1, unimplementedHandler}
	OpcodeExtRes5d  = opcode{0xAA, "RES 5,D", "Clear (reset) bit 5 of D", 0, 1, unimplementedHandler}
	OpcodeExtRes5e  = opcode{0xAB, "RES 5,E", "Clear (reset) bit 5 of E", 0, 1, unimplementedHandler}
	OpcodeExtRes5h  = opcode{0xAC, "RES 5,H", "Clear (reset) bit 5 of H", 0, 1, unimplementedHandler}
	OpcodeExtRes5l  = opcode{0xAD, "RES 5,L", "Clear (reset) bit 5 of L", 0, 1, unimplementedHandler}
	OpcodeExtRes5hl = opcode{0xAE, "RES 5,(HL)", "Clear (reset) bit 5 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtRes5a  = opcode{0xAF, "RES 5,A", "Clear (reset) bit 5 of A", 0, 1, unimplementedHandler}
	OpcodeExtRes6b  = opcode{0xB0, "RES 6,B", "Clear (reset) bit 6 of B", 0, 1, unimplementedHandler}
	OpcodeExtRes6c  = opcode{0xB1, "RES 6,C", "Clear (reset) bit 6 of C", 0, 1, unimplementedHandler}
	OpcodeExtRes6d  = opcode{0xB2, "RES 6,D", "Clear (reset) bit 6 of D", 0, 1, unimplementedHandler}
	OpcodeExtRes6e  = opcode{0xB3, "RES 6,E", "Clear (reset) bit 6 of E", 0, 1, unimplementedHandler}
	OpcodeExtRes6h  = opcode{0xB4, "RES 6,H", "Clear (reset) bit 6 of H", 0, 1, unimplementedHandler}
	OpcodeExtRes6l  = opcode{0xB5, "RES 6,L", "Clear (reset) bit 6 of L", 0, 1, unimplementedHandler}
	OpcodeExtRes6hl = opcode{0xB6, "RES 6,(HL)", "Clear (reset) bit 6 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtRes6a  = opcode{0xB7, "RES 6,A", "Clear (reset) bit 6 of A", 0, 1, unimplementedHandler}
	OpcodeExtRes7b  = opcode{0xB8, "RES 7,B", "Clear (reset) bit 7 of B", 0, 1, unimplementedHandler}
	OpcodeExtRes7c  = opcode{0xB9, "RES 7,C", "Clear (reset) bit 7 of C", 0, 1, unimplementedHandler}
	OpcodeExtRes7d  = opcode{0xBA, "RES 7,D", "Clear (reset) bit 7 of D", 0, 1, unimplementedHandler}
	OpcodeExtRes7e  = opcode{0xBB, "RES 7,E", "Clear (reset) bit 7 of E", 0, 1, unimplementedHandler}
	OpcodeExtRes7h  = opcode{0xBC, "RES 7,H", "Clear (reset) bit 7 of H", 0, 1, unimplementedHandler}
	OpcodeExtRes7l  = opcode{0xBD, "RES 7,L", "Clear (reset) bit 7 of L", 0, 1, unimplementedHandler}
	OpcodeExtRes7hl = opcode{0xBE, "RES 7,(HL)", "Clear (reset) bit 7 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtRes7a  = opcode{0xBF, "RES 7,A", "Clear (reset) bit 7 of A", 0, 1, unimplementedHandler}
	OpcodeExtSet0b  = opcode{0xC0, "SET 0,B", "Set bit 0 of B", 0, 1, unimplementedHandler}
	OpcodeExtSet0c  = opcode{0xC1, "SET 0,C", "Set bit 0 of C", 0, 1, unimplementedHandler}
	OpcodeExtSet0d  = opcode{0xC2, "SET 0,D", "Set bit 0 of D", 0, 1, unimplementedHandler}
	OpcodeExtSet0e  = opcode{0xC3, "SET 0,E", "Set bit 0 of E", 0, 1, unimplementedHandler}
	OpcodeExtSet0h  = opcode{0xC4, "SET 0,H", "Set bit 0 of H", 0, 1, unimplementedHandler}
	OpcodeExtSet0l  = opcode{0xC5, "SET 0,L", "Set bit 0 of L", 0, 1, unimplementedHandler}
	OpcodeExtSet0hl = opcode{0xC6, "SET 0,(HL)", "Set bit 0 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtSet0a  = opcode{0xC7, "SET 0,A", "Set bit 0 of A", 0, 1, unimplementedHandler}
	OpcodeExtSet1b  = opcode{0xC8, "SET 1,B", "Set bit 1 of B", 0, 1, unimplementedHandler}
	OpcodeExtSet1c  = opcode{0xC9, "SET 1,C", "Set bit 1 of C", 0, 1, unimplementedHandler}
	OpcodeExtSet1d  = opcode{0xCA, "SET 1,D", "Set bit 1 of D", 0, 1, unimplementedHandler}
	OpcodeExtSet1e  = opcode{0xCB, "SET 1,E", "Set bit 1 of E", 0, 1, unimplementedHandler}
	OpcodeExtSet1h  = opcode{0xCC, "SET 1,H", "Set bit 1 of H", 0, 1, unimplementedHandler}
	OpcodeExtSet1l  = opcode{0xCD, "SET 1,L", "Set bit 1 of L", 0, 1, unimplementedHandler}
	OpcodeExtSet1hl = opcode{0xCE, "SET 1,(HL)", "Set bit 1 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtSet1a  = opcode{0xCF, "SET 1,A", "Set bit 1 of A", 0, 1, unimplementedHandler}
	OpcodeExtSet2b  = opcode{0xD0, "SET 2,B", "Set bit 2 of B", 0, 1, unimplementedHandler}
	OpcodeExtSet2c  = opcode{0xD1, "SET 2,C", "Set bit 2 of C", 0, 1, unimplementedHandler}
	OpcodeExtSet2d  = opcode{0xD2, "SET 2,D", "Set bit 2 of D", 0, 1, unimplementedHandler}
	OpcodeExtSet2e  = opcode{0xD3, "SET 2,E", "Set bit 2 of E", 0, 1, unimplementedHandler}
	OpcodeExtSet2h  = opcode{0xD4, "SET 2,H", "Set bit 2 of H", 0, 1, unimplementedHandler}
	OpcodeExtSet2l  = opcode{0xD5, "SET 2,L", "Set bit 2 of L", 0, 1, unimplementedHandler}
	OpcodeExtSet2hl = opcode{0xD6, "SET 2,(HL)", "Set bit 2 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtSet2a  = opcode{0xD7, "SET 2,A", "Set bit 2 of A", 0, 1, unimplementedHandler}
	OpcodeExtSet3b  = opcode{0xD8, "SET 3,B", "Set bit 3 of B", 0, 1, unimplementedHandler}
	OpcodeExtSet3c  = opcode{0xD9, "SET 3,C", "Set bit 3 of C", 0, 1, unimplementedHandler}
	OpcodeExtSet3d  = opcode{0xDA, "SET 3,D", "Set bit 3 of D", 0, 1, unimplementedHandler}
	OpcodeExtSet3e  = opcode{0xDB, "SET 3,E", "Set bit 3 of E", 0, 1, unimplementedHandler}
	OpcodeExtSet3h  = opcode{0xDC, "SET 3,H", "Set bit 3 of H", 0, 1, unimplementedHandler}
	OpcodeExtSet3l  = opcode{0xDD, "SET 3,L", "Set bit 3 of L", 0, 1, unimplementedHandler}
	OpcodeExtSet3hl = opcode{0xDE, "SET 3,(HL)", "Set bit 3 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtSet3a  = opcode{0xDF, "SET 3,A", "Set bit 3 of A", 0, 1, unimplementedHandler}
	OpcodeExtSet4b  = opcode{0xE0, "SET 4,B", "Set bit 4 of B", 0, 1, unimplementedHandler}
	OpcodeExtSet4c  = opcode{0xE1, "SET 4,C", "Set bit 4 of C", 0, 1, unimplementedHandler}
	OpcodeExtSet4d  = opcode{0xE2, "SET 4,D", "Set bit 4 of D", 0, 1, unimplementedHandler}
	OpcodeExtSet4e  = opcode{0xE3, "SET 4,E", "Set bit 4 of E", 0, 1, unimplementedHandler}
	OpcodeExtSet4h  = opcode{0xE4, "SET 4,H", "Set bit 4 of H", 0, 1, unimplementedHandler}
	OpcodeExtSet4l  = opcode{0xE5, "SET 4,L", "Set bit 4 of L", 0, 1, unimplementedHandler}
	OpcodeExtSet4hl = opcode{0xE6, "SET 4,(HL)", "Set bit 4 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtSet4a  = opcode{0xE7, "SET 4,A", "Set bit 4 of A", 0, 1, unimplementedHandler}
	OpcodeExtSet5b  = opcode{0xE8, "SET 5,B", "Set bit 5 of B", 0, 1, unimplementedHandler}
	OpcodeExtSet5c  = opcode{0xE9, "SET 5,C", "Set bit 5 of C", 0, 1, unimplementedHandler}
	OpcodeExtSet5d  = opcode{0xEA, "SET 5,D", "Set bit 5 of D", 0, 1, unimplementedHandler}
	OpcodeExtSet5e  = opcode{0xEB, "SET 5,E", "Set bit 5 of E", 0, 1, unimplementedHandler}
	OpcodeExtSet5h  = opcode{0xEC, "SET 5,H", "Set bit 5 of H", 0, 1, unimplementedHandler}
	OpcodeExtSet5l  = opcode{0xED, "SET 5,L", "Set bit 5 of L", 0, 1, unimplementedHandler}
	OpcodeExtSet5hl = opcode{0xEE, "SET 5,(HL)", "Set bit 5 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtSet5a  = opcode{0xEF, "SET 5,A", "Set bit 5 of A", 0, 1, unimplementedHandler}
	OpcodeExtSet6b  = opcode{0xF0, "SET 6,B", "Set bit 6 of B", 0, 1, unimplementedHandler}
	OpcodeExtSet6c  = opcode{0xF1, "SET 6,C", "Set bit 6 of C", 0, 1, unimplementedHandler}
	OpcodeExtSet6d  = opcode{0xF2, "SET 6,D", "Set bit 6 of D", 0, 1, unimplementedHandler}
	OpcodeExtSet6e  = opcode{0xF3, "SET 6,E", "Set bit 6 of E", 0, 1, unimplementedHandler}
	OpcodeExtSet6h  = opcode{0xF4, "SET 6,H", "Set bit 6 of H", 0, 1, unimplementedHandler}
	OpcodeExtSet6l  = opcode{0xF5, "SET 6,L", "Set bit 6 of L", 0, 1, unimplementedHandler}
	OpcodeExtSet6hl = opcode{0xF6, "SET 6,(HL)", "Set bit 6 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtSet6a  = opcode{0xF7, "SET 6,A", "Set bit 6 of A", 0, 1, unimplementedHandler}
	OpcodeExtSet7b  = opcode{0xF8, "SET 7,B", "Set bit 7 of B", 0, 1, unimplementedHandler}
	OpcodeExtSet7c  = opcode{0xF9, "SET 7,C", "Set bit 7 of C", 0, 1, unimplementedHandler}
	OpcodeExtSet7d  = opcode{0xFA, "SET 7,D", "Set bit 7 of D", 0, 1, unimplementedHandler}
	OpcodeExtSet7e  = opcode{0xFB, "SET 7,E", "Set bit 7 of E", 0, 1, unimplementedHandler}
	OpcodeExtSet7h  = opcode{0xFC, "SET 7,H", "Set bit 7 of H", 0, 1, unimplementedHandler}
	OpcodeExtSet7l  = opcode{0xFD, "SET 7,L", "Set bit 7 of L", 0, 1, unimplementedHandler}
	OpcodeExtSet7hl = opcode{0xFE, "SET 7,(HL)", "Set bit 7 of value pointed by HL", 0, 1, unimplementedHandler}
	OpcodeExtSet7a  = opcode{0xFF, "SET 7,A", "Set bit 7 of A", 0, 1, unimplementedHandler}
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

func lookupOpcode(opcodeByte byte) opcode {
	return opcodeMap[opcodeByte]
}

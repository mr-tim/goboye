package cpu

import (
	"fmt"
	"strings"
)

type opcodeHandler func(opcode, *processor)

type Opcode interface {
	Code() uint8
	Disassembly() string
	DisassemblyWithArg(arg string) string
	Description() string
	PayloadLength() uint8
	Cycles() uint8
}

type OpcodeAndPayload struct {
	op      Opcode
	payload []byte
}

func (o *OpcodeAndPayload) Disassembly() string {
	return o.op.DisassemblyWithArg(o.FormatPayload())
}

func (o *OpcodeAndPayload) Opcode() Opcode {
	return o.op
}

func (o *OpcodeAndPayload) FormatPayload() string {
	argWidth := o.op.PayloadLength()
	if argWidth == 1 {
		return fmt.Sprintf("0x%02X", o.payload[0])
	} else if argWidth == 2 {
		return fmt.Sprintf("0x%04X", (uint16(o.payload[1])<<8)|uint16(o.payload[0]))
	}
	return ""
}

type opcode struct {
	code          uint8
	disassembly   string
	description   string
	payloadLength uint8
	cycles        uint8
	handler       opcodeHandler
}

func (o *opcode) Code() uint8 {
	return o.code
}

func (o *opcode) Disassembly() string {
	return o.disassembly
}

func (o *opcode) DisassemblyWithArg(arg string) string {
	d := o.disassembly
	d = strings.ReplaceAll(d, "nn", arg)
	d = strings.ReplaceAll(d, "n", arg)
	return d
}

func (o *opcode) Description() string {
	return o.description
}

func (o *opcode) PayloadLength() uint8 {
	return o.payloadLength
}

func (o *opcode) Cycles() uint8 {
	return o.cycles
}

var (
	OpcodeNop    = opcode{0x00, "NOP", "No Operation", 0, 4, nopHandler}
	OpcodeLdBcnn = opcode{0x01, "LD BC,nn", "Load 16-bit immediate into BC", 2, 12, load16BitToRegPair(RegisterPairBC)}
	OpcodeLdBca  = opcode{0x02, "LD (BC),A", "Save A to address pointed by BC", 0, 8, saveAToBCAddr}
	OpcodeIncBc  = opcode{0x03, "INC BC", "Increment 16-bit BC", 0, 8, incrementRegPair(RegisterPairBC)}
	OpcodeIncB   = opcode{0x04, "INC B", "Increment B", 0, 4, incrementReg(RegisterB)}
	OpcodeDecB   = opcode{0x05, "DEC B", "Decrement B", 0, 4, decrementReg(RegisterB)}
	OpcodeLdBn   = opcode{0x06, "LD B,n", "Load 8-bit immediate into B", 1, 8, load8BitToReg(RegisterB)}
	//TODO: should this one also reset FlagZ?
	OpcodeRlcA    = opcode{0x07, "RLC A", "Rotate A left with carry", 0, 4, rotateRegLeftWithCarry(RegisterA)}
	OpcodeLdNnsp  = opcode{0x08, "LD (nn),SP", "Save SP to given address", 2, 20, saveSPToAddr}
	OpcodeAddHlbc = opcode{0x09, "ADD HL,BC", "Add 16-bit BC to HL", 0, 8, addRegPairToHL(RegisterPairBC)}
	OpcodeLdAbc   = opcode{0x0A, "LD A,(BC)", "Load A from address pointed to by BC", 0, 8, loadAFromRegPairAddr(RegisterPairBC)}
	OpcodeDecBc   = opcode{0x0B, "DEC BC", "Decrement 16-bit BC", 0, 8, decrementRegPair(RegisterPairBC)}
	OpcodeIncC    = opcode{0x0C, "INC C", "Increment C", 0, 4, incrementReg(RegisterC)}
	OpcodeDecC    = opcode{0x0D, "DEC C", "Decrement C", 0, 4, decrementReg(RegisterC)}
	OpcodeLdCn    = opcode{0x0E, "LD C,n", "Load 8-bit immediate into C", 1, 8, load8BitToReg(RegisterC)}
	//TODO: should this one also reset FlagZ?
	OpcodeRrcA   = opcode{0x0F, "RRC A", "Rotate A right with carry", 0, 4, rotateRegRightWithCarry(RegisterA)}
	OpcodeStop   = opcode{0x10, "STOP", "Stop processor", 0, 4, unimplementedHandler}
	OpcodeLdDenn = opcode{0x11, "LD DE,nn", "Load 16-bit immediate into DE", 2, 12, load16BitToRegPair(RegisterPairDE)}
	OpcodeLdDea  = opcode{0x12, "LD (DE),A", "Save A to address pointed by DE", 0, 8, saveAToDEAddr}
	OpcodeIncDe  = opcode{0x13, "INC DE", "Increment 16-bit DE", 0, 8, incrementRegPair(RegisterPairDE)}
	OpcodeIncD   = opcode{0x14, "INC D", "Increment D", 0, 4, incrementReg(RegisterD)}
	OpcodeDecD   = opcode{0x15, "DEC D", "Decrement D", 0, 4, decrementReg(RegisterD)}
	OpcodeLdDn   = opcode{0x16, "LD D,n", "Load 8-bit immediate into D", 1, 8, load8BitToReg(RegisterD)}
	//TODO: should this one also reset FlagZ?
	OpcodeRlA     = opcode{0x17, "RL A", "Rotate A left", 0, 4, rotateRegLeft(RegisterA)}
	OpcodeJrN     = opcode{0x18, "JR n", "Relative jump by signed immediate", 1, 12, relativeJumpImmediate}
	OpcodeAddHlde = opcode{0x19, "ADD HL,DE", "Add 16-bit DE to HL", 0, 8, addRegPairToHL(RegisterPairDE)}
	OpcodeLdAde   = opcode{0x1A, "LD A,(DE)", "Load A from address pointed to by DE", 0, 8, loadAFromRegPairAddr(RegisterPairDE)}
	OpcodeDecDe   = opcode{0x1B, "DEC DE", "Decrement 16-bit DE", 0, 8, decrementRegPair(RegisterPairDE)}
	OpcodeIncE    = opcode{0x1C, "INC E", "Increment E", 0, 4, incrementReg(RegisterE)}
	OpcodeDecE    = opcode{0x1D, "DEC E", "Decrement E", 0, 4, decrementReg(RegisterE)}
	OpcodeLdEn    = opcode{0x1E, "LD E,n", "Load 8-bit immediate into E", 1, 8, load8BitToReg(RegisterE)}
	//TODO: should this one also reset FlagZ?
	OpcodeRrA       = opcode{0x1F, "RR A", "Rotate A right", 0, 4, rotateRegRight(RegisterA)}
	OpcodeJrNzn     = opcode{0x20, "JR NZ,n", "Relative jump by signed immediate if last result was not zero", 1, 8, relativeJumpImmediateIfFlag(FlagZ, false)}
	OpcodeLdHlnn    = opcode{0x21, "LD HL,nn", "Load 16-bit immediate into HL", 2, 12, load16BitToRegPair(RegisterPairHL)}
	OpcodeLdiHla    = opcode{0x22, "LDI (HL),A", "Save A to address pointed by HL, and increment HL", 0, 8, saveAToHLAddrInc}
	OpcodeIncHl     = opcode{0x23, "INC HL", "Increment 16-bit HL", 0, 8, incrementRegPair(RegisterPairHL)}
	OpcodeIncH      = opcode{0x24, "INC H", "Increment H", 0, 4, incrementReg(RegisterH)}
	OpcodeDecH      = opcode{0x25, "DEC H", "Decrement H", 0, 4, decrementReg(RegisterH)}
	OpcodeLdHn      = opcode{0x26, "LD H,n", "Load 8-bit immediate into H", 1, 8, load8BitToReg(RegisterH)}
	OpcodeDaa       = opcode{0x27, "DAA", "Adjust A for BCD addition", 0, 4, adjustAForBCDAddition}
	OpcodeJrZn      = opcode{0x28, "JR Z,n", "Relative jump by signed immediate if last result was zero", 1, 8, relativeJumpImmediateIfFlag(FlagZ, true)}
	OpcodeAddHlhl   = opcode{0x29, "ADD HL,HL", "Add 16-bit HL to HL", 0, 8, addRegPairToHL(RegisterPairHL)}
	OpcodeLdiAhl    = opcode{0x2A, "LDI A,(HL)", "Load A from address pointed to by HL, and increment HL", 0, 8, loadAFromHLAddrInc}
	OpcodeDecHl     = opcode{0x2B, "DEC HL", "Decrement 16-bit HL", 0, 8, decrementRegPair(RegisterPairHL)}
	OpcodeIncL      = opcode{0x2C, "INC L", "Increment L", 0, 4, incrementReg(RegisterL)}
	OpcodeDecL      = opcode{0x2D, "DEC L", "Decrement L", 0, 4, decrementReg(RegisterL)}
	OpcodeLdLn      = opcode{0x2E, "LD L,n", "Load 8-bit immediate into L", 1, 8, load8BitToReg(RegisterL)}
	OpcodeCpl       = opcode{0x2F, "CPL", "Complement (logical NOT) on A", 0, 4, complementOnA}
	OpcodeJrNcn     = opcode{0x30, "JR NC,n", "Relative jump by signed immediate if last result caused no carry", 1, 8, relativeJumpImmediateIfFlag(FlagC, false)}
	OpcodeLdSpnn    = opcode{0x31, "LD SP,nn", "Load 16-bit immediate into SP", 2, 12, load16BitToRegPair(RegisterPairSP)}
	OpcodeLddHla    = opcode{0x32, "LDD (HL),A", "Save A to address pointed by HL, and decrement HL", 0, 8, saveAToHLAddrDec}
	OpcodeIncSp     = opcode{0x33, "INC SP", "Increment 16-bit SP", 0, 8, incrementRegPair(RegisterPairSP)}
	OpcodeIncHlAddr = opcode{0x34, "INC (HL)", "Increment value pointed by HL", 0, 12, incrementHLAddr}
	OpcodeDecHlAddr = opcode{0x35, "DEC (HL)", "Decrement value pointed by HL", 0, 12, decrementHLAddr}
	OpcodeLdHln     = opcode{0x36, "LD (HL),n", "Load 8-bit immediate into address pointed by HL", 1, 12, load8BitToHLAddr}
	OpcodeScf       = opcode{0x37, "SCF", "Set carry flag", 0, 4, setCarryFlag}
	OpcodeJrCn      = opcode{0x38, "JR C,n", "Relative jump by signed immediate if last result caused carry", 1, 8, relativeJumpImmediateIfFlag(FlagC, true)}
	OpcodeAddHlsp   = opcode{0x39, "ADD HL,SP", "Add 16-bit SP to HL", 0, 8, addRegPairToHL(RegisterPairSP)}
	OpcodeLddAhl    = opcode{0x3A, "LDD A,(HL)", "Load A from address pointed to by HL, and decrement HL", 0, 8, loadAFromHLAddrDec}
	OpcodeDecSp     = opcode{0x3B, "DEC SP", "Decrement 16-bit SP", 0, 8, decrementRegPair(RegisterPairSP)}
	OpcodeIncA      = opcode{0x3C, "INC A", "Increment A", 0, 4, incrementReg(RegisterA)}
	OpcodeDecA      = opcode{0x3D, "DEC A", "Decrement A", 0, 4, decrementReg(RegisterA)}
	OpcodeLdAn      = opcode{0x3E, "LD A,n", "Load 8-bit immediate into A", 1, 8, load8BitToReg(RegisterA)}
	OpcodeCcf       = opcode{0x3F, "CCF", "Clear carry flag", 0, 4, clearCarryFlag}
	OpcodeLdBb      = opcode{0x40, "LD B,B", "Copy B to B", 0, 4, loadRegToReg(RegisterB, RegisterB)}
	OpcodeLdBc      = opcode{0x41, "LD B,C", "Copy C to B", 0, 4, loadRegToReg(RegisterB, RegisterC)}
	OpcodeLdBd      = opcode{0x42, "LD B,D", "Copy D to B", 0, 4, loadRegToReg(RegisterB, RegisterD)}
	OpcodeLdBe      = opcode{0x43, "LD B,E", "Copy E to B", 0, 4, loadRegToReg(RegisterB, RegisterE)}
	OpcodeLdBh      = opcode{0x44, "LD B,H", "Copy H to B", 0, 4, loadRegToReg(RegisterB, RegisterH)}
	OpcodeLdBl      = opcode{0x45, "LD B,L", "Copy L to B", 0, 4, loadRegToReg(RegisterB, RegisterL)}
	OpcodeLdBhl     = opcode{0x46, "LD B,(HL)", "Copy value pointed by HL to B", 0, 8, loadHLAddrToReg(RegisterB)}
	OpcodeLdBa      = opcode{0x47, "LD B,A", "Copy A to B", 0, 4, loadRegToReg(RegisterB, RegisterA)}
	OpcodeLdCb      = opcode{0x48, "LD C,B", "Copy B to C", 0, 4, loadRegToReg(RegisterC, RegisterB)}
	OpcodeLdCc      = opcode{0x49, "LD C,C", "Copy C to C", 0, 4, loadRegToReg(RegisterC, RegisterC)}
	OpcodeLdCd      = opcode{0x4A, "LD C,D", "Copy D to C", 0, 4, loadRegToReg(RegisterC, RegisterD)}
	OpcodeLdCe      = opcode{0x4B, "LD C,E", "Copy E to C", 0, 4, loadRegToReg(RegisterC, RegisterE)}
	OpcodeLdCh      = opcode{0x4C, "LD C,H", "Copy H to C", 0, 4, loadRegToReg(RegisterC, RegisterH)}
	OpcodeLdCl      = opcode{0x4D, "LD C,L", "Copy L to C", 0, 4, loadRegToReg(RegisterC, RegisterL)}
	OpcodeLdChl     = opcode{0x4E, "LD C,(HL)", "Copy value pointed by HL to C", 0, 8, loadHLAddrToReg(RegisterC)}
	OpcodeLdCa      = opcode{0x4F, "LD C,A", "Copy A to C", 0, 4, loadRegToReg(RegisterC, RegisterA)}
	OpcodeLdDb      = opcode{0x50, "LD D,B", "Copy B to D", 0, 4, loadRegToReg(RegisterD, RegisterB)}
	OpcodeLdDc      = opcode{0x51, "LD D,C", "Copy C to D", 0, 4, loadRegToReg(RegisterD, RegisterC)}
	OpcodeLdDd      = opcode{0x52, "LD D,D", "Copy D to D", 0, 4, loadRegToReg(RegisterD, RegisterD)}
	OpcodeLdDe      = opcode{0x53, "LD D,E", "Copy E to D", 0, 4, loadRegToReg(RegisterD, RegisterE)}
	OpcodeLdDh      = opcode{0x54, "LD D,H", "Copy H to D", 0, 4, loadRegToReg(RegisterD, RegisterH)}
	OpcodeLdDl      = opcode{0x55, "LD D,L", "Copy L to D", 0, 4, loadRegToReg(RegisterD, RegisterL)}
	OpcodeLdDhl     = opcode{0x56, "LD D,(HL)", "Copy value pointed by HL to D", 0, 8, loadHLAddrToReg(RegisterD)}
	OpcodeLdDa      = opcode{0x57, "LD D,A", "Copy A to D", 0, 4, loadRegToReg(RegisterD, RegisterA)}
	OpcodeLdEb      = opcode{0x58, "LD E,B", "Copy B to E", 0, 4, loadRegToReg(RegisterE, RegisterB)}
	OpcodeLdEc      = opcode{0x59, "LD E,C", "Copy C to E", 0, 4, loadRegToReg(RegisterE, RegisterC)}
	OpcodeLdEd      = opcode{0x5A, "LD E,D", "Copy D to E", 0, 4, loadRegToReg(RegisterE, RegisterD)}
	OpcodeLdEe      = opcode{0x5B, "LD E,E", "Copy E to E", 0, 4, loadRegToReg(RegisterE, RegisterE)}
	OpcodeLdEh      = opcode{0x5C, "LD E,H", "Copy H to E", 0, 4, loadRegToReg(RegisterE, RegisterH)}
	OpcodeLdEl      = opcode{0x5D, "LD E,L", "Copy L to E", 0, 4, loadRegToReg(RegisterE, RegisterL)}
	OpcodeLdEhl     = opcode{0x5E, "LD E,(HL)", "Copy value pointed by HL to E", 0, 8, loadHLAddrToReg(RegisterE)}
	OpcodeLdEa      = opcode{0x5F, "LD E,A", "Copy A to E", 0, 4, loadRegToReg(RegisterE, RegisterA)}
	OpcodeLdHb      = opcode{0x60, "LD H,B", "Copy B to H", 0, 4, loadRegToReg(RegisterH, RegisterB)}
	OpcodeLdHc      = opcode{0x61, "LD H,C", "Copy C to H", 0, 4, loadRegToReg(RegisterH, RegisterC)}
	OpcodeLdHd      = opcode{0x62, "LD H,D", "Copy D to H", 0, 4, loadRegToReg(RegisterH, RegisterD)}
	OpcodeLdHe      = opcode{0x63, "LD H,E", "Copy E to H", 0, 4, loadRegToReg(RegisterH, RegisterE)}
	OpcodeLdHh      = opcode{0x64, "LD H,H", "Copy H to H", 0, 4, loadRegToReg(RegisterH, RegisterH)}
	OpcodeLdHl      = opcode{0x65, "LD H,L", "Copy L to H", 0, 4, loadRegToReg(RegisterH, RegisterL)}
	OpcodeLdHhl     = opcode{0x66, "LD H,(HL)", "Copy value pointed by HL to H", 0, 8, loadHLAddrToReg(RegisterH)}
	OpcodeLdHa      = opcode{0x67, "LD H,A", "Copy A to H", 0, 4, loadRegToReg(RegisterH, RegisterA)}
	OpcodeLdLb      = opcode{0x68, "LD L,B", "Copy B to L", 0, 4, loadRegToReg(RegisterL, RegisterB)}
	OpcodeLdLc      = opcode{0x69, "LD L,C", "Copy C to L", 0, 4, loadRegToReg(RegisterL, RegisterC)}
	OpcodeLdLd      = opcode{0x6A, "LD L,D", "Copy D to L", 0, 4, loadRegToReg(RegisterL, RegisterD)}
	OpcodeLdLe      = opcode{0x6B, "LD L,E", "Copy E to L", 0, 4, loadRegToReg(RegisterL, RegisterE)}
	OpcodeLdLh      = opcode{0x6C, "LD L,H", "Copy H to L", 0, 4, loadRegToReg(RegisterL, RegisterH)}
	OpcodeLdLl      = opcode{0x6D, "LD L,L", "Copy L to L", 0, 4, loadRegToReg(RegisterL, RegisterL)}
	OpcodeLdLhl     = opcode{0x6E, "LD L,(HL)", "Copy value pointed by HL to L", 0, 8, loadHLAddrToReg(RegisterL)}
	OpcodeLdLa      = opcode{0x6F, "LD L,A", "Copy A to L", 0, 4, loadRegToReg(RegisterL, RegisterA)}
	OpcodeLdHlb     = opcode{0x70, "LD (HL),B", "Copy B to address pointed by HL", 0, 8, loadRegToHLAddr(RegisterB)}
	OpcodeLdHlc     = opcode{0x71, "LD (HL),C", "Copy C to address pointed by HL", 0, 8, loadRegToHLAddr(RegisterC)}
	OpcodeLdHld     = opcode{0x72, "LD (HL),D", "Copy D to address pointed by HL", 0, 8, loadRegToHLAddr(RegisterD)}
	OpcodeLdHle     = opcode{0x73, "LD (HL),E", "Copy E to address pointed by HL", 0, 8, loadRegToHLAddr(RegisterE)}
	OpcodeLdHlh     = opcode{0x74, "LD (HL),H", "Copy H to address pointed by HL", 0, 8, loadRegToHLAddr(RegisterH)}
	OpcodeLdHll     = opcode{0x75, "LD (HL),L", "Copy L to address pointed by HL", 0, 8, loadRegToHLAddr(RegisterL)}
	OpcodeHalt      = opcode{0x76, "HALT", "Halt processor", 0, 4, unimplementedHandler}
	OpcodeLdHla     = opcode{0x77, "LD (HL),A", "Copy A to address pointed by HL", 0, 8, loadRegToHLAddr(RegisterA)}
	OpcodeLdAb      = opcode{0x78, "LD A,B", "Copy B to A", 0, 4, loadRegToReg(RegisterA, RegisterB)}
	OpcodeLdAc      = opcode{0x79, "LD A,C", "Copy C to A", 0, 4, loadRegToReg(RegisterA, RegisterC)}
	OpcodeLdAd      = opcode{0x7A, "LD A,D", "Copy D to A", 0, 4, loadRegToReg(RegisterA, RegisterD)}
	OpcodeLdAe      = opcode{0x7B, "LD A,E", "Copy E to A", 0, 4, loadRegToReg(RegisterA, RegisterE)}
	OpcodeLdAh      = opcode{0x7C, "LD A,H", "Copy H to A", 0, 4, loadRegToReg(RegisterA, RegisterH)}
	OpcodeLdAl      = opcode{0x7D, "LD A,L", "Copy L to A", 0, 4, loadRegToReg(RegisterA, RegisterL)}
	OpcodeLdAhl     = opcode{0x7E, "LD A,(HL)", "Copy value pointed by HL to A", 0, 8, loadHLAddrToReg(RegisterA)}
	OpcodeLdAa      = opcode{0x7F, "LD A,A", "Copy A to A", 0, 4, loadRegToReg(RegisterA, RegisterA)}
	OpcodeAddAb     = opcode{0x80, "ADD A,B", "Add B to A", 0, 4, addRegToA(RegisterB)}
	OpcodeAddAc     = opcode{0x81, "ADD A,C", "Add C to A", 0, 4, addRegToA(RegisterC)}
	OpcodeAddAd     = opcode{0x82, "ADD A,D", "Add D to A", 0, 4, addRegToA(RegisterD)}
	OpcodeAddAe     = opcode{0x83, "ADD A,E", "Add E to A", 0, 4, addRegToA(RegisterE)}
	OpcodeAddAh     = opcode{0x84, "ADD A,H", "Add H to A", 0, 4, addRegToA(RegisterH)}
	OpcodeAddAl     = opcode{0x85, "ADD A,L", "Add L to A", 0, 4, addRegToA(RegisterL)}
	OpcodeAddAhl    = opcode{0x86, "ADD A,(HL)", "Add value pointed by HL to A", 0, 8, addHLAddrToA}
	OpcodeAddAa     = opcode{0x87, "ADD A,A", "Add A to A", 0, 4, addRegToA(RegisterA)}
	OpcodeAdcAb     = opcode{0x88, "ADC A,B", "Add B and carry flag to A", 0, 4, addRegAndCarryToA(RegisterB)}
	OpcodeAdcAc     = opcode{0x89, "ADC A,C", "Add C and carry flag to A", 0, 4, addRegAndCarryToA(RegisterC)}
	OpcodeAdcAd     = opcode{0x8A, "ADC A,D", "Add D and carry flag to A", 0, 4, addRegAndCarryToA(RegisterD)}
	OpcodeAdcAe     = opcode{0x8B, "ADC A,E", "Add E and carry flag to A", 0, 4, addRegAndCarryToA(RegisterE)}
	OpcodeAdcAh     = opcode{0x8C, "ADC A,H", "Add H and carry flag to A", 0, 4, addRegAndCarryToA(RegisterH)}
	OpcodeAdcAl     = opcode{0x8D, "ADC A,L", "Add L and carry flag to A", 0, 4, addRegAndCarryToA(RegisterL)}
	OpcodeAdcAhl    = opcode{0x8E, "ADC A,(HL)", "Add value pointed by HL and carry flag to A", 0, 8, addHLAddrAndCarryToA}
	OpcodeAdcAa     = opcode{0x8F, "ADC A,A", "Add A and carry flag to A", 0, 4, addRegAndCarryToA(RegisterA)}
	OpcodeSubAb     = opcode{0x90, "SUB A,B", "Subtract B from A", 0, 4, subtractRegFromA(RegisterB)}
	OpcodeSubAc     = opcode{0x91, "SUB A,C", "Subtract C from A", 0, 4, subtractRegFromA(RegisterC)}
	OpcodeSubAd     = opcode{0x92, "SUB A,D", "Subtract D from A", 0, 4, subtractRegFromA(RegisterD)}
	OpcodeSubAe     = opcode{0x93, "SUB A,E", "Subtract E from A", 0, 4, subtractRegFromA(RegisterE)}
	OpcodeSubAh     = opcode{0x94, "SUB A,H", "Subtract H from A", 0, 4, subtractRegFromA(RegisterH)}
	OpcodeSubAl     = opcode{0x95, "SUB A,L", "Subtract L from A", 0, 4, subtractRegFromA(RegisterL)}
	OpcodeSubAhl    = opcode{0x96, "SUB A,(HL)", "Subtract value pointed by HL from A", 0, 8, subtractHLAddrFromA}
	OpcodeSubAa     = opcode{0x97, "SUB A,A", "Subtract A from A", 0, 1, subtractRegFromA(RegisterA)}
	OpcodeSbcAb     = opcode{0x98, "SBC A,B", "Subtract B and carry flag from A", 0, 4, subtractRegAndCarryFromA(RegisterB)}
	OpcodeSbcAc     = opcode{0x99, "SBC A,C", "Subtract C and carry flag from A", 0, 4, subtractRegAndCarryFromA(RegisterC)}
	OpcodeSbcAd     = opcode{0x9A, "SBC A,D", "Subtract D and carry flag from A", 0, 4, subtractRegAndCarryFromA(RegisterD)}
	OpcodeSbcAe     = opcode{0x9B, "SBC A,E", "Subtract E and carry flag from A", 0, 4, subtractRegAndCarryFromA(RegisterE)}
	OpcodeSbcAh     = opcode{0x9C, "SBC A,H", "Subtract H and carry flag from A", 0, 4, subtractRegAndCarryFromA(RegisterH)}
	OpcodeSbcAl     = opcode{0x9D, "SBC A,L", "Subtract and carry flag L from A", 0, 4, subtractRegAndCarryFromA(RegisterL)}
	OpcodeSbcAhl    = opcode{0x9E, "SBC A,(HL)", "Subtract value pointed by HL and carry flag from A", 0, 8, subtractHLAddrAndCarryFromA}
	OpcodeSbcAa     = opcode{0x9F, "SBC A,A", "Subtract A and carry flag from A", 0, 4, subtractRegAndCarryFromA(RegisterA)}
	OpcodeAndB      = opcode{0xA0, "AND B", "Logical AND B against A", 0, 4, logicalAndRegAgainstA(RegisterB)}
	OpcodeAndC      = opcode{0xA1, "AND C", "Logical AND C against A", 0, 4, logicalAndRegAgainstA(RegisterB)}
	OpcodeAndD      = opcode{0xA2, "AND D", "Logical AND D against A", 0, 4, logicalAndRegAgainstA(RegisterB)}
	OpcodeAndE      = opcode{0xA3, "AND E", "Logical AND E against A", 0, 4, logicalAndRegAgainstA(RegisterB)}
	OpcodeAndH      = opcode{0xA4, "AND H", "Logical AND H against A", 0, 4, logicalAndRegAgainstA(RegisterB)}
	OpcodeAndL      = opcode{0xA5, "AND L", "Logical AND L against A", 0, 4, logicalAndRegAgainstA(RegisterB)}
	OpcodeAndHl     = opcode{0xA6, "AND (HL)", "Logical AND value pointed by HL against A", 0, 8, logicalAndHLAddrAgainstA}
	OpcodeAndA      = opcode{0xA7, "AND A", "Logical AND A against A", 0, 4, logicalAndRegAgainstA(RegisterA)}
	OpcodeXorB      = opcode{0xA8, "XOR B", "Logical XOR B against A", 0, 4, logicalXorRegAgainstA(RegisterB)}
	OpcodeXorC      = opcode{0xA9, "XOR C", "Logical XOR C against A", 0, 4, logicalXorRegAgainstA(RegisterC)}
	OpcodeXorD      = opcode{0xAA, "XOR D", "Logical XOR D against A", 0, 4, logicalXorRegAgainstA(RegisterD)}
	OpcodeXorE      = opcode{0xAB, "XOR E", "Logical XOR E against A", 0, 4, logicalXorRegAgainstA(RegisterE)}
	OpcodeXorH      = opcode{0xAC, "XOR H", "Logical XOR H against A", 0, 4, logicalXorRegAgainstA(RegisterH)}
	OpcodeXorL      = opcode{0xAD, "XOR L", "Logical XOR L against A", 0, 4, logicalXorRegAgainstA(RegisterL)}
	OpcodeXorHl     = opcode{0xAE, "XOR (HL)", "Logical XOR value pointed by HL against A", 0, 8, logicalXorHLAddrAgainstA}
	OpcodeXorA      = opcode{0xAF, "XOR A", "Logical XOR A against A", 0, 4, logicalXorRegAgainstA(RegisterA)}
	OpcodeOrB       = opcode{0xB0, "OR B", "Logical OR B against A", 0, 4, logicalOrRegAgainstA(RegisterB)}
	OpcodeOrC       = opcode{0xB1, "OR C", "Logical OR C against A", 0, 4, logicalOrRegAgainstA(RegisterC)}
	OpcodeOrD       = opcode{0xB2, "OR D", "Logical OR D against A", 0, 4, logicalOrRegAgainstA(RegisterD)}
	OpcodeOrE       = opcode{0xB3, "OR E", "Logical OR E against A", 0, 4, logicalOrRegAgainstA(RegisterE)}
	OpcodeOrH       = opcode{0xB4, "OR H", "Logical OR H against A", 0, 4, logicalOrRegAgainstA(RegisterH)}
	OpcodeOrL       = opcode{0xB5, "OR L", "Logical OR L against A", 0, 4, logicalOrRegAgainstA(RegisterL)}
	OpcodeOrHl      = opcode{0xB6, "OR (HL)", "Logical OR value pointed by HL against A", 0, 8, logicalOrHLAddrAgainstA}
	OpcodeOrA       = opcode{0xB7, "OR A", "Logical OR A against A", 0, 4, logicalOrRegAgainstA(RegisterA)}
	OpcodeCpB       = opcode{0xB8, "CP B", "Compare B against A", 0, 4, compareRegAgainstA(RegisterB)}
	OpcodeCpC       = opcode{0xB9, "CP C", "Compare C against A", 0, 4, compareRegAgainstA(RegisterC)}
	OpcodeCpD       = opcode{0xBA, "CP D", "Compare D against A", 0, 4, compareRegAgainstA(RegisterD)}
	OpcodeCpE       = opcode{0xBB, "CP E", "Compare E against A", 0, 4, compareRegAgainstA(RegisterE)}
	OpcodeCpH       = opcode{0xBC, "CP H", "Compare H against A", 0, 4, compareRegAgainstA(RegisterH)}
	OpcodeCpL       = opcode{0xBD, "CP L", "Compare L against A", 0, 4, compareRegAgainstA(RegisterL)}
	OpcodeCpHl      = opcode{0xBE, "CP (HL)", "Compare value pointed by HL against A", 0, 8, compareHLAddrAgainstA}
	OpcodeCpA       = opcode{0xBF, "CP A", "Compare A against A", 0, 4, compareRegAgainstA(RegisterA)}
	OpcodeRetNz     = opcode{0xC0, "RET NZ", "Return if last result was not zero", 0, 8, conditionalReturn(FlagZ, false)}
	OpcodePopBc     = opcode{0xC1, "POP BC", "Pop 16-bit value from stack into BC", 0, 12, popRegisterPair(RegisterPairBC)}
	OpcodeJpNznn    = opcode{0xC2, "JP NZ,nn", "Absolute jump to 16-bit location if last result was not zero", 2, 12, jumpTo16BitAddressIfFlag(FlagZ, false)}
	OpcodeJpNn      = opcode{0xC3, "JP nn", "Absolute jump to 16-bit location", 2, 16, jumpTo16BitAddress}
	OpcodeCallNznn  = opcode{0xC4, "CALL NZ,nn", "Call routine at 16-bit location if last result was not zero", 2, 12, conditionalCall16BitAddress(FlagZ, false)}
	OpcodePushBc    = opcode{0xC5, "PUSH BC", "Push 16-bit BC onto stack", 0, 16, pushRegisterPair(RegisterPairBC)}
	OpcodeAddAn     = opcode{0xC6, "ADD A,n", "Add 8-bit immediate to A", 1, 8, addImmediate}
	OpcodeRst0      = opcode{0xC7, "RST 0", "Call routine at address 0000h", 0, 16, callRoutineAtAddress(0x0000)}
	OpcodeRetZ      = opcode{0xC8, "RET Z", "Return if last result was zero", 0, 8, conditionalReturn(FlagZ, true)}
	OpcodeRet       = opcode{0xC9, "RET", "Return to calling routine", 0, 16, doReturn}
	OpcodeJpZnn     = opcode{0xCA, "JP Z,nn", "Absolute jump to 16-bit location if last result was zero", 2, 12, jumpTo16BitAddressIfFlag(FlagZ, true)}
	OpcodeExtOps    = opcode{0xCB, "Ext ops", "Extended operations (two-byte instruction code)", 0, 4, extendedOps}
	OpcodeCallZnn   = opcode{0xCC, "CALL Z,nn", "Call routine at 16-bit location if last result was zero", 2, 12, conditionalCall16BitAddress(FlagZ, true)}
	OpcodeCallNn    = opcode{0xCD, "CALL nn", "Call routine at 16-bit location", 2, 24, call16BitAddress}
	OpcodeAdcAn     = opcode{0xCE, "ADC A,n", "Add 8-bit immediate and carry to A", 1, 8, addCImmediate}
	OpcodeRst8      = opcode{0xCF, "RST 8", "Call routine at address 0008h", 0, 16, callRoutineAtAddress(0x0008)}
	OpcodeRetNc     = opcode{0xD0, "RET NC", "Return if last result caused no carry", 0, 8, conditionalReturn(FlagC, false)}
	OpcodePopDe     = opcode{0xD1, "POP DE", "Pop 16-bit value from stack into DE", 0, 12, popRegisterPair(RegisterPairDE)}
	OpcodeJpNcnn    = opcode{0xD2, "JP NC,nn", "Absolute jump to 16-bit location if last result caused no carry", 2, 12, jumpTo16BitAddressIfFlag(FlagC, false)}
	OpcodeXxD3      = opcode{0xD3, "XX", "Operation removed in this CPU", 0, 0, unsupportedHandler}
	OpcodeCallNcnn  = opcode{0xD4, "CALL NC,nn", "Call routine at 16-bit location if last result caused no carry", 2, 12, conditionalCall16BitAddress(FlagC, false)}
	OpcodePushDe    = opcode{0xD5, "PUSH DE", "Push 16-bit DE onto stack", 0, 16, pushRegisterPair(RegisterPairDE)}
	OpcodeSubAn     = opcode{0xD6, "SUB A,n", "Subtract 8-bit immediate from A", 1, 8, subtractImmediate}
	OpcodeRst10     = opcode{0xD7, "RST 10", "Call routine at address 0010h", 0, 16, callRoutineAtAddress(0x0010)}
	OpcodeRetC      = opcode{0xD8, "RET C", "Return if last result caused carry", 0, 8, conditionalReturn(FlagC, true)}
	OpcodeReti      = opcode{0xD9, "RETI", "Enable interrupts and return to calling routine", 0, 16, doReturnEnablingInterrupts}
	OpcodeJpCnn     = opcode{0xDA, "JP C,nn", "Absolute jump to 16-bit location if last result caused carry", 2, 12, jumpTo16BitAddressIfFlag(FlagC, true)}
	OpcodeXxDB      = opcode{0xDB, "XX", "Operation removed in this CPU", 0, 0, unsupportedHandler}
	OpcodeCallCnn   = opcode{0xDC, "CALL C,nn", "Call routine at 16-bit location if last result caused carry", 2, 12, conditionalCall16BitAddress(FlagC, true)}
	OpcodeXxDD      = opcode{0xDD, "XX", "Operation removed in this CPU", 0, 0, unsupportedHandler}
	OpcodeSbcAn     = opcode{0xDE, "SBC A,n", "Subtract 8-bit immediate and carry from A", 1, 8, subCImmediate}
	OpcodeRst18     = opcode{0xDF, "RST 18", "Call routine at address 0018h", 0, 16, callRoutineAtAddress(0x0018)}
	OpcodeLdhNa     = opcode{0xE0, "LDH (n),A", "Save A at address pointed to by (FF00h + 8-bit immediate)", 1, 12, saveAToFFPlusImmediateAddr}
	OpcodePopHl     = opcode{0xE1, "POP HL", "Pop 16-bit value from stack into HL", 0, 12, popRegisterPair(RegisterPairHL)}
	OpcodeLdhCa     = opcode{0xE2, "LDH (C),A", "Save A at address pointed to by (FF00h + C)", 0, 8, saveAToFFPlusCAddr}
	OpcodeXxE3      = opcode{0xE3, "XX", "Operation removed in this CPU", 0, 0, unsupportedHandler}
	OpcodeXxE4      = opcode{0xE4, "XX", "Operation removed in this CPU", 0, 0, unsupportedHandler}
	OpcodePushHl    = opcode{0xE5, "PUSH HL", "Push 16-bit HL onto stack", 0, 16, pushRegisterPair(RegisterPairHL)}
	OpcodeAndN      = opcode{0xE6, "AND n", "Logical AND 8-bit immediate against A", 1, 8, logicalAndImmediate}
	OpcodeRst20     = opcode{0xE7, "RST 20", "Call routine at address 0020h", 0, 16, callRoutineAtAddress(0x0020)}
	OpcodeAddSpd    = opcode{0xE8, "ADD SP,d", "Add signed 8-bit immediate to SP", 0, 16, add8BitSignedImmediateToSP}
	OpcodeJpHl      = opcode{0xE9, "JP (HL)", "Jump to 16-bit value pointed by HL", 0, 4, jumpToHLAddr}
	OpcodeLdNna     = opcode{0xEA, "LD (nn),A", "Save A at given 16-bit address", 2, 16, saveATo16BitAddr}
	OpcodeXxEB      = opcode{0xEB, "XX", "Operation removed in this CPU", 0, 0, unsupportedHandler}
	OpcodeXxEC      = opcode{0xEC, "XX", "Operation removed in this CPU", 0, 0, unsupportedHandler}
	OpcodeXxED      = opcode{0xED, "XX", "Operation removed in this CPU", 0, 0, unsupportedHandler}
	OpcodeXorN      = opcode{0xEE, "XOR n", "Logical XOR 8-bit immediate against A", 1, 8, logicalXorImmediate}
	OpcodeRst28     = opcode{0xEF, "RST 28", "Call routine at address 0028h", 0, 16, callRoutineAtAddress(0x0028)}
	OpcodeLdhAn     = opcode{0xF0, "LDH A,(n)", "Load A from address pointed to by (FF00h + 8-bit immediate)", 1, 12, loadAFromFFPlusImmediateAddr}
	OpcodePopAf     = opcode{0xF1, "POP AF", "Pop 16-bit value from stack into AF", 0, 12, popRegisterPair(RegisterPairAF)}
	OpcodeXxF2      = opcode{0xF2, "XX", "Operation removed in this CPU", 0, 8, unsupportedHandler}
	OpcodeDi        = opcode{0xF3, "DI", "Disable interrupts", 0, 4, disableInterrupts}
	OpcodeXxF4      = opcode{0xF4, "XX", "Operation removed in this CPU", 0, 0, unsupportedHandler}
	OpcodePushAf    = opcode{0xF5, "PUSH AF", "Push 16-bit AF onto stack", 0, 16, pushRegisterPair(RegisterPairAF)}
	OpcodeOrN       = opcode{0xF6, "OR n", "Logical OR 8-bit immediate against A", 1, 8, logicalOrImmediate}
	OpcodeRst30     = opcode{0xF7, "RST 30", "Call routine at address 0030h", 0, 16, callRoutineAtAddress(0x0030)}
	OpcodeLdhlSpd   = opcode{0xF8, "LDHL SP,d", "Add signed 8-bit immediate to SP and save result in HL", 0, 12, add8BitImmediateToSPSaveInHL}
	OpcodeLdSphl    = opcode{0xF9, "LD SP,HL", "Copy HL to SP", 0, 8, copyHLToSP}
	OpcodeLdAnn     = opcode{0xFA, "LD A,(nn)", "Load A from given 16-bit address", 2, 16, loadAFromAddr}
	OpcodeEi        = opcode{0xFB, "EI", "Enable interrupts", 0, 4, enableInterrupts}
	OpcodeXxFC      = opcode{0xFC, "XX", "Operation removed in this CPU", 0, 0, unsupportedHandler}
	OpcodeXxFD      = opcode{0xFD, "XX", "Operation removed in this CPU", 0, 0, unsupportedHandler}
	OpcodeCpN       = opcode{0xFE, "CP n", "Compare 8-bit immediate against A", 1, 8, compareImmediate}
	OpcodeRst38     = opcode{0xFF, "RST 38", "Call routine at address 0038h", 0, 16, callRoutineAtAddress(0x0038)}
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

func LookupOpcode(opcodeByte byte) opcode {
	return opcodeMap[opcodeByte]
}

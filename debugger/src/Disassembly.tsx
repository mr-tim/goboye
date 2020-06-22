import React from "react";
import styled from "styled-components";

interface DisassemblyProps {
    currentAddress: number;
    instructions: Instruction[]
}

interface Instruction {
    address: number
    disassembly: string
}

interface DisassemblyCellProps {
    isCurrentAddress?: boolean
}

const DisassemblyCell = styled.div<DisassemblyCellProps>`
    padding: 6px 12px;
    border-bottom: 1px solid #999;
    color: ${props => props.isCurrentAddress? '#000': '#444'};
`;

const toHex = (value: number, padding: number) => {
    let hexValue = value.toString(16);

    while (hexValue.length < padding) {
        hexValue = '0' + hexValue
    }
    return '0x' + hexValue;
}

const Disassembly: React.FC<DisassemblyProps> = (props) => {
    let cells = props.instructions.map((instruction, index) => {
        let isCurrentAddress = props.currentAddress === instruction.address;
        return (<DisassemblyCell className="monospaced" isCurrentAddress={isCurrentAddress}>
            {toHex(instruction.address, 4)}: {instruction.disassembly}
        </DisassemblyCell>);
    });

    return (
        <div className="disassembly">
            <DisassemblyCell>Disassembly</DisassemblyCell>
            {cells}
        </div>
    );
}

export default Disassembly;
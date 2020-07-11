import React, {ChangeEventHandler} from "react";
import styled from "styled-components";
import {toHex} from "../util/hex";

interface DisassemblyProps {
  currentAddress: number;
  instructions: Instruction[];
  breakpoints: number[];
  setBreakpoint: (address: number, isBreakpoint: boolean) => void;
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
    color: ${props => props.isCurrentAddress ? '#000' : '#888'};
`;

const Disassembly: React.FC<DisassemblyProps> = (props) => {
  let cells = props.instructions.map((instruction, idx) => {
    let isCurrentAddress = props.currentAddress === instruction.address;
    let isBreakpoint = props.breakpoints.includes(instruction.address);
    let toggleBreakpoint: ChangeEventHandler = (event) => {
      props.setBreakpoint(instruction.address, !isBreakpoint);
    };
    return (<DisassemblyCell key={idx} className="monospaced" isCurrentAddress={isCurrentAddress}>
      <input type="checkbox" checked={isBreakpoint} onChange={toggleBreakpoint}/>
      {toHex(instruction.address, 4)}: {instruction.disassembly}
    </DisassemblyCell>);
  });

  return (
      <div className="disassembly">
        <DisassemblyCell isCurrentAddress={true}>Disassembly</DisassemblyCell>
        {cells}
      </div>
  );
}

export default Disassembly;

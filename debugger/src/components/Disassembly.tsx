import React, {ChangeEventHandler, useState} from "react";
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

const DisassemblyCell = styled.tr<DisassemblyCellProps>`
    padding: 6px 12px;
    color: ${props => props.isCurrentAddress ? '#000' : '#888'};
`;

const CommentCell = () => {
  let [comment, setComment] = useState('');
  let [isMouseover, setIsMouseover] = useState(false);

  return <td style={{width: '500px'}} onMouseEnter={() =>setIsMouseover(true)} onMouseLeave={() => setIsMouseover(false)}>
    {comment === ''? isMouseover && '//' : comment}&nbsp;
  </td>
};

const Disassembly: React.FC<DisassemblyProps> = (props) => {
  let cells = props.instructions.map((instruction, idx) => {
    let isCurrentAddress = props.currentAddress === instruction.address;
    let isBreakpoint = props.breakpoints.includes(instruction.address);
    let toggleBreakpoint: ChangeEventHandler = (event) => {
      props.setBreakpoint(instruction.address, !isBreakpoint);
    };
    return (<DisassemblyCell key={idx} className="monospaced" isCurrentAddress={isCurrentAddress}>
      <td>
        <input type="checkbox" checked={isBreakpoint} onChange={toggleBreakpoint}/>
      </td>
      <td className="addr-instr">
      {toHex(instruction.address, 4)}: {instruction.disassembly}
      </td>
      <CommentCell/>
    </DisassemblyCell>);
  });

  return (
      <div className="disassembly">
        <h4>Disassembly</h4>
        <table>
        {cells}
        </table>
      </div>
  );
}

export default Disassembly;

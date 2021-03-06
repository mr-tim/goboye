import React from "react";
import styled from "styled-components";
import {toHex} from "../util/hex";

interface RegisterProps {
  registers: { [registerName: string]: number },
  flags: {
    Z: boolean,
    N: boolean,
    H: boolean,
    C: boolean,
  }
}

const SidebarCell = styled.div`
    padding: 6px 12px;
    border-bottom: 1px solid #999;
`

const tick = '\u2714';
const cross = '\u274e';

const flag = (value: boolean): string => value ? tick : cross;

const Registers: React.FC<RegisterProps> = (props) => {
  const {registers} = props;
  return (<div>
    <SidebarCell>Registers</SidebarCell>
    {Object.keys(registers).sort().map((register, idx) => (
        <SidebarCell key={idx}
                     className="monospaced"><b>{register}</b> {toHex(registers[register], 4)}
        </SidebarCell>
    ))}
    <SidebarCell/>
    <SidebarCell>Flags</SidebarCell>
    <SidebarCell className="monospaced"><b>Z</b> {flag(props.flags.Z)}</SidebarCell>
    <SidebarCell className="monospaced"><b>N</b> {flag(props.flags.N)}</SidebarCell>
    <SidebarCell className="monospaced"><b>H</b> {flag(props.flags.H)}</SidebarCell>
    <SidebarCell className="monospaced"><b>C</b> {flag(props.flags.C)}</SidebarCell>
  </div>);
}

export default Registers;

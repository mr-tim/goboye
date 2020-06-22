import React from "react";
import styled from "styled-components";
import { toHex } from "./util/hex";

interface RegisterProps {
    registers: { [registerName: string]: number }
};

const SidebarCell = styled.div`
    padding: 6px 12px;
    border-bottom: 1px solid #999;
`

const tick = '\u2714';
const cross = '\u274e';

const flag = (value: boolean): string => value ? tick : cross;

const Registers: React.FC<RegisterProps> = (props) => {
    const { registers } = props;
    return (<div>
        <SidebarCell>Registers</SidebarCell>
        {Object.keys(registers).sort().map((register, idx) => (
            <SidebarCell key={idx} className="monospaced"><b>{register}</b> {toHex(registers[register], 4)}</SidebarCell>
        ))}
        <SidebarCell />
        <SidebarCell>Flags</SidebarCell>
        <SidebarCell className="monospaced"><b>Z</b> {flag(false)}</SidebarCell>
        <SidebarCell className="monospaced"><b>N</b> {flag(false)}</SidebarCell>
        <SidebarCell className="monospaced"><b>H</b> {flag(false)}</SidebarCell>
        <SidebarCell className="monospaced"><b>C</b> {flag(false)}</SidebarCell>
    </div>);
}

export default Registers;
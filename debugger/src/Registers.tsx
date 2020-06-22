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

const Registers: React.FC<RegisterProps> = (props) => {
    const { registers } = props;
    return (<div>
        <SidebarCell>Registers</SidebarCell>
        {Object.keys(registers).map((register, idx) => (
            <SidebarCell className="monospaced">{register}: {toHex(registers[register], 4)}</SidebarCell>
        ))}
    </div>);
}

export default Registers;
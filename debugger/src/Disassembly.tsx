import React from "react";

interface DisassemblyProps {
    index: number;
    instructions: Instruction[]
}

interface Instruction {
    address: number
    disassembly: string
}

const Disassembly: React.FC<DisassemblyProps> = (props) => {
    return (<p>Disassembly here</p>);
}

export default Disassembly;
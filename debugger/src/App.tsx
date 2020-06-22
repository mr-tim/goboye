import React from 'react';
import './App.css';

import Disassembly from './Disassembly';
import Display from './Display';
import MemoryView from './MemoryView';
import Registers from './Registers';
import { useState, useEffect } from 'react';
import { w3cwebsocket as W3CWebSocket } from 'websocket';

interface instruction {
  address: number
  disassembly: string
}

interface Message {
  update: {
    instructions?: instruction[]
    registers?: { [key: string]: number }
  }
}

function App() {
  let [registers, setRegisters] = useState<{[k: string]: number}>({
    AF: 0,
    BC: 0,
    DE: 0,
    HL: 0,
    SP: 0,
    PC: 0
  });
  let [instructions, setInstructions] = useState<instruction[]>([]);

  useEffect(() => {
    const client = new W3CWebSocket('ws://127.0.0.1:8080/ws');
    client.onopen = () => {
      console.log('Websocket connected');
    }
    client.onmessage = (e) => {
      if (typeof e.data === "string") {
        let message: Message = JSON.parse(e.data);
        if ('update' in message) {
          let update = message.update;
          if (update.registers !== undefined) {
            setRegisters(update.registers);
          }
          if (update.instructions !== undefined) {
            setInstructions(update.instructions);
          }
        }
      }
    }
  }, []);

  return (
    <div className="app">
      <div className="left-column">
        <Disassembly currentAddress={registers.PC} instructions={instructions} />
      </div>
      <div className="central-column">
        <Display />
        <MemoryView />
      </div>
      <div className="right-column">
        <Registers registers={registers} />
      </div>
    </div>
  );
}

export default App;

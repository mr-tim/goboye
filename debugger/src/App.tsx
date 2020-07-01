import React from 'react';
import './App.css';

import Disassembly from './components/Disassembly';
import Display from './components/Display';
import MemoryView from './components/MemoryView';
import Registers from './components/Registers';
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
    memory_base64?: string
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
  let [memory, setMemory] = useState<Array<number>>(new Array(0xffff));

  useEffect(() => {
    const client = new W3CWebSocket('ws://127.0.0.1:8080/ws');
    client.onopen = () => {
      console.log('Websocket connected');
        let handleKeyDown = (event: KeyboardEvent) => {
            if (event.key === ' ') {
                console.log("Sending step command");
                client.send(JSON.stringify({
                    command: {
                        step: {}
                    }
                }));
            }
        };

        document.addEventListener("keydown", handleKeyDown);
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
          if (update.memory_base64 !== undefined) {
            let decoded = window.atob(update.memory_base64);
            let newBuffer = new Array(decoded.length);
            for (var i = 0; i < decoded.length; i++) {
              newBuffer[i] = decoded.charCodeAt(i);
            }
            setMemory(newBuffer);
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
        <MemoryView memory={memory} />
      </div>
      <div className="right-column">
        <Registers registers={registers} />
      </div>
    </div>
  );
}

export default App;

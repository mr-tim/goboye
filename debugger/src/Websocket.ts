import {useEffect, useState} from 'react';
import {w3cwebsocket as W3CWebSocket} from "websocket";

interface instruction {
  address: number
  disassembly: string
}

interface Message {
  update: {
    instructions?: instruction[]
    registers?: { [key: string]: number }
    memory_base64?: string
    breakpoints?: number[]
  }
}

export function useWebsocket(): [boolean, instruction[], {[key:string]:number}, Array<number>, number[], (command:any)=>void] {
  let [isConnected, setIsConnected] = useState(false);
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
  let [breakpoints, setBreakpoints] = useState<number[]>([]);
  let [sendCommand, setSendCommand] = useState<(command:any)=>void>(() => (command: any) => {});

  useEffect(() => {
    const client = new W3CWebSocket('ws://127.0.0.1:8080/ws');
    client.onopen = () => {
      console.log('Websocket connected');
      setIsConnected(true);

      setSendCommand(() => (command: any) => {
        client.send(JSON.stringify({
          command: command
        }));
      });
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
          if (update.breakpoints !== undefined) {
            setBreakpoints(update.breakpoints);
          }
        }
      }
    }
  }, []);

  return [isConnected, instructions, registers, memory, breakpoints, sendCommand];
}

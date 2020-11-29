import {useEffect, useState} from 'react';
import {w3cwebsocket as W3CWebSocket} from "websocket";

interface instruction {
  address: number
  disassembly: string
}

interface memory_update {
  memory_base64: string
  start: number
  length: number
}

interface Message {
  update: {
    instructions?: instruction[]
    registers?: { [key: string]: number }
    memory_updates?: memory_update[]
    breakpoints?: number[]
    debug_image?: string
  }
}

export function useWebsocket(): [boolean, instruction[], {[key:string]:number}, Array<number>, number[], string, (command:any)=>void] {
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
  let [debugImage, setDebugImage] = useState<string>('');

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
          if (update.memory_updates !== undefined) {
            let buffer = memory;
            for (let i=0; i<update.memory_updates.length; i++) {
              let m = update.memory_updates[i];
              let decoded = window.atob(m.memory_base64);
              for (let idx=0; idx<m.length; idx++) {
                buffer[m.start+idx] = decoded.charCodeAt(idx);
              }
            }
            setMemory(buffer);
          }
          if (update.breakpoints !== undefined) {
            setBreakpoints(update.breakpoints);
          }
          if (update.debug_image !== undefined) {
            setDebugImage(update.debug_image);
          }
        }
      }
    }

    return () => {
      console.log('Closing websocket...');
      client.close();
    };
  }, [memory]);

  return [isConnected, instructions, registers, memory, breakpoints, debugImage, sendCommand];
}

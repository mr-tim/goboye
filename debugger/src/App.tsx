import React from 'react';
import './App.css';

import Disassembly from './components/Disassembly';
import Display from './components/Display';
import MemoryView from './components/MemoryView';
import Registers from './components/Registers';

import {useWebsocket} from "./Websocket";

function App() {
  let [isConnected, instructions, registers, memory, breakpoints, flags, debugRender, sendCommand] = useWebsocket();

  let onKeyDown = (event: KeyboardEvent) => {
    if (event.key === ' ') {
      sendCommand({
        step: {}
      });
      event.preventDefault();
    } else if (event.key === 'Enter') {
      sendCommand({
        continue: {}
      })
    }
  };

  React.useEffect(() => {
    document.addEventListener('keydown', onKeyDown);
    return () => {
      document.removeEventListener('keydown', onKeyDown);
    };
  })

  let setBreakpoint = (address: number, isBreakpoint: boolean) => {
    sendCommand({
      breakpoint: {
        address: address,
        break: isBreakpoint
      }
    })
  }

  return (
      <div className="app">
        {isConnected && (
            <>
              <div className="central-column">
                <Display debugRender={debugRender}/>
                <Disassembly currentAddress={registers.PC} instructions={instructions} breakpoints={breakpoints} setBreakpoint={setBreakpoint}/>
              </div>
              <div className="right-column">
                <Registers registers={registers} flags={flags}/>
              </div>
            </>
        )}
        {!isConnected && (
            <p>Connecting...</p>
        )}
      </div>
  );
}

export default App;

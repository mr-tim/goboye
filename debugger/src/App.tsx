import React from 'react';
import './App.css';

import Disassembly from './components/Disassembly';
import Display from './components/Display';
import MemoryView from './components/MemoryView';
import Registers from './components/Registers';

import {useWebsocket} from "./Websocket";

function App() {
  let [isConnected, instructions, registers, memory, sendCommand] = useWebsocket();

  let onKeyDown = (event: KeyboardEvent) => {
    if (event.key === ' ') {
      sendCommand({
        step: {}
      });
      event.preventDefault();
    }
  };

  React.useEffect(() => {
    document.addEventListener('keydown', onKeyDown);
    return () => {
      document.removeEventListener('keydown', onKeyDown);
    };
  })

  return (
      <div className="app">
        {isConnected && (
            <>
              <div className="left-column">
                <Disassembly currentAddress={registers.PC} instructions={instructions}/>
              </div>
              <div className="central-column">
                <Display/>
                <MemoryView memory={memory}/>
              </div>
              <div className="right-column">
                <Registers registers={registers}/>
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

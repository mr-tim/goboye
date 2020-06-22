import React from 'react';
import './App.css';

import Disassembly from './Disassembly';
import Display from './Display';
import Registers from './Registers';

import state from './testState.json';

function App() {
  return (
    <div className="app">
      <div className="left-column">
        <Disassembly currentAddress={state.registers.PC} instructions={state.instructions} />
      </div>
      <div className="central-column">
        <Display />
      </div>
      <div className="right-column">
        <Registers registers={state.registers} />
      </div>
    </div>
  );
}

export default App;

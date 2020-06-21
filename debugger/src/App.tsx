import React from 'react';
import './App.css';

import Disassembly from './Disassembly';
import Display from './Display';
import Registers from './Registers';

function App() {
  const instructions = [
    {
      address: 0x0000,
      disassembly: 'LD SP,0xFFFE'
    }
  ];

  return (
    <div className="app">
      <div className="left-column">
        <Disassembly index={0} instructions={instructions} />
      </div>
      <div className="central-column">
        <Display />
      </div>
      <div className="right-column">
        <Registers />
      </div>
    </div>
  );
}

export default App;

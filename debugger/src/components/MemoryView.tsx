import React from 'react';
import styled from 'styled-components';
import {toHex} from '../util/hex';

interface range {
  start: number;
  end: number;
}

interface offsetProps {
  address: number
}

const Offset: React.FC<offsetProps> = (props) => {
  return (
      <MemoryViewCol>{toHex(props.address, 4)}</MemoryViewCol>
  );
}

interface hexViewProps {
  values: Array<number>
}

const HexView: React.FC<hexViewProps> = (props) => {
  return <MemoryViewCol>{props.values.map(v => toHex(v, 2, false)).join(' ')}</MemoryViewCol>
}

const MemoryViewRow = styled.div`
    display: flex;
    justify-content: flex-start;
    align-items: baseline
    background-color: #ccc;
`

const MemoryViewCol = styled.div`
    margin-right: 18px;
    &:nth-child(3n) {
        margin-right: 0px;
    }
`

const MemoryViewContainer = styled.div`
    padding: 12px;
    padding-top: 24px;
    height: 100%;
    width: 705px;
    overflow-y: scroll;
    margin-left: auto;
    margin-right: auto;
`

interface memoryViewProps {
  memory: Array<number>
};

const MemoryView: React.FC<memoryViewProps> = (props) => {
  let rows = [];
  for (let i = 0; i < 0x0400; i += 0x10) {
    rows.push(
        <MemoryViewRow key={i} className="monospaced">
          <Offset address={i}/>
          <HexView values={props.memory.slice(i, i + 16)}/>
        </MemoryViewRow>
    );
  }

  return (<MemoryViewContainer>
    {rows}
  </MemoryViewContainer>);
}

export default MemoryView;

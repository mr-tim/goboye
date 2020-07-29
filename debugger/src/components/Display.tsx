import React from "react";
import styled from "styled-components";

const ScreenDiv = styled.div`
    background-color: #9bbc0f;
    width: 256px;
    height: 256px;
    margin-left: auto;
    margin-right: auto;
    margin-top: 30px;
    border: 12px solid #ccc;
    flex-grow: 0;
    flex-shrink: 0;
`;

interface DisplayProps {
  debugRender: string
}

const Display: React.FC<DisplayProps> = (props) => {
  let imgSrc = props.debugRender === ''?
      undefined : 'data:image/png;base64,' + props.debugRender;
  return (<ScreenDiv>
    {props.debugRender !== undefined &&
      <img src={imgSrc}/>
    }
  </ScreenDiv>);
}

export default Display;

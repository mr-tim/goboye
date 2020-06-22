import React from "react";
import styled from "styled-components";

const ScreenDiv = styled.div`
    background-color: #9ca04c;
    width: 320px;
    height: 288px;
    margin-left: auto;
    margin-right: auto;
    margin-top: 30px;
    border: 12px solid #ccc;
`;

const Display: React.FC<{}> = (props) => {
    return (<ScreenDiv>
    </ScreenDiv>);
}

export default Display;
import styled from "styled-components";

import React, { FunctionComponent } from "react";

import { Card } from "antd";

const Span = styled.span`
  font-size: 20px;
  /* color: white; */
`;

const Div = styled.div`
  display: flex;
  justify-content: space-between;
  /* color: white; */
  align-items: center;
`;

interface IProps {
  children: React.ReactNodeArray;
  style: any;
}

// interface CardContent {
//   title: String
//   total: any
// }

const RenderCard: FunctionComponent<IProps> = ({ style, children }) => {
  return (
    <Card
      bordered={true}
      style={{
        width: 340,
        backgroundColor: style,
        borderRadius: 4,
        height: 185
      }}
    >
      {children}
    </Card>
  );
};

export const renderCardContent = (content: any, index: any) => (
  <RenderCard style={content.backgroundColor} key={index}>
    <Span> {content.title} </Span>
    <Div>
      <span style={{ paddingTop: 20, paddingBottom: 12, fontSize: 24 }}>
        {content.total }
      </span>
      <span>{content.pending || content.day}</span>
    </Div>
    {/* <Div>
      <p> {content.unlisted ? "unlisted" : "in last in week"}</p>
      <p>{content.unlisted || content.week}</p>
    </Div>
    <Div>
      <p>{content.live ? "live" : "in last in month"}</p>
      <p>{content.live || content.month}</p>
    </Div> */}
  </RenderCard>
);

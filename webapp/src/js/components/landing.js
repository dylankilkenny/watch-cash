import React from 'react';
import ReactDOM from 'react-dom';
import { Columns, Column, Title, Subtitle } from 'bloomer';

const signup = () => (
  <Columns style={{ margin: 10 }} isCentered>
    <Column isSize="1/2">
      <div>hey </div>
    </Column>
    <Column isSize="1/2">
      <Title style={{ color: 'white' }} isSize={3}>
        Watch a bitcoin cash address
      </Title>
      <Subtitle style={{ color: '#d9dce0' }} isSize={4}>
        Recieve an email notification for all incoming and outgoin transactions.
      </Subtitle>
    </Column>
  </Columns>
);

export default signup;

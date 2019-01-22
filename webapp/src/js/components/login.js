import React from 'react';
import ReactDOM from 'react-dom';
import {
  Columns,
  Column,
  Field,
  Box,
  Control,
  Label,
  Input,
  Button
} from 'bloomer';

const login = () => (
  <Columns style={{ marginTop: 10 }} isCentered>
    <Column isSize="1/4" />
    <Column isSize="1/2">
      <Box>
        <Field>
          <Label>Email</Label>
          <Control>
            <Input type="email" placeholder="Email" />
          </Control>
        </Field>
        <Field>
          <Label>Password</Label>
          <Control>
            <Input type="password" placeholder="Password" />
          </Control>
        </Field>
        <Field>
          <Control>
            <Button isColor="primary">Submit</Button>
          </Control>
        </Field>
      </Box>
    </Column>
    <Column isSize="1/4" />
  </Columns>
);

export default login;

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

const signup = () => (
  <Columns style={{ marginTop: 10 }} isCentered>
    <Column isSize="1/4" />
    <Column isSize="1/2">
      <Box>
        <Field>
          <Label>First Name</Label>
          <Control>
            <Input type="text" placeholder="First Name" />
          </Control>
        </Field>
        <Field>
          <Label>Last Name</Label>
          <Control>
            <Input type="text" placeholder="Last Name" />
          </Control>
        </Field>
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

export default signup;

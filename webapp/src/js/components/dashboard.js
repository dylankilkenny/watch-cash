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
  Button,
  Table,
  Icon
} from 'bloomer';

const dashboard = () => (
  <Columns style={{ marginTop: 10 }} isCentered>
    <Column isCentered isSize="2/3">
      <Column>
        <Box>
          <Field>
            <Control>
              <Columns isCentered>
                <Column>
                  <Input
                    isSize="large"
                    type="text"
                    placeholder="Bitcoin Cash Address"
                  />
                </Column>
                <Column isSize="narrow">
                  <Button
                    style={{ marginTop: 4 }}
                    isSize="medium"
                    isHovered
                    isColor="primary"
                  >
                    Add
                  </Button>
                </Column>
              </Columns>
            </Control>
          </Field>
          <Table style={{ width: '100%' }}>
            <thead>
              <tr>
                <th>Watched Adresses</th>
                <th />
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>qpt26jvyqs6gwhm92w9m9gu4r2h8x04n6skfhl35qg</td>
                <td>
                  <Button isColor="danger">Remove</Button>
                </td>
              </tr>
              <tr>
                <td>qpt26jvyqs6gwhm92w9m9gu4r2h8x04n6skfhl35qg</td>
                <td>
                  <Icon isSize="medium" className="fa fa-times  fa-2x" />
                </td>
              </tr>
              <tr>
                <td>qpt26jvyqs6gwhm92w9m9gu4r2h8x04n6skfhl35qg</td>
                <td>
                  <Icon isSize="medium" className="fa fa-times  fa-2x" />
                </td>
              </tr>
            </tbody>
          </Table>
        </Box>
      </Column>
    </Column>
  </Columns>
);

export default dashboard;

import React from 'react';
import ReactDOM from 'react-dom';
import {
  Panel,
  Grid,
  Row,
  Col,
  FormGroup,
  FormControl,
  Button
} from 'react-bootstrap';

const signup = () => (
  <Grid>
    <Row className="show-grid">
      <Col xs={12} md={3} />
      <Col xs={12} md={6}>
        <Panel>
          <Panel.Heading>
            <Panel.Title componentClass="h3">Login</Panel.Title>
          </Panel.Heading>
          <Panel.Body>
            <FormGroup>
              <FormControl type="text" placeholder="First Name" />
            </FormGroup>
            <FormGroup>
              <FormControl type="text" placeholder="Last Name" />
            </FormGroup>
            <FormGroup>
              <FormControl type="email" placeholder="Email" />
            </FormGroup>
            <FormGroup>
              <FormControl type="password" placeholder="Password" />
            </FormGroup>
            <FormGroup>
              <Button bsStyle="primary">Submit</Button>
            </FormGroup>
          </Panel.Body>
        </Panel>
      </Col>
      <Col xs={12} md={3} />
    </Row>
  </Grid>
);

export default signup;

import React from 'react';
import ReactDOM from 'react-dom';
import { Grid, Row, Col } from 'react-bootstrap';

const signup = () => (
  <Grid>
    <Row className="show-grid">
      <Col xs={12} md={6} />
      <Col xs={12} md={6}>
        <h1 class="landing-header">Watch a bitcoin cash address</h1>
        <h2 class="landing-subtitle">
          Recieve an email notification for all incoming and outgoin
          transactions.
        </h2>
      </Col>
    </Row>
  </Grid>
);

export default signup;

/* <Container>
    <Row>
      <Col md="6">.col-6 .col-sm-4</Col>
      <Col md="6">
        <h1 class="landing-header">Watch a bitcoin cash address</h1>
        <h2 class="landing-subtitle">
          Recieve an email notification for all incoming and outgoin
          transactions.
        </h2>
      </Col>
    </Row>
  </Container> */

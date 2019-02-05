import React from 'react';
import ReactDOM from 'react-dom';
import { Container, Row, Col } from 'react-bootstrap';
import NavBar from '../container/navbar';

const signup = () => (
  <div>
    <NavBar />
    <Container>
      <Row className="show-grid">
        <Col xs={12} md={6} />
        <Col xs={12} md={6}>
          <div className="card-pad">
            <h1 class="landing-header">Watch a bitcoin cash address</h1>
            <h2 class="landing-subtitle">
              Recieve an email notification for all incoming and outgoin
              transactions.
            </h2>
          </div>
        </Col>
      </Row>
    </Container>
  </div>
);

export default signup;

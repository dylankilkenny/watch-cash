import React from 'react';
import ReactDOM from 'react-dom';
import { Card, Container, Row, Col, Button } from 'react-bootstrap';
import Form from 'react-bootstrap/Form';
import NavBar from './navbar';
import User from '../../utility/user';
import ConditionalAlert from '../presentational/alert';

export default class SignUp extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      firstname: '',
      lastname: '',
      email: '',
      password: ''
    };
    this.submitSignUp = this.submitSignUp.bind(this);
  }

  updateFormValue(evt, element) {
    this.setState({
      [element]: evt.target.value
    });
  }

  async submitSignUp() {
    let success = await User.Signup(
      this.state.firstname,
      this.state.lastname,
      this.state.email,
      this.state.password
    );
    if (success == true) {
      this.props.history.push('/login');
    } else {
      if (success == 'invalid_form') {
        this.setState({
          warning: true,
          warningDetail: 'Invalid form - some required fields are missing.'
        });
      }
      if (success == 'email_taken') {
        this.setState({
          warning: true,
          warningDetail: 'An account with this email already exists.'
        });
      }
    }
  }

  render() {
    return (
      <div>
        <NavBar />
        <Container>
          <Row>
            <Col xs={12} md={3} />
            <Col xs={12} md={6}>
              <ConditionalAlert
                isActive={this.state.warning}
                hasStyle="danger"
                content={this.state.warningDetail}
              />
              <Card className="card-pad">
                <Card.Header>
                  <Card.Title as="h4">Sign Up</Card.Title>
                </Card.Header>
                <Card.Body>
                  <Form>
                    <Form.Group controlId="formGroupFirstName">
                      <Form.Label>First Name</Form.Label>
                      <Form.Control
                        value={this.state.firstname}
                        onChange={evt => this.updateFormValue(evt, 'firstname')}
                        type="text"
                        placeholder="Enter First Name"
                      />
                    </Form.Group>
                    <Form.Group controlId="formGroupLastName">
                      <Form.Label>Last Name</Form.Label>
                      <Form.Control
                        value={this.state.lastname}
                        onChange={evt => this.updateFormValue(evt, 'lastname')}
                        type="text"
                        placeholder="Enter Last Name"
                      />
                    </Form.Group>
                    <Form.Group controlId="formGroupEmail">
                      <Form.Label>Email address</Form.Label>
                      <Form.Control
                        value={this.state.email}
                        onChange={evt => this.updateFormValue(evt, 'email')}
                        type="email"
                        placeholder="Enter email"
                      />
                    </Form.Group>
                    <Form.Group controlId="formGroupPassword">
                      <Form.Label>Password</Form.Label>
                      <Form.Control
                        value={this.state.password}
                        onChange={evt => this.updateFormValue(evt, 'password')}
                        type="password"
                        placeholder="Password"
                      />
                    </Form.Group>
                  </Form>
                  <Button onClick={this.submitSignUp} variant="primary">
                    Submit
                  </Button>
                </Card.Body>
              </Card>
            </Col>
            <Col xs={12} md={3} />
          </Row>
        </Container>
      </div>
    );
  }
}

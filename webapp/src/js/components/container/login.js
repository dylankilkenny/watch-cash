import React from 'react';
import User from '../../utility/user';
import ConditionalAlert from '../presentational/alert';
import Container from 'react-bootstrap/Container';
import Card from 'react-bootstrap/Card';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Form from 'react-bootstrap/Form';
import FormControl from 'react-bootstrap/FormControl';
import Button from 'react-bootstrap/Button';
import NavBar from './navbar';

export default class Login extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      email: '',
      password: ''
    };
    this.submitLogin = this.submitLogin.bind(this);
    this.updateFormValue = this.updateFormValue.bind(this);
  }

  async submitLogin() {
    let success = await User.Login(this.state.email, this.state.password);
    if (success) {
      this.props.history.push('/dashboard');
    } else {
      this.setState({ warning: true });
    }
  }

  updateFormValue(evt, element) {
    this.setState({
      [element]: evt.target.value
    });
  }

  render() {
    return (
      <div>
        <NavBar logout={this.props.logout} />
        <Container>
          <Row className="show-grid">
            <Col xs={12} md={3} />
            <Col xs={12} md={6}>
              <ConditionalAlert
                isActive={this.state.warning}
                hasStyle="danger"
                content="Uh oh! Your email/password is incorrect"
              />
              <Card className="card-pad">
                <Card.Header>
                  <Card.Title as="h4">Login</Card.Title>
                </Card.Header>
                <Card.Body>
                  <Form>
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
                  <Button onClick={this.submitLogin} variant="primary">
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

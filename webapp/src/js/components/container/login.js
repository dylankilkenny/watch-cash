import React from 'react';
import ReactDOM from 'react-dom';
import validateToken from '../../utility/validateToken';
import ConditionalAlert from '../presentational/alert';
import {
  Panel,
  Grid,
  Row,
  Col,
  FormGroup,
  FormControl,
  Button
} from 'react-bootstrap';

export default class Login extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      email: '',
      password: ''
    };
    this.submitLogin = this.submitLogin.bind(this);
    this.updateEmailValue = this.updateEmailValue.bind(this);
    this.updatePasswordValue = this.updatePasswordValue.bind(this);
  }

  componentDidMount() {
    const token = sessionStorage.getItem('token');
    if (token != undefined) {
      var resp = validateToken(token);
      if (resp.status != 200) {
        this.props.history.push('/login');
      }
    }
  }

  submitLogin() {
    fetch('http://0.0.0.0:3001/api/login', {
      method: 'POST',
      body: JSON.stringify({
        email: this.state.email,
        password: this.state.password
      }),
      headers: { 'Content-Type': 'application/json' }
    })
      .then(response => {
        return response.json();
      })
      .then(json => {
        if (json.status == 200) {
          sessionStorage.setItem('token', json.token);
          sessionStorage.setItem('firstname', json.firstname);
          this.props.history.push('/dashboard');
        } else {
          this.setState({ warning: true });
        }
      });
  }

  updateEmailValue(evt) {
    this.setState({
      email: evt.target.value
    });
  }

  updatePasswordValue(evt) {
    this.setState({
      password: evt.target.value
    });
  }

  render() {
    return (
      <Grid>
        <Row className="show-grid">
          <Col xs={12} md={3} />
          <Col xs={12} md={6}>
            <ConditionalAlert
              isActive={this.state.warning}
              hasStyle="danger"
              content="Uh oh! Your email/password is incorrect"
            />
            <Panel>
              <Panel.Heading>
                <Panel.Title componentClass="h3">Login</Panel.Title>
              </Panel.Heading>
              <Panel.Body>
                <FormGroup>
                  <FormControl
                    value={this.state.email}
                    onChange={evt => this.updateEmailValue(evt)}
                    type="email"
                    placeholder="Email"
                  />
                </FormGroup>
                <FormGroup>
                  <FormControl
                    value={this.state.password}
                    onChange={evt => this.updatePasswordValue(evt)}
                    type="password"
                    placeholder="Password"
                  />
                </FormGroup>
                <FormGroup>
                  <Button onClick={this.submitLogin} bsStyle="primary">
                    Submit
                  </Button>
                </FormGroup>
              </Panel.Body>
            </Panel>
          </Col>
          <Col xs={12} md={3} />
        </Row>
      </Grid>
    );
  }
}

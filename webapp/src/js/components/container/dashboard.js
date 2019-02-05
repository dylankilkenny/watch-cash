import React from 'react';
import ReactDOM from 'react-dom';
import Container from 'react-bootstrap/Container';
import Card from 'react-bootstrap/Card';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import FormGroup from 'react-bootstrap/FormGroup';
import FormControl from 'react-bootstrap/FormControl';
import InputGroup from 'react-bootstrap/InputGroup';
import Button from 'react-bootstrap/Button';
import NavBar from './navbar';
import User from '../../utility/user';
import AddressTable from '../../components/presentational/table';

export default class Login extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      firstname: sessionStorage.getItem('firstname'),
      token: sessionStorage.getItem('token')
    };
  }

  componentDidMount() {
    (async () => {
      let resp = await User.Validate(this.state.token);
      console.log(resp);
      if (resp.status == 200) {
        let subscribedAddresses = await User.SubscribedAddresses(
          this.state.token
        );
        this.setState({
          subscribedAddresses: subscribedAddresses,
          fetched: true
        });
      } else {
        sessionStorage.removeItem('firstname');
        sessionStorage.removeItem('token');
        this.props.history.push('/login');
      }
    })().catch(error => {
      console.log(error);
      this.props.history.push('/login');
    });
  }

  render() {
    if (!this.state.fetched) {
      return <div />;
    }
    return (
      <div>
        <NavBar />
        <Container>
          <Row className="show-grid">
            <Col xs={12} md={2} />
            <Col xs={12} md={8}>
              <Card className="card-pad">
                <Card.Body>
                  <FormGroup bsSize="large">
                    <InputGroup className="mb-3">
                      <FormControl
                        placeholder="Bitcoin Cash Address"
                        aria-label="Bitcoin Cash Address"
                        aria-describedby="basic-addon2"
                      />
                      <InputGroup.Append>
                        <Button variant="outline-secondary">Watch</Button>
                      </InputGroup.Append>
                    </InputGroup>
                    <AddressTable
                      subscribedAddresses={this.state.subscribedAddresses}
                    />
                  </FormGroup>
                </Card.Body>
              </Card>
            </Col>
            <Col xs={12} md={2} />
          </Row>
        </Container>
      </div>
    );
  }
}

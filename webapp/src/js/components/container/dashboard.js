import React from 'react';
import ReactDOM from 'react-dom';
import Container from 'react-bootstrap/Container';
import Card from 'react-bootstrap/Card';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import FormControl from 'react-bootstrap/FormControl';
import InputGroup from 'react-bootstrap/InputGroup';
import Button from 'react-bootstrap/Button';
import NavBar from './navbar';
import User from '../../utility/user';
import AddressTable from '../../components/presentational/table';
import user from '../../utility/user';

export default class Login extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      firstname: sessionStorage.getItem('firstname'),
      token: sessionStorage.getItem('token')
    };
    this.watchAddress = this.watchAddress.bind(this);
    this.updateSubscribedAddresses = this.updateSubscribedAddresses.bind(this);
    this.removeAddress = this.removeAddress.bind(this);
  }

  componentDidMount() {
    (async () => {
      let resp = await User.Validate(this.state.token);
      console.log(resp);
      if (resp.status == 200) {
        this.updateSubscribedAddresses();
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

  updateSubscribedAddresses() {
    (async () => {
      let subscribedAddresses = await User.SubscribedAddresses(
        this.state.token
      );
      this.setState({
        subscribedAddresses: subscribedAddresses ? subscribedAddresses : [],
        fetched: true
      });
    })().catch(error => {
      console.log(error);
      this.props.history.push('/login');
    });
  }

  updateFormValue(evt) {
    this.setState({
      address: evt.target.value
    });
  }

  async watchAddress() {
    let resp = await user.WatchAddress(this.state.address, this.state.token);
    if (resp.status != 200) {
      console.log('Watch Address Err: ', resp.detail);
    }
    console.log(resp);
    this.updateSubscribedAddresses();
  }

  async removeAddress(address) {
    let resp = await user.RemoveAddress(address, this.state.token);
    if (resp.status != 200) {
      console.log('Remove Address Err: ', resp.detail);
    }
    this.updateSubscribedAddresses();
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
                  <InputGroup className="mb-3">
                    <FormControl
                      placeholder="Bitcoin Cash Address"
                      aria-label="Bitcoin Cash Address"
                      aria-describedby="basic-addon2"
                      onChange={evt => this.updateFormValue(evt)}
                    />
                    <InputGroup.Append>
                      <Button
                        onClick={this.watchAddress}
                        variant="outline-secondary"
                      >
                        Watch
                      </Button>
                    </InputGroup.Append>
                  </InputGroup>
                  <AddressTable
                    subscribedAddresses={this.state.subscribedAddresses}
                    removeAddress={this.removeAddress}
                  />
                </Card.Body>
              </Card>
            </Col>
            <Col xs={12} md={2} />
          </Row>
        </Container>
      </div>
    );
  }
  componentDidCatch(error, info) {
    console.log(error);
    console.log(info);
  }
}

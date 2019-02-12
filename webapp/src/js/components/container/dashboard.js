import React from 'react';
import Container from 'react-bootstrap/Container';
import Card from 'react-bootstrap/Card';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import FormControl from 'react-bootstrap/FormControl';
import ConditionalAlert from '../presentational/alert';
import InputGroup from 'react-bootstrap/InputGroup';
import Button from 'react-bootstrap/Button';
import NavBar from './navbar';
import User from '../../utility/user';
import Address from '../../utility/address';
import AddressTable from '../../components/presentational/table';

export default class Login extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            firstname: sessionStorage.getItem('firstname'),
            token: sessionStorage.getItem('token')
        };
        this.watchAddress = this.watchAddress.bind(this);
        this.updateSubscribedAddresses = this.updateSubscribedAddresses.bind(
            this
        );
        this.removeAddress = this.removeAddress.bind(this);
    }

    async componentDidMount() {
        let resp = await User.Validate(this.state.token);
        if (resp.status == 200) {
            this.updateSubscribedAddresses();
        } else {
            sessionStorage.removeItem('firstname');
            sessionStorage.removeItem('token');
            this.props.history.push('/login');
        }
    }

    async updateSubscribedAddresses() {
        let subscribedAddresses = await User.SubscribedAddresses(
            this.state.token
        );
        this.setState({
            subscribedAddresses: subscribedAddresses ? subscribedAddresses : [],
            fetched: true
        });
    }

    updateFormValue(evt) {
        this.setState({
            address: evt.target.value
        });
    }

    async watchAddress() {
        let isValidAddress = Address.ValidAddress(this.state.address);
        if (isValidAddress) {
            let address = Address.CashAddress(this.state.address);
            let resp = await User.WatchAddress(address, this.state.token);
            if (resp.status == 400 && resp.code == 'already_subscribed') {
                this.setState({
                    showAlert: true,
                    alertMessage: 'You are already subscribed to this address',
                    alertStyle: 'warning'
                });
            } else if (resp.status == 200) {
                this.setState({ showAlert: false, address: '' }, () => {
                    this.updateSubscribedAddresses();
                });
            }
        } else {
            this.setState({
                showAlert: true,
                alertMessage:
                    'You have not entered a valid bitcoin cash address',
                alertStyle: 'danger'
            });
        }
    }

    async removeAddress(address) {
        let resp = await User.RemoveAddress(address, this.state.token);
        if (resp.status != 200) {
            this.setState({
                showAlert: true,
                alertMessage:
                    'Something went wrong! Failed to remove address from watch list',
                alertStyle: 'danger'
            });
            console.log(resp);
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
                            <ConditionalAlert
                                isActive={this.state.showAlert}
                                hasStyle={this.state.alertStyle}
                                content={this.state.alertMessage}
                            />
                            <Card className="card-pad">
                                <Card.Body>
                                    <InputGroup className="mb-3">
                                        <FormControl
                                            placeholder="Bitcoin Cash Address"
                                            aria-label="Bitcoin Cash Address"
                                            aria-describedby="basic-addon2"
                                            onChange={evt =>
                                                this.updateFormValue(evt)
                                            }
                                            value={this.state.address}
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
                                        subscribedAddresses={
                                            this.state.subscribedAddresses
                                        }
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
        this.props.history.push('/login');
    }
}

import React from 'react';
import ReactDOM from 'react-dom';
import validateToken from '../../utility/validateToken';
import {
  Panel,
  Grid,
  Row,
  Col,
  FormGroup,
  FormControl,
  Button,
  InputGroup
} from 'react-bootstrap';

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
      let resp = await validateToken(this.state.token);
      console.log(resp);
      if (resp.status != 200) {
        this.props.history.push('/login');
      }
    })().catch(error => {
      console.log(error);
    });
  }

  render() {
    return (
      <Grid>
        <Row className="show-grid">
          <Col xs={12} md={3} />
          <Col xs={12} md={6}>
            <Panel>
              <Panel.Body>
                <FormGroup bsSize="large">
                  <InputGroup>
                    <FormControl
                      type="text"
                      placeholder="Bitcoin Cash Address"
                    />
                    <InputGroup.Button>
                      <Button bsSize="large" bsStyle="primary">
                        Watch
                      </Button>
                    </InputGroup.Button>
                  </InputGroup>
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

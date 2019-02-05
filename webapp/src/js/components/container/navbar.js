import React from 'react';
import Navbar from 'react-bootstrap/Navbar';
import Nav from 'react-bootstrap/Nav';

import { Link, Redirect } from 'react-router-dom';

export default class NavBar extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      LoggedIn: false
    };
    this.logout = this.logout.bind(this);
  }
  componentDidMount() {
    let user = sessionStorage.getItem('firstname');
    let token = sessionStorage.getItem('token');
    if ((user != undefined) & (token != undefined)) {
      this.setState({ LoggedIn: true, firstName: user });
    }
  }

  logout() {
    sessionStorage.removeItem('firstname');
    sessionStorage.removeItem('token');
    this.setState({ redirect: true });
  }

  render() {
    if (this.state.redirect) {
      return <Redirect to="/" />;
    }
    return (
      <Navbar collapseOnSelect expand="lg" bg="dark" variant="dark">
        <Navbar.Brand>
          <Link className="nav-text" to="/">
            watch.cash
          </Link>
        </Navbar.Brand>
        <Navbar.Toggle aria-controls="responsive-navbar-nav" />
        <Navbar.Collapse
          id="responsive-navbar-nav"
          className="justify-content-end"
        >
          <NavItems
            firstName={this.state.firstName}
            LoggedIn={this.state.LoggedIn}
            logout={this.logout}
          />
        </Navbar.Collapse>
      </Navbar>
    );
  }
}

function NavItems(props) {
  if (props.LoggedIn) {
    return (
      <Nav>
        <Navbar.Text>Welcome, {props.firstName}!</Navbar.Text>
        <Nav.Link eventKey={2}>
          <Link onClick={props.logout} className="nav-text" to="/">
            Logout
          </Link>
        </Nav.Link>
      </Nav>
    );
  } else {
    return (
      <Nav>
        <Nav.Link>
          <Link className="nav-text" to="/login">
            Login
          </Link>
        </Nav.Link>
        <Nav.Link eventKey={2}>
          <Link className="nav-text" to="/signup">
            Signup
          </Link>
        </Nav.Link>
      </Nav>
    );
  }
}

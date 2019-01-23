import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import {
  Navbar,
  NavbarBrand,
  NavbarBurger,
  NavbarMenu,
  Title,
  NavbarItem,
  NavbarStart,
  NavbarEnd,
  Field,
  Icon,
  Control,
  Button
} from 'bloomer';

class Navigationbar extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isActive: true
    };

    this.onClickNav = this.onClickNav.bind(this);
  }

  componentDidMount() {}

  componentDidCatch(error, info) {
    console.log(error);
    console.log(info);
  }

  onClickNav = () => {
    this.setState({ isActive: this.state.isActive == true ? false : true });
  };

  render() {
    return (
      <Navbar style={{ margin: '0', backgroundColor: 'rgb(0, 0, 0, 0)' }}>
        <NavbarBrand>
          <NavbarItem>
            <Link
              to={{
                pathname: '/'
              }}
            >
              <Title style={{ color: 'white' }} isSize={3}>
                watch.cash
              </Title>
            </Link>
          </NavbarItem>
          <NavbarBurger
            isActive={this.state.isActive}
            onClick={this.onClickNav}
          />
        </NavbarBrand>
        <NavbarMenu isActive={this.state.isActive} onClick={this.onClickNav}>
          <NavbarEnd>
            <NavbarItem>
              <Field>
                <Link
                  to={{
                    pathname: '/login'
                  }}
                >
                  <Button
                    style={{ margin: 10 }}
                    isColor="light"
                    isOutlined
                    isSize="medium"
                  >
                    Login
                  </Button>
                </Link>
                <Link
                  to={{
                    pathname: '/signup'
                  }}
                >
                  <Button
                    style={{ margin: 10 }}
                    isColor="light"
                    isOutlined
                    isSize="medium"
                  >
                    Sign Up
                  </Button>
                </Link>
              </Field>
            </NavbarItem>
          </NavbarEnd>
        </NavbarMenu>
      </Navbar>
    );
  }
}
export default Navigationbar;

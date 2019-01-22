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
      <Navbar style={{ border: 'solid 1px ', margin: '0' }}>
        <NavbarBrand>
          <NavbarItem>
            <Title isSize={3}>watch.cash</Title>
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
                  <Button isColor="success" isOutlined>
                    Login
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

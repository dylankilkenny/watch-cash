import 'bulma/css/bulma.css';
import ReactDOM from 'react-dom';
import React from 'react';
import {
  Switch,
  Route,
  withRouter,
  BrowserRouter as Router
} from 'react-router-dom';
import withAnalytics, { initAnalytics } from 'react-with-analytics';
import NavBar from './js/components/navbar';
import SignUp from './js/components/signup';
import Login from './js/components/login';

// initAnalytics(GA_KEY);

const Routes = () => (
  <Switch>
    <Route exact path="/" component={SignUp} />
    <Route exact path="/login" component={Login} />
  </Switch>
);

const AppWithRouter = withRouter(withAnalytics(Routes));

class Main extends React.Component {
  render() {
    return (
      <Router>
        <div>
          <NavBar />
          <AppWithRouter />
        </div>
      </Router>
    );
  }
}

ReactDOM.render(<Main />, document.getElementById('root'));

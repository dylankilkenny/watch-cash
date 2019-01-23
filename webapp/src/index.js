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
import Landing from './js/components/landing';
import Login from './js/components/login';
import Dashboard from './js/components/dashboard';

// initAnalytics(GA_KEY);

const Routes = () => (
  <Switch>
    <Route exact path="/" component={Landing} />
    <Route exact path="/login" component={Login} />
    <Route exact path="/signup" component={SignUp} />
    <Route exact path="/dashboard" component={Dashboard} />
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

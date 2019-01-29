import './css/index.css';
import ReactDOM from 'react-dom';
import React from 'react';
import {
  Switch,
  Route,
  withRouter,
  BrowserRouter as Router
} from 'react-router-dom';
import withAnalytics, { initAnalytics } from 'react-with-analytics';
import NavBar from './js/components/container/navbar';
import SignUp from './js/components/container/signup';
import Landing from './js/components/presentational/landing';
import Login from './js/components/container/login';
import Dashboard from './js/components/container/dashboard';

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

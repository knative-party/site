import React from "react";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";

import Events from './Events';
import PartyPage from './PartyPage';

function Routes() {
  return (
  <Router>
    <Switch>
      <Route path="/party">
        <PartyPage />
      </Route>
      <Route path="/">
        <Events />
      </Route>
    </Switch>
  </Router>
  );
}

export default Routes;
import React, { Suspense } from 'react';
import { Route, Switch, withRouter } from 'react-router-dom';

import Layout from './containers/Layout';
import Applications from './components/Applications';
import Profile from './components/Profile';

/* TODO Load data like UserStore.loadData() */
/*
  loadData () {
    if (document.getElementById('data')) {
      _setupData = JSON.parse(document.getElementById('data').innerHTML);
    }
  }
*/

const app = props => {
  const routes = (
    <Switch>
      <Route path="/applications" exact component={Applications} />
      <Route path="/profile" exact component={Profile} />
    </Switch>
  );

  return (
    <Layout>
      <Suspense fallback={<p>Pending...</p>}>{routes}</Suspense>
    </Layout>
  );
};

export default withRouter(app);

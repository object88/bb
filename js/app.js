// @flow
import 'babel-polyfill';

import React from 'react';
import ReactDOM from 'react-dom';
import ReactRouterRelay from 'react-router-relay';
import Relay from 'react-relay';
import { IndexRoute, Route, Router, browserHistory } from 'react-router';
import applyRouterMiddleware from 'react-router/lib/applyRouterMiddleware';
import useRelay from 'react-router-relay';

import AppContainer from './containers/AppContainer';
import PhotoListContainer from './containers/PhotoListContainer';
import ViewerQueries from './queries/ViewerQueries';

ReactDOM.render(
  <Router
      createElement={ReactRouterRelay.createElement}
      environment={Relay.Store}
      history={browserHistory}
      render={applyRouterMiddleware(useRelay)}>
    <Route path="/" component={AppContainer} queries={ViewerQueries}>
      <IndexRoute component={PhotoListContainer} queries={ViewerQueries} prepareParams={params => ({... params, status:'any' })} />
      <Route path=":status" component={PhotoListContainer} queries={ViewerQueries} />
    </Route>
  </Router>,
  document.getElementById('root')
);

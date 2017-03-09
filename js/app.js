import 'babel-polyfill';

import React from 'react';
import ReactDOM from 'react-dom';
import ReactRouterRelay from 'react-router-relay';
import Relay from 'react-relay';
import { IndexRoute, Route, Router, browserHistory } from 'react-router';
import applyRouterMiddleware from 'react-router/lib/applyRouterMiddleware';
import useRelay from 'react-router-relay';

import App from './components/App';
import AppHomeRoute from './routes/AppHomeRoute';
import PhotoList from './components/PhotoList';
import ViewerQueries from './queries/ViewerQueries';

ReactDOM.render(
  <Router
      createElement={ReactRouterRelay.createElement}
      environment={Relay.Store}
      history={browserHistory}
      render={applyRouterMiddleware(useRelay)}>
    <Route path="/" component={App} queries={ViewerQueries}>
      <IndexRoute component={PhotoList} queries={ViewerQueries} prepareParams={params => ({... params, status:'any' })} />
      <Route path=":status" component={PhotoList} queries={ViewerQueries} />
    </Route>
  </Router>,
  document.getElementById('root')
);

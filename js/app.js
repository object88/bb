import 'babel-polyfill';

import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import React from 'react';
import ReactDOM from 'react-dom';
import ReactRouterRelay from 'react-router-relay';
import Relay from 'react-relay';
import { IndexRoute, Route, Router, browserHistory } from 'react-router';
import applyRouterMiddleware from 'react-router/lib/applyRouterMiddleware';
import darkBaseTheme from 'material-ui/styles/baseThemes/darkBaseTheme';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import { black, grey100, gray300, gray500 } from 'material-ui/styles/colors';
import useRelay from 'react-router-relay';

import App from './components/App';
import AppHomeRoute from './routes/AppHomeRoute';
import PhotoList from './components/PhotoList';
import ViewerQueries from './queries/ViewerQueries';

// import schema from '../data/schema.json';
//
// Relay.injectNetworkLayer(
//   new RelayLocalSchema.NetworkLayer({ schema })
// );

// This replaces the textColor value on the palette
// and then update the keys for each component that depends on it.
// More on Colors: http://www.material-ui.com/#/customization/colors
const muiTheme = getMuiTheme(darkBaseTheme, {
  palette: {
    textColor: grey100,
    secondaryTextColor: gray300,
    alternateTextColor: gray500,
  },
  appBar: {
    color: black,
//    height: spacing.desktopKeylineIncrement,
//    padding: spacing.desktopGutter,
    height: 50,
    titleFontWeight: 900,
  },
});

ReactDOM.render(
  <MuiThemeProvider muiTheme={muiTheme}>
    <Router
        createElement={ReactRouterRelay.createElement}
        environment={Relay.Store}
        history={browserHistory}
        render={applyRouterMiddleware(useRelay)}>
      <Route path="/" component={App} queries={ViewerQueries}>
        <IndexRoute component={PhotoList} queries={ViewerQueries} prepareParams={params => ({... params, status:'any' })} />
        <Route path=":status" component={PhotoList} queries={ViewerQueries} />
      </Route>
    </Router>
  </MuiThemeProvider>,
  document.getElementById('root')
);

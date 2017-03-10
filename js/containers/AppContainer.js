// @flow
import Relay from 'react-relay';

import App from '../components/App';
import AppHeaderContainer from './AppHeaderContainer';

export default Relay.createContainer(App, {
  fragments: {
    viewer: () => Relay.QL`fragment on User { ${AppHeaderContainer.getFragment('viewer')} photos(first: 10) {edges {node {id}}}}`,
  },
});

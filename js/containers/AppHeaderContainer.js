// @flow
import Relay from 'react-relay';
import AppHeader from '../components/AppHeader';

export default Relay.createContainer(AppHeader, {
  fragments: {
    viewer: () => Relay.QL`fragment on User { id }`,
  },
});

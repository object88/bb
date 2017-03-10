// @flow
import Relay from 'react-relay';
import PhotoList from '../components/PhotoList';

export default Relay.createContainer(PhotoList, {
  initialVariables: {
    limit: 200,
  },

  // prepareVariables() {
  //   return {
  //     limit: 200
  //   };
  // },

  fragments: {
    viewer: () => Relay.QL`fragment on User { photos(first: 10) {edges {node { id }}}}`,
  },
});

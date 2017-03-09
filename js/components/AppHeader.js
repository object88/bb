// @flow
import React from 'react';
import Relay from 'react-relay';

type Props = {
  relay: Object,
  viewer: Object,
}

class AppHeader extends React.Component {
  props: Props;

  render() {
    return <div>BRIGHTER BLACKER</div>
  }
}

export default Relay.createContainer(AppHeader, {
  fragments: {
    viewer: () => Relay.QL`fragment on User { id }`
  }
});

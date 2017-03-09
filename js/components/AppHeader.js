// @flow
import React from 'react';
import Relay from 'react-relay';

class AppHeader extends React.Component {
  static propTypes = {
    relay: React.PropTypes.object.isRequired,
    viewer: React.PropTypes.object.isRequired
  };

  render() {
    return <div>BRIGHTER BLACKER</div>
  }
}

export default Relay.createContainer(AppHeader, {
  fragments: {
    viewer: () => Relay.QL`fragment on User { id }`
  }
});

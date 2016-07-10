import AppBar from 'material-ui/AppBar';
import React from 'react';
import Relay from 'react-relay';

class AppHeader extends React.Component {
  static propTypes = {
    relay: React.PropTypes.object.isRequired,
    viewer: React.PropTypes.object.isRequired
  };

  render() {
    return <AppBar title="BRIGHTER BLACKER"/>;
  }
}

export default Relay.createContainer(AppHeader, {
  fragments: {
    viewer: () => Relay.QL`fragment on User { totalCount }`
  }
});

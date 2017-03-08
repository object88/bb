import AppBar from 'material-ui/AppBar';
import IconButton from 'material-ui/IconButton';
import IconMenu from 'material-ui/IconMenu';
import MenuItem from 'material-ui/MenuItem';
import MoreVertIcon from 'material-ui/svg-icons/navigation/more-vert';
import NavigationClose from 'material-ui/svg-icons/navigation/close';
import React from 'react';
import Relay from 'react-relay';

class AppHeader extends React.Component {
  static propTypes = {
    relay: React.PropTypes.object.isRequired,
    viewer: React.PropTypes.object.isRequired
  };

//  iconClassNameRight="muidocs-icon-navigation-expand-more"
  render() {
    return <AppBar
      title="BRIGHTER BLACKER"
      iconElementRight={
      <IconMenu
        animated={false}
        iconButtonElement={
          <IconButton><MoreVertIcon /></IconButton>
        }
        targetOrigin={{horizontal: 'right', vertical: 'top'}}
        anchorOrigin={{horizontal: 'right', vertical: 'top'}}
        >
          <MenuItem primaryText="Help" />
        </IconMenu>
      }
    />;
  }
}

export default Relay.createContainer(AppHeader, {
  fragments: {
    viewer: () => Relay.QL`fragment on User { id }`
  }
});

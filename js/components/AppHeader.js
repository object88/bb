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

  onSignIn(googleUser) {
    console.log('Sign in returned');
    var profile = googleUser.getBasicProfile();
    console.log('ID: ' + profile.getId()); // Do not send to your backend! Use an ID token instead.
    console.log('Name: ' + profile.getName());
    console.log('Image URL: ' + profile.getImageUrl());
    console.log('Email: ' + profile.getEmail());
  }

  onTouchTap(event, menuItem, index) {
    gapi.load('client:auth2', () => {
      console.log('Finished client:auth2 load', this);
      gapi.auth2.init({}).then(() => {
        // Init completed!
        console.log('Init completed', this)
        gapi.auth2.getAuthInstance().signIn().then(this.onSignIn);
      });
    });
  }

//  iconClassNameRight="muidocs-icon-navigation-expand-more"
  render() {
    console.log('render', this);
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
          <MenuItem primaryText="Sign on" onTouchTap={this.onTouchTap.bind(this)} />
        </IconMenu>
      }
    />;
  }
}

export default Relay.createContainer(AppHeader, {
  fragments: {
    viewer: () => Relay.QL`fragment on User { totalCount }`
  }
});

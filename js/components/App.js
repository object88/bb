import React from 'react';
import Relay from 'react-relay';

import AppHeader from './AppHeader';

class App extends React.Component {
  static propTypes = {
    viewer: React.PropTypes.object.isRequired,
    children: React.PropTypes.node.isRequired,
    relay: React.PropTypes.object.isRequired,
  };

  render() {
    const { viewer, children } = this.props;

    return (
      <div data-framework="relay">
        <section className="photoapp">
          <AppHeader viewer={viewer}/>

          {children}

        </section>
        <footer className="info">
          <p>Words words words</p>
        </footer>
      </div>
    );
  }
}
//          <PhotoListFooter photos={this.props.viewer.photos} viewer={this.props.viewer}/>

export default Relay.createContainer(App, {
  fragments: {
    viewer: () => Relay.QL`fragment on User { ${AppHeader.getFragment('viewer')} photos(first: 10) {edges {node {id}}}}`,
  },
});

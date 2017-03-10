// @flow
import React from 'react';
import Relay from 'react-relay';

import AppHeader from './AppHeader';

type Props = {
  children: typeof React.PropTypes.node,
  relay: Object,
  viewer: Object,
};

class App extends React.Component {
  props: Props;

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

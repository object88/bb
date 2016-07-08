import React from 'react';
import Relay from 'react-relay';

class App extends React.Component {
  static propTypes = {
    viewer: React.PropTypes.object.isRequired,
    children: React.PropTypes.node.isRequired,
    relay: React.PropTypes.object.isRequired,
  };

  render() {
    return (
      <div data-framework="relay">
        <section className="photoapp">
          <header className="header">
            <h1>Brighter Blacker</h1>
          </header>

          {this.props.children}

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
    viewer: () => Relay.QL`fragment on User {photos(first: 10) {edges {node {id,name}}}}`,
  },
});

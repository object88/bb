// @flow
import React from 'react';
import Relay from 'react-relay';

import AppHeaderContainer from '../containers/AppHeaderContainer';

type Props = {
  children: typeof React.PropTypes.node,
  relay: Object,
  viewer: Object,
};

const displayName = "App";

export default class App extends React.Component {
  static displayName = displayName;

  props: Props;

  render() {
    const { viewer, children } = this.props;

    return (
      <div data-framework="relay">
        <section className="photoapp">
          <AppHeaderContainer viewer={viewer}/>

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

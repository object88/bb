// @flow
import React from 'react';
import Relay from 'react-relay';

type Props = {
  relay: Object,
  viewer: Object,
};

const displayName = "AppHeader";

export default class AppHeader extends React.Component {
  static displayName = displayName;

  props: Props;

  render() {
    return <div>BRIGHTER BLACKER</div>
  }
}

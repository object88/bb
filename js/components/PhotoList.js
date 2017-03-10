// @flow
import React from 'react'
import Relay from 'react-relay'

const displayName = "PhotoList";

type Props = {
  viewer: any,
};

export default class PhotoList extends React.Component {
  static displayName = displayName;

  props: Props;

  renderPhotos() {
    return this.props.viewer.photos.edges.map(edge =>
      <li key={edge.node.id}>{edge.node.id}</li>
    );
  }

  render() {
    return (
      <section className="main">
        <ul className="photo-list">
          {this.renderPhotos()}
        </ul>
      </section>
    );
  }
}

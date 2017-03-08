export default class LoginMutation extends Relay.Mutation {
  static fragments = {
    user: () => Relay.QL`
     fragment on User {
      id,
    }`,
  }

  getMutation() {
    return Relay.QL`mutation{Login}`;
  }

  getVariables() {
    return {
      id: this.props.user.id
    };
  }

  getConfigs() {
    return [{
      type: 'FIELDS_CHANGE',
      fieldIDs: {
        user: this.props.user.id,
      }
    }];
  }

  getOptimisticResponse() {
    return {
      id: this.props.user.id
    };
  }
  
  getFatQuery() {
    return Relay.QL`
    fragment on LoginPayload {
      user {
        userID,
      }
    }
    `;
  }
}

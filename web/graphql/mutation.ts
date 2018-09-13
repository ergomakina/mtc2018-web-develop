import gql from 'graphql-tag';

export const LIKE_SESSION_MUTATION = gql`
  mutation LikeSesstion($randomId: String!, $sessionId: ID!) {
    createLike(input: { clientMutationId: $randomId, sessionId: $sessionId }) {
      clientMutationId
      like {
        id
        sessionId
      }
    }
  }
`;

import { gql } from "graphql-tag";

export const MeDocument = gql`
  query Me {
    me {
      id
      uid
      createdAt
      updatedAt
    }
  }
`;
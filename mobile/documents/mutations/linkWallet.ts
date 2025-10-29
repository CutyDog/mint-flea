import { gql } from "graphql-tag";

export const LinkWalletDocument = gql`
  mutation LinkWallet($input: LinkWalletInput!) {
    linkWallet(input: $input) {
      id accountId address chainId isMain createdAt updatedAt
    }
  }
`;
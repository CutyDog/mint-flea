import { GraphQLClient, RequestOptions } from 'graphql-request';
import { gql } from 'graphql-tag';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
type GraphQLClientRequestHeaders = RequestOptions['requestHeaders'];
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
};

export type Account = {
  __typename?: 'Account';
  createdAt: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  mainWallet?: Maybe<Wallet>;
  uid: Scalars['String']['output'];
  updatedAt: Scalars['String']['output'];
  wallets: Array<Wallet>;
};

export type LinkWalletInput = {
  address: Scalars['String']['input'];
  chainId: Scalars['Int']['input'];
  isMain: Scalars['Boolean']['input'];
};

export type Mutation = {
  __typename?: 'Mutation';
  linkWallet: Wallet;
  setMainWallet: Wallet;
  unlinkWallet: Scalars['Boolean']['output'];
};


export type MutationLinkWalletArgs = {
  input: LinkWalletInput;
};


export type MutationSetMainWalletArgs = {
  input: SetMainWalletInput;
};


export type MutationUnlinkWalletArgs = {
  input: UnlinkWalletInput;
};

export type Query = {
  __typename?: 'Query';
  me?: Maybe<Account>;
};

export type SetMainWalletInput = {
  walletId: Scalars['ID']['input'];
};

export type UnlinkWalletInput = {
  walletId: Scalars['ID']['input'];
};

export type Wallet = {
  __typename?: 'Wallet';
  account: Account;
  accountId: Scalars['ID']['output'];
  address: Scalars['String']['output'];
  chainId: Scalars['Int']['output'];
  createdAt: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  isMain: Scalars['Boolean']['output'];
  updatedAt: Scalars['String']['output'];
};

export type LinkWalletMutationVariables = Exact<{
  input: LinkWalletInput;
}>;


export type LinkWalletMutation = { __typename?: 'Mutation', linkWallet: { __typename?: 'Wallet', id: string, accountId: string, address: string, chainId: number, isMain: boolean, createdAt: string, updatedAt: string } };

export type MeQueryVariables = Exact<{ [key: string]: never; }>;


export type MeQuery = { __typename?: 'Query', me?: { __typename?: 'Account', id: string, uid: string, createdAt: string, updatedAt: string } | null };


export const LinkWalletDocument = gql`
    mutation LinkWallet($input: LinkWalletInput!) {
  linkWallet(input: $input) {
    id
    accountId
    address
    chainId
    isMain
    createdAt
    updatedAt
  }
}
    `;
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

export type SdkFunctionWrapper = <T>(action: (requestHeaders?:Record<string, string>) => Promise<T>, operationName: string, operationType?: string, variables?: any) => Promise<T>;


const defaultWrapper: SdkFunctionWrapper = (action, _operationName, _operationType, _variables) => action();

export function getSdk(client: GraphQLClient, withWrapper: SdkFunctionWrapper = defaultWrapper) {
  return {
    LinkWallet(variables: LinkWalletMutationVariables, requestHeaders?: GraphQLClientRequestHeaders, signal?: RequestInit['signal']): Promise<LinkWalletMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<LinkWalletMutation>({ document: LinkWalletDocument, variables, requestHeaders: { ...requestHeaders, ...wrappedRequestHeaders }, signal }), 'LinkWallet', 'mutation', variables);
    },
    Me(variables?: MeQueryVariables, requestHeaders?: GraphQLClientRequestHeaders, signal?: RequestInit['signal']): Promise<MeQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<MeQuery>({ document: MeDocument, variables, requestHeaders: { ...requestHeaders, ...wrappedRequestHeaders }, signal }), 'Me', 'query', variables);
    }
  };
}
export type Sdk = ReturnType<typeof getSdk>;
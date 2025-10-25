import React, { useEffect, useState } from "react";
import {
  ApolloClient,
  InMemoryCache,
  HttpLink,
} from "@apollo/client";
import { setContext } from "@apollo/client/link/context";
import { ApolloProvider } from "@apollo/client/react";
import { loadErrorMessages, loadDevMessages } from "@apollo/client/dev";
import { useAuth } from "./AuthContext";
import { Platform } from "react-native";

interface ApolloWrapperProps {
  children: React.ReactNode;
}

// プラットフォームに応じてURIを設定
const getGraphQLUri = () => {
  if (process.env.EXPO_PUBLIC_ENV === 'production' || process.env.EXPO_PUBLIC_ENV === 'preview') {
    return `${process.env.EXPO_PUBLIC_SERVER_URL}/query`;
  }

  // 環境変数からサーバー情報を取得
  const serverIP = process.env.EXPO_PUBLIC_SERVER_IP || 'localhost';

  if (Platform.OS === 'ios' || Platform.OS === 'android') {
    // iOS/Android用: ホストマシンのIPアドレスを使用
    return `http://${serverIP}:8080/query`;
  } else {
    // Web用: localhostを使用
    return `http://localhost:8080/query`;
  }
};

const createApolloClient = (getIdToken: () => Promise<string | null>) => {
  // 認証ヘッダーを動的に設定するLink
  const authLink = setContext(async (_, { headers }) => {
    const token = await getIdToken();
    console.log('token', token);
    return {
      headers: {
        ...headers,
        ...(token ? { Authorization: `Bearer ${token}` } : {}),
      },
    };
  });

  // HTTPリンク
  const httpLink = new HttpLink({
    uri: getGraphQLUri(),
  });

  const apolloClient = new ApolloClient({
    link: authLink.concat(httpLink),
    cache: new InMemoryCache(),
  })

  return apolloClient;
}

export const ApolloWrapper = ({ children }: ApolloWrapperProps) => {
  const { getIdToken } = useAuth();
  const [client, setClient] = useState(createApolloClient(async () => null));

  useEffect(() => {
    setClient(createApolloClient(getIdToken));
  }, [getIdToken]);

  if (__DEV__) {
    loadDevMessages();
    loadErrorMessages();
  }

  return <ApolloProvider client={client}>{children}</ApolloProvider>;
}
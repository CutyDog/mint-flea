import "@walletconnect/react-native-compat";

import React from 'react';
import { SafeAreaProvider } from 'react-native-safe-area-context';
import { AppKitProvider, AppKit } from '@reown/appkit-react-native';
import { appKit } from '@/providers/AppKitConfig';
import { Stack } from "expo-router";
import { AuthProvider, ApolloWrapper } from "../hooks";
import { ThemeProvider } from "../theme";

export default function RootLayout() {
  return (
    <SafeAreaProvider>
      <AppKitProvider instance={appKit}>
        <ThemeProvider>
          <AuthProvider>
            <ApolloWrapper>
              <Stack />
            </ApolloWrapper>
          </AuthProvider>
        </ThemeProvider>

        <AppKit />
      </AppKitProvider>
    </SafeAreaProvider>
  )
}

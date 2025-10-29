// mobile/components/WalletHeaderBar.tsx
import React from 'react';
import { View } from 'react-native';
import {
  AppKitButton,
  // AccountButton,
  // ConnectButton,
  NetworkButton
} from '@reown/appkit-react-native';

export const WalletHeaderBar: React.FC = () => {
  return (
    <View style={{ padding: 12, gap: 8, flexDirection: 'row', alignItems: 'center' }}>
      {/* これ一つで「未接続→Connect」「接続後→Account」へ自動変化 */}
      <AppKitButton size="md" />

      {/* 必要に応じて個別ボタンも併用可 */}
      <NetworkButton />
      {/* 明示的に分けたいときは下の2つを切替レンダ */}
      {/* <ConnectButton label="Connect Wallet" loadingLabel="Opening…" /> */}
      {/* <AccountButton balance="hide" /> */}
    </View>
  );
};

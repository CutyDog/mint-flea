// mobile/components/LinkWalletButton.tsx
import React from 'react';
import { Button } from 'react-native';
import { useMutation } from '@apollo/client/react';
import { LinkWalletDocument } from '@/documents';
import { LinkWalletMutation } from '@/types/graphql';
import { useAccount, useAppKit } from '@reown/appkit-react-native';

export const LinkWalletButton: React.FC = () => {
  const { open } = useAppKit(); // モーダル開閉
  const { address, chainId, isConnected } = useAccount(); // 接続状態/アドレス
  const [linkWallet, { loading }] = useMutation<LinkWalletMutation>(LinkWalletDocument);

  const onPress = async () => {
    if (!isConnected) {
      await open(); // まず接続モーダルを開く
    }
    if (!address || !chainId) return;
    await linkWallet({
      variables: {
        input: {
          address: address.toLowerCase(),
          chainId: Number(chainId),
          isMain: true,
        }
      }
    });
  };

  return <Button title={loading ? 'Linking...' : 'Link this wallet'} onPress={onPress} />;
};

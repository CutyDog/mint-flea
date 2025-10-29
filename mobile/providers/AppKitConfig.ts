import "@walletconnect/react-native-compat";

import {
  createAppKit,
  bitcoin,
  solana,
  // type AppKitNetwork,
} from '@reown/appkit-react-native';
import { EthersAdapter } from '@reown/appkit-ethers-react-native';
import { SolanaAdapter, PhantomConnector } from '@reown/appkit-solana-react-native';
import { BitcoinAdapter } from '@reown/appkit-bitcoin-react-native';
import { storage } from './StorageUtil';

// You can use 'viem/chains' or define your own chains using `AppKitNetwork` type. Check Options/networks for more detailed info
import { mainnet, polygon } from 'viem/chains';

const projectId = process.env.WALLETCONNECT_PROJECT_ID!; // Obtain from https://dashboard.reown.com/

const ethersAdapter = new EthersAdapter();
const solanaAdapter = new SolanaAdapter();
const bitcoinAdapter = new BitcoinAdapter();

export const appKit = createAppKit({
  projectId,
  networks: [mainnet, polygon, solana, bitcoin],
  defaultNetwork: mainnet, // Optional: set a default network
  adapters: [ethersAdapter, solanaAdapter, bitcoinAdapter],
  storage,
  extraConnectors: [
    new PhantomConnector({ cluster: 'mainnet-beta' }) // Or 'devnet', 'testnet'
  ],

  // Other AppKit options (e.g., metadata for your dApp)
  metadata: {
    name: 'mint flea',
    description: 'NFT marketplace',
    url: 'https://mintflea.app',
    icons: ['https://mintflea.app/icon.png'],
    redirect: {
      native: "mintflea://",
      universal: "https://mintflea.app",
    },
  }
});
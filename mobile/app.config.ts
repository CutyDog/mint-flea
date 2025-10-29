import { ExpoConfig, ConfigContext } from '@expo/config';
import * as dotenv from 'dotenv';
import path from 'path';

// ★ .env.local を明示的に読む（なければ .env）
dotenv.config({ path: path.resolve(process.cwd(), '.env.local') }) || dotenv.config();

export default ({ config }: ConfigContext): ExpoConfig => ({
  ...config,
  name: "mint flea",
  slug: "mint-flea",
  version: "1.0.0",
  orientation: "portrait",
  icon: "./assets/images/icon.png",
  scheme: ["com.googleusercontent.apps.1053499135011-0du0opg6p3a7tsam4kk6ip0mugbge28r", "mintflea"],
  userInterfaceStyle: "automatic",
  newArchEnabled: true,
  ios: {
    supportsTablet: true,
    bundleIdentifier: "com.cutydog.mint-flea",
    associatedDomains: ["applinks:mintflea.com"],
    googleServicesFile: process.env.GOOGLE_SERVICES_INFO_PLIST ?? './GoogleService-Info.plist',
    infoPlist: {
      ITSAppUsesNonExemptEncryption: false,
      NSPhotoLibraryUsageDescription: "プロフィール写真を選択するためにカメラロールへのアクセスが必要です。",
      NSCameraUsageDescription: "プロフィール写真の撮影とQRコードのスキャンにカメラへのアクセスが必要です。",
      NSUserTrackingUsageDescription: "関連性の高い広告を表示するためにIDをトラッキングいたします。"
    }
  },
  android: {
    adaptiveIcon: {
      foregroundImage: "./assets/images/adaptive-icon.png",
      backgroundColor: "#ffffff"
    },
    edgeToEdgeEnabled: true,
    package: "com.cutydog.mint_flea",
    googleServicesFile: process.env.GOOGLE_SERVICES_JSON ?? './google-services.json',
    permissions: [
      "android.permission.CAMERA",
      "android.permission.READ_EXTERNAL_STORAGE",
      "android.permission.WRITE_EXTERNAL_STORAGE"
    ]
  },
  web: {
    bundler: "metro",
    output: "static",
    favicon: "./assets/images/favicon.png"
  },
  plugins: [
    "@react-native-firebase/app",
    "@react-native-firebase/auth",
    "@react-native-firebase/crashlytics",
    [
      "@react-native-google-signin/google-signin",
      {
        iosUrlScheme: process.env.GOOGLE_SIGNIN_IOS_URL_SCHEME
      }
    ],
    // [
    //   "react-native-google-mobile-ads",
    //   {
    //     iosAppId: "ca-app-pub-5614922645470689~3726530890",
    //     androidAppId: "ca-app-pub-5614922645470689~3693637653",
    //     userTrackingPermission: "関連性の高い広告を表示するためにIDをトラッキングいたします。"
    //   }
    // ],
    [
      "expo-build-properties",
      {
        ios: {
          useFrameworks: "static",
          buildReactNativeFromSource: true
        }
      }
    ],
    "expo-router",
    [
      "expo-splash-screen",
      {
        image: "./assets/images/splash-icon.png",
        imageWidth: 200,
        resizeMode: "contain",
        backgroundColor: "#ffffff",
        dark: {
          backgroundColor: "#000000"
        }
      }
    ],
    "expo-image-picker",
    "expo-camera",
    "expo-tracking-transparency"
  ],
  experiments: {
    typedRoutes: true
  },
  extra: {
    router: {},
    eas: {
      projectId: "91e7199e-3fe1-4e6d-a9b5-cd2789800fcb"
    }
  }
})
import React, { createContext, useContext, useState, useEffect, useMemo } from 'react';
import { useRouter, useSegments } from 'expo-router';
import { GoogleSignin } from "@react-native-google-signin/google-signin";
import {
  getAuth,
  onAuthStateChanged,
  signInWithCredential,
  signOut,
  GoogleAuthProvider,
  signInWithEmailAndPassword,
  createUserWithEmailAndPassword
} from "@react-native-firebase/auth";
import type { FirebaseAuthTypes } from "@react-native-firebase/auth";

// // Google Sign-In設定
// GoogleSignin.configure({
//   webClientId: process.env.GOOGLE_WEB_CLIENT_ID,
//   iosClientId: process.env.GOOGLE_IOS_CLIENT_ID,
// });

interface AuthContextType {
  user: FirebaseAuthTypes.User | null;
  loading: boolean;
  error: string;
  signInWithGoogle: () => Promise<string | undefined>;
  signInWithEmail: (email: string, password: string) => Promise<string | undefined>;
  signUpWithEmail: (email: string, password: string) => Promise<string | undefined>;
  signOut: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  // Firebase Auth状態
  const [user, setUser] = useState<FirebaseAuthTypes.User | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string>("");

  // ルーティング関連
  const router = useRouter();
  const segments = useSegments();

  // Firebase Authインスタンスを一度だけ作成
  const auth = useMemo(() => getAuth(), []);

  // 共通のエラーハンドリング関数
  const handleAuthError = (error: any) => {
    setError(error.message);
    return undefined;
  };

  // 成功時のエラークリア
  const clearError = () => {
    setError("");
  };

  // 認証状態の監視
  useEffect(() => {
    let isMounted = true;

    const unsubscribe = onAuthStateChanged(auth, async (u: FirebaseAuthTypes.User | null) => {
      if (!isMounted) return;

      setUser(u);
      setLoading(false);
    });

    return () => {
      isMounted = false;
      unsubscribe();
    };
  }, [auth]);

  // 認証状態に基づく自動遷移
  useEffect(() => {
    if (loading) return; // ローディング中は何もしない

    const inAuthGroup = segments[0] === '(auth)';

    if (!user && !inAuthGroup) {
      // ログインしていない場合、ログインページに遷移
      router.replace('/(auth)/login');
    } else if (user && inAuthGroup) {
      // ログイン済みの場合、ホームページに遷移
      router.replace('/');
    }
  }, [user, loading, segments, router]);

  const signInWithGoogle = async () => {
    try {
      await GoogleSignin.hasPlayServices({ showPlayServicesUpdateDialog: true });
      const signInResult = await GoogleSignin.signIn();
      const { data } = signInResult;
      const googleCredential = GoogleAuthProvider.credential(data?.idToken ?? "");
      const userCredential = await signInWithCredential(auth, googleCredential);
      const firebaseIdToken = await userCredential.user.getIdToken(true);
      clearError();
      return firebaseIdToken;
    } catch (e: any) {
      return handleAuthError(e);
    }
  };

  const signInWithEmail = async (email: string, password: string) => {
    try {
      const userCredential = await signInWithEmailAndPassword(auth, email, password);
      const firebaseIdToken = await userCredential.user.getIdToken(true);
      clearError();
      return firebaseIdToken;
    } catch (e: any) {
      return handleAuthError(e);
    }
  };

  const signUpWithEmail = async (email: string, password: string) => {
    try {
      const userCredential = await createUserWithEmailAndPassword(auth, email, password);
      const firebaseIdToken = await userCredential.user.getIdToken(true);
      clearError();
      return firebaseIdToken;
    } catch (e: any) {
      return handleAuthError(e);
    }
  };

  const handleSignOut = async () => {
    try {
      await signOut(auth);
    } catch (e: any) {
      handleAuthError(e);
    }
  };

  const value = {
    user,
    loading,
    error,
    signInWithGoogle,
    signInWithEmail,
    signUpWithEmail,
    signOut: handleSignOut,
  };

  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  );
};
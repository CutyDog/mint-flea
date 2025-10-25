import { Stack } from "expo-router";
import { AuthProvider, ApolloWrapper } from "../hooks";
import { ThemeProvider } from "../theme";

export default function RootLayout() {
  return (
    <ThemeProvider>
      <AuthProvider>
        <ApolloWrapper>
          <Stack />
        </ApolloWrapper>
      </AuthProvider>
    </ThemeProvider>
  )
}

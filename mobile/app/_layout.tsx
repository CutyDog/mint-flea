import { Stack } from "expo-router";
import { AuthProvider } from "../hooks";
import { ThemeProvider } from "../theme";

export default function RootLayout() {
  return (
    <ThemeProvider>
      <AuthProvider>
        <Stack />
      </AuthProvider>
    </ThemeProvider>
  )
}

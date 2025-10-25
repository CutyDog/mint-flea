import { Text, View } from "react-native";
import { useQuery } from "@apollo/client/react";
import { MeDocument } from "../documents/queries/me";
import { MeQuery } from "../types/graphql";

export default function Index() {
  const { data, loading, error } = useQuery<MeQuery>(MeDocument);
  if (loading) return <Text>Loading...</Text>;
  if (error) return <Text>Error: {error.message}</Text>;

  return (
    <View
      style={{
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Text>{data?.me?.uid}</Text>
    </View>
  );
}

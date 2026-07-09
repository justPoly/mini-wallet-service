import { useQuery } from "@tanstack/react-query";
import { getAccount } from "../services/accountService";

export function useAccount(id: string) {
  return useQuery({
    queryKey: ["account", id],
    queryFn: () => getAccount(id),
    enabled: !!id,
  });
}
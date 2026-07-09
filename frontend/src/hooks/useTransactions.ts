import { useQuery } from "@tanstack/react-query";
import { getTransactions } from "../services/accountService";

export function useTransactions(accountId: string) {
  return useQuery({
    queryKey: ["transactions", accountId],
    queryFn: () => getTransactions(accountId),
    enabled: !!accountId,
  });
}
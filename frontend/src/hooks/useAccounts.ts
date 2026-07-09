import { useQuery } from "@tanstack/react-query";
import { getAccounts } from "../services/accountService";

export function useAccounts() {
  return useQuery({
    queryKey: ["accounts"],
    queryFn: getAccounts,
  });
}
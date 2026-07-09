import { useMutation, useQueryClient } from "@tanstack/react-query";
import { transferMoney } from "../services/accountService";

export function useTransfer() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: transferMoney,

    onSuccess: () => {
      // Refresh account balances
      queryClient.invalidateQueries({
        queryKey: ["accounts"],
      });

      // Refresh account details
      queryClient.invalidateQueries({
        queryKey: ["account"],
      });

      // Refresh transaction history
      queryClient.invalidateQueries({
        queryKey: ["transactions"],
      });
    },
  });
}
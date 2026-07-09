import { useMutation, useQueryClient } from "@tanstack/react-query";
import { transferMoney } from "../services/accountService";

export function useTransfer() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: transferMoney,

    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["accounts"],
      });

      queryClient.invalidateQueries({
        queryKey: ["account"],
      });

      queryClient.invalidateQueries({
        queryKey: ["transactions"],
      });
    },
  });
}
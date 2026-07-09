import client from "../api/client";
import type { Transaction } from "../types/transaction";

export const getTransactions = async (
  id: string
): Promise<Transaction[]> => {
  const response = await client.get<Transaction[]>(
    `/accounts/${id}/transactions`
  );

  return response.data;
};
import client from "../api/client";
import type { Account } from "../types/account";
import type { Transaction } from "../types/transaction";

export const getAccounts = async (): Promise<Account[]> => {
  const response = await client.get<Account[]>("/accounts");
  return response.data;
};

export const getAccount = async (id: string): Promise<Account> => {
  const response = await client.get<Account>(`/accounts/${id}`);
  return response.data;
};

export const getTransactions = async (
  accountId: string
): Promise<Transaction[]> => {
  const response = await client.get<Transaction[]>(
    `/accounts/${accountId}/transactions`
  );

  return response.data;
};
import client from "../api/client";
import type { Account } from "../types/account";

export const getAccounts = async (): Promise<Account[]> => {
  const response = await client.get<Account[]>("/accounts");
  return response.data;
};

export const getAccount = async (id: string): Promise<Account> => {
  const response = await client.get<Account>(`/accounts/${id}`);
  return response.data;
};
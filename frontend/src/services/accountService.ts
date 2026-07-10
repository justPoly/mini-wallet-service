import client from "../api/client";
import type { Account } from "../types/account";
import type { Transaction } from "../types/transaction";
import type { TransferRequest } from "../types/transfer";

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

export async function transferMoney(data: TransferRequest) {
  const response = await client.post(
    "/transfers",
    data,
    {
      headers: {
        "Idempotency-Key": Date.now().toString(),
      },
    }
  );

  return response.data;
}
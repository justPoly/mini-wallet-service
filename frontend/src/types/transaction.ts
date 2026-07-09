export interface Transaction {
  id: string;
  accountId: string;
  type: string;
  amount: number;
  description: string;
  createdAt: string;
}
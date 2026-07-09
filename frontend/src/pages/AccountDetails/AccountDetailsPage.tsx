import { useParams } from "react-router-dom";

import Loading from "../../components/Loading/Loading";
import Error from "../../components/Error/Error";

import { useAccount } from "../../hooks/useAccount";
import { useTransactions } from "../../hooks/useTransactions";

export default function AccountDetailsPage() {
  const { id } = useParams();

  const {
    data: account,
    isLoading: accountLoading,
    isError: accountError,
  } = useAccount(id ?? "");

  const {
    data: transactions,
    isLoading: transactionLoading,
    isError: transactionError,
  } = useTransactions(id ?? "");

  if (accountLoading || transactionLoading) {
    return <Loading />;
  }

  if (accountError || transactionError) {
    return <Error message="Failed to load account." />;
  }

  if (!account) {
    return <p>Account not found.</p>;
  }

  return (
    <div>
      <h2>{account.name}</h2>

      <p>
        <strong>Currency:</strong> {account.currency}
      </p>

      <p>
        <strong>Balance:</strong>{" "}
        {account.currency} {account.balance.toFixed(2)}
      </p>

      <hr />

      <h3>Transaction History</h3>

      {!transactions || transactions.length === 0 ? (
        <p>No transactions yet.</p>
      ) : (
        <table
          style={{
            width: "100%",
            borderCollapse: "collapse",
          }}
        >
          <thead>
            <tr>
              <th align="left">Type</th>
              <th align="left">Amount</th>
              <th align="left">Description</th>
              <th align="left">Date</th>
            </tr>
          </thead>

          <tbody>
            {transactions.map((transaction) => (
              <tr key={transaction.id}>
                <td>{transaction.type}</td>

                <td>{transaction.amount}</td>

                <td>{transaction.description}</td>

                <td>
                  {new Date(
                    transaction.createdAt
                  ).toLocaleString()}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}
import { useEffect, useState } from "react";

import { useAccounts } from "../../hooks/useAccounts";
import { useTransfer } from "../../hooks/useTransfer";
import { exchangeRates } from "../../utils/exchangeRates";

export default function TransferPage() {
  const { data: accounts } = useAccounts();

  const transfer = useTransfer();

  const [fromAccountId, setFromAccountId] = useState("");
  const [toAccountId, setToAccountId] = useState("");
  const [amount, setAmount] = useState("");

  useEffect(() => {
    if (transfer.isSuccess) {
        setFromAccountId("");
        setToAccountId("");
        setAmount("");
    }
  }, [transfer.isSuccess]);

  const fromAccount = accounts?.find(
    (account) => account.id === fromAccountId
  );

  const toAccount = accounts?.find(
    (account) => account.id === toAccountId
  );

  const convertedAmount =
    fromAccount &&
    toAccount &&
    amount
      ? Number(amount) *
        exchangeRates[fromAccount.currency][toAccount.currency]
      : null;

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    if (!fromAccountId || !toAccountId) {
      alert("Please select both accounts.");
      return;
    }

    if (fromAccountId === toAccountId) {
      alert("Source and destination accounts cannot be the same.");
      return;
    }

    if (Number(amount) <= 0) {
      alert("Amount must be greater than zero.");
      return;
    }

    transfer.mutate({
      fromAccountId,
      toAccountId,
      amount: Number(amount),
    });
  };

  return (
    <>
      <h2>Transfer Money</h2>

      <form onSubmit={handleSubmit}>
        <div>
          <label>From Account</label>

          <br />

          <select
            value={fromAccountId}
            onChange={(e) => setFromAccountId(e.target.value)}
          >
            <option value="">Select account</option>

            {accounts?.map((account) => (
              <option
                key={account.id}
                value={account.id}
              >
                {account.name} - {account.currency} ({account.balance.toFixed(2)})
              </option>
            ))}
          </select>
        </div>

        <br />

        <div>
          <label>To Account</label>

          <br />

          <select
            value={toAccountId}
            onChange={(e) => setToAccountId(e.target.value)}
          >
            <option value="">Select account</option>

            {accounts?.map((account) => (
              <option
                key={account.id}
                value={account.id}
              >
                {account.name} ({account.currency})
              </option>
            ))}
          </select>
        </div>

        <br />

        <div>
          <label>Amount</label>

          <br />

          <input
            type="number"
            min="0"
            step="0.01"
            value={amount}
            onChange={(e) => setAmount(e.target.value)}
          />
        </div>

        {fromAccount &&
          toAccount &&
          amount &&
          fromAccount.currency !== toAccount.currency && (
            <>
              <br />

              <div
                style={{
                  padding: "12px",
                  border: "1px solid #ddd",
                  borderRadius: "8px",
                  background: "#f5f8ff",
                }}
              >
                <strong>Conversion Preview</strong>

                <p>
                  {Number(amount).toFixed(2)} {fromAccount.currency} ≈{" "}
                  {convertedAmount?.toFixed(2)} {toAccount.currency}
                </p>
              </div>
            </>
          )}

        <br />

        <button
          type="submit"
          disabled={
            transfer.isPending ||
            !fromAccountId ||
            !toAccountId ||
            fromAccountId === toAccountId ||
            Number(amount) <= 0
            }
        >
          {transfer.isPending ? "Transferring..." : "Transfer"}
        </button>

        {transfer.isSuccess && (
        <p style={{ color: "green" }}>
            ✅ Transfer completed successfully.
        </p>
        )}

        {transfer.isError && (
        <p style={{ color: "red" }}>
            ❌ {
            transfer.error instanceof Error
                ? transfer.error.message
                : "Transfer failed."
            }
        </p>
        )}
      </form>
    </>
  );
}
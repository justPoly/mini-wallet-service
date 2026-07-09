import { Link } from "react-router-dom";
import type { Account } from "../../types/account";

interface Props {
  account: Account;
}

export default function AccountCard({ account }: Props) {
  console.log(account);
  return (
    <Link
      to={`/accounts/${account.id}`}
      style={{
        display: "block",
        border: "1px solid #ddd",
        borderRadius: "8px",
        padding: "16px",
        marginBottom: "16px",
        textDecoration: "none",
        color: "inherit",
      }}
    >
        <h3>{account.name ?? "Unknown"}</h3>

        <p>Currency: {account.currency ?? "-"}</p>

        <h2>
        {account.currency ?? "-"} {(account.balance ?? 0).toFixed(2)}
        </h2>
    </Link>
  );
}
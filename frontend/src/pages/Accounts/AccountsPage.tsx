import AccountCard from "../../components/AccountCard/AccountCard";
import Error from "../../components/Error/Error";
import Loading from "../../components/Loading/Loading";
import { useAccounts } from "../../hooks/useAccounts";

export default function AccountsPage() {
  const { data, isLoading, isError } = useAccounts();

  if (isLoading) {
    return <Loading />;
  }

  if (isError) {
    return <Error message="Failed to load accounts." />;
  }

  if (!data || data.length === 0) {
    return <p>No accounts found.</p>;
  }

  return (
    <>
      <h2>Accounts</h2>

      {data.map((account) => (
        <AccountCard
          key={account.id}
          account={account}
        />
      ))}
    </>
  );
}
import { Link, Outlet } from "react-router-dom";

export default function MainLayout() {
  return (
    <div style={{ maxWidth: "1000px", margin: "0 auto", padding: "20px" }}>
      <h1>Mini Wallet</h1>

      <nav style={{ marginBottom: "30px" }}>
        <Link to="/">Accounts</Link>{" "}
        |{" "}
        <Link to="/transfer">Transfer</Link>
      </nav>

      <Outlet />
    </div>
  );
}
import { Routes, Route } from "react-router-dom";

import MainLayout from "./layouts/MainLayout";

import AccountsPage from "./pages/Accounts/AccountsPage";
import AccountDetailsPage from "./pages/AccountDetails/AccountDetailsPage";
import TransferPage from "./pages/Transfer/TransferPage";

function App() {
  return (
    <Routes>
      <Route element={<MainLayout />}>
        <Route path="/" element={<AccountsPage />} />

        <Route
          path="/accounts/:id"
          element={<AccountDetailsPage />}
        />

        <Route
          path="/transfer"
          element={<TransferPage />}
        />
      </Route>
    </Routes>
  );
}

export default App;
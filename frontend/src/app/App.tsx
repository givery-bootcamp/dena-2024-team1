import { AppRoute } from "./AppRoute";
import "./App.scss";

import { Header } from "~/shared/components/Header";

function App() {
  return (
    <div>
      <Header />
      <main>
        <AppRoute />
      </main>
    </div>
  );
}

export default App;

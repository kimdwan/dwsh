import { BrowserRouter as Routers, Routes, Route } from "react-router-dom"
import { Main, SignUpOne } from "./pkgs"

function App() {
  return (
    <div className="App">
      <Routers>

        <Routes>
          <Route path = "/" element = {<Main />} />
          <Route path = "/signup/" element = {<SignUpOne />} />
        </Routes>

      </Routers>
    </div>
  );
}

export default App;

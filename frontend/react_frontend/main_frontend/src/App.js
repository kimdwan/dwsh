import { BrowserRouter as Routers, Routes, Route } from "react-router-dom"
import { Main, SignUpDs, SignUpGuest, SignUpOne, SignUpTwo } from "./pkgs"

function App() {
  return (
    <div className="App">
      <Routers>

        <Routes>
          <Route path = "/" element = {<Main />} />
          <Route path = "/signup/" element = {<SignUpOne />} />
          <Route path = "/signup/term/ds/&/*" element = {<SignUpDs />} />
          <Route path = "/signup/term/guest/&/*" element = {<SignUpGuest />} />
          <Route path = "/signup/term/*" element = {<SignUpTwo />} />
        </Routes>

      </Routers>
    </div>
  );
}

export default App;

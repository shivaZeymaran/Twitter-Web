import "./App.css";
import Signup from "./Components/Signup";
import Login from "./Components/Login";

import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import { Button } from "antd";
import { useState } from "react";

function App() {
  const [clicked, setClicked] = useState(false);

  const handleClick = () => {
    setClicked(true);
  };

  return (
    <Router>
      {/* {!clicked && (
        <> */}
          <div className="app-container">
            <Link to="/signup">
              <Button type="primary" onClick={handleClick}>
                Sign up
              </Button>
            </Link>
            <Link to="/login">
              <Button type="default" onClick={handleClick}>
                Log in
              </Button>
            </Link>
            <Switch>
              <Route path="/login" component={Login} />
              <Route path="/signup" component={Signup} />
            </Switch>
          </div>
        {/* </>
      )} */}
    </Router>
  );
}

export default App;

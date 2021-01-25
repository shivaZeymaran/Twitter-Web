import "./App.css";
import Signup from "./Signup";
import Login from "./Login";

import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import { Button } from "antd";

function App() {
  return (
    <Router>
      <div>
        <Link to="/signup">
          <Button type="primary">Sign up</Button>
        </Link>
        <Link to="/login">
          <Button type="default">Log in</Button>
        </Link>
        <Switch>
          <Route path="/login" component={Login} />
          <Route path="/signup" component={Signup} />
        </Switch>
      </div>
    </Router>
  );
}

export default App;

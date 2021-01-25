import "./App.css";
import Signup from "./Signup";
import Login from "./Login";

import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import { Button } from "antd";

function App() {
  return (
    <Router>
      <div>
        <Link to="/sign-up">
          <Button type="primary">Sign up</Button>
        </Link>
        <Link to="/log-in">
          <Button type="default">Log in</Button>
        </Link>
        <Switch>
          <Route path="/log-in" component={Login} />
          <Route path="/sign-up" component={Signup} />
        </Switch>
      </div>
    </Router>
  );
}

export default App;

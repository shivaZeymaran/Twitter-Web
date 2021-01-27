import "../App.css";
import Signup from "./Signup";
import Login from "./Login";

import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import { Button } from "antd";
import * as React from "react";

export default class Authentication extends React.Component {
  render() {
    return (
      <div className="app-container">
        <Router>
          {!this.props.isAuth && (
            <>
              <Route exact path="/login" component={Login} />
              <Route exact path="/signup" component={Signup} />

              <Link to="/signup">
                <Button type="primary">Sign up</Button>
              </Link>
              <Link to="/login">
                <Button type="default">Log in</Button>
              </Link>
            </>
          )}
        </Router>
      </div>
    );
  }
}

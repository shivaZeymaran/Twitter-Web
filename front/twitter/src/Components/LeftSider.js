import "../App.css";
import Profile from "./Profile";

import { Menu, Layout } from "antd";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";

const { Sider } = Layout;

function LeftSider() {
  return (
    <div className="Menu">
      <Router>
        <Route path="/profile" component={Profile} />
        {/* <Route path="/explore" component={Explore} /> */}

        <Sider
          style={{
            overflow: "auto",
            height: "100vh",
            position: "fixed",
            left: 0,
            backgroundColor: "white",
          }}
        >
          {/* <div className="logo" /> */}
          <Menu style={{ width: 300 }}>
            <Link to="/home">
              <Menu.Item key="1">Home</Menu.Item>
            </Link>
            <Menu.Item key="2">Explore</Menu.Item>
            <Menu.Item key="3">Notifications</Menu.Item>
            <Link to="/profile">
              <Menu.Item key="4">Profile</Menu.Item>
            </Link>
            <Menu.Item key="5">Log out</Menu.Item>
          </Menu>
        </Sider>
      </Router>
    </div>
  );
}

export default LeftSider;

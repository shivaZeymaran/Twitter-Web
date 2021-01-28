import "../App.css";
import Home from "./Home";
import Explore from "./Explore";
import Notifications from "./Notifications";
import Profile from "./Profile";
import Logout from "./Logout";

import { Menu, Layout } from "antd";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";

const { Sider } = Layout;

function LeftSider() {
  return (
    <div className="Menu">
      <Router>
        <Route exact path="/home" component={Home} />
        <Route exact path="/explore" component={Explore} />
        <Route exact path="/notification" component={Notifications} />
        <Route exact path="/profile" component={Profile} />
        <Route exact path="/logout" component={Logout} />

        <Sider
          style={{
            overflow: "auto",
            height: "100vh",
            position: "fixed",
            left: 0,
            backgroundColor: "white",
            width: "30%",
          }}
        >
          {/* <div className="logo" /> */}
          <Menu style={{ width: 300, padding: 100 }}>
            <div className="menu_item">
              <Link to="/home">
                <Menu.Item key="1">Home</Menu.Item>
              </Link>
            </div>
            <div className="menu_item">
              <Link to="/explore">
                <Menu.Item key="2">Explore</Menu.Item>
              </Link>
            </div>
            <div className="menu_item">
              <Link to="/notifications">
                <Menu.Item key="3">Notifications</Menu.Item>
              </Link>
            </div>
            <div className="menu_item">
              <Link to="/profile">
                <Menu.Item key="4">Profile</Menu.Item>
              </Link>
            </div>
            <div className="menu_item">
              <Link to="/logout">
                <Menu.Item key="5">Log out</Menu.Item>
              </Link>
            </div>
          </Menu>
        </Sider>
      </Router>
    </div>
  );
}

export default LeftSider;

import "./App.css";
import { Menu, Layout } from "antd";

const { Sider } = Layout;

function LeftSider() {
  return (
    <div className="Menu">
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
          <Menu.Item key="1">Home</Menu.Item>
          <Menu.Item key="2">Explore</Menu.Item>
          <Menu.Item key="3">Notifications</Menu.Item>
          <Menu.Item key="4">Profile</Menu.Item>
          <Menu.Item key="5">Log out</Menu.Item>
        </Menu>
      </Sider>
    </div>
  );
}

export default LeftSider;

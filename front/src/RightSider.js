import "./App.css";
import { Input, Layout, List } from "antd";

const { Sider } = Layout;
const { Search } = Input;

const data = [
    'Azin liked your tweet.',
    'Ali Followed you.',
    'Azin liked your tweet.',
    'sara retweeted your tweet.',
    'Azin liked your tweet.',
  ];

function RightSider() {
  return (
    <div className="righ-sider">
      <Sider
        style={{
          overflow: "auto",
          height: "100vh",
          position: "fixed",
          right: 0,
          backgroundColor:"white",
        //   width: "",
        }}
      >
        <Search placeholder="input search text" style={{ width: 200 }} />
        <List
          size="large"
          header={<div>Notifications</div>}
          bordered
          dataSource={data}
          renderItem={(item) => <List.Item>{item}</List.Item>}
        />
      </Sider>
    </div>
  );
}

export default RightSider;

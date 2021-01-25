import "./App.css";
import { Layout, Card } from "antd";

import LeftSider from "./LeftSider";
import NewTweet from "./NewTweet";
import RightSider from "./RightSider";

const { Header, Content, Footer } = Layout;

function Home() {
  return (
    <div>
      <Layout>
        <LeftSider />

        <RightSider />
        <Layout className="site-layout" style={{ marginLeft: 200 }}>
          <Header className="site-layout-background" style={{ padding: 0 }} />
          <Content style={{ margin: "24px 16px 0", overflow: "initial" }}>
            <div
              className="site-layout-background"
              style={{ padding: 24, textAlign: "center" }}
            >
              <NewTweet />

              <div className="site-card-border-less-wrapper">
                <Card
                  title="Card title"
                  bordered={false}
                  style={{ width: 900 }}
                >
                  <p>Card content</p>
                  <p>Card content</p>
                  <p>Card content</p>
                </Card>
                <Card
                  title="Card title"
                  bordered={false}
                  style={{ width: 900 }}
                >
                  <p>Card content</p>
                  <p>Card content</p>
                  <p>Card content</p>
                </Card>
                <Card
                  title="Card title"
                  bordered={false}
                  style={{ width: 900 }}
                >
                  <p>Card content</p>
                  <p>Card content</p>
                  <p>Card content</p>
                </Card>
              </div>
            </div>
          </Content>

          <Footer style={{ textAlign: "center" }}>
            Sweeter ©2021 Created by Vahide & Shiva
          </Footer>
        </Layout>
      </Layout>
    </div>
  );
}

export default Home;

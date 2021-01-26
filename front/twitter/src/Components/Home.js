import "../App.css";
import { Layout, Card } from "antd";
import * as React from "react";

import LeftSider from "./LeftSider";
import NewTweet from "./NewTweet";
import RightSider from "./RightSider";
import { Route, BrowserRouter as Router } from "react-router-dom";

const { Header, Content, Footer } = Layout;

export default class Home extends React.Component {
  render() {
    return (
      <div>
        <Router>
          <Layout>
            <LeftSider />

            <RightSider />
            <Layout className="site-layout" style={{ marginLeft: 200 }}>
              <Content style={{ margin: "24px 16px 0", overflow: "initial" }}>
                <div
                  className="site-layout-background"
                  style={{ padding: 24, textAlign: "center" }}
                >
                  <NewTweet token={this.props.token} />

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
        </Router>
      </div>
    );
  }
}

import "../App.css";
import * as React from "react";
import { Layout, Card } from "antd";
import { Route, BrowserRouter as Router } from "react-router-dom";

import LeftSider from "./LeftSider";
import NewTweet from "./NewTweet";
import RightSider from "./RightSider";
import Tweet from "./Tweet";
import Profile from "./Profile";


const { Header, Content, Footer } = Layout;

export default class Home extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      tweets: [],
      notifs: [],
    };
  }

  render() {
    return (
      <div>
        <Router>
          <Route path="/profile" component={Profile} />
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
                    {this.state.tweets.map((tweet) => (
                      <Tweet />
                    ))}
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

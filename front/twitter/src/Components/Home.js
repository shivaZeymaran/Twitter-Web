import "../App.css";
import { AuthContext } from "../App";

import * as React from "react";
import { Layout } from "antd";
import { BrowserRouter as Router, Link, Route } from "react-router-dom";

import LeftSider from "./LeftSider";
import NewTweet from "./NewTweet";
import RightSider from "./RightSider";
import Tweet from "./Tweet";
import Profile from "./Profile";
import Login from "./Login";

const { Content } = Layout;

const Home = () => {
  const [tweets, setTweets] = React.useState([]);
  const [notifs, setNotifs] = React.useState([]);

  const [timer, setTimer] = React.useState(null);
  const [isMounted, setIsMounted] = React.useState(false);

  const { state: authState } = React.useContext(AuthContext);

  React.useEffect(() => {
    if (!isMounted) {
      getTweets();
      setIsMounted(true);
    }
  });

  const getTweets = async () => {
    await fetch("http://localhost:8090/timeline", {
      method: "post",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        token: authState.token,
      }),
    })
      .then((res) => {
        console.log(res.status);
        if (res.status === 200) return res.json();
      })
      .then((resJson) => {
        setTweets(resJson.timeline);
        console.log("data" + JSON.stringify(resJson));
      })
      .catch((errorMessage) => {
        console.log(errorMessage);
      });

    clearTimeout(timer);
    setTimer(setTimeout(getTweets, 200));
  };

  return (
    <div>
      {/* {!authState.isAuthenticated ? (
        <Login />
      ) : ( */}
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
                <NewTweet />

                <div className="site-card-border-less-wrapper">
                  {tweets.map((tweet) => (
                    <Tweet tweet={tweet} /> 
                  ))}
                </div>
              </div>
            </Content>
          </Layout>
        </Layout>
      </Router>
    </div>
  );
};

export default Home;

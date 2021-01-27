import "../App.css";

import { Form, Input, Button } from "antd";
import * as React from "react";
import { AuthContext } from "../App";

const layout = {
  labelCol: {
    span: 8,
  },
  wrapperCol: {
    span: 16,
  },
};

const NewTweet = () => {
  const { state: authState } = React.useContext(AuthContext);
  const [tweetText, setTweetText] = React.useState("");

  const handleSubmit = () => {
    console.log(authState.token);
    fetch("http://localhost:8090/tweet", {
      method: "post",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        text: tweetText,
        token: authState.token,
      }),
    })
      .then((response) => {
        if (response.status === 201) console.log(response.status + " created!");
        return response.json();
      })
      .then((data) => {
        console.log(data);
      });
  };

  return (
    <div className="newTweet">
      <h1>Latest Tweets</h1>
      <Form {...layout} name="nest-messages">
        <Form.Item name={["user", "introduction"]}>
          <Input.TextArea
            placeholder="What's happening?"
            onChange={(e) => setTweetText(e.target.value)}
          />
        </Form.Item>
        <Form.Item wrapperCol={{ ...layout.wrapperCol, offset: 8 }}>
          <Button type="primary" htmlType="submit" onClick={handleSubmit}>
            Tweet
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};

export default NewTweet;

import "../App.css";

import { Form, Input, Button } from "antd";
import * as React from "react";

const layout = {
  labelCol: {
    span: 8,
  },
  wrapperCol: {
    span: 16,
  },
};

export default class NewTweet extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      tweetText: "",
    };
  }

  setText = (text) => {
    this.setState({
      tweetText: text,
    });
  };

  handleSubmit = () => {
    console.log(this.props.token);
    fetch("http://localhost:8090/tweet", {
      method: "post",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        text: this.state.tweetText,
        token: this.props.token,
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

  render() {
    return (
      <div className="newTweet">
        <h1>Latest Tweets</h1>
        <Form {...layout} name="nest-messages">
          <Form.Item name={["user", "introduction"]}>
            <Input.TextArea
              placeholder="What's happening?"
              onChange={(e) => this.setText(e.target.value)}
            />
          </Form.Item>
          <Form.Item wrapperCol={{ ...layout.wrapperCol, offset: 8 }}>
            <Button
              type="primary"
              htmlType="submit"
              onClick={this.handleSubmit}
            >
              Tweet
            </Button>
          </Form.Item>
        </Form>
      </div>
    );
  }
}

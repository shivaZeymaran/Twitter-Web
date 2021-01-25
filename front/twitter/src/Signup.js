import "./App.css";
import * as React from "react";
import { Layout, Form, Input, Button, Checkbox } from "antd";

import { BrowserRouter as Router } from "react-router-dom";

const { Header, Content, Footer } = Layout;

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};
const tailLayout = {
  wrapperCol: { offset: 8, span: 16 },
};

export default class Signup extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      username: "",
      email: "",
      password: "",
    };
  }

  getUsername = (username) => {
    this.setState({ username: username });
  };

  getPassword = (password) => {
    this.setState({ password: password });
  };

  getEmailAddress = (email) => {
    this.setState({ email: email });
  };

  handleSubmit = (event) => {
    event.preventDefault();
    fetch("http://localhost:8090/signup", {
      method: "post",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        username: this.state.username,
        email: this.state.email,
        password: this.state.password,
      }),
    }).then((res) => {
      var status = res.status;
      console.log(status);
      console.log(res);
    });
  };

  render() {
    return (
      <Router>
        <div>
          <Layout>
            <Header>Header</Header>
            <Content>
              <Form
                {...layout}
                name="basic"
                initialValues={{ remember: true }}
                onSubmit={this.handleSubmit}
                // onFinish={onFinish}
                // onFinishFailed={onFinishFailed}
              >
                <Form.Item
                  label="Username"
                  name="username"
                  rules={[
                    { required: true, message: "Please input your username!" },
                  ]}
                >
                  <Input onChange={(e) => this.getUsername(e.target.value)} />
                </Form.Item>

                <Form.Item
                  name={["user", "email"]}
                  label="Email Address"
                  rules={[{ required: true, type: "email" }]}
                >
                  <Input
                    onChange={(e) => this.getEmailAddress(e.target.value)}
                  />
                </Form.Item>

                <Form.Item
                  label="Password"
                  name="password"
                  rules={[
                    { required: true, message: "Please input your password!" },
                  ]}
                >
                  <Input.Password
                    onChange={(e) => this.getPassword(e.target.value)}
                  />
                </Form.Item>

                <Form.Item
                  {...tailLayout}
                  name="remember"
                  valuePropName="checked"
                >
                  <Checkbox>Remember me</Checkbox>
                </Form.Item>

                <Form.Item {...tailLayout}>
                  <Button onClick={this.handleSubmit} type="primary" htmlType="submit">
                    Sign up
                  </Button>
                </Form.Item>
              </Form>
            </Content>
            <Footer>Footer</Footer>
          </Layout>
        </div>
      </Router>
    );
  }
}

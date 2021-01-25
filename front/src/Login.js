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

export default class Login extends React.Component {

constructor(props) {
  super(props);
  this.state = {
    username: "",
    password: "",
  }
}

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
                  <Input />
                </Form.Item>

                <Form.Item
                  label="Password"
                  name="password"
                  rules={[
                    { required: true, message: "Please input your password!" },
                  ]}
                >
                  <Input.Password />
                </Form.Item>

                <Form.Item
                  {...tailLayout}
                  name="remember"
                  valuePropName="checked"
                >
                  <Checkbox>Remember me</Checkbox>
                </Form.Item>

                <Form.Item {...tailLayout}>
                  <Button type="primary" htmlType="submit">
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


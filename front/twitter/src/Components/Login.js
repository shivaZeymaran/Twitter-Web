import "../App.css";
import Home from "./Home";
import * as React from "react";

import { Layout, Form, Input, Button, Checkbox } from "antd";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";

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
      token: "",
      isAuthenticated: false,
    };
  }

  setUsername = (username) => {
    this.setState({ username: username });
  };

  setPassword = (password) => {
    this.setState({ password: password });
  };

  handleSubmit = () => {
    fetch("http://localhost:8090/login", {
      method: "post",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: this.state.username,
        password: this.state.password,
      }),
    })
      .then((response) => {
        if (response.status === 200) console.log(response.status + " OK!");
        return response.json();
      })
      .then((data) => {
        this.setState({
          token: data.user.token,
          isAuthenticated: true,
        });
        console.log(data.user.token);
      });
  };

  render() {
    return (
      <Router>
        <Route exact path="/home" component={Home} />
        {!this.state.isAuthenticated && (
          <>
            <div>
              <Layout>
                <Header>Header</Header>
                <Content>
                  <Form
                    {...layout}
                    name="basic"
                    initialValues={{ remember: true }}
                  >
                    <Form.Item
                      label="Username"
                      name="username"
                      rules={[
                        {
                          required: true,
                          message: "Please input your username!",
                        },
                      ]}
                    >
                      <Input
                        onChange={(e) => this.setUsername(e.target.value)}
                      />
                    </Form.Item>

                    <Form.Item
                      label="Password"
                      name="password"
                      rules={[
                        {
                          required: true,
                          message: "Please input your password!",
                        },
                      ]}
                    >
                      <Input.Password
                        onChange={(e) => this.setPassword(e.target.value)}
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
                      <Link to="/home">
                        <Button
                          type="primary"
                          htmlType="submit"
                          onClick={this.handleSubmit}
                        >
                          Log in
                        </Button>
                      </Link>
                    </Form.Item>
                  </Form>
                </Content>
                <Footer>Footer</Footer>
              </Layout>
            </div>
          </>
        )}
      </Router>
    );
  }
}

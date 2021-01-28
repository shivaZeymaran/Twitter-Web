import "../App.css";
import Home from "./Home";
import React from "react";
import { AuthContext } from "../App";

import { Layout, Form, Input, Button, Checkbox } from "antd";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";

const { Content } = Layout;

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};
const tailLayout = {
  wrapperCol: { offset: 8, span: 16 },
};

const Login = () => {
  const { dispatch } = React.useContext(AuthContext);
  const initialState = {
    username: "",
    password: "",
    token: "",
    isSubmitting: false,
  };

  const [data, setData] = React.useState(initialState);

  const setUsername = (username) => {
    setData({
      ...data,
      username: username,
    });
  };

  const setPassword = (password) => {
    setData({
      ...data,
      password: password,
    });
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    setData({
      ...data,
      isSubmitting: true,
      errorMessage: null,
    });
    fetch("http://localhost:8090/login", {
      method: "post",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: data.username,
        password: data.password,
      }),
    })
      .then((res) => {
        console.log(res.status);
        if (res.status === 200) return res.json();
      })
      .then((resJson) => {
        console.log(resJson.user);
        dispatch({
          type: "LOGIN",
          payload: resJson.user,
        });
      })
      .catch((errorMessage) => {
        console.log(errorMessage);
      });
  };

  return (
    <div>
      <Router>
        <Route exact path="/home" component={Home} />

        <Layout className="log-in">
          <Content>
            <Form {...layout} name="basic" initialValues={{ remember: true }}>
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
                <Input onChange={(e) => setUsername(e.target.value)} />
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
                <Input.Password onChange={(e) => setPassword(e.target.value)} />
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
                    onClick={handleSubmit}
                  >
                    Log in
                  </Button>
                </Link>
              </Form.Item>
            </Form>
          </Content>
        </Layout>
      </Router>
    </div>
  );
};

export default Login;

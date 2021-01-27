import "../App.css";
import Home from "./Home";
import { AuthContext } from "../App";

import * as React from "react";
import { Layout, Form, Input, Button, Checkbox } from "antd";
import { BrowserRouter as Router, Link, Route } from "react-router-dom";

const { Header, Content, Footer } = Layout;

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};
const tailLayout = {
  wrapperCol: { offset: 8, span: 16 },
};

const Signup = () => {
  const { dispatch } = React.useContext(AuthContext);
  const initialState = {
    username: "",
    emailAddress: "",
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

  const setEmailAddress = (email) => {
    setData({
      ...data,
      emailAddress: email,
    });
  };

  const handleSubmit = async () => {
    await fetch("http://localhost:8090/signup", {
      method: "post",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: data.username,
        email: data.emailAddress,
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
          type: "SIGNUP",
          payload: resJson.user,
        });
      })
      .catch((errorMessage) => {
        console.log(errorMessage);
      });
  };

  return (
    <Router>
      <Route exact path="/home" component={Home} />
      {/* {!this.state.isAuthenticated && (
          <> */}
      <div>
        <Layout>
          <Header>Header</Header>
          <Content>
            <Form
              {...layout}
              name="basic"
              initialValues={{ remember: true }}
              onSubmit={handleSubmit}
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
                <Input onChange={(e) => setUsername(e.target.value)} />
              </Form.Item>

              <Form.Item
                name={["user", "email"]}
                label="Email Address"
                rules={[{ required: true, type: "email" }]}
              >
                <Input onChange={(e) => setEmailAddress(e.target.value)} />
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
                    Sign up
                  </Button>
                </Link>
              </Form.Item>
            </Form>
          </Content>
          <Footer>Footer</Footer>
        </Layout>
      </div>
      {/* </>
        )} */}
    </Router>
  );
};

export default Signup;

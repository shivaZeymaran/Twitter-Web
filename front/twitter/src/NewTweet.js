import "./App.css";
import { Form, Input, Button } from "antd";

const layout = {
  labelCol: {
    span: 8,
  },
  wrapperCol: {
    span: 16,
  },
};

const validateMessages = {
  required: "${label} is required!",
  types: {
    email: "${label} is not a valid email!",
    number: "${label} is not a valid number!",
  },
  number: {
    range: "${label} must be between ${min} and ${max}",
  },
};

function NewTweet() {
  return (
    <div className="newTweet">
      <h1>Latest Tweets</h1>
      <Form
        {...layout}
        name="nest-messages"
        validateMessages={validateMessages}
      >
        <Form.Item name={["user", "introduction"]}>
          <Input.TextArea placeholder="What's happening?" />
        </Form.Item>
        <Form.Item wrapperCol={{ ...layout.wrapperCol, offset: 8 }}>
          <Button type="primary" htmlType="submit">
            Tweet
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
}

export default NewTweet;

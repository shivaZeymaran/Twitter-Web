import { Card } from "antd";

const Tweet = ({ text }) => {
  return (
    <Card title="tweet zanande" bordered={false} style={{ width: 900 }}>
      <p>{text}</p>
    </Card>
  );
};

export default Tweet;

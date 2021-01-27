import { Card } from "antd";

const Tweet = ({ tweet }) => {
  return (
    <Card title={tweet.owner.username} bordered={true} style={{ width: 900 }}>
      <p>{tweet.text}</p>
    </Card>
  );
};

export default Tweet;

import React, { useState } from "react";

import ImgCrop from "antd-img-crop";
import { Input, Upload } from "antd";

const Profile = () => {
  const [file, setFile] = useState({
    uid: "-1",
    name: "image.png",
    status: "done",
    url:
      "https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png",
  });

  const [username, setUsername] = useState("");

  const changeUsername = (username) => {
    setUsername(username);
  };

  const onChange = ({ file: newFile }) => {
    setFile(newFile);
  };

  const onPreview = async (file) => {
    let src = file.url;
    if (!src) {
      src = await new Promise((resolve) => {
        const reader = new FileReader();
        reader.readAsDataURL(file.originFileObj);
        reader.onload = () => resolve(reader.result);
      });
    }
    const image = new Image();
    image.src = src;
    const imgWindow = window.open(src);
    imgWindow.document.write(image.outerHTML);
  };

  return (
    <div>
      <h1>hi vahideeeeeeee</h1>
      <ImgCrop rotate>
        <Upload
          action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
          listType="picture-card"
          fileList={file}
          onChange={onChange}
          onPreview={onPreview}
        >
          {file.length < 5 && "+ Upload"}
        </Upload>
      </ImgCrop>
      <Input
        placeholder="username"
        value={username}
        onChange={(e) => changeUsername(e.target.value)}
      />
    </div>
  );
};

export default Profile;

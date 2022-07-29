import axios from "axios";

/**
 * 自定义上传图片插件
 */
class MyUploadAdapter {
  constructor(loader) {
    this.loader = loader;
  }

  async upload() {
    const data = new FormData();
    data.append("img", await this.loader.file);

    const res = await axios({
      url: `/api/upload/img`,
      method: "POST",
      data,
      withCredentials: true // true 为不允许带 token, false 为允许
    });
    return {
      default: res.data.data.url
    };
  }
}

export default MyUploadAdapter;

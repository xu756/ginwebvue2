//封装axios
import axios from "axios";
import router from "../router/index";
import { Message, Loading } from "element-ui";
import moment from "moment";
import VueCookies from "vue-cookies";
let loadingInstance; // 加载全局的loading
const baseurl = "https://cs.xu756.top/api/vue2/user/";
const request = axios.create({
  baseURL: baseurl,
  timeout: 5000,
  headers: {
    sn: moment().format("YYYY-MM-DD hh:mm:ss"),
    "Content-Type": "application/x-www-form-urlencoded"
  }
});
//请求拦截器
request.interceptors.request.use(
  config => {
    //在请求发送之前做一些处理
    //判断路由path是否是登录页面
    if (router.currentRoute.path !== "/login") {
      //获取token
      const token = VueCookies.get("token");
      //如果token存在，则每个请求都带上token
      if (token) {
        config.headers.token = token;
        config.headers.username = VueCookies.get("pad");
      } else {
        //如果token不存在，则跳转到登录页面
        Message({
          message: "请重新登录",
          type: "error"
        });
        router.push("/login");
      }
    }
    //显示加载动画
    loadingInstance = Loading.service({
      text: "加载中...",
      spinner: "el-icon-loading",
      background: "rgba(0, 0, 0, 0.7)"
    });
    return config;
  },
  error => {
    //请求错误时做一些处理
    Message({
      message: "系统错误,请稍后再试",
      type: "error"
    });
    return Promise.reject(error);
  }
);
//请求拦截器
request.interceptors.response.use(
  response => {
    //在请求成功之后做一些处理
    //关闭加载动画
    loadingInstance.close();
    //判断返回的状态码是否是200
    if (response.status === 200) {
      //如果是200，则返回数据
      return response.data;
    }
  },
  error => {
    //请求错误时做一些处理
    //关闭加载动画
    loadingInstance.close();
    Message({
      message: "请求错误,请稍后再试",
      type: "error"
    });
    return Promise.reject(error);
  }
);
/**
 * 封装post请求
 * @param url
 * @param data
 * @returns {Promise}
 */

export function post(url, data = {}) {
  return new Promise((resolve, reject) => {
    request.post(url, data).then(
      res => {
        if (res.type === "loginout") {
          Message({
            message: "登录超时，请重新登录",
            type: "error"
          });
          router.push("/login");
          return;
        }
        if (res.type === "success") {
          Message({
            message: res.msg,
            type: "success"
          });
        }
        resolve(res);
      },
      err => {
        reject(err);
      }
    );
  });
}

/**
 * 封装get请求
 * @param url
 * @param data
 * @returns {Promise}
 */
export function get(url, data = {}) {
  return new Promise((resolve, reject) => {
    request.get(url, { params: data }).then(
      res => {
        if (res.type === "loginout") {
          Message({
            message: "登录超时，请重新登录",
            type: "error"
          });
          router.push("/login");
          return;
        }
        if (res.type === "success") {
          Message({
            message: res.msg,
            type: "success"
          });
        }
        resolve(res);
      },
      err => {
        reject(err);
      }
    );
  });
}

export default {
  get,
  post
};

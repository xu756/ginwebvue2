//封装axios
import axios from 'axios';
import qs from 'qs';
import router from '../router/index';
let loadingInstance // 加载全局的loading
var that
that=this
const baseurl = 'http://localhost:7000/api/user/';
const request = axios.create({
	baseURL: baseurl,
	timeout: 5000,
	headers: {
		sn: new Date().getTime()
	}
});
export default request
//请求拦截器
request.interceptors.request.use(
	(config) => {
		//在请求发送之前做一些处理
		//判断路由path是否是登录页面
		if (router.currentRoute.path !== '/login') {
			//获取token
			const token = localStorage.getItem('token');
			//如果token存在，则每个请求都带上token
			if (token) {
				config.headers.token = token;
			} else {
				//如果token不存在，则跳转到登录页面
				this.$message({
					message: '请重新登录',
					type: 'error'
				});
				router.push('/login');
			}
		}
		//显示加载动画
		loadingInstance = Loading.service({
			text: '加载中...',
			spinner: 'el-icon-loading',
			background: 'rgba(0, 0, 0, 0.7)'
		});
		return config;
	},
	(error) => {
		//请求错误时做一些处理
		that.$message({
			message: '系统错误,请稍后再试',
			type: 'error'
		});
		return Promise.reject(error);
	}
);
//请求拦截器
request.interceptors.response.use(
	(response) => {
		//在请求成功之后做一些处理
		//关闭加载动画
		// loadingInstance.close();
		//判断返回的状态码是否是200
		if (response.status === 200) {
			//如果是200，则返回数据
			return response.data;
		}
	},
	(error) => {
		//请求错误时做一些处理
		//关闭加载动画
		// loadingInstance.close();
		that.$message({
			message: '请求错误,请稍后再试',
			type: 'error'
		});
		return Promise.reject(error);
	}
);
/**
 * 
 * @param {string} url 请求地址
 * @param {object} params 	请求参数
 * @returns 					Promise对象
 */
export const get = (url, params) => {
	return new Promise((resolve, reject) => {
		request
			.get(url, {
				params: params
			})
			.then((response) => {
				resolve(response);
			})
			.catch((error) => {
				reject(error);
			});
	});
};
/** 
*	@param {string} url 请求地址
*	@param {object} params 	请求参数
*/
export const post = (url, params) => {
	return new Promise((resolve, reject) => {
		request
			.post(url, qs.stringify(params))
			.then((response) => {
				resolve(response);
			})
			.catch((error) => {
				reject(error);
			});
	});
};

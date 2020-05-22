/* Axios 方式 */
import axios from 'axios';
import qs from 'qs'; // axios发送的数据不是json格式，若需要json格式，添加此库

let baseUrl = '';
if (process.env.NODE_ENV === 'production') {
  baseUrl = 'http://127.0.0.1:8888';
  // baseUrl = 'https://book.missyo.vip';
} else {
  baseUrl = 'http://127.0.0.1:8888';
  // baseUrl = 'https://book.missyo.vip';
}
// Add a request interceptor
axios.interceptors.request.use(config => {
  // Do something before request is sent
  config.withCredentials = false; // 需要跨域打开此配置
  // post提交 data存在 并且 data不是FormData对象时对数据进行json化处理
  if (config.method === 'post' && config.data && config.data.constructor !== FormData) {
    config.data = qs.stringify(config.data);
    config.headers['Content-Type'] = 'application/x-www-form-urlencoded';
  }
  return config;
}, function (error) {
  // Do something with request error
  return Promise.reject(error);
});

// Add a response interceptor
axios.interceptors.response.use(response => {
  // Do something with response data
  return response.data;
}, function (error) {
  return Promise.reject(error);
});

const CancelToken = axios.CancelToken;

export default {
  get(url, params) {
    return axios({
      method: 'get',
      url: baseUrl + url,
      params,
      timeout: 10000,
      withCredentials: false,
      headers: {
        'Content-Type': 'application/x-www-form-encoding'
      }
    });
  },
  weichat_get(url) {
    return axios({
      method: 'get',
      url: url,
      timeout: 50000,
      withCredentials: false,
      headers: {
        'Content-Type': 'application/x-www-form-encoding'
      }
    });
  },
  post(url, data) {
    return axios({
      method: 'post',
      url: baseUrl + url,
      data: data,
      timeout: 10000,
    });
  },
  options(url, data) {
    return axios({
      method: 'options',
      url: baseUrl + url,
      data: data,
      timeout: 10000,
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
      }
    });
  },
  put(url, data) {
    return axios({
      method: 'put',
      url: baseUrl + url,
      data: data,
      // withCredentials: true,
      timeout: 10000,

      headers: {
        'Content-Type': ' '
      }
    });
  },
};
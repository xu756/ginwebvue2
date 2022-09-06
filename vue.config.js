// const webpack = require('webpack')

module.exports = {
  publicPath: "./",
  assetsDir: "static",
  transpileDependencies: true,
  // 跨域支持
  devServer: {
    proxy: {
      "/api": {
        target: "https://cs.xu756.top/api/",
        changeOrigin: true, // 允许跨域
        pathRewrite: {
          "^/api": ""
        }
      }
    }
  }
};

// const webpack = require('webpack')


module.exports = {
  // configureWebpack: {
  //   plugins: [
  //     new webpack.ProvidePlugin({
  //       'window.Quill': 'quill/dist/quill.js',
  //       'Quill': 'quill/dist/quill.js'
  //     })
  //   ]
  // },
  transpileDependencies: true,
  // 跨域支持
  devServer: {
    proxy: {
      "/api/vue2/user": {
        target: "http://localhost:7000/api/vue2/user/",
        changeOrigin: true, // 允许跨域
        pathRewrite: {
          "^/api/vue2/user": ""
        }
      },
      "/upload": {
        target: "http://localhost:7000/upload/",
        changeOrigin: true, // 允许跨域
        pathRewrite: {
          "^/upload": ""
        }
      },
      "/get": {
        target: "http://localhost:7000/get/upload",
        changeOrigin: true, // 允许跨域
        pathRewrite: {
          "^/get": ""
        }
      }
    }
  }
}

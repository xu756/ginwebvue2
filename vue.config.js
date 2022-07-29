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
      "/api": {
        target: "http://localhost:7000/api/",
        changeOrigin: true, // 允许跨域
        pathRewrite: {
          "^/api": ""
        }
      }
    }
  }
}

/*
 * @Author: Firefly
 * @Date: 2021-04-04 22:46:34
 * @Descripttion:
 * @LastEditTime: 2021-04-24 15:21:35
 */
const { createProxyMiddleware } = require("http-proxy-middleware");

module.exports = (app) => {
  app.use(
    "/todo",
    createProxyMiddleware({
      target: "https://todo.firego.cn",
      changeOrigin: true,
      headers: {
        cookie:
          "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJhZ2w4cDh4Ymswc2ciLCJleHAiOjE2MTk4NzkzMDN9.tQjxhEO-MU0Ef98_31h_RLHWqLFxUqdRptJY9N3jJ7s",
      },
    })
  );
};

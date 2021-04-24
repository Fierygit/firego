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
      // target: "https://todo.firego.cn",
      target: "http://localhost:8716",
      changeOrigin: true,
      headers: {
        cookie:
          "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJhZ2w4cDh4Ymswc2ciLCJleHAiOjE2MTk1MjU5MDV9.2-IusUVLrZqkaZD6jCKLhxNIfkupGb9HwY9hZJMfKlA",
      },
    })
  );
};

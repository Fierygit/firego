/*
 * @Author: Firefly
 * @Date: 2021-04-04 22:46:34
 * @Descripttion:
 * @LastEditTime: 2021-04-23 15:31:25
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
          "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJhaXk0eGQwMjlneHMiLCJleHAiOjE2MTk3Njc1MTF9.RLHFfjkWozaSuWCu-hw_Wi-rxqsG90cuuYAs_b1ZYxE",
      },
    })
  );
};

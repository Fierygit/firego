const { createProxyMiddleware } = require("http-proxy-middleware");

module.exports = (app) => {
  app.use(
    "/todo",
    createProxyMiddleware({
      target: "https://todo.firego.cn",
      // target: "http://localhost:8716",
      changeOrigin: true,
      headers: {
        cookie:
          "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJhZ2w4cDh4Ymswc2ciLCJleHAiOjE2MTk4NzkzMDN9.tQjxhEO-MU0Ef98_31h_RLHWqLFxUqdRptJY9N3jJ7s",
          // "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJhZ2tyM3hybXYyZjQiLCJleHAiOjE2MjAxMDU4NjZ9.aTtlen0-4uJ6OXJiJIIoQAdgaB1eEI-_-TIgucg_edk",
      },
    })
  );
};

const { createProxyMiddleware } = require("http-proxy-middleware");

module.exports = app => {
    app.use(
        "/todo",
        createProxyMiddleware({
            target: "https://todo.firego.cn",
            changeOrigin: true,
            headers: {
                "cookie": "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJhZ2w4cDh4Ymswc2ciLCJleHAiOjE2MTk1MjU5MDV9.2-IusUVLrZqkaZD6jCKLhxNIfkupGb9HwY9hZJMfKlA"
            }
        })
    );
};
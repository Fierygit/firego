const { createProxyMiddleware } = require("http-proxy-middleware");

module.exports = app => {
    app.use(
        "/todo",
        createProxyMiddleware({
            target: "https://todo.firego.cn",
            changeOrigin: true,
            headers: {
                "cookie": "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJhZ2w4cDh4Ymswc2ciLCJleHAiOjE2MTkzNTc3OTR9.u6fGao6zRiZrgqG1d8P3kYdXBVm4Y1cLP2KhIZJf5ng"
            }
        })
    );
};
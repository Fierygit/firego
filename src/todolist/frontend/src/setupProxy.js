const { createProxyMiddleware } = require("http-proxy-middleware");

module.exports = app => {
    app.use(
        "/todo",
        createProxyMiddleware({
            target: "https://todo.firego.cn",
            changeOrigin: true,
            headers: {
                "cookie": "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJhZ2w4cDh4Ymswc2ciLCJleHAiOjE2MTc1NDQ0NzZ9.Q00QKC_Mxbjpw7tHm0AjdD4EX3m91PxOv8LGFpUB4d8"
            }
        })
    );
};
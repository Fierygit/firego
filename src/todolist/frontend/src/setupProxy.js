const { createProxyMiddleware } = require("http-proxy-middleware");

module.exports = app => {
    app.use(
        "/todo",
        createProxyMiddleware({
            target: "https://todo.firego.cn",
            changeOrigin: true,
            headers: {
                "cookie": "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJhZ2w4cDh4Ymswc2ciLCJleHAiOjE2MTg3NTI5MTl9.QCvz5nuNaia_oZgNw0bwlzswG-Qpe85Sf2_IT_Vxa_c"
            }
        })
    );
};
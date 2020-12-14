const Koa = require('koa');

const app = new Koa();
const serverPort = process.env.SERVER_PORT || 80;

app.use((ctx, next) => {
});

app.listen(serverPort);

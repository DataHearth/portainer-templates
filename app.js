const Koa = require('koa');
const { createWriteStream } = require('fs');
const { join } = require('path');
const morgan = require('koa-morgan');
const logger = require('./src/logger');

const app = new Koa();
const serverPort = process.env.SERVER_PORT || 80;

app.use(morgan('combined', process.env.PRODUCTION ? { stream: createWriteStream(join(__dirname, 'logs', 'routing.log')) } : {}));
app.use(require('./src/api').routes());

app.listen(serverPort, () => {
  logger.info('Server listening', { port: serverPort });
});

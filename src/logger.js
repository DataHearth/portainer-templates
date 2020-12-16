const {
  createLogger, transports, format: {
    combine, timestamp, json,
  },
} = require('winston');
const { PRODUCTION, LOG_LEVEL } = require('./env');

const logger = createLogger({
  level: LOG_LEVEL,
  format: combine(
    json(),
    timestamp({ format: 'DD-MM-YYYY HH:mm:ss' }),
  ),
  transports: [new transports.Console({
    format: combine(
      json(),
      timestamp({ format: 'DD-MM-YYYY HH:mm:ss' }),
    ),
  })],
});

if (PRODUCTION) {
  logger.add(new transports.File({
    filename: 'logs/application.log',
    maxFiles: 100000,
    handleExceptions: true,
    format: combine(
      json(),
      timestamp({ format: 'DD-MM-YYYY HH:mm:ss' }),
    ),
  }));
}

module.exports = logger;

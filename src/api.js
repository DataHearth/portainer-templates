const Router = require('@koa/router');
const fs = require('fs');
const path = require('path');
const logger = require('./logger');
const { gatherTemplates } = require('./main');

const router = new Router();

router.get('/api/templates/:templateType', async (ctx) => {
  if (!['docker', 'swarm', 'composer'].includes(ctx.params.templateType)) {
    logger.error('templateType parameter is invalid', { templateType: ctx.params.templateType });
    ctx.throw(
      {
        message: 'provided template type is not recognized',
        details: {
          templateType: ctx.params.templateType,
          validTypes: ['docker', 'swarm', 'composer'],
        },
      },
    );
  }
  const templatesPath = path.join(__dirname, '../', 'templates');
  const templates = [];

  try {
    const directoriesPath = fs.readdirSync(templatesPath, { withFileTypes: true })
      .filter((content) => content.isDirectory())
      .map((directory) => directory.name)
      .map((name) => path.join(templatesPath, name));
    logger.debug('template directory scanned', { path: templatesPath, directories: directoriesPath });

    const templatePromises = [];
    for (let i = 0; i < directoriesPath.length; i += 1) {
      // eslint-disable-next-line no-await-in-loop
      templatePromises.push(gatherTemplates(directoriesPath[i], 'docker'));
    }

    const values = await Promise.all(templatePromises);
    values.forEach((fileTemplates) => {
      templates.push(...fileTemplates);
    });
  } catch (error) {
    logger.error(error.message, { ...error.details });
    ctx.throw(error);
    return;
  }

  ctx.body = {
    version: '2',
    templates,
  };
});

router.get('/api/templates/:templateType/:category', () => {

});

module.exports = router;

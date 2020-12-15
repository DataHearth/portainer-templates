const Router = require('@koa/router');
const fs = require('fs');
const path = require('path');
const logger = require('./logger');
const { gatherTemplates } = require('./main');

const router = new Router();

router.get('/api/templates/:templateType', async (ctx) => {
  // check if template type is available in available in portainer or implemented
  if (!['docker', 'swarm'].includes(ctx.params.templateType)) {
    logger.error('templateType parameter is invalid', { templateType: ctx.params.templateType });
    ctx.throw(
      {
        message: 'provided template type is not recognized',
        details: {
          templateType: ctx.params.templateType,
          validTypes: ['docker', 'swarm'],
        },
      },
    );
  }
  logger.debug('template type is valid and implemented');

  const templatesPath = path.join(__dirname, '../', process.env.TEMPLATES_FOLDER || 'templates', ctx.params.templateType);
  const templates = [];

  try {
    // gather all directories in template folder
    const directoriesPath = fs.readdirSync(templatesPath, { withFileTypes: true })
      .filter((content) => content.isDirectory())
      .map((directory) => directory.name)
      .map((name) => path.join(templatesPath, name));
    logger.debug('template directory scanned', { path: templatesPath, directories: directoriesPath });

    // create a promise for each directory
    const templatePromises = [];
    for (let i = 0; i < directoriesPath.length; i += 1) {
      // eslint-disable-next-line no-await-in-loop
      templatePromises.push(gatherTemplates(directoriesPath[i], 'docker'));
    }

    // resolve promises array in paralelle for performance
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

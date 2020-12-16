const Router = require('@koa/router');
const fs = require('fs');
const path = require('path');
const logger = require('./logger');
const { gatherTemplates, checkImplementation } = require('./utils');
const { TEMPLATES_FOLDER } = require('./env');

const router = new Router();

router.get('/api/templates/:templateType', async (ctx) => {
  // check if template type is available in available in portainer or implemented
  try {
    checkImplementation(ctx.params.templateType);
  } catch (error) {
    ctx.throw(400, error);
  }

  const templatesPath = path.join(__dirname, '../', TEMPLATES_FOLDER, ctx.params.templateType);
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
    directoriesPath.forEach((dirPath) => {
      templatePromises.push(gatherTemplates(dirPath, 'docker'));
    });

    // resolve promises array in paralelle for performance
    const values = await Promise.all(templatePromises);
    values.forEach((fileTemplates) => {
      templates.push(...fileTemplates);
    });
  } catch (error) {
    logger.error(error.message, { ...error.details });
    ctx.throw(500, error);
    return;
  }

  ctx.body = {
    version: '2',
    templates,
  };
});

router.get('/api/templates/:templateType/:category', async (ctx) => {
  const { templateType, category } = ctx.params;
  // check if template type is available in available in portainer or implemented
  try {
    checkImplementation(templateType);
  } catch (error) {
    ctx.throw(400, error);
  }

  let templates;
  try {
    const dirPath = path.join(__dirname, '../', TEMPLATES_FOLDER, templateType, category);
    templates = await gatherTemplates(dirPath, 'docker');
  } catch (error) {
    ctx.throw(500, error);
  }

  ctx.body = {
    version: '2',
    templates,
  };
});

module.exports = router;

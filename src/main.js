const path = require('path');
const fs = require('fs');
const { dockerTemplate, swarmTemplate } = require('./default-template');
const logger = require('./logger');

const readDirContent = (dirPath) => {
  // list directory's files
  const files = fs.readdirSync(dirPath);
  for (let index = 0; index < files.length; index += 1) {
    // set a path like
    files[index] = path.join(dirPath, files[index]);
  }

  return files;
};

const readFileContent = async (filePath, template) => {
  if (!fs.existsSync(filePath)) {
    // eslint-disable-next-line no-throw-literal
    throw { message: 'file doesn\'t exist' };
  }
  logger.debug('file exist', { path: filePath });

  let fileTemplate;
  const content = fs.readFileSync(filePath, { encoding: 'utf8', flag: 'r' });
  try {
    fileTemplate = JSON.parse(content);
  } catch (error) {
    // eslint-disable-next-line no-throw-literal
    throw { message: `invalid JSON for file: ${filePath}`, details: error.message };
  }
  logger.debug('parsing completed');

  const validationResults = [];
  const validationPromise = [];
  const errors = [];
  if (template === 'docker') {
    logger.debug('docker template selected');
    // eslint-disable-next-line no-restricted-syntax
    for (const data of fileTemplate.templates) {
      validationPromise.push(
        new Promise((resolve) => {
          const result = dockerTemplate.validate(data);
          if (result.errors) {
            errors.push({
              templateTitle: data.Title,
              errorName: result.errors.name,
              errorDetails: result.errors.details,
            });
          } else {
            validationResults.push(result.value);
            resolve();
          }
        }),
      );
    }
  } else if (template === 'swarm') {
    // eslint-disable-next-line no-restricted-syntax
    for (const data of fileTemplate.templates) {
      validationPromise.push(
        new Promise((resolve) => {
          const result = swarmTemplate.validate(data);
          if (result.errors) {
            errors.push({
              templateTitle: data.Title,
              errorName: result.errors.name,
              errorDetails: result.errors.details,
            });
          } else {
            validationResults.push(result.value);
            resolve();
          }
        }),
      );
    }
  } // todo: add compose template

  await Promise.all(validationPromise);

  if (errors.length !== 0) {
    // eslint-disable-next-line no-throw-literal
    throw { message: 'some templates are invalid', details: errors };
  }

  return validationResults;
};

const gatherTemplates = async (dirPath, template) => {
  const templates = [];
  const filesPath = readDirContent(dirPath);
  logger.debug('directory scanned', { path: dirPath, files: filesPath });

  // eslint-disable-next-line no-restricted-syntax
  for (const filePath of filesPath) {
    // eslint-disable-next-line no-await-in-loop
    templates.push(await readFileContent(filePath, template));
  }

  return templates;
};

module.exports = {
  readDirContent, readFileContent, gatherTemplates,
};

/* eslint-disable no-throw-literal */
const path = require('path');
const fs = require('fs');
const { dockerTemplate, swarmTemplate } = require('./default-template');
const logger = require('./logger');

const readDirContent = (dirPath) => {
  // list directory's files
  let files;
  try {
    files = fs.readdirSync(dirPath);
  } catch (error) {
    throw { message: `invalid directory or path: ${dirPath}` };
  }

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

  let content;
  let fileTemplate;
  try {
    content = fs.readFileSync(filePath, { encoding: 'utf8', flag: 'r' });
  } catch (error) {
    throw { message: `invalid path or file: ${filePath}`, details: error.message };
  }
  try {
    fileTemplate = JSON.parse(content);
  } catch (error) {
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
    templates.push(...await readFileContent(filePath, template));
  }

  return templates;
};

module.exports = {
  readDirContent, readFileContent, gatherTemplates,
};

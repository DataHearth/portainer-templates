const joi = require('joi');

// * Default Portainer-CE templates definitions: https://raw.githubusercontent.com/portainer/templates/master/templates-2.0.json

const portRegex = /^((6553[0-5])|(655[0-2][0-9])|(65[0-4][0-9]{2})|(6[0-4][0-9]{3})|([1-5][0-9]{4})|([0-5]{0,5})|([0-9]{1,4})):((6553[0-5])|(655[0-2][0-9])|(65[0-4][0-9]{2})|(6[0-4][0-9]{3})|([1-5][0-9]{4})|([0-5]{0,5})|([0-9]{1,4}))\/(tcp|udp)$/m;

const dockerTemplate = joi.object({
  type: 1, // * don't need to be set 👌
  title: joi.string().length(50).message('title can\'t be more than 50 characters long')
    .required(),
  description: joi.string().required(),
  image: joi.string().required(), // todo: add a custom image regex
  administrator_only: joi.bool().default(false),
  name: joi.string(),
  logo: joi.string().uri().message('logo needs to be a valid URI'),
  registry: joi.string(), // todo: add a custom registery regex
  command: joi.string(),
  env: joi.array().items(joi.object({
    name: joi.string().required(),
    label: joi.alternatives()
      .conditional('preset', {
        is: true,
        then: joi.string().required(),
        otherwise: joi.string(),
      }),
    description: joi.string(),
    default: joi.any(),
    preset: joi.bool().default(false),
    select: joi.array().items(joi.object({
      text: joi.string().required(),
      value: joi.string().required(),
      default: joi.bool().default(false),
    })),
  })),
  network: joi.string().default('bridge'),
  volumes: joi.array().items(joi.object({
    container: joi.string().required(),
    bind: joi.string(),
    readonly: joi.bool().default(true),
  })),
  ports: joi.array().items(
    joi.string().regex(portRegex).message('ports\'s item doesn\'t match format (2-65535:2-65535/tcp|udp)'),
  ),
  labels: joi.array().items(joi.object({
    name: joi.string().required(),
    value: joi.string().required(),
  })),
  privileged: joi.bool().default(false),
  interactive: joi.bool().default(false),
  restart_policy: joi.string().valid('no', 'unless-stopped', 'on-failure', 'always').default('always'),
  hostname: joi.string().length(50).message('hostname can\'t be more than 50 characters long'),
  note: joi.string(),
  platform: joi.string().valid('linux', 'windows')
    .default('linux'),
  categories: joi.array().items(joi.string()),
});

const swarmTemplate = joi.object({
  type: joi.number().default(2), // * don't need to be set 👌
  title: joi.string().length(50).required(),
  description: joi.string().required(),
  repository: joi.object({
    url: joi.string().uri().required(),
    stackfile: joi.string().required(),
  }).required(),
  administrator_only: joi.bool().default(false), // ? Optional
  name: joi.string().length(50),
  logo: joi.string().uri(),
  env: joi.array().items(joi.object({
    name: joi.string().required(),
    label: joi.alternatives().conditional('preset', { is: true, then: joi.string().required(), otherwise: joi.string() }),
    description: joi.string(),
    default: joi.any(),
    preset: joi.bool().default(false),
    select: joi.array().items(joi.object({
      text: joi.string().required(),
      value: joi.string().required(),
      default: joi.bool().default(false),
    })),
  })),
  note: joi.string(),
  platform: joi.string().valid('linux', 'windows').default('linux'),
  categories: joi.array().items(joi.string()),
});

module.exports = {
  dockerTemplate, swarmTemplate,
};

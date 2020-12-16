FROM node:15-alpine3.10

WORKDIR /home/portainer-templates

ADD . .

RUN npm i

ENV NPM_CONFIG_LOGLEVEL info 
ENV LOG_LEVEL info
ENV TEMPLATES_FOLDER templates
ENV PRODUCTION true

CMD [ "npm", "start" ]
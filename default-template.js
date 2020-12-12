// * Default Portainer-CE templates definitions: https://raw.githubusercontent.com/portainer/templates/master/templates-2.0.json

const dockerTemplate = {
  "types": 1, // ! Mandatory
  "title": "Title of the template", // ! Mandatory
  "description": "Template description", // ! Mandatory
  "image": "Docker image", // ! Mandatory
  "administrator_only": false, // ? Optional
  "name": "Name of the template in the UI", // ? Optional
  "logo": "Logo URL", // ? Optional
  "registry": "Registry where the image is stored (default dockerhub)", // ? Optional
  "command": "Command to run in the container", // ? Optional
  "env": [ // ? Optional
    {
      "name": "Name of the variable", // ! Mandatory
      "label": "Label for the UI", // ! Mandatory unless preset is present
      "description": "Description for the UI", // ? Optional
      "default": "Default value", // ? Optional
      "preset": false, // ? Optional : if false => UI will generate an input
      "select": [ // ? Optional : Select input
        {
          "text": "Yes, I agree",
          "value": "Y",
          "default": true
        }
      ]
    }
  ],
  "network": "Network to use (default auto-select the network (if it exists) in the templates view)", // ? Optional
  "volumes": [ // ? Optional
    {
      "container": "Container path", // ! Mandatory
      "bind": "/var/www", // ? Optional
      "readonly": true
    }
  ],
  "ports": [ // ? Optional
    "host:container/protocol" // * e.g: 8080:80/tcp, 4443:443/udp
  ],
  "labels": [ // ? Optional
    {
      "name": "Label name", // ! Mandatory
      "value": "Label value" // ! Mandatory
    }
  ],
  "privileged": false, // ? Optional : default "false"
  "interactive": false, // ? Optional : should container be started in foreground (default "false")
  "restart_policy": "no | unless-stopped | on-failure | always", // ? Optional : default "always"
  "hostname": "Name of the container", // ? Optional : default generated docker hostname
  "note": "Notes about the template", // ? Optional
  "platform": "linux | windows", // ? Optional
  "categories": ["category name"] // ? Optional
};

const stackTemplate = {
  "type": 2, // ! Mandatory
  "title": "Title of the template", // ! Mandatory
  "description": "Template description", // ! Mandatory
  "repository": { // ! Mandatory
    "url": "https://github.com/portainer/templates",
    "stackfile": "stacks/cockroachdb/docker-stack.yml"
  },
  "administrator_only": false, // ? Optional
  "name": "Name of the template", // ? Optional
  "logo": "https://cloudinovasi.id/assets/img/logos/cockroachdb.png", // ? Optional
  "env": [ // ? Optional
    {
      "name": "Name of the variable", // ! Mandatory
      "label": "Label for the UI", // ! Mandatory unless set is present
      "description": "Description for the UI", // ? Optional
      "default": "Default value", // ? Optional
      "preset": false, // ? Optional : if false => UI will generate an input
      "select": [ // ? Optional : Select input
        {
          "text": "Yes, I agree",
          "value": "Y",
          "default": true
        }
      ]
    }
  ],
  "note": "Note about the template", // ? Optional
  "platform": "linux | windows", // ? Optional
  "categories": ["category name"] // ? Optional
};
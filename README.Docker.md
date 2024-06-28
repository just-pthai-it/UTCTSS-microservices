### Building and running your application on local

When you're ready, start your application by running:

`docker compose -f compose-local.yaml --env-file .env.local up -d --build`

Stop your application by running:

`docker compose -f compose-local.yaml --env-file .env.local down`


### Building and running your application on production

When you're ready, start your application by running:

`docker compose -f compose-production.yaml --env-file .env.production up -d --build`

Stop your application by running:

`docker compose -f compose-production.yaml --env-file .env.production up down`


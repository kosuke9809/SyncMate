#!/bin/sh
# entrypoint.sh

if [ "$NEXT_ENV" = "dev" ]; then
    echo "Running in dev mode"
    yarn start:dev
elif [ "$NEXT_ENV" = "prod" ]; then
    echo "Running in prod mode"
    yarn start:prod
else
    echo "NEXT_ENV not set"
    exit 1
fi
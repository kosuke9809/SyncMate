FROM node:20.12.1-bullseye

WORKDIR /app

COPY ./frontend/package.json .
COPY ./frontend/yarn.lock .

RUN yarn install

COPY ./frontend/ .

COPY ./frontend/entrypoint.sh /user/local/bin/

ENTRYPOINT [ "/user/local/bin/entrypoint.sh" ]
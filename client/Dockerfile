# syntax=docker/dockerfile:1

# バージョンを指定
ARG NODE_VERSION=22

FROM node:${NODE_VERSION}-alpine AS base



FROM base As final

WORKDIR /opt/app

# Expose the port that the application listens on.
EXPOSE 8080

CMD ["sh", "./setup.sh"]
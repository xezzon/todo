FROM node:lts-alpine AS base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
ENV NODE_ENV=production
RUN corepack enable
WORKDIR /app

FROM base AS builder
COPY pnpm-workspace.yaml .
COPY todo-proto ./todo-proto
COPY todo-fe ./todo-fe
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm i
WORKDIR /app/todo-proto
RUN pnpm run generate
WORKDIR /app/todo-fe
RUN pnpm run build
EXPOSE 3000
CMD ["pnpm", "start"]

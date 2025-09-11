FROM node:lts-alpine AS generator
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
WORKDIR /todo-proto
COPY todo-proto .
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm i
RUN pnpm run generate

FROM golang:1.25 AS builder
WORKDIR /todo-service
COPY todo-service .
COPY --from=generator /todo-service/gen ./gen
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o todo-server ./cmd/server/main.go
CMD tail -f /dev/null

FROM scratch
WORKDIR /app
COPY --from=builder /todo-service/todo-server .
EXPOSE 8080
ENTRYPOINT ["/app/todo-server"]

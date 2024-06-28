FROM node:20-alpine as frontend

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

WORKDIR /app

COPY . .
WORKDIR /app/client

RUN pnpm i --frozen-lockfile
RUN pnpm build

FROM golang:1.22-alpine as backend

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o ./server-out

FROM alpine

WORKDIR /app

COPY --from=frontend /app/client/dist client/dist
COPY --from=backend /app/server-out server-out

CMD [ "./server-out" ]


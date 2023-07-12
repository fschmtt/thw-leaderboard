FROM node:18-alpine as node-builder

WORKDIR /app

COPY public ./public
COPY src ./src
COPY index.html ./
COPY env.d.ts ./
COPY package.json ./
COPY package-lock.json ./
COPY tsconfig.app.json ./
COPY tsconfig.json ./
COPY tsconfig.node.json ./
COPY vite.config.ts ./

RUN npm ci && \
    npm run build

FROM golang:1.20-alpine as go-builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY main.go ./
COPY api ./api
COPY db ./db
COPY migrations ./migrations

COPY --from=node-builder /app/dist ./frontend

RUN go mod download && \
    go build -o leaderboard-api main.go

EXPOSE 80
CMD ["./leaderboard-api"]
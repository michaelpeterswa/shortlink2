# -=-=-=-=-=-=- Compile Image -=-=-=-=-=-=-

FROM node AS builder

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci

COPY . .

RUN npm run build

# -=-=-=-=-=-=- Final Image -=-=-=-=-=-=-

FROM node:alpine

WORKDIR /app
COPY --from=builder /app/build .
COPY . .

EXPOSE 3000

ENTRYPOINT ["node", "."]
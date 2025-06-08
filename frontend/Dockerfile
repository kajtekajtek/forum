# stage 1: build next.js in the node environment
ARG NODE_VERSION=18
FROM node:${NODE_VERSION}-alpine AS builder

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm ci

COPY . .

ARG KEYCLOAK_URL
ARG KEYCLOAK_REALM
ARG KEYCLOAK_CLIENT_ID
ARG API_URL
ENV NEXT_PUBLIC_KEYCLOAK_URL=${KEYCLOAK_URL}
ENV NEXT_PUBLIC_KEYCLOAK_REALM=${KEYCLOAK_REALM}
ENV NEXT_PUBLIC_KEYCLOAK_CLIENT_ID=${KEYCLOAK_CLIENT_ID}
ENV NEXT_PUBLIC_API_URL=${API_URL}

RUN npm run build

FROM node:${NODE_VERSION}-alpine AS runner

RUN addgroup -S app && adduser -S -G app app

WORKDIR /home/app

COPY --from=builder /app/package.json /app/package-lock.json ./
RUN npm ci --only=production

COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/next.config.mjs ./

USER app

EXPOSE 3000

CMD ["npm", "run", "start"]

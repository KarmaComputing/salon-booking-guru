# Install dependencies only when needed
FROM node:alpine AS deps
# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY ./front-end/package.json ./front-end/yarn.lock ./
RUN yarn install --frozen-lockfile

# Rebuild the source code only when needed
FROM node:alpine AS builder
WORKDIR /app
COPY ./front-end .
COPY --from=deps /app/node_modules ./node_modules
RUN yarn build && yarn install --production --ignore-scripts --prefer-offline

# Production image, copy all the files and run an apache server
FROM httpd:alpine AS runner
WORKDIR /app
EXPOSE 80

RUN echo -e "<IfModule mod_rewrite.c>\n \
  RewriteEngine On\n \
  RewriteBase /\n \
  RewriteRule ^index\.html$ - [L]\n \
  RewriteCond %{REQUEST_FILENAME} !-f\n \
  RewriteCond %{REQUEST_FILENAME} !-d\n \
  RewriteRule . /index.html [L]\n \
</IfModule>" > /usr/local/apache2/htdocs/.htaccess
RUN sed -i 's/#LoadModule rewrite_module modules\/mod_rewrite.so/LoadModule rewrite_module modules\/mod_rewrite.so/g' /usr/local/apache2/conf/httpd.conf
RUN sed -i 's/AllowOverride None/AllowOverride All/g' /usr/local/apache2/conf/httpd.conf

COPY --from=builder /app/dist /usr/local/apache2/htdocs

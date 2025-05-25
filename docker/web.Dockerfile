# ---- Build Stage ----
FROM node:20-alpine AS builder

WORKDIR /app

# Copy package files and install dependencies (cached)
COPY apps/web/package*.json ./
RUN npm install

# Copy all source files
COPY apps/web ./

# Build the React app (outputs static files to /app/dist)
RUN npm run build

# ---- Serve Stage ----
FROM nginx:alpine AS runner

# Remove default nginx static assets
RUN rm -rf /usr/share/nginx/html/*

# Copy built React app from builder
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose port 80 (nginx default)
EXPOSE 80

# Start nginx in foreground
CMD ["nginx", "-g", "daemon off;"]

# Base image
FROM node

# Make folder to put our files in
RUN mkdir -p /usr/app-ais
RUN mkdir -p /usr/app-ais/frontend

# Set working directory so that all subsequent command runs in this folder
WORKDIR /usr/app-ais/frontend

# Copy package json and install dependencies
COPY package*.json ./
RUN npm install

# Copy our app
COPY . .

# Expose port to access server
EXPOSE 8081

# Command to run our app
CMD [ "npm", "start"]

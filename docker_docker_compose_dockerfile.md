
# Docker, Docker Compose, and Dockerfile - Key Concepts and Commands

## 1. Docker YAML Example

Docker allows you to package and run applications in isolated containers. It abstracts away the operating system and provides a simple and efficient way to run applications in any environment.

### Docker Command Example:

```bash
# Build a Docker image from a Dockerfile
docker build -t myapp:latest .

# Run a container from the image
docker run -d -p 8080:80 --name mycontainer myapp:latest

# List all running containers
docker ps

# Stop a running container
docker stop mycontainer
```

### Key Docker Commands:
- **docker build**: Build an image from a Dockerfile.
  - `-t myapp:latest`: Tag the image as `myapp:latest`.
- **docker run**: Run a container from a specified image.
  - `-d`: Run the container in detached mode (in the background).
  - `-p 8080:80`: Map port 8080 on the host to port 80 inside the container.
- **docker ps**: List all running containers.
- **docker stop**: Stop a running container.

---

## 2. Dockerfile Example

A **Dockerfile** is a script containing instructions on how to build a Docker image. It automates the process of building Docker images with specific configurations.

### Dockerfile Example:

```Dockerfile
# Use the official Node.js image as a base
FROM node:14

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and install dependencies
COPY package.json /app
RUN npm install

# Copy the rest of the application code
COPY . /app

# Expose port 3000
EXPOSE 3000

# Command to run the application
CMD ["npm", "start"]
```

### Key Components of Dockerfile:
- **FROM**: Specifies the base image to use for the container.
- **WORKDIR**: Sets the working directory inside the container.
- **COPY**: Copies files from the host into the container.
- **RUN**: Executes commands inside the container during image build.
- **EXPOSE**: Exposes a port to the outside world.
- **CMD**: Specifies the command to run when the container starts.

---

## 3. Docker Compose YAML Example

**Docker Compose** is a tool for defining and running multi-container Docker applications. You define the services that make up your app in a YAML file, and with a single command, you spin up all the services.

### Docker Compose Example:

```yaml
version: '3'
services:
  web:
    image: nginx:latest
    ports:
      - "8080:80"
    volumes:
      - ./data:/usr/share/nginx/html
  app:
    build: ./app
    ports:
      - "3000:3000"
    environment:
      - NODE_ENV=production
    depends_on:
      - db
  db:
    image: postgres:alpine
    environment:
      - POSTGRES_USER=user
      - POSTGRES_PASSWORD=password
      - POSTGRES_DB=mydb
```

### Key Components of Docker Compose YAML:
- **version**: Specifies the version of the Compose file format.
- **services**: Defines the list of services (containers) to be used.
  - **web**: A service that runs an Nginx container.
    - `image`: Specifies the image to use.
    - `ports`: Maps container ports to host ports.
    - `volumes`: Mounts local directories or files to container paths.
  - **app**: A service for the application.
    - `build`: Builds an image from the Dockerfile in the specified directory.
    - `environment`: Sets environment variables for the container.
    - `depends_on`: Specifies dependencies between services.
  - **db**: A service that runs a PostgreSQL container.
    - `image`: Specifies the image to use.
    - `environment`: Sets environment variables for database configuration.


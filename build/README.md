# Medicare Framework Service - Docker Setup

This guide explains how to **build**, **create**, and **run** a Docker container for the Medicare Framework Service application.  
The project uses a `Dockerfile` located in the `/build` directory.

---

## 🧱 1. Build Docker Image

Make sure you are in the project root directory.  
Run the following command to build the Docker image using the Dockerfile inside the `/build` folder:

```bash
docker build -t medicare-framework-service:latest -f build/Dockerfile .
```

## 🚀 2. Create Docker Container

Run the following command to create a Docker container using the image you built in the previous step:

```bash
docker run -d --name medicare-framework-service --env-file .env -p 8080:8080 medicare-framework-service:latest
```

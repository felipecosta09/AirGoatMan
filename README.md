# My Dockerized Web Application

This is a guide on how to build and run the frontend and backend services of my web application using Docker Compose.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

1. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/JustinDPerkins/AirGoatMan.git
    cd AirGoatMan
    ```

2. **Frontend**:

    - Navigate to the frontend directory:

        ```bash
        cd main-home-page
        ```

    - Build the frontend Docker image:

        ```bash
        docker build -t frontend_java:v1.0 .
        ```

3. **Backend**:

    - Navigate to the backend directory:

        ```bash
        cd ../backend
        ```

    - Build the backend Docker image:

        ```bash
        docker build -t go_backend:v1.0 .
        ```

4. Return to the project root directory:

    ```bash
    cd ..
    ```

5. Run the application using Docker Compose:

    ```bash
    docker-compose up
    ```

    This command will start both the frontend and backend services and connect them to a shared network.

6. Access the web application in your web browser:

    - Frontend: http://localhost:8080
    - Backend: Your backend API is now accessible via its respective endpoints.

## Stopping the Application

To stop the running containers and remove the associated resources, use the following command:

```bash
docker-compose down
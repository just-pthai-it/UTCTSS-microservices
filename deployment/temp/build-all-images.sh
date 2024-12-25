#!/bin/bash

# Set the base directory containing the services
SERVICES_DIR="./services"

# Loop through each subdirectory in the services directory
for SERVICE in "$SERVICES_DIR"/*; do
    if [ -d "$SERVICE" ]; then
        # Extract the name of the subdirectory (service name)
        SERVICE_NAME=$(basename "$SERVICE")

        # Run the docker build command with the service name as a build argument
        echo "Building Docker image for service: $SERVICE_NAME"
        docker build -f Dockerfile-development --build-arg service_name="$SERVICE_NAME" -t "pthaiit210501/utconnect-go-$SERVICE_NAME:dev-v1" .

        # Check if the build succeeded
        # shellcheck disable=SC2181
        if [ $? -ne 0 ]; then
            echo "Docker build failed for service: $SERVICE_NAME"
            exit 1
        fi
    fi
done

echo "Docker images built successfully for all services."

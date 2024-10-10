#!/bin/bash

# Set the root directory where the services are located
SERVICES_DIR="./services"

# Initialize an empty matrix array
MATRIX="["

# Loop through each subfolder in the 'services' directory
for SERVICE in $(ls -d ${SERVICES_DIR}/*/); do
  # Extract the folder name without the path
  SERVICE_NAME=$(basename $SERVICE)

  # Append to the matrix JSON object
  MATRIX+="{\"service-name\": \"$SERVICE_NAME\"},"
done

# Remove the last comma and close the JSON array
MATRIX=${MATRIX%,}
MATRIX+="]"

# Output the matrix in the format that GitHub Actions expects using Environment Files
echo "matrix=$MATRIX" >> $GITHUB_OUTPUT

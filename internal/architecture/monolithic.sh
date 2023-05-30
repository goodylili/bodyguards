#!/bin/bash

# Create the directory structure
mkdir -p cmd/app cmd/admin internal/api/handlers internal/api/middleware internal/api/models internal/api/repositories internal/api/services internal/config internal/database/migrations internal/logging internal/utils web/static/css web/static/js web/templates config migrations scripts

# Create empty placeholder files
touch cmd/app/main.go cmd/app/config.go cmd/admin/main.go cmd/admin/config.go internal/config/config.go internal/database/connection.go internal/logging/logger.go internal/utils/validation.go web/templates/base.html config/development.yaml config/production.yaml migrations/migration1.sql scripts/deploy.sh README.md

# Display success message
echo "Directory tree created successfully."

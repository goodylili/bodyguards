#!/bin/bash

# Create the main directory
mkdir myapp
cd myapp

# Create the app directory
mkdir app
cd app

# Create subdirectories inside the app directory
mkdir controllers
mkdir models
mkdir views
mkdir middlewares
mkdir routers
mkdir services
mkdir utils

# Create subdirectories inside the views directory
cd views
mkdir home
mkdir user

# Create subdirectories inside the static directory
cd ../..
mkdir static
cd static
mkdir css
mkdir js
mkdir img

# Create subdirectories inside the templates directory
cd ..
mkdir templates

# Return to the root directory
cd ../..

echo "Directory structure created successfully!"

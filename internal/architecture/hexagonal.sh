#!/bin/bash

# create top level directories
mkdir cmd internal pkg vendor
z
# create cmd/http directory
mkdir cmd/http

# create internal directories
mkdir internal/adapters internal/app internal/app/domain internal/app/ports internal/app/ports/input internal/app/ports/output internal/app/usecases

# create internal/adapters directories
mkdir internal/adapters/api internal/adapters/database

# create internal/app/domain directory

# create internal/app/ports directories
mkdir internal/app/ports/input internal/app/ports/output


# create vendor directories
mkdir vendor/module1 vendor/module2

# print success message
echo "Directory structure created successfully."
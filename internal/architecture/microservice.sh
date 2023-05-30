#!/bin/bash

# Create main project directory
mkdir my_microservice_app
cd my_microservice_app

# Create directories
mkdir cmd configs docs internal pkg local mocks proto pb test

# Create cmd/main.go file
touch cmd/main.go

# Create configs/config.go file
touch configs/config.go

# Create docs/api_docs.md and docs/design_docs.md files
touch docs/api_docs.md docs/design_docs.md

# Create internal/package1/file1.go and internal/package1/file2.go files
mkdir internal/package1
touch internal/package1/file1.go internal/package1/file2.go

# Create internal/package2/file3.go and internal/package2/file4.go files
mkdir internal/package2
touch internal/package2/file3.go internal/package2/file4.go

# Create pkg/package3/file5.go and pkg/package3/file6.go files
mkdir pkg/package3
touch pkg/package3/file5.go pkg/package3/file6.go

# Create pkg/package4/file7.go and pkg/package4/file8.go files
mkdir pkg/package4
touch pkg/package4/file7.go pkg/package4/file8.go

# Create local/docker-compose.yaml and local/local_env.sh files
mkdir local
touch local/docker-compose.yaml local/local_env.sh

# Create mocks/package1/mock_file1.go and mocks/package1/mock_file2.go files
mkdir mocks/package1
touch mocks/package1/mock_file1.go mocks/package1/mock_file2.go

# Create mocks/package2/mock_file3.go and mocks/package2/mock_file4.go files
mkdir mocks/package2
touch mocks/package2/mock_file3.go mocks/package2/mock_file4.go

# Create proto/service1.proto and proto/service2.proto files
mkdir proto
touch proto/service1.proto proto/service2.proto

# Create pb/service1.pb.go and pb/service2.pb.go files
mkdir pb
touch pb/service1.pb.go pb/service2.pb.go

# Create test/integration/test_file1.go and test/integration/test_file2.go files
mkdir test/integration
touch test/integration/test_file1.go test/integration/test_file2.go

# Create test/unit/test_file3.go and test/unit/test_file4.go files
mkdir test/unit
touch test/unit/test_file3.go test/unit/test_file4.go

# Create Dockerfile and .dockerignore files
touch Dockerfile .dockerignore

# Create .git directory
mkdir .git

# Create .gitignore file
touch .gitignore

# Create .golangci.yaml, mockery.yaml, Makefile, README.md, buf.gen.yaml, buf.work.yaml, tools.go files
touch .golangci.yaml mockery.yaml Makefile README.md buf.gen.yaml buf.work.yaml tools.go

# Create go.mod and go.sum files
go mod init my_microservice_app
go mod tidy

echo "Directory structure created successfully!"

# minio-golang

This repository provides examples and utilities for integrating MinIO object storage into Go applications. It includes sample code and configurations for managing storage buckets, uploading, downloading, and securing data in MinIO using Go.

## Quick Start

Clone the repository to get started with MinIO and Go:

```bash
git clone https://github.com/hajbabaeim/minio-golang.git
cd minio-golang

## Prerequisites

- Go 1.22.0+
- Docker (for running a local MinIO server)
- MinIO Client (mc) installed and configured
- Make

## Running a MinIO Server Locally

Use Docker to run a MinIO server for development:

```bash
docker pull minio/minio
docker run -p 9000:9000 -p 9001:9001 minio/minio server /data --console-address ":9001"
```

## MinIO Client Setup

You can manage your MinIO server using the provided Makefile commands:

- **Create mc alias**: Sets up a MinIO client alias for easy access.
  ```make
  make create-mc ENDPOINT=<your-minio-endpoint> ACCESS_KEY=<your-access-key> SECRET_KEY=<your-secret-key>
  ```

- **Add Policy**: Adds a predefined policy from `readwrite-perm.json`.
  ```make
  make add-policy ALIAS=<your-alias> POLICY_NAME=<your-policy-name>
  ```

- **Create User**: Creates a new user with specified access and secret keys.
  ```make
  make create-user ALIAS=<your-alias> USERNAME=<your-username> ACCESS_KEY=<your-access-key> SECRET_KEY=<your-secret-key>
  ```

- **Assign Policy**: Assigns the specified policy to a user.
  ```make
  make assign-policy ALIAS=<your-alias> POLICY_NAME=<your-policy-name> USERNAME=<your-username>
  ```

Each of these commands can be customized with your MinIO instance details.

## Examples

In the `examples` directory, you will find Go programs that demonstrate:

- Connecting to MinIO
- Uploading files
- Downloading files
- Listing buckets

Each example is self-contained and includes comments explaining the code.

## Usage

To run any example, use:

```bash
go run examples/upload.go
```

Replace `upload.go` with the filename of the example you want to execute.

## Contributing

Contributions to `minio-golang` are welcome! Please refer to the CONTRIBUTING.md file for guidelines on how to make contributions.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
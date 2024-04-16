# minio-golang

This repository provides examples and utilities for integrating MinIO object storage into Go applications. It includes sample code and configurations for managing storage buckets, uploading, downloading, and securing data in MinIO using Go.

## Quick Start

Clone the repository to get started with MinIO and Go:

```bash
git clone https://github.com/hajbabaeim/minio-golang.git
cd minio-golang
```

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

- **Create User**: Creates a new user with specified access and secret keys.
  ```make
  make create-user ALIAS=<your-alias> USERNAME=<your-username> SECRET_KEY=<your-secret-key>
  ```

- **Add Policy**: Adds a predefined policy from a json policy-file.
  ```make
  make add-policy ALIAS=<your-alias> POLICY_NAME=<your-policy-name> POLICY_FILE_PATH=<your-policy-file-path>
  ```

- **Assign Policy**: Assigns the specified policy to a user.
  ```make
  make assign-policy ALIAS=<your-alias> POLICY_NAME=<your-policy-name> USERNAME=<your-username>
  ```

- **Get User Information**: Retrieves information about a specific user.
  ```make
  make user-info ALIAS=<your-alias> USERNAME=<your-username>
  ```

- **Create Bucket**: Creates a new bucket in MinIO.
  ```make
  make create-bucket ALIAS=<your-alias> BUCKET_NAME=<bucket-name>
  ```

- **List Buckets**: Lists all buckets on the MinIO server.
  ```make
  make buckets-list ALIAS=<your-alias>
  ```

- **List Bucket Objects**: Lists all objects in a specified bucket, including versions.
  ```make
  make bucket-objects ALIAS=<your-alias> BUCKET_NAME=<bucket-name>
  ```

Each of these commands can be customized with your MinIO instance details.

## Examples

In the `test_app` directory, you will find Go programs that demonstrate:

- [x] Connecting to MinIO with HTTP
- [x] Connecting to MinIo with TLS
- [x] Uploading a file
- [ ] Downloading a file
- [ ] Listing buckets

Each example is self-contained and includes comments explaining the code.

## Usage

#### To run example through http, use:

```bash
go run test_app/main.go --mode="http" --certFile="<path to your public.cert file>"
```

#### To run example through TLS, use:

```bash
openssl genrsa -out private.key 2048
openssl req -new -x509 -key private.key -sha256 -days 3650 -out public.crt -subj "/CN=minio" -extensions v3_req -config <(printf "[req]\ndistinguished_name=req\n[req]\nreq_extensions=v3_req\n[v3_req]\nbasicConstraints=CA:TRUE\nkeyUsage=nonRepudiation, digitalSignature, keyEncipherment\nsubjectAltName=DNS:minio, IP:<SERVER_IP_ADDRESS>")
```

Copy both **private.key** & **public.crt** to **/root/.minio/cert**

```bash
go run test_app/main.go --mode="secure" --certFile="<path to your public.cert file>"
```


## Contributing

Contributions to `minio-golang` are welcome! Please refer to the CONTRIBUTING.md file for guidelines on how to make contributions.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
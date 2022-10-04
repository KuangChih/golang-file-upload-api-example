# Golang File Upload Api Example

This is an example for an uploading file-api in golang.

## Usage

### Setting the environment

Set the environment in `./configYml/application.yml`, specifies the `profiles.active` to develop, stage, or production

### Build code

```
go build main.go
```

### Run code in develop mode

```
go run main.go
```

### Api Spec

| HTTP | URI | Description |
| --- | --- | --- |
| GET | /v1/file | List all uploaded files. |
| POST | /v1/file | Upload file. |
| GET | /v1/file/:fileName | Download file. |
| DELETE | /v1/file/:fileName | Delete file. |

# email-microservice
gRPC based emailing service for microservice based architecture.

## Methods
### **SendEmail**

Send Standard emails with following parameters:
- `Recipients` *[Recipients](#recipients)* - List of recipients
- `Subject` *string* - Subject of the email
- `Body` *string* - Email body
- `ContentType` *string*: If `text/html` then HTML otherwise Plain text.
- `Attachments` *[Attachment](#attachment)* - Attachments to be sent with email.

[Example](./examples/Std.go)
### **SendRawEmail**

Send Raw [RFC822](https://www.w3.org/Protocols/rfc822/) based emails:
- `Recipients` *[string]* - List of recipients
- `Body` *Bytes* - Email body

[Example](./examples/Raw.go)

### **SendEmailWithTemplate**

Send Templated emails (templates should exist beforehand):
- `Recipients` *[Recipients](#recipients)* - List of recipients
- `Subject` *string* - Subject of the email
- `TemplateName` *string* - Name of the template
- `Attachments` *[Attachment](#attachment)* - Attachments to be sent with email.
- `TemplateParams` *map(string  -> string)* - Template data to be used in the template.

[Example](./examples/Template.go)

## Usage

### **As a service**

### *Server Environment Variables*
```
SMTP_HOST:                  SMTP host
SMTP_PORT:                  SMTP port
SMTP_SENDER:                SMTP user
SMTP_PASSWORD:              SMTP password
PORT:                       Port to listen on
PROMETHEUS_PORT:            Port to expose metrics on
```

### *Templates*
For templates to be used in the email, they should be stored in the `./templates` directory relative to the binary. Naming scheme for files is `<template_name>.html`, where `template_name` is the name of the template to be used by the client.

> For docker workdir in /app, hence, templates should be stored in `/app/templates`

### Docker
```
docker run -d --name email-microservice --env-file app.env -p 5555:5555 crossphoton/email-microservice
```

### Kubernetes

> TODO

### Locally
1. Clone the repository
2. After setting up the environment variables, run `go run main.go`


### **As a client**

> **See examples in the [examples](./examples/) directory.**

Generate the client code using the proto file [email.proto](./email.proto)

> In examples directory run `./gen_proto.sh`

# Prometheus

Port `9090` is exposed for Prometheus metrics. Can be changed using [environment variable](#server-environment-variables).

## License

GNU General Public License v3.0

## Authors

- [Aditya Agrawal (crossphoton)](https://github.com/crossphoton)


## Types

### Recipients
`Recipients`:
 - `To` *[string]* - Recipient email address
 - `Cc` *[string]* - Carbon copy email address
 - `Bcc` *[string]* - Blind carbon copy email address

> `Name <address>` formatting is supported.

### Attachment
`Attachment`:
  - base64data *string* - Base64 encoded data of attachment.
  - filename *string* - Name of attachment.

## Dependencies

- [simple-email-service](https://github.com/xhit/go-simple-mail)
- [gRPC](https://google.golang.org/grpc)
- [protobuf](https://google.golang.org/protobuf)
- [prometheus](https://github.com/prometheus/client_golang)
- [go-grpc-prometheus](https://github.com/grpc-ecosystem/go-grpc-prometheus)
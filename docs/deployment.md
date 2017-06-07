# Space

> A user management microservice; OAuth 2 provider

## Tips and general commands for deployment in the Google Cloud Platform

### Generating a 256-bit string in RFC 4648 standard base64

```sh
$ bin/gen-key
```

### Connecting to VM instance

```sh
$ gcloud compute ssh ${MACHINE_ID}
```

### Add the remote `production` to git project

```sh
$ git remote add production {REMOTE_USER}@{REMOTE_SERVER}:space
```

## License

[MIT License](http://earaujoassis.mit-license.org/) &copy; Ewerton Assis

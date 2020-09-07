## Commands for generate certificates

private certificate: 
```bash
openssl genrsa -out app.rsa 1024
```

Public certificate:
```bash
openssl rsa -in app.rsa -pubout > app.rsa.pub
```
Para generar el certificado privado : 
```bash
openssl genrsa -out app.rsa 1024
```

Para generar el certificado publico : 
```bash
openssl rsa -in app.rsa -pubout > app.rsa.pub```
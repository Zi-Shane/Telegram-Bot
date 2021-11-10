Telegram Bot
===

preparation
---

### With domain name

1. Generate the Self-Signed certificate and private key for your domain name.

```sh
openssl genrsa -out tls.key 2048

# .csr for domain
openssl req -nodes -new -key tls.key -subj "/CN=<your.domain.com>/" -out tls.csr

openssl x509 -req -sha256 -days 3650 -in tls.csr -signkey tls.key -out tls.crt
```

put `tls.crt` and `tls.key` into the folder `./secret/<DOMAINNAME>` 

2. Environment variables

```
PRODUCTION=true
PORT=8443
TELEGRAM_APITOKEN=<BotToken:xxxxxxxxxxxx>
DOMAINNAME=<your.domain.com>
```

---

### Without domain name

1. Start `ngrok`

```
./ngrok http 8443
```

2. Environment variables

```
PRODUCTION=false
PORT=8443
TELEGRAM_APITOKEN=<BotToken:xxxxxxxxxxxx>
DOMAINNAME=<xxxx.ngrok.io>
```

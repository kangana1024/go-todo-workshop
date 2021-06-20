# install utile

## Air

```bash
cd /tmp
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
air -v
```

## Install gofiber

```bash
go get -u github.com/gofiber/fiber/v2
```

# website

The base website template based on golang.

## Environment

```SHELL
export POSTGRES_HOST=127.0.0.1
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=postgres
```

## Docker Compose

Make the directory for volume first.

```SHELL
docker-compose up --build
```

## Lauch Postgres Database

```SHELL
docker run -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres
```

## Tencent Hub

```SHELL
docker build -t hub.tencentyun.com/semiscone/website .
docker login -u semiscone
docker push hub.tencentyun.com/semiscone/website
```

## Install Google Library

```SHELL
mkdir -p src/github.com/golang
mkdir -p src/golang.org
git clone https://github.com/golang/sys.git
git clone https://github.com/golang/net.git
git clone https://github.com/golang/text.git
git clone https://github.com/golang/lint.git
git clone https://github.com/golang/tools.git
git clone https://github.com/golang/crypto.git
ln -s src/github.com/golang/ src/golang.org/x
```


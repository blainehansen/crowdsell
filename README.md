```bash
# turns off microk8s
snap disable microk8s

# turns on microk8s
snap enable microk8s
```



# go-http-parser-bot

> A bot for parsing http

## Build Setup

``` bash
# make sure you have git secret
http://git-secret.io/

# pre-commit hooks
https://www.npmjs.com/package/pre-commit


# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run dev

# build for production with minification
npm run build

# build for production and view the bundle analyzer report
npm run build --report
```

For a detailed explanation on how things work, check out the [guide](http://vuejs-templates.github.io/webpack/) and [docs for vue-loader](http://vuejs.github.io/vue-loader).


## Go Stuff

https://github.com/cespare/reflex

```bash
sudo docker build -t go-anywhere .

go run $(ls -1 *.go | grep -v _test.go)
```

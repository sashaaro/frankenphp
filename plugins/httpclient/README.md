```shell
docker run --rm -v ${PWD}:/app -u $(id -u) -w /app composer install
```

```shell
docker run --rm -v $PWD/default.conf:/etc/nginx/conf.d/default.conf -p 180:80 -p 1444:444 --name=frankenphp-test1  yannrobert/docker-nginx
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' frankenphp-test1
```

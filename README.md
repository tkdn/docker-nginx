# quick start

```sh
docker compose up
```

# nginx reverse proxy

- http://localhost
  - -> go-app routing
- http://localhost/api/v1/path
  - -> go-app routing
- http://localhost/api/v2/foo/bar
  - -> node-app routing
- http://localhost/api/v1/content/path?name=foobar
  - -> redirect to http://localhost/api/v2/content/path?name=foobar

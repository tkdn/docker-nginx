# nginx redirect

```sh
docker compose up
```

- http://localshot
- http://localhost/unreachable/path -> redirect to http://localhost/redirected/path
- http://localhost/unreachable/path?name=foobar -> redirect to http://localhost/redirected/path?name=foobar
- http://localhost/unreachable/path2 -> redirect to http://localhost/redirected/path2

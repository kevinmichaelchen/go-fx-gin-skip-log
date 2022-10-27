Barebones reproduction of bug in Gin where disabling logs for certain endpoints
is not working.

```shell
# Gin will log this endpoint.
curl -XPOST http://localhost:8081/query

# Gin will NOT log this endpoint.
curl http://localhost:8081/health
```

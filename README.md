Barebones reproduction of bug in Gin where disabling logs for certain endpoints
is not working.

```shell
# Gin should log this endpoint. That is expected.
curl -XPOST http://localhost:8081/query

# Gin should not be logging this endpoint.
curl http://localhost:8081/health
```
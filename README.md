## ntfy-brdige
ntfy-bridge converts paylods to [ntfy](https://ntfy.sh) payload and sends the it to ntfy specified topic.

## Usage
the pattern in `/grafana/{topic}` and for custom ntfy servers you must pass `X-Endpoint` header:

```
https://bridge.service.ir/grafana/devops-28xg2nk
X-Endpoint: https://ntfy.service.ir
```

### Why another package
actually there are many solutions. but we wanted to use a different approach by getting topic in url (not args)


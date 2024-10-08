# Kamal Proxy Sidecar

A ridiculously simple sidecar logger for Kamal Proxy that forwards all logs as JSON to an endpoint of your choice, set by the SIDECAR_ENDPOINT environment variable.

Just add a section like the below to your Kamal deploy.yml:

```yml
accessories:
  sidecar:
    image: carldaws/kamal-proxy-sidecar
    hosts:
      - 192.168.0.1
    env:
      clear:
        SIDECAR_ENDPOINT: http://my-sidecar-endpoint.com/api
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
```

## Development & Testing

Clone the repository and use the `docker-compose.yml` file in the `/test` directory to launch a demo app, Kamal Proxy, a demo endpoint app for handling logs and the sidecar:

```bash
docker compose up --build
```

Don't forget to route Kamal Proxy traffic to the app:

```bash
docker compose exec proxy kamal-proxy deploy test-app --target app:8080
```
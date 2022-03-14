# Go Hasher
🔒 Simple HTTP service for data hashing

## Deploying
First of all, docker must be installed to deploy the project:

```bash
sudo curl https://get.docker.com/ | bash
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```

After that, all you have to do is to run:

```bash
docker-compose up -d
```

## Documentation
Documentation to API is available at `/api/docs`.

## Demo
Demo instance is available at https://hasher.exp101t.me/.

# Logging Service

Will handle all of the logging for each of the services.

Idea:
  Listen to kafka and insert records into postgres db

```
docker run -d -p 8081:8081 --name logger github.com/jmnelson12/distributed-world/logger:tag
```
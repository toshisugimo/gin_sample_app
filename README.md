# Gin + Gorm Sample web app

## Start app
```
docker-compose up
```

## Migrate DB
```
docker exec -it gin_sample_app go run /go/src/app/migrate/migrate.go
```

## Connect DB

```
docker exec -it gin_sample_db mysql gin_sample_app

mysql> show tables;
+--------------------------+
| Tables_in_gin_sample_app |
+--------------------------+
| todos                    |
+--------------------------+
1 row in set (0.01 sec)
```

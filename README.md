# Hotel reservation backend 
This is a booking.com clone like backend

## Project outline
- users -> book a room from a hotel
- admins -> going to check reservations/bookings
- auth -> JWT tokens
- hotels -> CRUD API -> JSON
- rooms -> CRUD API -> JSON
- scripts -> database management -> seeding, migration 


## Resources
### gofiber
Documentation
```
https://gofiber.io
```

Installing gofiber
```
go get github.com/gofiber/fiber/v2
```

### Task
Documentation
```
https://taskfile.dev/
```

Installing Task
```
yay -S go-task-bin
```
or
```
https://taskfile.dev/installation/
```

### Mongodb
Documentation
```
https://mongodb.com/docs/drivers/go/current/quick-start
```

Installing mongodb client
```
go get go.mongodb.org/mongo-driver/mongo
```
and running mongodb in a docker container
```
docker run -d -p 27017:27017 --name mongodb-booking mongo:latest
```
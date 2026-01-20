
#  Distributed Health Monitor System

A scalable distributed system built with GoLang to monitor service health in real-time.  
The system uses RabbitMQ, WebSockets and PostgreSQL to perform scheduled health checks, process workers asynchronously, and notify clients instantly when service status changes.


# Features

- Distributed worker-based health check processing
- Real-time status updates using WebSocket
- Scheduler-based periodic service monitoring
- Message queue with RabbitMQ
- Persistent health logs in PostgreSQL
- State change detection
- Unit & integration test support

# System Architecture

Scheduler  -> RabbitMQ Queue -> worker 


# Technologies
-------------------
Go (Golang) : Backend services
RabbitMQ :  Message queue
PostgreSQL : Database
WebSocket : Real-time notifications
GORM : ORM
Docker : Containerization
REST API : Service management

# Run dependancies

Make sure Docker is running:  docker-compose up -d
this will start postgres - RabbitMQ

# Configuration
Create a .env file in the root directory and add the following variables:
- `DB_HOST`: localhost
- `DB_USER`: user
- `DB_PASSWORD`: password
- `DB_NAME`: health_monitor
- `RABBITMQ_URL`: amqp://guest:guest@localhost:5672/

# Run Unit Test 
go test ./... -v



# Run application 
go run cmd/main.go


# API EndPoint 

The system provides a RESTful API to manage the services you want to monitor. By default, the API runs on `http://localhost:8088`.

1- Register new service ( Register monitored services)
  Adds a new service to the monitoring queue.

- URL: http://localhost:8088/services
- Method: POST

{
    "name": "test register ",
    "url": "https://www.google.com",
    "interval": 30,
     "Timeout": 20
}

RESPONSE

{
  "ID": 12,
  "Name": "test register ",
  "URL": "https://www.google.com",
  "Interval": 30,
  "LastStatus": "",
  "CreatedAt": "2026-01-19T23:28:41.5044326+02:00",
  "LastCheck": "0001-01-01T00:00:00Z",
  "Timeout": 20,
  "HealthLogs": null
}

 


2- Get All Services ( list monitored services)
Returns a list of all registered services and their current status.
- URL: http://localhost:8088/services
- Method: GET


RESPONSE 

[
  {
    "id": 7,
    "name": "NonExistent",
    "url": "http://nonexistent.example.com",
    "interval": 90,
    "last_status": "DOWN",
    "created_at": "2026-01-19T16:36:10.986295+02:00",
    "last_state": "2026-01-20T03:30:41.762837+02:00",
    "timeout": 30
  },
]


3- Get Service Logs (Query historical health data)

- URL: http://localhost:8088/services/1/logs
- Method: GET

RESPONSE : 
[{
    "ID": 426,
    "ServiceID": 8,
    "Status": "DOWN",
    "LatencyMs": 5000,
    "CheckedAt": "2026-01-19T23:27:37.359725+02:00",
    "Service": {
      "ID": 0,
      "Name": "",
      "URL": "",
      "Interval": 0,
      "LastStatus": "",
      "CreatedAt": "0001-01-01T00:00:00Z",
      "LastCheck": "0001-01-01T00:00:00Z",
      "Timeout": 0,
      "HealthLogs": null
    }
  },
  ]


  # Real time update 

  The system pushes live health status updates to connected clients whenever a service check is performed or a status change is detected.

  open HOPPSCOTCH OR POSTMAN => realtime => websocket
  
  Connect to this endpoint to receive real-time JSON updates.
  
  URL:  ws://localhost:8088/ws

  RESPONSE 
  
{
  "latency": 5000,
  "name": "DelayLinkToTest",
  "new_status": "DOWN",
  "old_status": "UP",
  "service_id": 8
}

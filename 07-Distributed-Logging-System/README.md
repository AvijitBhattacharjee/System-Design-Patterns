


                           ┌────────────────────┐
                           │  Service A         │
                           └─────────┬──────────┘
                                     │

                           ┌────────────────────┐
                           │  Service B         │
                           └─────────┬──────────┘
                                     │

                           ┌────────────────────┐
                           │  Service C         │
                           └─────────┬──────────┘
                                     │

                                     ▼

                     ┌─────────────────────────────┐
                     │      Logger SDK             │
                     └─────────────┬───────────────┘
                                   │
                                   ▼

                     ┌─────────────────────────────┐
                     │         Load Balancer       │
                     └─────────────┬───────────────┘
                                   │

                ┌──────────────────┼──────────────────┐
                │                  │                  │

                ▼                  ▼                  ▼

      ┌────────────────┐ ┌────────────────┐ ┌────────────────┐
      │ Collector #1   │ │ Collector #2   │ │ Collector #3   │
      └───────┬────────┘ └───────┬────────┘ └───────┬────────┘
              │                  │                  │
              └──────────────────┼──────────────────┘
                                 │

                                 ▼

                 ┌────────────────────────────┐
                 │       Kafka Cluster         │
                 └─────────────┬──────────────┘
                               │

         ┌─────────────────────┼─────────────────────┐
         │                     │                     │

         ▼                     ▼                     ▼

 ┌───────────────┐   ┌─────────────────┐   ┌─────────────────┐
 │ OpenSearch    │   │ Alert Consumer  │   │ Metrics Consumer│
 │ Consumer      │   │                 │   │                 │
 └───────┬───────┘   └────────┬────────┘   └────────┬────────┘
         │                    │                     │

         ▼                    ▼                     ▼

 ┌───────────────┐   ┌─────────────────┐   ┌─────────────────┐
 │ OpenSearch    │   │ PagerDuty       │   │ Prometheus      │
 │ Cluster       │   │ Slack           │   │ Grafana         │
 └───────┬───────┘   └─────────────────┘   └─────────────────┘
         │

         ▼

 ┌────────────────────┐
 │ Query Service      │
 └─────────┬──────────┘
           │

           ▼

 ┌────────────────────┐
 │ UI Dashboard       │
 └─────────┬──────────┘
           │

           ▼

 ┌────────────────────┐
 │ S3 Cold Storage    │
 └────────────────────┘





Project Structure
distributed-logger/

├── cmd/
│   └── main.go
│
├── internal/
│
│   ├── logger/
│   │   ├── logger.go
│   │   ├── level.go
│   │   ├── entry.go
│   │   └── singleton.go



This is the version I'd memorize for interviews.

Logger
 ├── Info(msg)
 │    → Create INFO log entry and push to async queue
 │
 ├── Error(msg)
 │    → Create ERROR log entry and push to async queue
 │
 └── Log(level, msg)
      → Generic method used by all log levels
Appender
 └── Write(logEntry)
      → Persist/send log to destination
         (Console/File/Kafka)
Formatter
 └── Format(logEntry)
      → Convert LogEntry into JSON/Text format

Example:

{
  "service":"payment",
  "level":"ERROR",
  "message":"db timeout"
}
Transport
 └── Send(payload)
      → Send formatted log batch to Collector API
         using HTTP/gRPC

Example:

POST /v1/logs
Batcher
 ├── Add(logEntry)
 │    → Add log into in-memory buffer
 │
 └── Flush()
      → Send accumulated logs as one batch

Example:

1000 logs
   ↓
1 HTTP request
CollectorService
 └── Ingest(logRequest)
      → Validate request and publish log to Kafka

Responsibilities:

Validate
Authenticate
Rate Limit
Publish
KafkaPublisher
 └── Publish(logEntry)
      → Write log event into Kafka topic

Example:

logs-topic

Partition Key:

traceId
serviceId
KafkaConsumer
 └── Consume()
      → Read logs from Kafka for processing

Example:

Kafka
  ↓
Consumer
  ↓
OpenSearch
Repository
 ├── Save(logEntry)
 │    → Store log in OpenSearch
 │
 └── Search(filter)
      → Retrieve logs matching criteria

Example:

service=payment
level=ERROR
QueryService
 ├── SearchLogs(filter)
 │    → Search logs by service, level, time range
 │
 └── GetTrace(traceId)
      → Fetch all logs belonging to a distributed request

Example:

traceId = abc123

Order Service
Payment Service
Inventory Service

Returns complete request flow.

Complete Business Flow
Info()
Error()
Log()
    │
    ▼

Write()
    │
    ▼

Format()
    │
    ▼

Add()
Flush()
    │
    ▼

Send()
    │
    ▼

Ingest()
    │
    ▼

Publish()
    │
    ▼

Kafka
    │
    ▼

Consume()
    │
    ▼

Save()
    │
    ▼

OpenSearch
    │
    ▼

SearchLogs()
GetTrace()
    │
    ▼

UI
One-line explanation for every component
Component	Responsibility
Logger	Create log events
Appender	Decide where logs go
Formatter	Convert logs into JSON/Text
Batcher	Reduce network calls
Transport	Send logs to collector
CollectorService	Entry point of logging platform
KafkaPublisher	Put logs into Kafka
KafkaConsumer	Read logs from Kafka
Repository	Store and query logs
QueryService	Serve logs to UI
Interview questions - 



1. What if Kafka goes down?

Kafka Replication
Partition-1

Leader  -> Broker1
Replica -> Broker2
Replica -> Broker3

If Broker1 dies:

Broker2 becomes leader


2. If consumer crashes, how to recover?

Kafka Consumer Offset Management. 



3. How to avoid duplicate logs?

Idempotency - LogID UUID.


4. How do you balance load?

load balancer in front of collectors. Each collector is stateless and can be scaled horizontally. Use consistent hashing to route logs to collectors based on LogID or source service ID.

Kafka partitioning - each collector can write to multiple partitions in Kafka, and Kafka will handle distributing the load across brokers.


5. How do you preserve ordering?

Kafka guarantees ordering within a partition. To preserve ordering, logs from the same source service should be sent to the same partition. This can be achieved by using a consistent hashing function on the LogID or source service ID to determine the partition.


6. What if OpenSearch goes down?

OpenSearch Replication and Sharding. OpenSearch can be configured with multiple nodes and replicas to ensure high availability. If a node goes down, the cluster will automatically promote a replica to become the new primary for the affected shard.

7. How do you scale the system?

Collectors can be scaled horizontally by adding more instances behind the load balancer. Kafka can be scaled by adding more brokers and partitions. OpenSearch can be scaled by adding more nodes and shards. The system should be designed to handle increased load by distributing logs across multiple collectors, Kafka partitions, and OpenSearch shards. 

8. Why not write directly to OpenSearch?

OpenSearch is not designed to handle high write throughput directly from multiple sources. Writing directly to OpenSearch can lead to performance bottlenecks and increased latency. By using Kafka as a buffer, we can decouple the log ingestion from the storage layer, allowing for better scalability and fault tolerance. Additionally, Kafka provides features like message retention and replay, which can be useful for recovering from failures or reprocessing logs if needed.


9. How do you search efficiently?

Indexing and Sharding - OpenSearch uses inverted indexes to allow for fast full-text searches. Logs can be indexed based on relevant fields (e.g., timestamp, log level, service name) to enable efficient querying. Sharding allows the data to be distributed across multiple nodes, improving search performance and scalability. Additionally, caching frequently accessed queries can further enhance search efficiency.


10. How do you archive old logs?

Hot and Cold Storage - Logs can be stored in a hot storage (OpenSearch) for recent logs that require fast access and search capabilities. Older logs can be archived to cold storage (e.g., S3, Glacier) for long-term retention. This can be achieved by implementing a log retention policy that moves logs from hot storage to cold storage after a certain period of time or based on specific criteria (e.g., log age, log level).

11. How would you handle a noisy service?

Rate Limiting and Sampling - If a particular service is generating an excessive amount of logs, it can overwhelm the logging system. To handle this, we can implement rate limiting to restrict the number of logs ingested from that service within a specific time frame. Additionally, we can use sampling techniques to only log a subset of the messages from that service, ensuring that we still capture relevant information without overloading the system.

12. How would you correlate logs across microservices?

TraceID and SpanID - To correlate logs across different microservices, we can use a unique TraceID that is generated at the start of a request and passed along to all services involved in processing that request. Each service can also generate a SpanID for individual operations within the request. By including the TraceID and SpanID in the log entries, we can trace the flow of a request across multiple services and identify any issues or bottlenecks in the system.


13. What if Query API becomes slow?

Pagination and Caching - If the Query API becomes slow due to a large number of logs being returned, we can implement pagination to limit the number of results returned in a single query. This allows clients to request logs in smaller chunks, improving response times. Additionally, we can implement caching for frequently accessed queries to reduce the load on the OpenSearch cluster and improve query performance.


14. How do you secure the system?

API Key and JWT Authentication - To secure the logging system, we can implement API key-based authentication for clients that are sending logs to the collectors. Each client can be assigned a unique API key that must be included in the log request headers. Additionally, we can use JWT (JSON Web Token) authentication for clients that need to query logs from the system. This ensures that only authorized users can access the logs and perform queries.




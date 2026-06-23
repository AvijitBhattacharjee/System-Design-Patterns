1. Functional Requirements
Producer can publish messages
Consumer can consume messages
Messages delivered in order
Multiple consumers supported
Acknowledgement support
Retry failed messages
Dead Letter Queue (DLQ)

2. Non Functional Requirements
High throughput
Low latency
Durable storage
Scalable
Fault tolerant

Interview Questions 


1. Why do we need Partitions?
Without partitions, all messages would be sent to a single broker, which can become a bottleneck. Partitions allow for parallel processing of messages across multiple brokers, improving throughput and scalability. Each partition can be consumed independently by different consumers, enabling better load distribution.

2. What is an Offset?

Offset is a unique identifier for each message within a partition. It represents the position of the message in the partition and is used by consumers to track which messages have been consumed. Consumers can commit offsets to keep track of their progress, allowing them to resume consumption from the last committed offset in case of a crash or restart.

3. What is a Broker?
Broker is a server that stores and manages messages in a message queue system. It receives messages from producers, stores them in partitions, and serves them to consumers. Brokers handle message persistence, replication, and delivery guarantees.

4. What is partitions?
Partitions are a way to divide a topic into multiple segments, allowing for parallel processing of messages. Each partition is an ordered, immutable sequence of messages that is continually appended to. Partitions enable scalability and fault tolerance by allowing multiple consumers to read from different partitions simultaneously, and they also allow for message ordering within a partition.

5. What is a Topic?
A topic is a logical channel to which producers send messages and from which consumers receive messages. Topics can have multiple partitions, and each partition can be consumed independently by different consumers. Topics provide a way to organize messages based on their content or purpose, allowing for better message routing and filtering.

6. What is a Zookeeper?
Zookeeper is a centralized service for maintaining configuration information, naming, providing distributed synchronization, and providing group services. In the context of Kafka, Zookeeper is used to manage and coordinate the Kafka brokers, maintain metadata about topics and partitions, and handle leader election for partitions. It ensures that the Kafka cluster remains consistent and available.

7. How is a partition selected?
When a producer sends a message to a topic, it can specify a partition key (e.g., LogID or source service ID). Kafka uses a consistent hashing function on the partition key to determine which partition the message should be sent to. This ensures that messages with the same partition key are always sent to the same partition, preserving ordering for those messages.

8. Can Kafka guarantee global ordering?
No, Kafka can only guarantee ordering within a partition. If messages are sent to different partitions, there is no guarantee of global ordering across those partitions. To preserve ordering for related messages, they should be sent to the same partition using a consistent partition key.

9. What happens if Leader crashes?
If the leader of a partition crashes, Kafka will automatically promote one of the replicas to become the new leader for that partition. This process is handled by Zookeeper, which monitors the health of the brokers and manages leader election. The new leader will take over responsibility for handling all read and write requests for that partition, ensuring continued availability and durability of messages.

10. What happens if a consumer crashes?
If a consumer crashes, it can recover by resuming consumption from the last committed offset. When the consumer restarts, it will read the last committed offset for each partition it was consuming from and continue processing messages from that point onward. This allows for fault tolerance and ensures that messages are not lost due to consumer failures.

11. What happens if a producer crashes?
If a producer crashes while sending a message, the message may not be successfully sent to the broker. However, Kafka provides durability guarantees, so if the producer receives an acknowledgment from the broker that the message has been received and replicated, it can be confident that the message will not be lost even if the producer crashes afterward. If the producer does not receive an acknowledgment, it can retry sending the message until it succeeds or reaches a retry limit.

12. How do you handle message retries?
Message retries can be handled by the consumer. If a consumer fails to process a message successfully, it can choose to not commit the offset for that message, allowing it to be re-delivered by Kafka. Alternatively, the consumer can implement a retry mechanism where it attempts to process the message a certain number of times before giving up and sending it to a Dead Letter Queue (DLQ) for further analysis.

13. What is a Dead Letter Queue (DLQ)?
A Dead Letter Queue (DLQ) is a special queue where messages that cannot be processed successfully after a certain number of retries are sent. This allows for further analysis and debugging of problematic messages without affecting the main processing flow. Consumers can monitor the DLQ to identify and address issues with specific messages or patterns of failures.

14. What happens if 4 consumers and 2 partitions?
Then, two consumers will be assigned to each partition. Kafka uses a consumer group mechanism to distribute partitions among consumers. In this case, each partition will be consumed by two consumers, allowing for parallel processing of messages. However, since there are more consumers than partitions, some consumers may remain idle if there are no additional partitions available for them to consume from.

15. How would you handle 1 Million messages/sec?
To handle 1 million messages per second, you would need to consider several factors:
- **Partitioning**: Increase the number of partitions to allow for parallel processing and better load distribution across consumers.
- **Scaling**: Scale the number of brokers and consumers to handle the increased load. Use horizontal scaling to add more resources as needed.
- **Batching**: Implement batching of messages to reduce the overhead of individual message processing and improve throughput.
- **Compression**: Use message compression to reduce the size of messages being sent and received, which can help improve network efficiency and reduce latency.
- **Monitoring and Tuning**: Continuously monitor the system's performance and tune configurations such as buffer sizes, timeouts, and retry policies to optimize throughput and latency.

16. What limits Kafka throughput?
Several factors can limit Kafka throughput, including:
- **Network Bandwidth**: The available network bandwidth can limit the rate at which messages can be sent and received between producers, brokers, and consumers.
- **Disk I/O**: The speed of the underlying storage system can affect how quickly messages can be written to and read from disk. Slow disk I/O can become a bottleneck for high-throughput scenarios.
- **Broker Configuration**: Kafka broker configurations, such as the number of partitions, replication factor, and batch sizes, can impact throughput. Improperly configured brokers may not be able to handle high message rates effectively.

17. Design Kafka from Scratch

Expected components:
Producer
Broker
Partition
Offset
Consumer
Replication
Leader Election

18. How would you prevent duplicate messages?
To prevent duplicate messages in a Kafka-like message queue system, you can implement the following strategies:
- **Idempotent Producers**: Ensure that producers are idempotent, meaning that sending the same message multiple times will have the same effect as sending it once. This can be achieved by assigning a unique identifier (e.g., UUID) to each message and having the broker recognize and discard duplicates based on this identifier.
- **Consumer Deduplication**: Consumers can maintain a record of processed message IDs to avoid processing the same message multiple times. This can be done using a local cache or a distributed store like Redis to track processed message IDs.
- **Acknowledgment Mechanism**: Implement an acknowledgment mechanism where consumers acknowledge the successful processing of messages. If a consumer fails to acknowledge a message, the broker can re-deliver it, but the consumer can check if it has already processed that message based on its unique identifier.

19. How would you handle a noisy producer?
To handle a noisy producer that generates an excessive number of messages, you can implement the following strategies:
- **Rate Limiting**: Implement rate limiting on the producer side to control the number of messages sent to the broker within a specific time frame. This can help prevent overwhelming the system and ensure fair usage among multiple producers.
- **Backpressure Mechanism**: Introduce a backpressure mechanism where the broker can signal to the producer to slow down or pause message production when the system is under heavy load. This can help prevent message loss and maintain system stability.
- **Message Filtering**: Apply filtering logic on the producer side to discard unnecessary or low-priority messages before they are sent to the broker. This can help reduce the overall message volume and focus on processing only relevant messages.

20. Explain ACK=0, ACK=1, ACK=ALL

- **ACK=0**: The producer does not wait for any acknowledgment from the broker. It sends the message and moves on, providing the highest throughput but with no guarantee of delivery.
- **ACK=1**: The producer waits for an acknowledgment from the leader broker. This ensures that the message is written to the leader's log but does not guarantee that it has been replicated to all in-sync replicas.
- **ACK=ALL**: The producer waits for an acknowledgment from all in-sync replicas. This provides the strongest durability guarantee, ensuring that the message is replicated across all available brokers in the cluster.

21. How does hash(key)%N work?
The hash(key)%N operation is used to distribute messages across partitions in a Kafka topic. Here's how it works: 
- The key of each message is hashed using a hash function.
- The resulting hash value is then taken modulo N (the number of partitions).
- This ensures that messages with the same key are consistently routed to the same partition, while messages with different keys are distributed across different partitions.

22. How does leader-follower replication work?
Leader-follower replication in Kafka works as follows:
- One broker is designated as the leader for a particular partition, while others are followers.
- All writes to the partition are directed to the leader broker, which appends the message to its log.
- The leader then replicates the message to all follower brokers in the cluster.
- If the leader fails, one of the followers is automatically elected as the new leader, ensuring high availability and fault tolerance.

23. How do consumer groups achieve parallelism?
Consumer groups achieve parallelism by allowing multiple consumers to read from the same topic simultaneously. Each consumer in the group is assigned a subset of the partitions for the topic, enabling them to process messages in parallel. This allows for better load distribution and increased throughput, as multiple consumers can work on different partitions concurrently. If a consumer fails, the remaining consumers in the group can take over its partitions, ensuring continued processing of messages.

24. Difference between At-most-once, At-least-once, Exactly-once?
At-most-once: Messages are delivered at most once, meaning that they may be lost but will never be duplicated. This is the fastest delivery guarantee but does not ensure reliability.
At-least-once: Messages are delivered at least once, meaning that they may be duplicated but will never be lost. This provides a balance between reliability and performance, but consumers must handle potential duplicates.
Exactly-once: Messages are delivered exactly once, meaning that they are neither lost nor duplicated. This is the most reliable delivery guarantee but may introduce additional complexity and overhead in the system to ensure idempotency and deduplication.

25. Difference between AWS SQS and Kafka and RabbitMQ?
- **AWS SQS**: A fully managed message queuing service that provides a simple and scalable way to decouple and coordinate microservices, distributed systems, and serverless applications. It offers at-least-once delivery and supports message retention, but does not provide strong ordering guarantees or exactly-once semantics.
- **Kafka**: A distributed streaming platform that provides high-throughput, low-latency message delivery with strong ordering guarantees within partitions. It supports at-least-once and exactly-once semantics, making it suitable for real-time data processing and event-driven architectures.
- **RabbitMQ**: A message broker that implements the Advanced Message Queuing Protocol (AMQP). It provides flexible routing, message acknowledgment, and support for various messaging patterns. RabbitMQ offers at-least-once delivery and can be configured for high availability, but it may not achieve the same throughput as Kafka for large-scale streaming use cases.
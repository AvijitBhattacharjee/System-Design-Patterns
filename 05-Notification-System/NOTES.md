
Q: Why use a message queue instead of calling providers directly?
A: Direct calls couple your API latency to provider latency — if Twilio is slow, your whole /send endpoint slows down. A queue decouples ingestion from delivery, absorbs traffic bursts, and gives you retry semantics. If a worker crashes mid-delivery, Kafka replays the message.
Q: How do you handle retries without sending a notification twice?
A: Two layers. First, the dedup cache (SetNX with a 24h TTL) prevents re-enqueuing the same logical event. Second, workers use idempotency keys when calling providers (e.g. SendGrid's X-Message-Id header). So even if a worker processes a message twice, the provider side-steps the duplicate.
Q: How do you scale to millions of notifications per minute?
A: Partition Kafka topics by userID — this keeps per-user ordering while letting you add partitions horizontally. Scale worker pods independently per channel (email is slower than push, so you might run 10× more email workers). The Notification Service itself is stateless, so it scales behind a load balancer.
Q: How do you prioritize urgent notifications (e.g. 2FA OTPs) over marketing blasts?
A: Use separate Kafka topics or priority queues — notifications.sms.critical vs notifications.sms.bulk. Workers for critical topics have dedicated pods and higher throughput SLAs.
Q: What happens if the DB is down when a notification is sent?
A: Accept the request anyway — the queue is the source of truth. Workers write to the DB asynchronously after delivery. You trade a brief inconsistency in logs for higher availability at the write path.

=======================================================================================================

4. Corner cases
Cache (Redis for dedup + rate limiting)
Use SET key 1 NX EX 86400 for idempotency. Rate limiting uses a sliding window counter per (callerID, minute). If Redis is unavailable, fail open for dedup (accept the risk of a duplicate) but fail closed for rate limiting (reject the request) to protect downstream providers from being hammered.
Rate limiting
Apply at two levels: the API gateway limits requests per API key (e.g. 1000/min), and per-user limits prevent spamming a single recipient. Token bucket or sliding window log algorithms work well. Store counters in Redis with TTL equal to the window size.
Auth for sending
API keys (for server-to-server callers) scoped to a tenant, validated at the gateway. JWTs (for user-facing clients) verified by the Auth Service. Never let raw notification payloads bypass auth — always verify the caller has permission to notify the target userID.
Message queue design
Per-channel topics let you independently scale and throttle each channel. Use a dead-letter queue (DLQ) for messages that fail after N retries — wire the DLQ to an alert and a manual review process. Set message TTLs so stale notifications (e.g. a 5-minute OTP sent 2 hours ago) are dropped instead of delivered.
Worker nodes
Workers should be stateless and horizontally scalable. Use consumer groups so multiple worker instances share load. Implement graceful shutdown — finish in-flight messages before killing the pod. Add circuit breakers (e.g. using the gobreaker library) around provider calls so a flaky SMS provider doesn't block the whole worker.
Database
A notifications table stores (id, user_id, type, channel, status, sent_at, payload). Index on (user_id, created_at) for "fetch my notification history" queries. Use soft deletes — never hard-delete notification logs; they're audit trails. For very high volume, shard by user_id or use a time-series-friendly store like Cassandra.
Handling duplicity
Three mechanisms work together: (1) the Redis dedup key at ingestion, (2) idempotency keys passed to each provider API, and (3) the status column in the DB checked before re-attempting. This gives exactly-once delivery semantics end-to-end even under worker crashes and retries.

===============================================================================================
Rate Limiter is one of the most frequently asked backend design topics in Senior SWE, API Gateway, Kubernetes Controller, Distributed Systems, and Microservices interviews.

1. Design a Rate Limiter for an API Gateway that can handle millions of requests per second.
Client
  |
  v
Rate Limiter
  |
  +--> Extract User/IP/API Key
  |
  +--> Find Rate Limit Rule
  |
  +--> Fetch Current Count
  |
  +--> Compare Against Limit
           |
      +----+----+
      |         |
   Allow      Reject
      |         |
      v         v
 Backend      429
 Service




2. func RateLimiterMiddleware(next http.Handler) http.Handler

Everything is inside one process.

+--------------------------------+
| RateLimiter Middleware         |
|                                |
| Extract UserID                 |
| Load Rule                      |
| Check Counter                  |
| Decision                       |
+--------------------------------+

No separate services.
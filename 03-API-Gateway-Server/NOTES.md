
1. Why do we need API Gateway?
   - To provide a single entry point for all client requests
   - To handle cross-cutting concerns like authentication, authorization, and rate limiting
   - To provide load balancing and failover capabilities
   - To enable API management features like monitoring, logging, and analytics  


2. How does API Gateway work?
   - API Gateway receives client requests and routes them to the appropriate backend services
   - It can perform request validation, authentication, and authorization before forwarding the request
   - It can also perform response transformation and aggregation before sending the response back to the client

3. API Gateway vs Reverse Proxy?
   - Both API Gateway and Reverse Proxy act as intermediaries between clients and backend services
   - Reverse Proxy primarily focuses on load balancing and caching, while API Gateway provides additional features like authentication, authorization, and request/response transformation
   - API Gateway is more suitable for microservices architecture, while Reverse Proxy is commonly used in monolithic applications

4. How does request routing work?
   - API Gateway uses routing rules to determine which backend service should handle a particular request
   - Routing can be based on URL path, HTTP method, headers, or query parameters
   - API Gateway can also perform load balancing across multiple instances of a backend service

5. How does API Gateway handle authentication and authorization?
   - API Gateway can integrate with identity providers to authenticate clients using tokens (e.g., JWT, OAuth)
   - It can enforce authorization policies based on user roles or permissions
   - API Gateway can also support API key-based authentication for server-to-server communication

6. How does API Gateway handle rate limiting and throttling?
   - API Gateway can enforce rate limits on a per-client or per-API basis
   - It can use token bucket or leaky bucket algorithms to control the rate of incoming requests
   - Throttling can be applied to prevent abuse and ensure fair usage of backend services

7. How does API Gateway handle caching?
   - API Gateway can cache responses from backend services to reduce latency and improve performance
   - It can use in-memory caching or integrate with external caching systems like Redis or Memcached
   - Cache invalidation strategies can be implemented to ensure data consistency

8. How does API Gateway handle logging and monitoring?
   - API Gateway can log incoming requests and outgoing responses for auditing and debugging purposes
   - It can integrate with monitoring tools to track metrics like request count, latency, and error rates
   - Alerts can be configured to notify administrators of any issues or anomalies in the system

9. How does API Gateway handle failover and retries?
   - API Gateway can implement retry policies for failed requests to backend services
   - It can use exponential backoff and jitter to avoid overwhelming the backend during failures
   - Failover mechanisms can be implemented to route requests to healthy instances of backend services

10. How does API Gateway handle versioning and backward compatibility?
   - API Gateway can support multiple versions of an API by routing requests to the appropriate backend service based on the version specified in the request
   - It can provide backward compatibility by allowing older clients to continue using previous versions of the API while new clients use the latest version
   - Deprecation policies can be implemented to phase out older versions over time


11. How does API Gateway handle security and compliance?
   - API Gateway can enforce security policies like HTTPS, CORS, and input validation to protect against common web vulnerabilities
   - It can integrate with security tools to perform threat detection and prevention
   - Compliance requirements like GDPR or HIPAA can be addressed by implementing data protection and privacy measures in the API Gateway


12. How does API Gateway handle service discovery and dynamic routing?
   - API Gateway can integrate with service discovery mechanisms to dynamically route requests to available backend services
   - It can use health checks to monitor the status of backend services and update routing rules accordingly
   - Dynamic routing can be implemented to support blue-green deployments, canary releases, and A/B testing



13. How does API Gateway handle request and response transformation?
   - API Gateway can modify incoming requests before forwarding them to backend services (e.g., adding headers, changing payload format)
   - It can also transform responses from backend services before sending them back to clients (e.g., filtering fields, changing data structure)
   - Transformation rules can be defined using templates or scripting languages like Lua or JavaScript

14. How does API Gateway handle cross-origin resource sharing (CORS)?
   - API Gateway can configure CORS policies to allow or restrict access to APIs from different origins
   - It can set appropriate HTTP headers (e.g., Access-Control-Allow-Origin, Access-Control-Allow-Methods) in responses to enable or block cross-origin requests
   - CORS policies can be applied globally or on a per-API basis    


15. How does API Gateway handle request validation and schema enforcement?
   - API Gateway can validate incoming requests against predefined schemas (e.g., JSON Schema, OpenAPI)
   - It can enforce required fields, data types, and value constraints to ensure that requests are well-formed
   - Invalid requests can be rejected with appropriate error responses before reaching backend services


16. API Gateway vs Ingress?
   - Ingress is a Kubernetes resource that manages external access to services within a cluster, typically using HTTP/HTTPS
   - API Gateway provides additional features like authentication, authorization, rate limiting, and request/response transformation
   - Ingress is more focused on routing and load balancing, while API Gateway is designed for managing APIs and microservices in a more comprehensive manner

17. How does API Gateway handle WebSocket and gRPC traffic?
   - API Gateway can support WebSocket connections by acting as a reverse proxy, forwarding messages between clients and backend services
   - It can handle connection upgrades, maintain session state, and manage message routing for WebSocket traffic
   - For gRPC traffic, API Gateway can perform protocol translation, allowing clients to communicate with backend services using gRPC while exposing RESTful APIs to external clients


18. how to scale API Gateway?
   - API Gateway can be scaled horizontally by adding more instances behind a load balancer
   - It can use auto-scaling policies based on metrics like CPU usage, request count, or latency to dynamically adjust the number of instances
   - Caching, rate limiting, and other optimizations can be implemented to reduce the load on the API Gateway and improve performance

19. What if API Gateway goes down?
   - Implement high availability by deploying multiple instances of the API Gateway across different availability zones or regions
   - Use a load balancer to distribute traffic among healthy instances and automatically route requests away from failed instances
   - Implement health checks and monitoring to detect failures and trigger failover mechanisms to ensure continuous availability of the API Gateway

20. How does API gateway handles error like timeout or service unavailable or retry?
   - API Gateway can implement retry policies for failed requests to backend services, using exponential backoff and jitter to avoid overwhelming the backend
   - It can return appropriate error responses (e.g., 504 Gateway Timeout, 503 Service Unavailable) to clients when backend services are unavailable or unresponsive
   - Circuit breaker patterns can be implemented to prevent cascading failures and allow the system to recover gracefully from errors

21. How would you load balance across 100 pods?
    - Use a load balancer (e.g., NGINX, HAProxy, or cloud provider's load balancer) to distribute incoming requests across the 100 pods
    - Implement round-robin, least connections, or IP hash algorithms to determine how requests are routed to pods
    - Monitor pod health and performance metrics to ensure that traffic is directed to healthy pods and avoid overloading any single pod

22. How would you implement JWT validation?
   - API Gateway can validate JWTs by verifying the signature using a public key or secret key, depending on the signing algorithm used (e.g., RS256, HS256)
   - It can check the token's expiration time (exp claim) and ensure that it is still valid
   - The gateway can also verify other claims (e.g., issuer, audience) to ensure that the token is intended for the API Gateway and the specific backend service being accessed
   - If the JWT is invalid or expired, the API Gateway can reject the request with an appropriate error response (e.g., 401 Unauthorized)


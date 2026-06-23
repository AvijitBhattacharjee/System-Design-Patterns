1. High Level Discussion

An API Gateway is a server that sits between clients and backend services.

                    CLIENTS
        +----------------------------+
        | Mobile | Web | Third Party |
        +----------------------------+
                     |
                     v

          +----------------------+
          |     API GATEWAY      |
          +----------------------+

          Authentication
          Authorization
          Rate Limiting
          SSL Termination
          Request Validation
          Routing
          Load Balancing
          Caching
          Monitoring
          Logging
          Aggregation
          Protocol Translation
          Circuit Breaker
          Retry Logic

                     |
     -----------------------------------
     |              |                 |
     v              v                 v

+----------+   +----------+    +----------+
| UserSvc  |   | OrderSvc |    |PaymentSvc|
+----------+   +----------+    +----------+

     |              |                 |
     -----------------------------------
                     |
                Database


2. Popular API Gateway Products
Kong
NGINX
Traefik
Envoy
HAProxy
Apache APISIX
Spring Cloud Gateway

Cloud:

AWS API Gateway
Azure API Management
Google Cloud API Gateway


3. API Gateway vs Load Balancer
Interview Favorite

Feature	Load Balancer	API Gateway
Routing	Yes	Yes
Authentication	No	Yes
Rate Limit	No	Yes
Caching	No	Yes
Aggregation	No	Yes
Transformation	No	Yes
Protocol Conversion	No	Yes
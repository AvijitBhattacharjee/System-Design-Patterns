# 🏗️ System Design Patterns

A structured collection of popular system design problems with high-level design, architecture diagrams, Go implementations, corner cases, scaling and caching discussions.

> Built for interview prep, learning, and reference.

---

## 📚 Patterns Covered

| # | Pattern | Key Concepts |
|---|---------|-------------|
| 01 | [URL Shortener](./01-URL-Shortener/) | Hashing, redirects, DB choice |
| 02 | [Rate Limiter](./02-Rate-Limiter/) | Token bucket, sliding window |
| 03 | [Consistent Hashing](./03-Consistent-Hashing/) | Virtual nodes, ring |
| 04 | [LRU Cache](./04-LRU-Cache/) | Doubly linked list + hashmap |
| 05 | [Message Queue](./05-Message-Queue/) | Kafka-like, pub/sub |
| 06 | [Distributed Lock](./06-Distributed-Lock/) | Redis, TTL, deadlock |
| 07 | [Search Autocomplete](./07-Search-Autocomplete/) | Trie, top-k, CDN |
| 08 | [Notification System](./08-Notification-System/) | Fan-out, push/pull |
| 09 | [News Feed](./09-News-Feed/) | Timeline, ranking |
| 10 | [Distributed ID Generator](./10-Distributed-ID-Gen/) | Snowflake, clock skew |

---

## 📁 Structure

Each pattern follows the same structure:

```
XX-Pattern-Name/
├── README.md       # Problem statement + high level design
├── diagram.md      # Architecture diagram
├── main.go         # Core logic implemented in Go
└── NOTES.md        # Corner cases, scaling, caching discussions
```

---

## 🚀 Add a New Pattern

```bash
DIR="11-Your-Pattern-Name" && mkdir -p "$DIR" && touch "$DIR/README.md" "$DIR/diagram.md" "$DIR/main.go" "$DIR/NOTES.md" && ls "$DIR"
```

---

## 🧠 What Each Pattern Covers

| Section | What's inside |
|---|---|
| **High Level Design** | Requirements, capacity estimation, API design |
| **Architecture Diagram** | Component diagram with data flow |
| **Go Implementation** | Clean working code for the core logic |
| **Corner Cases** | Edge cases, failure scenarios, race conditions |
| **Scaling** | Horizontal scaling, sharding, load balancing |
| **Caching** | What to cache, where (Redis/CDN), TTL strategy |

---

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Diagrams:** Markdown + ASCII / draw.io
- **Concepts:** Distributed systems, databases, caching, messaging

---

## 🤝 Contributing

Contributions are welcome! If you'd like to add a new pattern or improve an existing one:

1. Fork the repo
2. Create a branch: `git checkout -b add/pattern-name`
3. Add your pattern following the folder structure above
4. Open a Pull Request

---

## 👨‍💻 Author

**Avijit Bhattacharjee**  
IBM Cloud • Backend Engineer • Open Source Contributor  
[GitHub](https://github.com/AvijitBhattacharjee) •

---


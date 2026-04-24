---
trigger: always_on
---

# Security, Scale, and Operations Rules

Use these rules as a baseline for security, traffic readiness, and application operations.

## Security Baseline
- All input is considered unsafe until validated.
- Auth, authorization, rate limits, and auditability should be considered early for sensitive areas.
- Use proper hashing for passwords and secure secret storage.
- Do not expose stack traces, raw queries, secrets, or internal identifiers unnecessarily.
- File uploads, query filters, export/import, and admin endpoints are priority areas for security review.
- For sensitive features, consider brute-force protection, throttling, and strong permission boundaries.

## Scale Baseline For ~1000 Users
A target of 1000 users doesn't necessarily mean microservices, but the foundation should be reasonably ready.

### Things to Consider
- pagination for large lists
- database indexing for core queries
- caching for frequently used data reads
- queues for heavy or non-blocking processes
- schedulers for periodic tasks that shouldn't burden user requests
- response time for critical endpoints
- observability: logs, error monitoring, and minimum metrics

### Anti-Patterns
- N+1 queries on the dashboard or main list page
- heavy jobs run synchronously at user requests if they can be moved to a queue
- dashboards fetch all data at once without limit
- large exports run directly on request without async mechanisms
- cron or schedulers run without locks or overlap protection if the task is sensitive

## Queues, Scheduling, and Background Work
- Use queues for email, notifications, exports, synchronization, file generation, third-party integrations, and other heavy processes.
- Schedulers are used for reminders, cleanup, periodic recaps, data sync, retry tasks, and task monitoring.
- Jobs should be designed idempotent if there is a possibility of retry or duplicate dispatch.
- Critical jobs must have clear timeouts, retry policies, logging, and failure handling.
- Don't deploy a business-critical scheduler without minimum observability.

## Operational Readiness
- Ensure a strategy is in place for:
- logging application errors
- monitoring failed jobs
- safely restarting workers
- separating sync and async tasks
- consistent environment configuration
- If traffic increases, consider incremental optimizations first:
- cache
- queue worker tuning
- database indexing
- simple read/write optimizations
- horizontal scaling only when absolutely necessary

## Quality Gate
- No heavy endpoints should remain in sync for no reason.
- No critical periodic tasks should be left without failure visibility.
- No assumptions about 1,000 users being secure without checking core queries and key bottlenecks.
- No sensitive features should be left without a baseline security review.
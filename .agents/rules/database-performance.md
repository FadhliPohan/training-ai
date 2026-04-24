---
trigger: always_on
---

# Database Performance Rules

Use these rules as a baseline for database design, queries, and data growth in new projects and advanced feature development.

## Basic Principles
- The database structure should be clear, consistent, and support business rules.
- Queries should be designed for real needs, not assumptions about small data forever.
- Database optimization starts with the correct schema and data access patterns, not patching them together when the system starts to slow down.

## Schema Design
- Use consistent, clear, and easy-to-understand table and column names.
- Define data types as sparingly as possible without sacrificing domain requirements.
- Use foreign keys or other integrity boundaries when appropriate for the project architecture.
- Add unique constraints for business rules that must be unique.
- Use nullables only when they make sense for the domain.

## Indexing Baseline
- Every column frequently used for searches, filters, joins, sorting, or lookups should be evaluated for index needs.
- Don't add excessive indexes; Indexes should follow the pattern of the main query.
- Consider compound indexes for queries that frequently use specific column combinations.
- Evaluate indexes on foreign keys, slugs, emails, statuses, dates, or primary relational/filter columns.

## Query Design
- Avoid `SELECT *` in queries that don't require all columns.
- Use pagination for potentially large data lists.
- Avoid repeating queries within loops.
- Avoid N+1 queries, especially on dashboards, tables, reports, and relational pages.
- Only retrieve relations and columns that are truly needed by the UI or API consumer.

## Growth Data Awareness
- Agents should consider how tables will grow after months of active data.
- Log, activity, transaction, notification, and history tables are all potential candidates for rapid growth.
- For fast-growing data, consider early:
- indexing strategy
- retention policy if relevant
- archive or cleanup if needed
- separation of operational queries and reporting queries when the system becomes overloaded

## Laravel Guidance
- Consciously use eager loading for relations that are actually used.
- Avoid loading unboundedly large relations if only a summary is needed.
- For complex queries, consider more explicit query builders, scopes, or service queries.
- Don't hide heavy query logic in accessors or patterns where the performance impact is difficult to see.
- For dashboards and reports, prioritize queries that are scalable and easy to optimize.

## Golang / SQL Guidance
- Write queries with explicit and understandable access patterns.
- Ensure the repository or data-access layer doesn't unknowingly create duplicate queries.
- For critical endpoints, evaluate timeouts, row counts, and payload sizes.
- Separate light-weight read queries from heavy-weight aggregation queries when the load starts to feel heavy.

## Reporting, Export, and Aggregation
- Large report queries or exports should not be assumed to be safe to run synchronously continuously.
- For large exports, consider async jobs.
- For heavy aggregations, consider caching, pre-compute, or summary tables if necessary.
- Don't make dashboards load all raw statistics when only a subset is displayed.

## Data Integrity and Performance Tradeoff
- Normalization remains the default, but don't be overly strict if real-world reads require more efficient reads with measurable tradeoffs.
- Denormalization should only be done for a clear, documented, and consistent reason.
- Don't copy data across tables without a justifiable business or performance reason.

## Red Flags
- Large tables without indexes on primary filter columns
- List queries without pagination
- Chained relationships loaded entirely on a regular page
- Dashboards that trigger multiple heavy queries per request
- Reports that retrieve the entire history without time limits or filters
- Use of accessors/helpers that silently trigger additional queries

## Database Quality Gate
- No critical data features without considering basic indexes.
- No large lists without pagination or explicit restrictions.
- No critical relational queries without N+1 evaluation.
- No large exports/reports without considering async processing.
- No assumptions about safety for 1,000 users without checking core query patterns.
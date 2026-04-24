---
trigger: always_on
---

# Backend API Rules

Use these rules for Laravel or Golang backends when building internal or public APIs.

## Basic Principles
- APIs should be consistent, predictable, and secure.
- Endpoint names, request payloads, and response shapes should not change without good reason.
- Separate concerns between the transport layer, business logic, and persistence layers.
- Use the appropriate HTTP method based on the purpose of the operation.

## HTTP Method Semantics
- `GET` is for reading data only and should not have business side effects that change state.
- `POST` is for creating new resources or triggering non-idempotent actions such as login, checkout, specific dispatches, or submitting processes.
- `PUT` is for full updates/replaces when the resource representation is completely updated.
- `PATCH` is for partial updates when only some fields are changed.
- `DELETE` is for deleting resources or executing a clear delete intent.
- Don't use `GET` for mutation operations.
- Don't mix `POST`, `PUT`, and `PATCH` without a clear reason in the same contract.

## Endpoint Design
- Use clear and consistent resource naming.
- Use HTTP methods according to their actions: `GET`, `POST`, `PUT`, `PATCH`, `DELETE`.
- Avoid endpoints with mixed actions if they can be broken down into clearer contracts.
- For filtering, sorting, pagination, and search, use stable query parameters.

## Request Validation
- All input must be validated at the request boundary.
- Validation should be explicit for type, required fields, enumerations, format, file size, and permissions.
- Don't rely solely on front-end validation.
- Sanitize input if any fields are at risk of being reused in queries, views, or exports.

## Response Contract
- Responses should be uniform and easy for consumers to read.
- At a minimum, clearly differentiate between success, message, and data.
- Validation errors, unauthorized, forbidden, not found, conflicts, and server errors must be clearly distinguishable.
- Do not leak stack traces or internal details to the client.

## Session and Auth Security
- If using a built-in auth framework like Laravel, continue to follow session and CSRF best practices.
- If not using a built-in auth framework, sessions must be explicitly designed and secure.
- Cookie-based sessions must use HttpOnly, Secure, and SameSite, as appropriate for the application context.
- Session IDs or tokens must not be easily guessed and must be rotatable during login, privilege changes, or sensitive events.
- For cookie/session-based auth, protect the mutation flow with CSRF protection.
- Sessions must have a clear expiration date, invalidation upon logout, and a revoke mechanism if needed.
- Don't store auth secrets in localStorage if a more secure flow with HttpOnly cookies is possible.
- If the frontend and backend are separated, the auth boundaries between `Auth.js`, session cookies, access tokens, and backend verification must be clear from the start.

## Split Frontend/Backend Guidance
- For a split architecture, the default frontend is `Next.js Pages Router` with `Auth.js`.
- The default backend for a split API is `Go Fiber`.
- Determine from the outset whether the backend accepts session-backed identities from the frontend server, bearer tokens, or signed tokens resulting from auth integration.
- Avoid confusing auth flows involving session cookies, bearer tokens, and manual headers without a clear contract.

## Data Integrity
- Interdependent multi-step operations must consider transactions.
- Database constraints must support critical business rules.
- Avoid silent failures when updating, deleting, or syncing data.
- For destructive actions, consider soft deletes, audit trails, or explicit confirmation flows.

## Auth & Authorization
- Auth must be consistent across sensitive endpoints.
- Clearly separate not logged in, no access, and resource not found.
- Authorization should not rely solely on a hidden button on the frontend.
- Admin, owner, and operator endpoints must have clear boundaries.

## Performance
- Always evaluate pagination for data lists.
- Avoid N+1 queries and eagerly load only relevant data.
- Don't send excessive payloads if the consumer doesn't need them.
- Add indexing if the core query pattern is clear.

## Laravel Guidance
- Use Form Requests for significant request validation.
- Use API Resources if the response transformation becomes complex.
- Use Policies or Gates for authorization beyond just auth checks.
- Pull business logic out of the controller if the flow is more than just a simple CRUD.

## Golang Guidance
- Handlers focus on request parsing, validation, and responses.
- Services focus on business rules.
- Repositories focus on data access when necessary.
- Use context, timeouts, and error wrapping disciplinedly for critical flows.

## API Quality Gate
- No critical endpoints without validation.
- No ambiguous error responses.
- No sensitive data leaks.
- No large lists without pagination or clear boundaries.
- No write endpoints that bypass auth and authorization.
- No non-built-in session flows left undone without clear cookie security, expiration, and invalidation.
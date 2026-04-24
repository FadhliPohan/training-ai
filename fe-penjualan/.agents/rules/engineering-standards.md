---
trigger: always_on
---

# Engineering Standards

## General
- Write code that is easy to read, not just quick to complete.
- Separate business logic from controllers, routes, or UI handlers once the logic has begun to evolve.
- Every change must maintain consistency in naming, structure, and data flow.
- Avoid magic values; use constants, enums, configs, or value objects when necessary.
- Prioritize reusability, scalability, clean code, and dry code, and structures that remain simple yet robust for growth.
- Deep code is allowed, but it doesn't mean complex; deep structure should provide clarity of responsibility, not add unnecessary layers.

## Architecture Direction
- Create structures that are easily reusable across pages, features, and modules.
- Avoid duplicating the same UI patterns, query patterns, and business flows.
- Abstractions are created when patterns are truly repetitive or potentially repetitive, not from the start for everything.
- Shared building blocks should be designed to be stable, easy to invoke, and not overly tied to a specific page.

## Frontend Component System
- Basic components such as buttons, modals, tabs, tables, datatables, inputs, selects, badges, cards, drawers, and alerts should be created as reusable internal components if they are intended to be used across pages.
- Pages should build more reusable components rather than rewriting the same UI elements repeatedly.
- Separate the `ui/base`, `shared`, `feature`, and `page` layers if the project size requires it.
- Reusable components should have a clear API: variant, size, state, loading, disabled, and reasonable extensibility.
- Avoid creating reusable components that are so specific to one use case that they are not truly reusable.

## Laravel & PHP
- Use explicit validation for all input requests.
- Use migrations and seeders that can be safely repeated.
- Use Form Requests, Policies, Middleware, Jobs, Events, and Services only when they provide clear clarity.
- Queries should be performance-conscious: avoid N+1, eager load only when necessary, and paginate large data sets.
- Auth and authorization must be clear; don't mix role checks haphazardly in views.
- Recurrent business logic should be pulled into a more appropriate service, action, domain helper, or class to prevent random scattering.

## React, TypeScript, JavaScript
- Strict TypeScript is preferred over loose typing.
- Components should have clear responsibilities.
- Forms, tables, modals, filters, and async state should not be mixed into one large component if they become difficult to maintain.
- Avoid excessive prop drilling; use healthy composition or state boundaries.
- All API integrations should have handling for loading, errors, and empty states.
- Separate UI primitives, composed components, feature sections, and page shells when complexity increases.
- Don't copy-paste logic between pages if it can be used as a hook, utility, service client, or reusable component.

## Golang
- Use a simple structure: handler, service, repository if necessary.
- Input validation should be done at the request boundary.
- Context, timeout, and error propagation should be clear for critical endpoints.
- Don't increase concurrency if there's no measurable benefit.
- API responses should be consistent and easily observable.
- Avoid putting all logic in handlers; break it out to the service/domain layer if the flow becomes congested.

## Security
- All input is considered untrusted until validated.
- Protect auth flows, uploads, query filters, and file access.
- Don't expose secrets in source code.
- Rate limits, audit logging, and permission boundaries need to be considered early on for sensitive endpoints.
- When a potential vulnerability is identified, agents should actively flag it.

## Quality Gate
- For the backend: validation, auth, error handling, and data integrity are the bare minimum.
- For the frontend: typed props/data, state flow, and UX error states are the bare minimum.
- For schema/data: migration paths must be clear.
- For critical features: provide minimal tests for critical logic areas.
- For UI systems: common components should not be implemented repeatedly across multiple pages without a compelling reason.
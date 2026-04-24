---
trigger: always_on
---

# Testing Strategy Rules

Use these rules as a baseline for testing quality for all new projects.

## Basic Principles
- Testing should protect the riskiest areas, not just focus on the number of tests.
- Primary focus: business logic, validation, authorization, data integrity, and flows that are easily broken during refactoring.
- Don't write fragile tests just to appear high coverage.

## Test Priorities
1. Validation and request rules
2. Authorization and permission boundaries
3. Core business logic
4. Critical query/data transformations
5. Critical API contracts
6. Important jobs, schedulers, and side effects
7. Important UI interactions if the frontend is complex

## Laravel Testing Guidance
- Use feature tests for endpoints, auth flows, and critical domain integrations.
- Use unit tests only for logic that is truly independent and worth testing separately.
- Tests should cover success cases, validation failures, unauthorized/forbidden cases, and important edge cases.
- For queues, events, notifications, mail, and jobs, use fakes/mocks as needed without making the test meaningless.
- For database-heavy logic, ensure assertions truly verify the business impact.

## React / TypeScript Testing Guidance
- Test the behavior of critical components, not internal implementation details.
- Prioritize form behavior, error states, loading states, conditional rendering, and important user actions.
- Don't over-test simple presentational components with low value.
- If the frontend is only a thin consumer of the backend, don't overburden UI tests unnecessarily.

## Golang Testing Guidance
- Test service logic, validation, mappers, and critical handlers.
- Use table-driven tests when appropriate for multiple input scenarios.
- Separate quick logic tests from integration tests that require databases or external dependencies.
- Don't let concurrency logic run untested if it's used in critical flows.

## Job & Scheduler Testing
- Ensure critical jobs are tested for:
- valid payload
- retry-safe behavior
- idempotency if the job can be re-executed
- failure handling and minimal logging
- The scheduler should be tested at the scheduling level and its business impact, especially if it touches billing, reminders, synchronization, or data cleanup.

## Test Style
- Test names should describe the behavior being protected.
- Each test should have a clear primary goal.
- Arrange, act, and assert should remain legible.
- Avoid nested mocks that make the test unreliable.
- Test data should be relevant to the domain, not random.

## Minimum Quality Gate
- No critical feature should be tested in a critical logic area.
- No critical auth/permission flow should be tested minimally.
- No critical job should be tested without verifying the success and failure paths.
- No major refactoring should be done without evaluating the impact on existing tests.
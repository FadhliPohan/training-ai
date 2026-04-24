# Project Init Workflow

This document is the default flow when starting a new project before full website or feature details are provided.

## 1. Initial Objectives
The agent should first help lay the foundation for the project, rather than jumping straight into random implementation.

Expected initial outputs:
- concise product objectives
- application type
- key system actors
- initial modules/MVP
- most reasonable stack choice
- basic project structure

## 2. Default Work Order
1. Identify the project objectives and primary users.
2. Determine whether the project is best suited to full-stack Laravel, Laravel + React, or a separate backend with Go/React.
3. Compile the MVP module list.
4. Define core entities and database relationships.
5. Define auth flow and roles/permissions.
6. Establish folder patterns, naming, and coding conventions.
7. Define baseline API/UI flow.
8. Define initial environment, testing, and deployment checklist.

## 3. Default Stack Decision
### Choose full-stack Laravel if:
- project admin/internal tools
- CRUD flow is dominant
- high need for fast delivery
- no need for complex frontend-backend separation

### Choose Laravel + React TypeScript if:
- more interactive UI
- many dynamic forms, dashboards, filtering, or client state
- want to maintain a more scalable frontend

### Choose Go services if:
- high-performance API requirements
- dedicated worker/background services
- want to separate certain bounded contexts from the main app

## 4. Baseline Folder Thinking
### Laravel project
- `app/Http/Controllers`
- `app/Http/Requests`
- `app/Models`
- `app/Services` if business logic starts to thicken
- `app/Policies` for authorization
- `database/migrations`
- `database/seeders`
- `resources/js` for React/TypeScript if used

### Go service
- `cmd/`
- `internal/handlers`
- `internal/services`
- `internal/repositories`
- `internal/middleware`
- `internal/config`
- `pkg/` only if reusable across services

## 5. Minimum Project Init Outcomes
Before any additional features are introduced, the agent should be able to help prepare:
- initial architecture
- MVP module list
- initial database design
- API response standards
- auth and permission patterns
- basic coding rules
- subsequent feature implementation workflow
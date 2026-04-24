---
trigger: always_on
---

# Default Stack Priority

Use the following preferences as a baseline for all new projects, unless the user directs otherwise or the existing codebase dictates otherwise.

## Top Priorities
1. Laravel for backend or full-stack web apps.
2. React + TypeScript for interactive frontends.
3. JavaScript/TypeScript for tooling and shared utilities.
4. Golang for high-performance APIs/services.
5. Non-Laravel PHP for compatibility purposes only.
6. CodeIgniter is not the default stack and should only be used for legacy maintenance or explicit requests.

## Architectural Preferences
- Start with a clean monolith before breaking down into microservices.
- Use separate APIs only if absolutely necessary for deployment, mobile apps, third-party integrations, or isolated workloads.
- For admin panels, dashboards, and internal applications, prioritize delivery speed without sacrificing structure.
- For product websites that later evolve into systems, prepare the foundation for modules and auth from the start.
- If the frontend and backend are separated, the default is Next.js Pages Router for the frontend and Go Fiber for the backend.
- For auth in a separated frontend-backend architecture, the default is Auth.js.

## Practical Technology Preferences
- Modern frontend: React + TypeScript + Vite.
- If React is used for the main frontend app, the default UI component is shadcn/ui as a base that can then be customized.
- Laravel frontend bridge: Inertia if you need a fast, full-stack approach with a SPA feel.
- Laravel server-rendered: Blade + Tailwind + Alpine + reusable custom Blade components if interaction is light and delivery is faster.
- For Blade UI, prioritize building your own internal components rather than relying too heavily on large UI libraries.
- Go Backend: Fiber as a lightweight default, unless there's a compelling reason to choose another framework.
- Default Database: MySQL or PostgreSQL.
- Default Styling: Tailwind if you need fast delivery and a systematic UI.

## Default Anti-Pattern
- Don't choose CodeIgniter for a new project without a compelling reason.
- Don't break services too early.
- Don't add packages just because they're popular; there should be a technical reason.
- Don't create layers of abstractions if the logic is still simple.
# Website Discovery Workflow

Use this workflow when a new project doesn't yet have complete website specifications. The goal is to help agents explore requirements in a structured way before moving on to technical design or implementation.

## 1. Outcomes to Lock
Before coding, agents should strive to clarify the following:
- Who is this website for?
- The business or primary purpose of the website.
- The primary action expected from the user.
- The type of website: company profile, landing page, dashboard, internal system, marketplace, portal, or service application.
- Whether the website is public, private, or a mix.

## 2. Discovery Checklist
### Business Context
- The name or context of the project.
- The problem to be solved.
- The primary value proposition.
- The primary and secondary target users.
- The priority of the MVP over advanced features.

### Website Scope
- The main pages required.
- Core features required in the first version.
- Features that can be postponed.
- Login, dashboard, or user role requirements.
- CMS requirements, file uploads, or content management requirements.

### Data & System
- Core entities to be stored.
- What data can users input?
- What data can only be managed by admins.
- Relationships between key entities.
- Is payment, WhatsApp, email, maps, AI, or third-party API integration required? Third

### UI/UX Direction
- Desired visual style
- Website references, if any
- Brand tone: formal, modern, playful, premium, minimal, etc.
- Device priority: mobile-first, desktop-first, or balanced
- Accessibility or multilingual needs

### Operational Needs
- Who is the system admin/operator?
- Audit log or change history needs
- Data export/import needs
- Notification needs
- SEO, analytics, or conversion tracking needs

## 3. Mapping Technical Decisions
Once the initial requirements are clear, the agent must map them to technical decisions:

### Use full-stack Laravel if:
- The website is dominated by CRUD, forms, auth, admin dashboard, and server-side logic
- The team or owner wants faster delivery while maintaining a clean structure

### Use Laravel + React + TypeScript if:
- The website requires more complex UI interactions
- There are heavy dashboards, filters, dynamic tables, form wizards, or a fairly active client state

### Use additional Go services if:
- There are heavy processes, public APIs, worker, webhook processor, or integration that needs to be separated

## 4. Required Discovery Outputs
After discovery, the agent should be able to compile at a minimum:
- a summary of the website's goals
- a list of key users/roles
- a list of MVP pages and features
- initial database entities and relationships
- a recommended stack
- key user flows
- a list of third-party integrations
- initial technical risks

## 5. Limitations During Discovery
- Don't over-engineer before the scope is clear.
- Don't force microservices if the website doesn't yet need them.
- Don't overcomplicate the UI if the underlying problem hasn't been proven.
- Don't assume all features must be present in the first phase.

## 6. Deliverables After Discovery
The agent should ideally transform the discovery results into the following working documents:
- a short project brief
- MVP modules
- a sitemap or list of pages
- an initial schema design
- a phase 1 and phase 2 backlog
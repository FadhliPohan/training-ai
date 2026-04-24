# Workflow Module Planning

Use this workflow after discovery is clear and before major feature implementation begins.

## 1. Objective
Break website requirements into realistic modules that can be implemented incrementally, without clumping all concerns into one random sprint.

## 2. How to Break Down Modules
Separate modules based on one or a combination of the following:
- business domain
- user actors
- key workflows
- data boundaries
- permission boundaries

Examples of common modules:
- authentication
- dashboard
- user management
- master data
- core transactions or processes
- reports
- system settings
- notifications
- audit logs

## 3. Minimum Content of Each Module
Each module should have this definition:
- module purpose
- users who use it
- related pages or endpoints
- involved entities
- key business rules
- dependencies to other modules
- implementation priority

## 4. Default Priority Order
1. authentication and authorization
2. core master data
3. key business processes
4. dashboard and monitoring
5. reports, export/import, and operational optimization
6. non-critical enhancements

## 5. How to Define Phases
### Phase 1
- everything required for the product to be usable
- it doesn't have to be complete, but the main flow should run from start to finish

### Phase 2
- operational efficiency
- automation
- additional reporting
- UI/UX polishing
- additional integrations

### Phase 3
- scale improvements
- advanced analytics
- AI assistance
- multi-tenant or multi-branch if needed

## 6. Expected Module Output
After planning, the agent should be able to produce:
- a list of prioritized modules
- dependencies between modules
- a list of entities per module
- an implementation backlog per phase
- technical risks per module
- which module is most secure to start first

## 7. Limitations
- Do not create each page as a separate module if the business domain is the same.
- Do not mix auth, master data, and core processes into one unscalable module.
- Do not start the reporting module first if the data source process is not yet mature.
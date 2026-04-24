# Feature Delivery Workflow

Use this workflow when new website specifications or features are added on top of the initial project foundation.

## 1. Feature Intake
Before coding, agents must lock in the following:
- feature purpose
- who will use the feature
- primary inputs and outputs
- database impact
- auth/permission impact
- UI and API impact

## 2. Implementation Order
1. Validate the feature scope to ensure it's unambiguous.
2. Check the impact on existing schemas and data.
3. Design the request/response contract.
4. Implement backend logic and validation.
5. Implement the UI flow or consumer API.
6. Add error handling, loading, and empty states.
7. Review security, data integrity, and regression risks.
8. Add tests for critical logic.

## 3. Review Before Completion
Make sure the agent checks:
- whether there are any redundant or N+1 queries
- whether input validation is complete
- whether roles/permissions are correct
- whether edge cases of empty/null/error data are present
- whether naming and structure remain consistent
- whether new dependencies are worthy of being added

## 4. Expected Output
After a feature is completed, the agent should be able to summarize:
- files that have changed
- schema changes
- new routes or endpoints
- remaining technical risks
- the most reasonable next steps
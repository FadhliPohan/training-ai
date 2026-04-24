# Workflow Jobs and Scheduling

Use this workflow when a project requires background processing, scheduling, reminders, or periodic tasks.

## 1. When to Use a Queue / Job
Use a queue or job if the process:
- is heavy or time-consuming
- does not have to complete at the user's request
- involves email, notifications, exports, imports, file generation, webhooks, synchronization, or third-party integration
- has the potential to fail and requires a retry policy

## 2. When to Use a Scheduler
Use a scheduler if the process:
- runs periodically
- is based on a specific time
- requires daily/weekly/monthly recaps
- requires data cleanup, monitoring, reminders, or scheduled retry

## 3. Job Design Checklist
Before implementing a job, the agent must clarify:
- what the job inputs are
- what the outputs or side effects are
- whether the job must be idempotent
- how many retry times
- how long the timeout is
- what happens if the job fails
- does the user need to know the job status
- is the job safe to run in parallel or requires a lock

## 4. Scheduler Design Checklist
Before implementing a scheduler, the agent must clarify:
- when the scheduler will run
- whether there are any risks Overlap
- How to monitor success/failure
- Does the result trigger another job
- How to verify the business outcome

## 5. Recommended Implementation Pattern
- User requests only dispatch jobs, not heavy processes directly.
- Small schedulers can directly execute light commands, but heavy processes should be dispatched to a queue.
- Separate jobs per responsibility if the flow is complex.
- Save process status if the user needs tracking.
- For critical tasks, prepare minimal retry and alerts.

## 6. Things to Test
- Job success path
- Job failure path
- Duplicate execution or retry scenario
- Scheduler trigger
- Impact to primary data
- Task security if it touches sensitive data

## 7. General Risks
- Duplicate jobs due to retry without idempotency
- Scheduler overlap and processing the same data twice
- Heavy processes continue to block user requests
- Failed jobs are not monitored
- Periodic commands consume excessive resources
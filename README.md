# Task Management System

You are tasked with designing a Task Management System for a team where users can:
Create Tasks with attributes like:

- Title
- Description
- Priority (High, Medium, Low)
- Status (To-Do, In Progress, Done)
- Assignee

- Implement features for:
  - Assigning tasks to team members.
  - Filtering tasks based on priority and status.
  - Notifying users when a task is updated (Observer Pattern).

Ensure the system:
- Follows SOLID principles.
- Allows adding new task statuses or priorities without modifying the - - existing code.
- Supports different task creation modes (Factory Pattern).
- Allows switching between strategies for filtering (e.g., by priority or status).
- Implements notifications using the Observer Pattern.
- Implements a Singleton for managing the notification service (to ensure only one instance exists).

| **Feature**            | **Static Input**                               | **Expected Output**                                                      |
|-------------------------|-----------------------------------------------|--------------------------------------------------------------------------|
| **List all tasks**      | Hardcoded list of tasks                      | All tasks with their details                                             |
| **Filter by Assignee**  | Predefine a user to filter for (e.g., Alice)  | Tasks assigned to that user                                              |
| **Update Task Status**  | Predefine which task's status to update       | Confirmation of updated task status                                      |
| **Assign Priority**     | Predefine a priority decorator for tasks      | Tasks updated with priority, e.g., "High Priority: Design UI Mockups"    |
| **Notifications**       | Predefine task updates                       | Messages sent to the assignees for those updates                         |
| **Task Metrics**        | Static list of tasks                        | Summary of the task counts by status (e.g., To Do, In Progress, Completed) |
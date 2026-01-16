# NeuroSync: Circadian-Aware Task Scheduler

NeuroSync is a Go microservice that optimizes task scheduling by aligning workload difficulty with predicted biological alertness levels based on the Two-Process Model of sleep regulation.

## Methodology
Unlike standard calendars that allocate time slots based on availability, NeuroSync utilizes a bio-mathematical model to estimate cognitive capacity.

* **Process S (Homeostatic Sleep Drive):** Models fatigue accumulation relative to wake time.
* **Process C (Circadian Rhythm):** Models the oscillation of alertness, accounting for variations such as the afternoon energy dip.

The system accepts a task list with associated difficulty ratings and applies a greedy algorithm to assign high-load tasks to peak alertness windows and low-load tasks to periods of lower energy.

## Scheduling Comparison
The following table compares a standard time-based schedule against the NeuroSync capacity-based schedule for a high-focus task (System Design) and a low-focus task (Administrative Updates).

| Time Window | Standard Allocation | NeuroSync Allocation |
| :--- | :--- | :--- |
| **10:00 AM** | Administrative Updates | **System Design** (Peak Focus) |
| **03:00 PM** | System Design | **Administrative Updates** (Recovery) |

## Technical Architecture
The application is structured using **Hexagonal Architecture** to separate business logic from external interfaces.

* **Language:** Go 1.25
* **Database:** SQLite
* **Containerization:** Docker (Alpine Linux)
* **Routing:** Chi

### Directory Structure
* `internal/core/domain`: Business logic entities and cognitive load models.
* `internal/core/ports`: Input/Output interfaces.
* `internal/adapters`: Implementation of ports (HTTP handlers, SQL repositories).

## Usage

### Docker Build & Run
Run the following commands to build and start the container:

```bash
# 1. Build the image
docker build -t neurosync-backend .

# 2. Run the container (Access at localhost:8080)
docker run -p 8080:8080 neurosync-backend

API Endpoints
1. Generate Schedule (POST)

Submit a user's wake time and tasks to generate an optimized schedule.

Endpoint: POST /schedule

Body:
JSON

{
  "user_id": "test_user",
  "wake_time": "2026-01-14T07:00:00Z",
  "tasks": [
    { "title": "System Design", "duration": 120, "difficulty": 3 },
    { "title": "Email Catchup", "duration": 30, "difficulty": 1 }
  ]
}

2. Retrieve Schedule (GET)

Fetch the stored schedule for a specific user.

Endpoint: GET /schedule/{user_id}

Example: GET /schedule/test_user
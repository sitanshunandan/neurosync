# NeuroSync Engine 🧠⚡

> A biologically-aware job scheduler that optimizes productivity based on circadian rhythms and homeostatic sleep pressure.

**Status:** _Pre-Alpha (MVP)_ | **Stack:** Go (Golang)

---

## 📖 The Problem
Traditional calendars are **static**. They schedule tasks based on *time* (e.g., "Meeting at 2:00 PM"), ignoring the biological reality that human cognitive capacity fluctuates dramatically throughout the day.

## 💡 The Solution
**NeuroSync** is a headless backend engine that:
1.  **Ingests Biological Data:** Wake time, sleep quality, and caffeine intake.
2.  **Models Cognitive State:** Uses the **Two-Process Model of Sleep Regulation** (Process C + Process S) to predict real-time alertness.
3.  **Optimizes Schedules:** Automatically assigns high-load tasks (creative/analytical) to peak energy windows and administrative tasks to "slump" periods.

---

## 🔬 The Science (How it Works)
The engine calculates a **Cognitive Capacity Score (0-100)** for every hour of the day using a weighted algorithm:

* **Process S (Sleep Pressure):** A linear decay function representing the buildup of adenosine (fatigue) the longer you stay awake.
* **Process C (Circadian Drive):** A sinusoidal oscillation mimicking the body's cortisol/melatonin release cycles.
* **Ultradian Modulation:** (Coming Soon) 90-minute energy cycles for granular task batching.

---

## 🏗️ Architecture
This project follows **Hexagonal Architecture (Ports & Adapters)** to decouple biological logic from infrastructure.

```text
neurosync/
├── cmd/                # Application entry points
├── internal/
│   ├── core/           # PURE LOGIC (The "Brain")
│   │   ├── domain/     # Structs: BioRhythm, Task, CognitiveLoad
│   │   └── ports/      # Interfaces for repositories/services
│   ├── logic/          # Algorithms: Circadian math, Scheduler
│   └── adapters/       # INFRASTRUCTURE (The "Limbs")
│       └── handlers/   # HTTP / CLI handlers
└── docs/               # Architecture Decision Records (ADRs)
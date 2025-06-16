# Schedule Tester

## Overview

This project is designed for developing and testing task scheduling algorithms. The primary goal is to evaluate
different scheduling strategies for distributing tasks over time efficiently. Successful algorithm implementations from
this project will eventually be integrated into another project called CoreBGP.

## Purpose

The Schedule Tester provides a framework to:

- Test various scheduling algorithms with different task pools
- Visualize the load distribution over time
- Measure performance metrics like computation duration
- Compare different scheduling strategies

## Project Structure

```
schedule-tester/
├── input-data/           # JSON files with task definitions
│   ├── smallpool.json    # Small set of tasks
│   ├── sametimepool.json # Tasks arriving at the same time
│   └── bigpool.json      # Large set of tasks
├── internal/
│   └── engine/           # Core testing functionality
│       ├── engine.go     # Test data structures and naive load computation
│       └── plot.go       # Visualization of load distributions
├── pkg/
│   └── scheduler/        # Scheduler implementations
│       ├── scheduler.go  # Interface and task definitions
│       ├── v1/           # Version 1 implementation
│       │   └── init.go   # Scheduling algorithm implementation
│       └── v..           # Other version implementation
└── main.go               # Main application entry point
```

## Task Model

Tasks are defined with the following properties:

- `ID`: Unique identifier for the task
- `Period`: Time interval for the task to run again
- `Arrived`: Time when the task arrives

This is a temporary structure describing incoming tasks. Later it will be adapted for real use.

## Usage

To run the scheduler tests:

```bash
go run main.go
```

This will:

1. Load task definitions from the input-data directory
2. Run the scheduling algorithm on each task pool
3. Generate a visualization of the load distribution in `loads.png`
4. Display performance metrics for each test case

## Scheduling Algorithms

### V1

The first implementation (v1) uses loops to calculate the startup with the least load. No additional logic is used.

1. For each new task, it calculates the best starting time within its arrival period
2. The best time is determined by finding the time slot that minimizes the maximum load
3. The algorithm uses the Least Common Multiple (LCM) of task periods to determine the scheduling horizon

### V2

The second implementation (v2) is functionally almost identical to v1 and follows the same approach for scheduling
tasks.

1. The least common multiple (LCM) of periods is used for the planning horizon
2. The enumeration of possible starts is stopped early if the maximum load for the current start coincides with the
   value at the previous step.
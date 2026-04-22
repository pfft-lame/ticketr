# TicketR

**Status:** In progress (early stage)

## About

TicketR is a backend service for a movie ticket booking system built using Go. The goal is to handle high-concurrency booking scenarios reliably, preventing issues like double-booking of seats during peak traffic.

The project focuses on building a scalable and robust backend using modern tools and practices.

## Current Features

- Basic movie, city, and theater management APIs
- Initial booking flow design (work in progress)
- Structured service and handler layers

## Goals

- Handle concurrent seat bookings safely
- Ensure data consistency under high load
- Build a clean and maintainable backend architecture

## Roadmap

Planned features and integrations:

- Redis (caching, distributed locking)
- Auth system (OAuth 2.0 / OpenAuth)
- Fine-grained authorization (OpenFGA)
- Observability:
  - Prometheus (metrics)
  - Grafana (visualization)

- Improved booking workflow with concurrency control

## Notes

This is an evolving learning/project codebase. Some planned features (auth, Redis, observability, etc.) are not yet implemented and are part of the roadmap.

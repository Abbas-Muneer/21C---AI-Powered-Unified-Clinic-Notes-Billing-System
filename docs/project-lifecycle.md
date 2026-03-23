# Project Lifecycle

## Overview

This project was designed as an assessment-ready modular monolith that can be inspected quickly by reviewers while still demonstrating practical production decisions.

## Delivery Sequence

1. Define folder structure and backend/frontend boundaries.
2. Model the relational schema for consultations, notes, prescriptions, lab requests, billing, and pricing references.
3. Implement Go services, repositories, handlers, and AI provider abstraction.
4. Build the React workflow around a single consultation workspace.
5. Add print document routes and structured output endpoints.
6. Seed demo data and document the system for evaluators.

## Final Outcomes

- end-to-end clinic workflow from patient selection to printable bill
- mandatory AI/NLP integration with mock and external provider modes
- clear separation between parsing, normalization, billing, and persistence
- assessment-friendly documentation and demo script

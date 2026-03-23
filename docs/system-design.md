# System Design

## Problem Statement

ABC Health Clinic needs one consultation workspace where doctors can capture mixed narrative input once, classify it into medications, investigations, and notes with AI assistance, store the structured result, and produce operational outputs including the bill.

## Functional Requirements

- unified consultation input screen
- AI/NLP extraction of drugs, tests, and observations
- doctor review and edit before save
- structured PostgreSQL persistence
- prescription, lab request, notes, and bill outputs
- billing calculation using reference pricing and service fees

## Non-Functional Requirements

- easy local setup for assessors
- modular monolith architecture
- replaceable AI provider
- predictable demo mode without external AI keys
- print-friendly UI

## High-Level Architecture

```mermaid
flowchart LR
    Doctor[Doctor] --> UI[React Frontend]
    UI --> API[Go REST API]
    API --> Parser[AI Provider Layer]
    API --> Billing[Billing Service]
    API --> Repo[Repository Layer]
    Repo --> DB[(PostgreSQL)]
    API --> Docs[Document Endpoints]
    Docs --> UI
```

## Consultation Workflow

```mermaid
flowchart TD
    A[Select or create patient] --> B[Enter unified consultation note]
    B --> C[Parse with AI]
    C --> D[Normalize and validate structured JSON]
    D --> E[Show editable review panel]
    E --> F[Calculate billing summary]
    F --> G[Save consultation transaction]
    G --> H[Open printable prescription, lab, notes, bill views]
```

## AI Parsing Flow

```mermaid
flowchart TD
    Input[Unified note] --> Provider[AI provider abstraction]
    Provider -->|mock| Mock[Deterministic parser]
    Provider -->|openai| LLM[OpenAI-compatible structured extraction]
    Mock --> Normalize[Normalization rules]
    LLM --> Normalize
    Normalize --> Pricing[Reference pricing lookup]
    Pricing --> Review[Editable frontend review state]
```

## Database Relationship Overview

```mermaid
erDiagram
    patients ||--o{ consultations : has
    consultations ||--|| clinical_notes : owns
    consultations ||--o| prescriptions : owns
    prescriptions ||--o{ prescription_items : contains
    consultations ||--o| lab_requests : owns
    lab_requests ||--o{ lab_request_items : contains
    consultations ||--o{ billing_items : contains
```

## Design Notes

- The modular monolith keeps evaluation simple while preserving clean boundaries.
- AI parsing is intentionally separated from billing and persistence so model changes do not ripple through business logic.
- The mock provider keeps the project runnable without an API key, which matters in assessment review environments.

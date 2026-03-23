# Demo Video and Documentation Link
https://drive.google.com/drive/folders/1A8TzlgiWLpJ_RpdF55oLl1QmwG3qI_Ej?usp=sharing




# AI-Powered Unified Clinic Notes & Billing System

Assessment-ready full-stack application for a Software Engineer technical evaluation. The system gives doctors one unified consultation workspace, uses AI/NLP to extract structured medical and billing information, stores it in PostgreSQL, and produces printable operational outputs.

## Stack

- Backend: Go + Gin + sqlx
- Database: PostgreSQL
- Frontend: React + TypeScript + Vite + Tailwind CSS
- AI parsing: OpenAI-compatible provider or built-in mock parser

## Core Features

- patient registration and selection
- unified consultation note entry
- AI parsing into medications, lab tests, and clinical notes
- editable review panel before save
- transactional persistence of structured consultation records
- billing calculation from drug prices, lab test prices, and consultation fee
- print views for prescription, lab request, notes, and bill
- demo seed data for immediate evaluator use

## Repository Layout

```text
/
  backend/
  database/
  docs/
  frontend/
  .env.example
  docker-compose.yml
  README.md
```

## Architecture Summary

- Modular monolith for fast inspection during assessment
- Layered Go backend:
  - handlers
  - services
  - AI provider layer
  - billing service
  - repositories
- React frontend with pages, reusable UI cards, consultation workflow components, services, hooks, and print routes
- PostgreSQL schema normalized around consultations and structured child records

Detailed design material:

- [System Design](/c:/Users/MSII/Desktop/21C---AI-Powered-Unified-Clinic-Notes-Billing-System/docs/system-design.md)
- [API Design](/c:/Users/MSII/Desktop/21C---AI-Powered-Unified-Clinic-Notes-Billing-System/docs/api.md)
- [AI/NLP Design](/c:/Users/MSII/Desktop/21C---AI-Powered-Unified-Clinic-Notes-Billing-System/docs/ai-nlp-design.md)
- [Demo Walkthrough](/c:/Users/MSII/Desktop/21C---AI-Powered-Unified-Clinic-Notes-Billing-System/docs/demo-walkthrough.md)
- [Project Lifecycle](/c:/Users/MSII/Desktop/21C---AI-Powered-Unified-Clinic-Notes-Billing-System/docs/project-lifecycle.md)

## Database Setup

### Option 1: Docker Compose

1. Copy `.env.example` to `.env`.
2. Run `docker compose up --build`.
3. Open the frontend at `http://localhost:5173`.

The PostgreSQL container automatically runs SQL files mounted from `backend/migrations` and `backend/seeds`.

### Option 2: Manual Local Run

1. Create PostgreSQL database `clinic_notes`.
2. Apply [backend/migrations/001_init.sql](/c:/Users/MSII/Desktop/21C---AI-Powered-Unified-Clinic-Notes-Billing-System/backend/migrations/001_init.sql).
3. Apply [backend/seeds/001_demo_seed.sql](/c:/Users/MSII/Desktop/21C---AI-Powered-Unified-Clinic-Notes-Billing-System/backend/seeds/001_demo_seed.sql).

## Environment Variables

Use [.env.example](/c:/Users/MSII/Desktop/21C---AI-Powered-Unified-Clinic-Notes-Billing-System/.env.example) as the source of truth.

Important variables:

- `BACKEND_DATABASE_URL`
- `BACKEND_AI_PROVIDER`
- `BACKEND_AI_API_KEY`
- `BACKEND_AI_MODEL`
- `BACKEND_CONSULTATION_FEE`
- `VITE_API_BASE_URL`

## Running Backend

```bash
cd backend
go mod tidy
go run ./cmd/server
```

## Running Frontend

```bash
cd frontend
npm install
npm run dev
```

## AI Configuration

### Mock mode

- Set `BACKEND_AI_PROVIDER=mock`
- No external API key required
- Best for assessment demos and first-time local setup

### OpenAI-compatible mode

- Set `BACKEND_AI_PROVIDER=openai`
- Set `BACKEND_AI_API_KEY`
- Optionally override `BACKEND_AI_BASE_URL` and `BACKEND_AI_MODEL`

## Example Workflow

1. Open `Patients` and review seeded patients.
2. Open `New Consultation`.
3. Enter one mixed note.
4. Click `Parse with AI`.
5. Review and edit extracted medications, tests, and notes.
6. Save or finalize consultation.
7. Open print views from the consultation detail page.

## Testing

Backend tests included:

- billing summary calculations
- parser normalization behavior
- service-level parser contract coverage

Frontend tests included:

- billing summary component rendering

Commands:

```bash
cd backend && go test ./...
cd frontend && npm test
```

## Assumptions

- Single-clinic deployment
- Doctor identity is captured per consultation as a free text field
- Billing uses reference prices and a flat consultation fee
- Printable HTML views are acceptable instead of server-side PDF generation

## Future Improvements

- richer clinical terminology normalization
- speech-to-text transcript timestamping
- authentication and doctor profiles
- inventory-aware dispensing quantities
- PDF export and email delivery

## Verification Note

This environment did not have the `go` toolchain installed, so backend test execution could not be run here. The project is structured to be runnable in a normal local Go/Node/PostgreSQL environment or through Docker Compose.

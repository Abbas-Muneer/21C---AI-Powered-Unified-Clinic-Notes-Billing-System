# API Design

Base path: `/api`

## Patients

- `POST /patients`
  - Create a patient record.
- `GET /patients`
  - List patients for selection in the consultation workspace.
- `GET /patients/:id`
  - Retrieve one patient.

## Consultation and AI

- `POST /consultations/parse`
  - Accepts `patient_id`, `doctor_name`, and `raw_input_text`.
  - Returns structured medications, lab tests, clinical notes, metadata, and billing preview.
- `POST /consultations`
  - Saves the reviewed consultation and billing outputs.
- `GET /consultations/:id`
  - Returns saved consultation details for the detail screen.
- `PUT /consultations/:id`
  - Updates structured records after review changes.

## Documents

- `GET /consultations/:id/prescription`
- `GET /consultations/:id/lab-request`
- `GET /consultations/:id/notes`
- `GET /consultations/:id/bill`

Each document endpoint returns print-oriented JSON with a common header block and the relevant structured content.

## Response Shape

Successful responses use:

```json
{
  "data": {}
}
```

Error responses use:

```json
{
  "error": "validation_failed",
  "message": "..."
}
```

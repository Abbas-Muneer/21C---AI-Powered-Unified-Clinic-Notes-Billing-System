# Demo Walkthrough

## 5-10 Minute Demo Flow

1. Open the dashboard and explain the modular monolith architecture.
2. Visit the patient list and show the seeded demo patients.
3. Open `New Consultation`.
4. Paste or dictate a mixed note containing symptoms, medications, dosage instructions, and test requests.
5. Click `Parse with AI` and explain the provider mode currently active.
6. Review the structured extraction panel.
7. Edit one medication or note field to show clinician-in-the-loop correction.
8. Finalize the consultation.
9. Open the saved detail page.
10. Open each print view: prescription, lab request, clinical notes, bill.
11. Highlight how the bill is derived from reference drug prices, lab test prices, and the consultation fee.

## Points To Explain

- AI provider abstraction and mock mode for evaluator readiness
- structured PostgreSQL schema instead of storing everything as raw JSON
- transaction-based save flow
- print-friendly frontend routes backed by dedicated document endpoints

## Suggested Example Input

Patient presents with fever, sore throat, and body aches. Start Amoxicillin 500 mg three times daily for 5 days after meals. Add Paracetamol 500 mg twice daily for 3 days. Order Full Blood Count and CRP. Encourage oral hydration and rest.

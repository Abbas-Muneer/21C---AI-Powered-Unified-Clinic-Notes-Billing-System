import { BillingSummary, ClinicalNotes, LabTest, Medication, PrintHeader } from "./api";

export type PrescriptionDocument = {
  header: PrintHeader;
  medications: Medication[];
};

export type LabRequestDocument = {
  header: PrintHeader;
  lab_tests: LabTest[];
};

export type NotesDocument = {
  header: PrintHeader;
  clinical_notes: ClinicalNotes;
};

export type BillDocument = {
  header: PrintHeader;
  billing: BillingSummary;
};

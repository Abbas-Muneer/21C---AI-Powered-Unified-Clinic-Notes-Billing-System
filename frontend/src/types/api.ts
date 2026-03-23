export type Patient = {
  id: string;
  full_name: string;
  date_of_birth: string;
  gender: string;
  phone: string;
  email: string;
  address: string;
  created_at: string;
};

export type Medication = {
  drug_name: string;
  dosage: string;
  frequency: string;
  duration: string;
  route: string;
  instructions: string;
  quantity: number;
  unit_price: number;
  line_total: number;
};

export type LabTest = {
  test_name: string;
  instructions: string;
  unit_price: number;
  line_total: number;
};

export type ClinicalNotes = {
  observations: string;
  additional_notes: string;
};

export type BillingItem = {
  item_type: string;
  item_name: string;
  quantity: number;
  unit_price: number;
  line_total: number;
};

export type BillingSummary = {
  items: BillingItem[];
  grand_total: number;
};

export type ParseMetadata = {
  provider: string;
  model: string;
  confidence: string;
  recovery_applied: boolean;
  warnings: string[];
  normalization_notes: string[];
};

export type ParseResult = {
  medications: Medication[];
  lab_tests: LabTest[];
  clinical_notes: ClinicalNotes;
  billing: BillingSummary;
  metadata: ParseMetadata;
};

export type ConsultationDetail = {
  id: string;
  patient: Patient;
  doctor_name: string;
  status: string;
  raw_input_text: string;
  consultation_date: string;
  parsed_result: ParseResult;
};

export type PrintHeader = {
  clinic_name: string;
  document_title: string;
  consultation_id: string;
  patient_name: string;
  patient_identifier: string;
  doctor_name: string;
  date: string;
};

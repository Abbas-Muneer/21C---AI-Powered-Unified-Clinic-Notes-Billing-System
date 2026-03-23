import { useEffect, useMemo, useState } from "react";
import { useNavigate } from "react-router-dom";

import { useSpeechRecognition } from "../../hooks/useSpeechRecognition";
import { apiRequest } from "../../services/api";
import { ParseResult, Patient } from "../../types/api";
import { BillingSummaryCard } from "./BillingSummaryCard";
import { SectionCard } from "../ui/SectionCard";
import { StatusPill } from "../ui/StatusPill";

type Props = {
  patients: Patient[];
};

const emptyParseResult: ParseResult = {
  medications: [],
  lab_tests: [],
  clinical_notes: {
    observations: "",
    additional_notes: ""
  },
  billing: {
    items: [],
    grand_total: 0
  },
  metadata: {
    provider: "",
    model: "",
    confidence: "",
    recovery_applied: false,
    warnings: [],
    normalization_notes: []
  }
};

export function ConsultationWorkspace({ patients }: Props) {
  const navigate = useNavigate();
  const [patientId, setPatientId] = useState(patients[0]?.id ?? "");
  const [doctorName, setDoctorName] = useState("Dr. Maya Perera");
  const [rawInputText, setRawInputText] = useState(
    "Patient complains of fever, sore throat and body aches. Start Amoxicillin 500 mg three times daily for 5 days after meals. Add Paracetamol 500 mg twice daily for 3 days. Order Full Blood Count and CRP."
  );
  const [parseResult, setParseResult] = useState<ParseResult>(emptyParseResult);
  const [saving, setSaving] = useState(false);
  const [parsing, setParsing] = useState(false);
  const [message, setMessage] = useState("");

  const selectedPatient = useMemo(() => patients.find((patient) => patient.id === patientId), [patients, patientId]);
  const speech = useSpeechRecognition((transcript) => setRawInputText((current) => `${current} ${transcript}`.trim()));

  useEffect(() => {
    if (!patientId && patients.length > 0) {
      setPatientId(patients[0].id);
    }
  }, [patientId, patients]);

  async function handleParse() {
    setParsing(true);
    setMessage("");
    try {
      const parsed = await apiRequest<ParseResult>("/consultations/parse", {
        method: "POST",
        body: JSON.stringify({
          patient_id: patientId,
          doctor_name: doctorName,
          raw_input_text: rawInputText
        })
      });
      setParseResult(parsed);
      setMessage("AI parsing complete. Review and refine the structured data before saving.");
    } catch (error) {
      setMessage(error instanceof Error ? error.message : "Parse failed");
    } finally {
      setParsing(false);
    }
  }

  async function handleSave(status: "draft" | "finalized") {
    setSaving(true);
    setMessage("");
    try {
      const detail = await apiRequest<{ id: string }>("/consultations", {
        method: "POST",
        body: JSON.stringify({
          patient_id: patientId,
          doctor_name: doctorName,
          raw_input_text: rawInputText,
          status,
          parsed_result: parseResult
        })
      });
      navigate(`/consultations/${detail.id}`);
    } catch (error) {
      setMessage(error instanceof Error ? error.message : "Save failed");
    } finally {
      setSaving(false);
    }
  }

  function updateMedication(index: number, field: string, value: string) {
    const next = [...parseResult.medications];
    const record = next[index];
    next[index] = {
      ...record,
      [field]: field === "quantity" || field === "unit_price" || field === "line_total" ? Number(value) : value
    } as typeof record;
    setParseResult((current) => ({ ...current, medications: next }));
  }

  function updateLabTest(index: number, field: string, value: string) {
    const next = [...parseResult.lab_tests];
    next[index] = {
      ...next[index],
      [field]: field === "unit_price" || field === "line_total" ? Number(value) : value
    } as typeof next[number];
    setParseResult((current) => ({ ...current, lab_tests: next }));
  }

  return (
    <div className="grid gap-6 xl:grid-cols-[1.15fr_0.85fr]">
      <SectionCard title="Unified Consultation Entry" description="Capture the full consultation once, then let the parser separate medications, investigations, notes, and billable lines.">
        <div className="grid gap-4">
          <div className="grid gap-4 md:grid-cols-2">
            <label className="grid gap-2">
              <span className="text-sm font-semibold text-slate-600">Patient</span>
              <select className="rounded-2xl border border-slate-200 px-4 py-3" value={patientId} onChange={(event) => setPatientId(event.target.value)}>
                {patients.map((patient) => (
                  <option key={patient.id} value={patient.id}>
                    {patient.full_name}
                  </option>
                ))}
              </select>
            </label>
            <label className="grid gap-2">
              <span className="text-sm font-semibold text-slate-600">Doctor</span>
              <input className="rounded-2xl border border-slate-200 px-4 py-3" value={doctorName} onChange={(event) => setDoctorName(event.target.value)} />
            </label>
          </div>

          <label className="grid gap-2">
            <span className="text-sm font-semibold text-slate-600">Consultation note</span>
            <textarea className="min-h-[260px] rounded-[24px] border border-slate-200 px-4 py-4 text-base leading-7" value={rawInputText} onChange={(event) => setRawInputText(event.target.value)} />
          </label>

          <div className="flex flex-wrap gap-3">
            <button type="button" onClick={handleParse} disabled={parsing || !patientId} className="rounded-full bg-accent px-5 py-3 font-semibold text-white hover:bg-teal-700">
              {parsing ? "Parsing..." : "Parse with AI"}
            </button>
            {speech.supported ? (
              <button type="button" onClick={speech.listening ? speech.stop : speech.start} className="rounded-full border border-slate-300 px-5 py-3 font-semibold text-slate-700">
                {speech.listening ? "Stop dictation" : "Start dictation"}
              </button>
            ) : (
              <div className="rounded-full border border-dashed border-slate-300 px-5 py-3 text-sm text-slate-500">Speech input depends on browser SpeechRecognition support.</div>
            )}
          </div>

          {selectedPatient ? (
            <div className="rounded-2xl bg-mist p-4 text-sm text-slate-700">
              <p className="font-semibold text-ink">{selectedPatient.full_name}</p>
              <p>{selectedPatient.phone} · {selectedPatient.gender} · {selectedPatient.date_of_birth}</p>
              <p>{selectedPatient.address}</p>
            </div>
          ) : null}
        </div>
      </SectionCard>

      <div className="space-y-6">
        <SectionCard title="Extraction Review" description="Every extracted field remains editable before it is stored in PostgreSQL and exposed in print views.">
          <div className="space-y-5">
            <div className="flex items-center gap-3">
              <StatusPill tone={parseResult.metadata.confidence === "high" ? "teal" : parseResult.metadata.confidence === "medium" ? "amber" : "slate"}>
                {parseResult.metadata.confidence || "Not parsed"}
              </StatusPill>
              {parseResult.metadata.provider ? <span className="text-sm text-slate-500">{parseResult.metadata.provider} · {parseResult.metadata.model}</span> : null}
            </div>

            {message ? <div className="rounded-2xl bg-slate-100 px-4 py-3 text-sm text-slate-700">{message}</div> : null}

            <div className="space-y-3">
              <h3 className="font-display text-lg font-semibold text-ink">Medications</h3>
              {parseResult.medications.map((item, index) => (
                <div key={`${item.drug_name}-${index}`} className="grid gap-3 rounded-2xl border border-slate-200 p-4">
                  <div className="grid gap-3 md:grid-cols-2">
                    <input className="rounded-2xl border border-slate-200 px-4 py-3" value={item.drug_name} onChange={(event) => updateMedication(index, "drug_name", event.target.value)} />
                    <input className="rounded-2xl border border-slate-200 px-4 py-3" value={item.dosage} onChange={(event) => updateMedication(index, "dosage", event.target.value)} />
                    <input className="rounded-2xl border border-slate-200 px-4 py-3" value={item.frequency} onChange={(event) => updateMedication(index, "frequency", event.target.value)} />
                    <input className="rounded-2xl border border-slate-200 px-4 py-3" value={item.duration} onChange={(event) => updateMedication(index, "duration", event.target.value)} />
                    <input className="rounded-2xl border border-slate-200 px-4 py-3" value={item.route} onChange={(event) => updateMedication(index, "route", event.target.value)} />
                    <input className="rounded-2xl border border-slate-200 px-4 py-3" value={item.instructions} onChange={(event) => updateMedication(index, "instructions", event.target.value)} />
                  </div>
                </div>
              ))}
            </div>

            <div className="space-y-3">
              <h3 className="font-display text-lg font-semibold text-ink">Lab tests</h3>
              {parseResult.lab_tests.map((item, index) => (
                <div key={`${item.test_name}-${index}`} className="grid gap-3 rounded-2xl border border-slate-200 p-4 md:grid-cols-2">
                  <input className="rounded-2xl border border-slate-200 px-4 py-3" value={item.test_name} onChange={(event) => updateLabTest(index, "test_name", event.target.value)} />
                  <input className="rounded-2xl border border-slate-200 px-4 py-3" value={item.instructions} onChange={(event) => updateLabTest(index, "instructions", event.target.value)} />
                </div>
              ))}
            </div>

            <div className="space-y-3">
              <h3 className="font-display text-lg font-semibold text-ink">Clinical notes</h3>
              <textarea className="min-h-28 rounded-2xl border border-slate-200 px-4 py-3" value={parseResult.clinical_notes.observations} onChange={(event) => setParseResult((current) => ({ ...current, clinical_notes: { ...current.clinical_notes, observations: event.target.value } }))} />
              <textarea className="min-h-24 rounded-2xl border border-slate-200 px-4 py-3" value={parseResult.clinical_notes.additional_notes} onChange={(event) => setParseResult((current) => ({ ...current, clinical_notes: { ...current.clinical_notes, additional_notes: event.target.value } }))} />
            </div>

            <BillingSummaryCard billing={parseResult.billing} />

            <div className="flex flex-wrap gap-3">
              <button type="button" disabled={saving || parseResult.metadata.provider === ""} onClick={() => handleSave("draft")} className="rounded-full border border-slate-300 px-5 py-3 font-semibold text-slate-700">
                {saving ? "Saving..." : "Save draft"}
              </button>
              <button type="button" disabled={saving || parseResult.metadata.provider === ""} onClick={() => handleSave("finalized")} className="rounded-full bg-ink px-5 py-3 font-semibold text-white">
                Finalize consultation
              </button>
            </div>
          </div>
        </SectionCard>
      </div>
    </div>
  );
}

import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

import { PageHeader } from "../components/ui/PageHeader";
import { apiRequest } from "../services/api";
import { BillDocument, LabRequestDocument, NotesDocument, PrescriptionDocument } from "../types/documents";
import { formatCurrency } from "../utils/format";

type Mode = "prescription" | "lab-request" | "notes" | "bill";

export function PrintDocumentPage() {
  const params = useParams<{ id: string; mode: string }>();
  const id = params.id ?? "";
  const mode = (params.mode as Mode | undefined) ?? "prescription";
  const [data, setData] = useState<PrescriptionDocument | LabRequestDocument | NotesDocument | BillDocument | null>(null);

  useEffect(() => {
    apiRequest(`/consultations/${id}/${mode}`).then(setData).catch(console.error);
  }, [id, mode]);

  if (!data) {
    return <div className="rounded-[24px] bg-white p-8 shadow-soft">Loading document...</div>;
  }

  return (
    <div className="space-y-6">
      <PageHeader
        eyebrow="Print View"
        title={data.header.document_title}
        description={`${data.header.patient_name} · ${data.header.date}`}
        actions={
          <button className="rounded-full bg-ink px-5 py-3 font-semibold text-white" onClick={() => window.print()}>
            Print document
          </button>
        }
      />

      <article className="rounded-[28px] border border-slate-200 bg-white p-8 shadow-soft">
        <header className="border-b border-slate-200 pb-6">
          <p className="font-display text-sm uppercase tracking-[0.3em] text-accent">{data.header.clinic_name}</p>
          <h2 className="mt-2 font-display text-3xl font-semibold text-ink">{data.header.document_title}</h2>
          <div className="mt-4 grid gap-2 text-sm text-slate-600 md:grid-cols-2">
            <p>Patient: {data.header.patient_name}</p>
            <p>Patient ID: {data.header.patient_identifier}</p>
            <p>Doctor: {data.header.doctor_name}</p>
            <p>Date: {data.header.date}</p>
          </div>
        </header>

        {"medications" in data ? (
          <div className="mt-6 space-y-4">
            {data.medications.map((item) => (
              <div key={`${item.drug_name}-${item.dosage}`} className="rounded-2xl border border-slate-200 p-4">
                <p className="font-semibold text-ink">{item.drug_name}</p>
                <p className="text-slate-600">{item.dosage} · {item.frequency} · {item.duration}</p>
                <p className="text-sm text-slate-500">{item.route} · {item.instructions}</p>
              </div>
            ))}
          </div>
        ) : null}

        {"lab_tests" in data ? (
          <div className="mt-6 space-y-4">
            {data.lab_tests.map((item) => (
              <div key={item.test_name} className="rounded-2xl border border-slate-200 p-4">
                <p className="font-semibold text-ink">{item.test_name}</p>
                <p className="text-sm text-slate-500">{item.instructions}</p>
              </div>
            ))}
          </div>
        ) : null}

        {"clinical_notes" in data ? (
          <div className="mt-6 space-y-4">
            <p className="whitespace-pre-wrap text-slate-700">{data.clinical_notes.observations}</p>
            <p className="text-sm text-slate-500">{data.clinical_notes.additional_notes}</p>
          </div>
        ) : null}

        {"billing" in data ? (
          <div className="mt-6 space-y-4">
            {data.billing.items.map((item) => (
              <div key={`${item.item_type}-${item.item_name}`} className="flex items-center justify-between rounded-2xl border border-slate-200 p-4">
                <div>
                  <p className="font-semibold text-ink">{item.item_name}</p>
                  <p className="text-sm text-slate-500">{item.quantity} x {formatCurrency(item.unit_price)}</p>
                </div>
                <p className="font-semibold text-ink">{formatCurrency(item.line_total)}</p>
              </div>
            ))}
            <div className="flex items-center justify-between border-t border-slate-200 pt-4 text-lg font-semibold text-ink">
              <span>Grand total</span>
              <span>{formatCurrency(data.billing.grand_total)}</span>
            </div>
          </div>
        ) : null}
      </article>
    </div>
  );
}

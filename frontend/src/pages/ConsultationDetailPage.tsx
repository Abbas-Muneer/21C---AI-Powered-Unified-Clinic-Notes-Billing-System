import { Link, useParams } from "react-router-dom";
import { useEffect, useState } from "react";

import { PageHeader } from "../components/ui/PageHeader";
import { SectionCard } from "../components/ui/SectionCard";
import { BillingSummaryCard } from "../components/consultation/BillingSummaryCard";
import { apiRequest } from "../services/api";
import { ConsultationDetail } from "../types/api";
import { formatDate } from "../utils/format";

export function ConsultationDetailPage() {
  const { id = "" } = useParams();
  const [detail, setDetail] = useState<ConsultationDetail | null>(null);

  useEffect(() => {
    apiRequest<ConsultationDetail>(`/consultations/${id}`).then(setDetail).catch(console.error);
  }, [id]);

  if (!detail) {
    return <div className="rounded-[24px] bg-white p-8 shadow-soft">Loading consultation...</div>;
  }

  return (
    <div className="space-y-6">
      <PageHeader
        eyebrow="Consultation Detail"
        title={`${detail.patient.full_name} · ${detail.status}`}
        description={`Saved on ${formatDate(detail.consultation_date)} by ${detail.doctor_name}. Use the print routes below for assessment demo and output validation.`}
        actions={
          <div className="flex flex-wrap gap-2">
            <Link className="rounded-full bg-accent px-4 py-2 font-semibold text-white" to={`/consultations/${id}/prescription`}>
              Prescription
            </Link>
            <Link className="rounded-full bg-ink px-4 py-2 font-semibold text-white" to={`/consultations/${id}/lab-request`}>
              Lab request
            </Link>
            <Link className="rounded-full border border-slate-300 px-4 py-2 font-semibold text-slate-700" to={`/consultations/${id}/notes`}>
              Notes
            </Link>
            <Link className="rounded-full border border-slate-300 px-4 py-2 font-semibold text-slate-700" to={`/consultations/${id}/bill`}>
              Bill
            </Link>
          </div>
        }
      />

      <div className="grid gap-6 lg:grid-cols-[1.1fr_0.9fr]">
        <SectionCard title="Original consultation input">
          <p className="whitespace-pre-wrap text-slate-700">{detail.raw_input_text}</p>
        </SectionCard>

        <SectionCard title="Clinical notes">
          <p className="text-slate-700">{detail.parsed_result.clinical_notes.observations}</p>
          <p className="mt-4 text-sm text-slate-500">{detail.parsed_result.clinical_notes.additional_notes}</p>
        </SectionCard>
      </div>

      <div className="grid gap-6 lg:grid-cols-2">
        <SectionCard title="Prescription lines">
          <div className="space-y-3">
            {detail.parsed_result.medications.map((item) => (
              <div key={`${item.drug_name}-${item.dosage}`} className="rounded-2xl border border-slate-200 p-4">
                <p className="font-semibold text-ink">{item.drug_name}</p>
                <p className="text-sm text-slate-600">
                  {item.dosage} · {item.frequency} · {item.duration} · {item.route}
                </p>
                <p className="text-sm text-slate-500">{item.instructions}</p>
              </div>
            ))}
          </div>
        </SectionCard>

        <SectionCard title="Requested investigations">
          <div className="space-y-3">
            {detail.parsed_result.lab_tests.map((item) => (
              <div key={item.test_name} className="rounded-2xl border border-slate-200 p-4">
                <p className="font-semibold text-ink">{item.test_name}</p>
                <p className="text-sm text-slate-500">{item.instructions}</p>
              </div>
            ))}
          </div>
        </SectionCard>
      </div>

      <BillingSummaryCard billing={detail.parsed_result.billing} />
    </div>
  );
}

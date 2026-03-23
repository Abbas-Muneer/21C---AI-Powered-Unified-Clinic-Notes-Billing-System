import { Link } from "react-router-dom";

import { PageHeader } from "../components/ui/PageHeader";
import { SectionCard } from "../components/ui/SectionCard";

const highlights = [
  {
    title: "Unified input workflow",
    body: "Doctors capture symptoms, notes, prescriptions, and investigations once, then review structured extraction before saving."
  },
  {
    title: "AI parsing with fallback",
    body: "The backend supports a replaceable AI provider interface with OpenAI-compatible parsing and a deterministic demo parser."
  },
  {
    title: "Printable operational outputs",
    body: "Prescription, lab request, clinical notes, and bill views are all print-friendly routes backed by dedicated APIs."
  }
];

export function DashboardPage() {
  return (
    <div className="space-y-6">
      <PageHeader
        eyebrow="Clinic Workspace"
        title="One consultation screen for notes, prescriptions, lab requests, and billing."
        description="This assessment build optimizes the exact clinic workflow from mixed consultation input to structured storage, print outputs, and final bill generation."
        actions={
          <Link className="rounded-full bg-accent px-5 py-3 font-semibold text-white" to="/consultations/new">
            Start consultation
          </Link>
        }
      />

      <div className="grid gap-6 lg:grid-cols-3">
        {highlights.map((item) => (
          <SectionCard key={item.title} title={item.title}>
            <p className="text-slate-600">{item.body}</p>
          </SectionCard>
        ))}
      </div>
    </div>
  );
}

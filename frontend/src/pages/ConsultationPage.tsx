import { useEffect, useState } from "react";

import { ConsultationWorkspace } from "../components/consultation/ConsultationWorkspace";
import { PageHeader } from "../components/ui/PageHeader";
import { apiRequest } from "../services/api";
import { Patient } from "../types/api";

export function ConsultationPage() {
  const [patients, setPatients] = useState<Patient[]>([]);

  useEffect(() => {
    apiRequest<Patient[]>("/patients").then(setPatients).catch(console.error);
  }, []);

  return (
    <div className="space-y-6">
      <PageHeader
        eyebrow="Consultation"
        title="Unified consultation workspace"
        description="Enter one mixed note, let the parser classify medications, tests, and clinical observations, then finalize structured records and billable outputs."
      />
      <ConsultationWorkspace patients={patients} />
    </div>
  );
}

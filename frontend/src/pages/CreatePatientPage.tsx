import { useNavigate } from "react-router-dom";

import { PatientForm } from "../components/patients/PatientForm";
import { PageHeader } from "../components/ui/PageHeader";
import { SectionCard } from "../components/ui/SectionCard";
import { apiRequest } from "../services/api";

export function CreatePatientPage() {
  const navigate = useNavigate();

  return (
    <div className="space-y-6">
      <PageHeader eyebrow="Patients" title="Register patient" description="Create a patient profile with the core demographic details needed for consultations and printed documents." />
      <SectionCard title="Patient information">
        <PatientForm
          onSubmit={async (values) => {
            await apiRequest("/patients", {
              method: "POST",
              body: JSON.stringify(values)
            });
            navigate("/patients");
          }}
        />
      </SectionCard>
    </div>
  );
}

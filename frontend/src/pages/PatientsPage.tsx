import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

import { PageHeader } from "../components/ui/PageHeader";
import { SectionCard } from "../components/ui/SectionCard";
import { apiRequest } from "../services/api";
import { Patient } from "../types/api";

export function PatientsPage() {
  const [patients, setPatients] = useState<Patient[]>([]);

  useEffect(() => {
    apiRequest<Patient[]>("/patients").then(setPatients).catch(console.error);
  }, []);

  return (
    <div className="space-y-6">
      <PageHeader
        eyebrow="Patients"
        title="Patient directory"
        description="Use seed data immediately, or register new patients before starting a consultation."
        actions={
          <Link to="/patients/new" className="rounded-full bg-ink px-5 py-3 font-semibold text-white">
            Add patient
          </Link>
        }
      />

      <SectionCard title="Registered patients">
        <div className="overflow-hidden rounded-2xl border border-slate-200">
          <table className="min-w-full divide-y divide-slate-200">
            <thead className="bg-slate-50 text-left text-sm font-semibold text-slate-600">
              <tr>
                <th className="px-4 py-3">Name</th>
                <th className="px-4 py-3">Phone</th>
                <th className="px-4 py-3">Gender</th>
                <th className="px-4 py-3">DOB</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-slate-100 bg-white">
              {patients.map((patient) => (
                <tr key={patient.id}>
                  <td className="px-4 py-3 font-semibold text-ink">{patient.full_name}</td>
                  <td className="px-4 py-3 text-slate-600">{patient.phone}</td>
                  <td className="px-4 py-3 text-slate-600 capitalize">{patient.gender}</td>
                  <td className="px-4 py-3 text-slate-600">{patient.date_of_birth}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </SectionCard>
    </div>
  );
}

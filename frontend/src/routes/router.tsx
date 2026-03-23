import { createBrowserRouter } from "react-router-dom";

import { AppShell } from "../layouts/AppShell";
import { ConsultationDetailPage } from "../pages/ConsultationDetailPage";
import { ConsultationPage } from "../pages/ConsultationPage";
import { CreatePatientPage } from "../pages/CreatePatientPage";
import { DashboardPage } from "../pages/DashboardPage";
import { PatientsPage } from "../pages/PatientsPage";
import { PrintDocumentPage } from "../pages/PrintDocumentPage";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <AppShell />,
    children: [
      { index: true, element: <DashboardPage /> },
      { path: "patients", element: <PatientsPage /> },
      { path: "patients/new", element: <CreatePatientPage /> },
      { path: "consultations/new", element: <ConsultationPage /> },
      { path: "consultations/:id", element: <ConsultationDetailPage /> },
      { path: "consultations/:id/:mode", element: <PrintDocumentPage /> }
    ]
  }
]);

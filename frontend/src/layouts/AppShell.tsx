import { NavLink, Outlet } from "react-router-dom";

const navItems = [
  { to: "/", label: "Overview" },
  { to: "/patients", label: "Patients" },
  { to: "/patients/new", label: "New Patient" },
  { to: "/consultations/new", label: "New Consultation" }
];

export function AppShell() {
  return (
    <div className="min-h-screen px-4 py-6 sm:px-6 lg:px-8">
      <div className="mx-auto max-w-7xl">
        <div className="mb-6 flex flex-col gap-4 rounded-[28px] border border-white/60 bg-ink px-6 py-5 text-white shadow-soft sm:flex-row sm:items-center sm:justify-between">
          <div>
            <p className="font-display text-xs uppercase tracking-[0.35em] text-teal-200">Assessment Build</p>
            <h1 className="mt-2 font-display text-2xl font-semibold">AI-Powered Unified Clinic Notes & Billing System</h1>
          </div>
          <nav className="flex flex-wrap gap-2">
            {navItems.map((item) => (
              <NavLink
                key={item.to}
                to={item.to}
                className={({ isActive }) =>
                  `rounded-full px-4 py-2 text-sm font-semibold transition ${isActive ? "bg-white text-ink" : "bg-white/10 text-white hover:bg-white/20"}`
                }
              >
                {item.label}
              </NavLink>
            ))}
          </nav>
        </div>
        <Outlet />
      </div>
    </div>
  );
}

import { ReactNode } from "react";

type Props = {
  title: string;
  description?: string;
  children: ReactNode;
};

export function SectionCard({ title, description, children }: Props) {
  return (
    <section className="rounded-[24px] border border-slate-200 bg-white p-6 shadow-soft">
      <div className="mb-5">
        <h2 className="font-display text-xl font-semibold text-ink">{title}</h2>
        {description ? <p className="mt-1 text-sm text-slate-500">{description}</p> : null}
      </div>
      {children}
    </section>
  );
}

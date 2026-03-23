import { ReactNode } from "react";

type Props = {
  eyebrow: string;
  title: string;
  description: string;
  actions?: ReactNode;
};

export function PageHeader({ eyebrow, title, description, actions }: Props) {
  return (
    <div className="flex flex-col gap-4 rounded-[28px] border border-white/60 bg-white/75 p-8 shadow-soft backdrop-blur sm:flex-row sm:items-end sm:justify-between">
      <div className="max-w-2xl">
        <p className="font-display text-sm font-semibold uppercase tracking-[0.3em] text-accent">{eyebrow}</p>
        <h1 className="mt-3 font-display text-3xl font-semibold text-ink sm:text-4xl">{title}</h1>
        <p className="mt-3 text-lg text-slate-600">{description}</p>
      </div>
      {actions ? <div className="no-print">{actions}</div> : null}
    </div>
  );
}

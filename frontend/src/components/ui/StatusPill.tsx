type Props = {
  tone: "teal" | "amber" | "slate";
  children: string;
};

const toneClassMap = {
  teal: "bg-emerald-100 text-emerald-800",
  amber: "bg-amber-100 text-amber-800",
  slate: "bg-slate-100 text-slate-700"
};

export function StatusPill({ tone, children }: Props) {
  return <span className={`rounded-full px-3 py-1 text-xs font-semibold ${toneClassMap[tone]}`}>{children}</span>;
}

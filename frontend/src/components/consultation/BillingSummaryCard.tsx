import { BillingSummary } from "../../types/api";
import { formatCurrency } from "../../utils/format";

type Props = {
  billing: BillingSummary;
};

export function BillingSummaryCard({ billing }: Props) {
  return (
    <div className="rounded-[24px] border border-slate-200 bg-slate-50 p-5">
      <div className="mb-4 flex items-center justify-between">
        <h3 className="font-display text-lg font-semibold text-ink">Billing Summary</h3>
        <span className="rounded-full bg-ink px-3 py-1 text-sm font-semibold text-white">{formatCurrency(billing.grand_total)}</span>
      </div>
      <div className="space-y-3">
        {billing.items.map((item) => (
          <div key={`${item.item_type}-${item.item_name}`} className="flex items-center justify-between border-b border-dashed border-slate-200 pb-3 text-sm">
            <div>
              <p className="font-semibold text-slate-700">{item.item_name}</p>
              <p className="text-slate-500">
                {item.quantity} x {formatCurrency(item.unit_price)}
              </p>
            </div>
            <span className="font-semibold text-ink">{formatCurrency(item.line_total)}</span>
          </div>
        ))}
      </div>
    </div>
  );
}

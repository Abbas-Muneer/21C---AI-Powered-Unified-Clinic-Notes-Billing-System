import { describe, expect, it } from "vitest";
import { render, screen } from "@testing-library/react";

import { BillingSummaryCard } from "./BillingSummaryCard";

describe("BillingSummaryCard", () => {
  it("renders billing totals and line items", () => {
    render(
      <BillingSummaryCard
        billing={{
          grand_total: 5000,
          items: [
            { item_type: "drug", item_name: "Amoxicillin", quantity: 10, unit_price: 45, line_total: 450 },
            { item_type: "service", item_name: "Consultation Fee", quantity: 1, unit_price: 4550, line_total: 4550 }
          ]
        }}
      />
    );

    expect(screen.getByText("Amoxicillin")).toBeInTheDocument();
    expect(screen.getByText(/Consultation Fee/)).toBeInTheDocument();
    expect(screen.getByText(/LKR/)).toBeInTheDocument();
  });
});

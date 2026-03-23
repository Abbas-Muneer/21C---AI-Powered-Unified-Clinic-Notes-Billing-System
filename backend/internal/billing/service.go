package billing

import (
	"math"

	"clinicnotes/backend/internal/dto"
)

type Service struct {
	consultationFee float64
}

func NewService(consultationFee float64) *Service {
	return &Service{consultationFee: consultationFee}
}

type PricingResolver interface {
	FindDrugPrice(name string) (float64, bool)
	FindLabTestPrice(name string) (float64, bool)
}

func (s *Service) BuildSummary(parsed dto.ParseConsultationResponse, resolver PricingResolver) dto.BillingSummaryDTO {
	items := make([]dto.BillingItemDTO, 0, len(parsed.Medications)+len(parsed.LabTests)+1)
	total := 0.0

	for i := range parsed.Medications {
		item := &parsed.Medications[i]
		price, ok := resolver.FindDrugPrice(item.DrugName)
		if !ok {
			price = item.UnitPrice
		}
		if item.Quantity == 0 {
			item.Quantity = 1
		}
		item.UnitPrice = round(price)
		item.LineTotal = round(item.UnitPrice * item.Quantity)
		items = append(items, dto.BillingItemDTO{
			ItemType:  "drug",
			ItemName:  item.DrugName,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
			LineTotal: item.LineTotal,
		})
		total += item.LineTotal
	}

	for i := range parsed.LabTests {
		item := &parsed.LabTests[i]
		price, ok := resolver.FindLabTestPrice(item.TestName)
		if !ok {
			price = item.UnitPrice
		}
		if price == 0 {
			price = 1
		}
		item.UnitPrice = round(price)
		item.LineTotal = round(price)
		items = append(items, dto.BillingItemDTO{
			ItemType:  "lab",
			ItemName:  item.TestName,
			Quantity:  1,
			UnitPrice: item.UnitPrice,
			LineTotal: item.LineTotal,
		})
		total += item.LineTotal
	}

	items = append(items, dto.BillingItemDTO{
		ItemType:  "service",
		ItemName:  "Consultation Fee",
		Quantity:  1,
		UnitPrice: round(s.consultationFee),
		LineTotal: round(s.consultationFee),
	})
	total += s.consultationFee

	return dto.BillingSummaryDTO{
		Items:      items,
		GrandTotal: round(total),
	}
}

func round(value float64) float64 {
	return math.Round(value*100) / 100
}

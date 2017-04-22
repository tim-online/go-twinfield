package twinfield

import "github.com/tim-online/go-twinfield/soap"

type VatDeclaration struct {
	Summary  *soap.DeclarationSummary
	Aangifte *soap.AangifteOmzetbelasting
}

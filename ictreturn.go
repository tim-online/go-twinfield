package twinfield

import "github.com/tim-online/go-twinfield/soap"

type IctDeclaration struct {
	Summary   *soap.DeclarationSummary
	OpgaafICP *soap.OpgaafICP
}

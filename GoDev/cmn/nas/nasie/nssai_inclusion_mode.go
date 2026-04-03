package nasie

type NSSAIInclusionModeType struct {
	NSSAIInclusionMode NSSAIInclusionModeIE
}

type NSSAIInclusionModeIE byte

const (
	NSSAIInclusionModeA NSSAIInclusionModeIE = 0
	NSSAIInclusionModeB NSSAIInclusionModeIE = 1
	NSSAIInclusionModeC NSSAIInclusionModeIE = 2
	NSSAIInclusionModeD NSSAIInclusionModeIE = 3
)

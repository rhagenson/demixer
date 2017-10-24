package ami

import (
	"strings"

	"bitbucket.org/rhagenson/demixer/dna"
)

// NucProbs takes a Sequence and computes the Probability of each Nuc
func NucProbs(seq *dna.Sequence) map[dna.Nuc]Probability {
	prnucs := make(map[dna.Nuc]Probability)

	for _, v := range dna.ValidNucs {
		prnucs[v] = Probability(strings.Count((*seq).ToString(), string(v)) / len(*seq))
	}

	return prnucs
}

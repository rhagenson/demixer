package lzdistance

import (
	"bytes"

	"github.com/rhagenson/demixer/dna"
)

// NormSumDistance is method four of computing LZ-based distance between
// two sequences. ds*(S,Q) = ds(S,Q) / (1/2[c(SQ)+c(QS)])
func NormSumDistance(s, q dna.Sequence) Distance {
	if bytes.Equal(s.Bytes(), q.Bytes()) {
		return Distance(0)
	}

	cs, cq, csq, cqs := LZ76Factors(s, q)
	return NormSumDistanceFromFactors(cs, cq, csq, cqs)
}

// NormSumDistanceFromFactors implements method three of computing
// LZ-based distance given the corresponding factors
func NormSumDistanceFromFactors(cs, cq, csq, cqs int) Distance {
	return Distance(SumDistanceFromFactors(cs, cq, csq, cqs) /
		Distance(0.5*float64(csq+cqs)))
}

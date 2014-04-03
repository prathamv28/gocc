package parser

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : RR	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RR : A	<<  >>`,
		Id:         "RR",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RR : B	<<  >>`,
		Id:         "RR",
		NTType:     1,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `B : a	<< "B ", nil >>`,
		Id:         "B",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "B ", nil
		},
	},
	ProdTabEntry{
		String: `A : a	<< "A0 ", nil >>`,
		Id:         "A",
		NTType:     3,
		Index:      4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "A0 ", nil
		},
	},
	ProdTabEntry{
		String: `A : A a	<< "A1 ", nil >>`,
		Id:         "A",
		NTType:     3,
		Index:      5,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "A1 ", nil
		},
	},
	ProdTabEntry{
		String: `A : c	<< "A2 ", nil >>`,
		Id:         "A",
		NTType:     3,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "A2 ", nil
		},
	},
}

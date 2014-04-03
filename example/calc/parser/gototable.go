/*
 */
package parser

const numNTSymbols = 5

type (
	gotoTable [numStates]gotoRow
	gotoRow   [numNTSymbols]int
)

var gotoTab = gotoTable{
	gotoRow{ // S0

		-1, // S'
		1,  // Calc
		2,  // Expr
		3,  // Term
		4,  // Factor

	},
	gotoRow{ // S1

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S2

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S3

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S4

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S5

		-1, // S'
		-1, // Calc
		9,  // Expr
		10, // Term
		11, // Factor

	},
	gotoRow{ // S6

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S7

		-1, // S'
		-1, // Calc
		-1, // Expr
		14, // Term
		4,  // Factor

	},
	gotoRow{ // S8

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		15, // Factor

	},
	gotoRow{ // S9

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S10

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S11

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S12

		-1, // S'
		-1, // Calc
		19, // Expr
		10, // Term
		11, // Factor

	},
	gotoRow{ // S13

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S14

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S15

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S16

		-1, // S'
		-1, // Calc
		-1, // Expr
		20, // Term
		11, // Factor

	},
	gotoRow{ // S17

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S18

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		21, // Factor

	},
	gotoRow{ // S19

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S20

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S21

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S22

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
}

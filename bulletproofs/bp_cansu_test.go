package bulletproofs

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
)

func TestBP(t *testing.T) {
	fmt.Println("Hello World")
	// Set up the range, [18, 200) in this case.
	// We want to prove that we are over 18, and less than 200 years old.
	// This information is shared between the prover and the verifier.
	params, _ := SetupGeneric(18, 200)

	// Our secret age is 40
	bigSecret := new(big.Int).SetInt64(int64(40))

	// Create the zero-knowledge range proof
	proof, _ := ProveGeneric(bigSecret, params)

	// Encode the proof to JSON
	jsonEncoded, _ := json.Marshal(proof)

	// It this stage, the proof is passed to the verifier, possibly over a network.

	// Decode the proof from JSON
	var decodedProof ProofBPRP
	_ = json.Unmarshal(jsonEncoded, &decodedProof)

	// Verify the proof
	ok, _ := decodedProof.Verify()

	if ok == true {
		println("Age verified to be [18, 200)")
	}
}

package server

import (
	"testing"

	"github.com/SOMAS2020/SOMAS2020/internal/common/baseclient"
	"github.com/SOMAS2020/SOMAS2020/internal/common/gamestate"
	"github.com/SOMAS2020/SOMAS2020/internal/common/shared"
)

// Check that getPredictions() does not recieve predictions from dead islands
func TestGetPredictionsDead(t *testing.T) {
	client1 := baseclient.NewClient(1)
	testServer := &SOMASServer{
		gameState: gamestate.GameState{
			ClientInfos: map[shared.ClientID]gamestate.ClientInfo{
				1: gamestate.ClientInfo{
					LifeStatus: shared.Dead,
				},
			},
		},
		clientMap: map[shared.ClientID]baseclient.Client{
			1: client1,
		},
	}

	resultIIFO, err := testServer.getPredictions()
	if err != nil {
		t.Errorf("An error occurred: '%v'", err)
	}

	if value, ok := resultIIFO[1]; ok {
		t.Errorf("dead island prediction present: '%v'", value)
	}
}

// Check that getPredictions() obtains the correct predictions from alive or critical islands
func TestGetPredictionsLiving(t *testing.T) {
	client1 := baseclient.NewClient(1)
	client2 := baseclient.NewClient(2)
	testServer := &SOMASServer{
		gameState: gamestate.GameState{
			ClientInfos: map[shared.ClientID]gamestate.ClientInfo{
				1: gamestate.ClientInfo{
					LifeStatus: shared.Alive,
				},
				2: gamestate.ClientInfo{
					LifeStatus: shared.Critical,
				},
			},
		},
		clientMap: map[shared.ClientID]baseclient.Client{
			1: client1,
			2: client2,
		},
	}

	resultClient1, err := client1.MakePrediction()
	if err != nil {
		t.Errorf("An error occurred: '%v'", err)
	}
	resultClient2, err := client2.MakePrediction()
	if err != nil {
		t.Errorf("An error occurred: '%v'", err)
	}
	resultIIFO, err := testServer.getPredictions()
	if err != nil {
		t.Errorf("An error occurred: '%v'", err)
	}

	if (resultClient1.PredictionMade != resultIIFO[1].PredictionMade) || (!Equal(resultClient1.TeamsOfferedTo, resultIIFO[1].TeamsOfferedTo)) {
		t.Errorf("client1 prediction: '%v' Recieved in IIFO: '%v'", resultClient1, resultIIFO[1])
	}
	if (resultClient2.PredictionMade != resultIIFO[2].PredictionMade) || (!Equal(resultClient2.TeamsOfferedTo, resultIIFO[2].TeamsOfferedTo)) {
		t.Errorf("client2 prediction: '%v' Recieved in IIFO: '%v'", resultClient2, resultIIFO[2])
	}
}

func Equal(a, b []shared.ClientID) bool {
	if len(a) != len(b) {
		return false
	}
	for index, value := range a {
		if value != b[index] {
			return false
		}
	}
	return true
}

// Check if the predictions that islands recieve are correct by altering
// the islands that island1 trusts
/*
func TestDistributePredictions(t *testing.T) {
	cases := []struct {
		testName       string
		trustedIslands []shared.ClientID
		expectedOutput shared.RecievedPredictionsDict
	}{
		{
			testName:       "everyone is trustworthy",
			trustedIslands: []shared.ClientID{1, 2, 3},
			expectedOutput: shared.RecievedPredictionsDict{
				1: shared.PredictionInfoDict{
					2: {
						PredictionMade: shared.Prediction{2, 2, 2, 2, 2},
						TeamsOfferedTo: nil,
					},
					3: {
						PredictionMade: shared.Prediction{3, 3, 3, 3, 3},
						TeamsOfferedTo: nil,
					},
				},
				2: shared.PredictionInfoDict{
					1: {
						PredictionMade: shared.Prediction{1, 1, 1, 1, 1},
						TeamsOfferedTo: nil,
					},
					3: {
						PredictionMade: shared.Prediction{3, 3, 3, 3, 3},
						TeamsOfferedTo: nil,
					},
				},
				3: shared.PredictionInfoDict{
					1: {
						PredictionMade: shared.Prediction{1, 1, 1, 1, 1},
						TeamsOfferedTo: nil,
					},
					2: {
						PredictionMade: shared.Prediction{2, 2, 2, 2, 2},
						TeamsOfferedTo: nil,
					},
				},
			},
		},
		{
			testName:       "island 1 does not trust 2",
			trustedIslands: []shared.ClientID{1, 3},
			expectedOutput: shared.RecievedPredictionsDict{
				1: shared.PredictionInfoDict{
					2: {
						PredictionMade: shared.Prediction{2, 2, 2, 2, 2},
						TeamsOfferedTo: nil,
					},
					3: {
						PredictionMade: shared.Prediction{3, 3, 3, 3, 3},
						TeamsOfferedTo: nil,
					},
				},
				2: shared.PredictionInfoDict{
					3: {
						PredictionMade: shared.Prediction{3, 3, 3, 3, 3},
						TeamsOfferedTo: nil,
					},
				},
				3: shared.PredictionInfoDict{
					1: {
						PredictionMade: shared.Prediction{1, 1, 1, 1, 1},
						TeamsOfferedTo: nil,
					},
					2: {
						PredictionMade: shared.Prediction{2, 2, 2, 2, 2},
						TeamsOfferedTo: nil,
					},
				},
			},
		},
		{
			testName:       "island 1 trusts no one",
			trustedIslands: []shared.ClientID{1},
			expectedOutput: shared.RecievedPredictionsDict{
				1: shared.PredictionInfoDict{
					2: {
						PredictionMade: shared.Prediction{2, 2, 2, 2, 2},
						TeamsOfferedTo: nil,
					},
					3: {
						PredictionMade: shared.Prediction{3, 3, 3, 3, 3},
						TeamsOfferedTo: nil,
					},
				},
				2: shared.PredictionInfoDict{
					3: {
						PredictionMade: shared.Prediction{3, 3, 3, 3, 3},
						TeamsOfferedTo: nil,
					},
				},
				3: shared.PredictionInfoDict{
					2: {
						PredictionMade: shared.Prediction{2, 2, 2, 2, 2},
						TeamsOfferedTo: nil,
					},
				},
			},
		},
		{
			testName:       "check empty trusted islands slice",
			trustedIslands: []shared.ClientID{},
			expectedOutput: shared.RecievedPredictionsDict{
				1: shared.PredictionInfoDict{
					2: {
						PredictionMade: shared.Prediction{2, 2, 2, 2, 2},
						TeamsOfferedTo: nil,
					},
					3: {
						PredictionMade: shared.Prediction{3, 3, 3, 3, 3},
						TeamsOfferedTo: nil,
					},
				},
				2: shared.PredictionInfoDict{
					3: {
						PredictionMade: shared.Prediction{3, 3, 3, 3, 3},
						TeamsOfferedTo: nil,
					},
				},
				3: shared.PredictionInfoDict{
					2: {
						PredictionMade: shared.Prediction{2, 2, 2, 2, 2},
						TeamsOfferedTo: nil,
					},
				},
			},
		},
	}
}*/

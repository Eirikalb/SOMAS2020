package team2

import (
	"testing"

	"github.com/SOMAS2020/SOMAS2020/internal/common/baseclient"
	"github.com/SOMAS2020/SOMAS2020/internal/common/gamestate"
	"github.com/SOMAS2020/SOMAS2020/internal/common/shared"
)

// type mockClient struct {
// 	baseclient.BaseClient
// 	opinionHist OpinionHist
// 	giftHist    GiftHist
// 	gamestate.GameState
// }

// func (c *mockClient) GetGiftRequests() shared.GiftRequestDict {
// 	return c.requests
// }

// // Test function for GetGiftRequest()
// //check if initialisation works
// func CreateTestClient(gamestate gamestate.ClientGameState) client {
// 	c := NewClient(shared.Team2)

// 	// Instantiate the clientgamestate ?

// 	return *c.(*client)
// }

type mockClient struct {
	gameState gamestate.ClientGameState
	confidences map[shared.ClientID]int
	GiftHist map[shared.ClientID]GiftExchange 
}

func  determineAllocation(c *mockClient) shared.Resources {
	return 100
}

func  criticalStatus(c *mockClient) bool {
	return gameState.status
}

func  (c *mockClient) gameState() gamestate.ClientGameState {
	return gameState
}

func (c *mockClient) confidence(situation Situation, island shared.ClientID) int {
	return c.confidence[island]
}

func (c *mockClient) updateGiftConfidence(island shared.ClientID){ }

func TestGetGiftRequest(t *testing.T) {

	/*we need:
		turn := c.gameState().Turn
		ourAgentCritical := shared.Critical == shared.ClientLifeStatus(1)
		requestAmount := determineAllocation(c) * 0.6
		range c.gameState().ClientLifeStatuses

	*/
	// Mock a bunch of clients

	c := &mockClient{
		gameState: gameState,
		confidences: map[shared.ClientID]int[]{
			shared.Team1 : [50, 60, 45],
			shared.Team2 : [50, 70, 80],
			shared.Team3 : [50, 40, 30],
		},
		GiftHist : GiftHist,

	}
	gameState := ClientGameState {
		turn: 2,
		ClientInfo: ClientInfo {
			Resources: 100,
			LifeStatus: shared.Alive,
		},
		ClientLifeStatuses: map[shared.ClientID]gamestate.ClientInfo{
			shared.Team1: {
				LifeStatus: shared.Alive,
			},
			shared.Team2: {
				LifeStatus: shared.Critical,
			},
			shared.Team3: {
				LifeStatus: shared.Dead,
			},
		},
	}

	GiftHist := map[shared.ClientID]GiftExchange {
		shared.Team1: GiftExchange{
			IslandRequest: map[uint]GiftInfo{
				1: GiftInfo{requested: 100, gifted: 80},
				2: GiftInfo{requested: 100, gifted: 80},
				3: GiftInfo{requested: 100, gifted: 80},
			},
			OurRequest: map[uint]GiftInfo{
				1: GiftInfo{requested: 100, gifted: 80},
				2: GiftInfo{requested: 100, gifted: 80},
			}
		},
		shared.Team2: GiftExchange{
			IslandRequest: map[uint]GiftInfo{
				1: GiftInfo{requested: 100, gifted: 80},
				2: GiftInfo{requested: 100, gifted: 80},
				3: GiftInfo{requested: 100, gifted: 80},
			},
			OurRequest: map[uint]GiftInfo{
				1: GiftInfo{requested: 100, gifted: 80},
				2: GiftInfo{requested: 100, gifted: 80},
			}
		},
		shared.Team3: GiftExchange{
			IslandRequest: map[uint]GiftInfo{
				1: GiftInfo{requested: 100, gifted: 80},
				2: GiftInfo{requested: 100, gifted: 80},
				3: GiftInfo{requested: 100, gifted: 80},
			},
			OurRequest: map[uint]GiftInfo{
				1: GiftInfo{requested: 100, gifted: 80},
				2: GiftInfo{requested: 100, gifted: 80},
			}
		},
	}
	
	
	clientMap := map[shared.ClientID]baseclient.Client{
		// Team 1 makes 1 valid request: 50 to team 2.
		shared.Team1: &mockClientIITO{requests: shared.GiftRequestDict{shared.Team1: 50, shared.Team2: 50, shared.Team3: 50}},
		// Team 2 makes no valid requests: a zero'ed entry, one to itself and one to a dead team.
		shared.Team2: &mockClientIITO{requests: shared.GiftRequestDict{shared.Team1: 0, shared.Team2: 50, shared.Team3: 50}},
		// Team 3 is dead boi
		shared.Team3: &mockClientIITO{},
	}

	want := map[shared.ClientID]shared.GiftRequestDict{
		shared.Team1: {shared.Team2: 50},
	}

	opinionHist := map[shared.ClientID]Opinion{
		shared.Team1: &Opinion{
			Histories: map[Situation][]int{
				"Gifts": [50, 60, 55, 70],
			},
			Performances: map[Situation]ExpectationReality{
				"Gifts": ExpectationReality{exp: 60,},
			},
		}
		shared.Team2: &Opinion{
			Histories: map[Situation][]int{
				"Gifts": [50, 60, 55, 70],
			},
			Performances: map[Situation]ExpectationReality{
				"Gifts": ExpectationReality{exp: 50,},
			},
		}
		shared.Team3: &Opinion{
			Histories: map[Situation][]int{
				"Gifts": [50, 60, 55, 70],
			},
			Performances: map[Situation]ExpectationReality{
				"Gifts": ExpectationReality{exp: 40,},
			},
		}
	}

	c := &baseclient.BaseClient{
		gameState : gamestate.GameState{
			ClientInfos: clientInfos,
			turn: turn,
		},
		opinionHist: opinionHist,

	}

}

// Test function for GetGiftOffers()
func TestGetGiftOffers(t *testing.T) {

	c := MakeTestClient(gamestate.ClientGameState{
		ClientInfo: gamestate.ClientInfo{
			LifeStatus: shared.Critical,
		},
	})

	clientInfos := map[shared.ClientID]gamestate.ClientInfo{
		shared.Team1: {
			LifeStatus: shared.Alive,
		},
		shared.Team2: {
			LifeStatus: shared.Critical,
		},
		shared.Team3: {
			LifeStatus: shared.Dead,
		},
		shared.Team4: {
			LifeStatus: shared.Alive,
		},
	}

	// should be a GiftRequestDict
	testInput := GiftRequestDict{
		GiftRequest{ClientID: shared.Team1, Resources: 100},
		GiftRequest{ClientID: shared.Team3, Resources: 0},
		GiftRequest{ClientID: shared.Team4, Resources: 100},
	}

	// GiftOffer contains the details of a gift offer from an island to another
	// type GiftOffer Resources
	// GiftOfferDict contains the details of an island's gift offers to everyone else.
	// type GiftOfferDict map[ClientID]GiftOffer

	// should be a GiftOfferDict
	expOutput := GiftOfferDict{
		GiftOffer{ClientID: shared.Team1, Resources: 100},
		GiftOffer{ClientID: shared.Team3, Resources: 100},
		GiftOffer{ClientID: shared.Team4, Resources: 100},
		GiftOffer{ClientID: shared.Team5, Resources: 100},
		GiftOffer{ClientID: shared.Team6, Resources: 100},
	}

	actualOutput = c.GetGiftOffers(testInput)

	if output != expOutput {
		t.Errorf("expected '%v' got '%v'", expected, actual)
	}
}

// Test function for GetGiftResponse()
func TestGetGiftResponses(t *testing.T) {

}

// Test function for UpdateGiftInfo()
func TestUpdateGiftInfo(t *testing.T) {

}

// Test function for SentGift()
func TestSentGift(t *testing.T) {

}

// Test function for ReceivedGift()
func TestReceivedGift(t *testing.T) {

}

// Test function for DecidedGiftAmount()
func TestDecideGiftAmount(t *testing.T) {

}

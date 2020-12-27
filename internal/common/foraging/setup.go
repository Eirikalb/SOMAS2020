package foraging

import (
	"github.com/SOMAS2020/SOMAS2020/internal/common/config"
	"github.com/SOMAS2020/SOMAS2020/internal/common/shared"
)

// CreateDeerHunt receives hunt participants and their contributions and returns a DeerHunt
func CreateDeerHunt(teamResourceInputs map[shared.ClientID]shared.Resources) (DeerHunt, error) {
	fConf := config.GameConfig().ForagingConfig
	params := deerHuntParams{p: fConf.BernoulliProb, lam: fConf.ExponentialRate}
	return DeerHunt{ParticipantContributions: teamResourceInputs, params: params}, nil // returning error too for future use
}

// CreateDeerPopulationModel returns the target population model. The formulation of this model should be changed here before runtime
func CreateDeerPopulationModel() DeerPopulationModel {
	return createBasicDeerPopulationModel()
}

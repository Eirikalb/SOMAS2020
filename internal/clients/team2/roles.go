package team2

import (
	"sort"

	"github.com/SOMAS2020/SOMAS2020/internal/common/baseclient"
	"github.com/SOMAS2020/SOMAS2020/internal/common/roles"
	"github.com/SOMAS2020/SOMAS2020/internal/common/shared"
)

func (c *client) GetClientSpeakerPointer() roles.Speaker {
	// return &c.currSpeaker
	return &Speaker{c: c, BaseSpeaker: &baseclient.BaseSpeaker{GameState: c.ServerReadHandle.GetGameState()}}
}

func (c *client) GetClientJudgePointer() roles.Judge {
	// return &c.currJudge
	return &Judge{c: c, BaseJudge: &baseclient.BaseJudge{GameState: c.ServerReadHandle.GetGameState()}}
}

func (c *client) GetClientPresidentPointer() roles.President {
	// return &c.currPresident
	return &President{c: c, BasePresident: &baseclient.BasePresident{GameState: c.ServerReadHandle.GetGameState()}}
}

func (c *client) VoteForElection(roleToElect shared.Role, candidateList []shared.ClientID) []shared.ClientID {
	var situation Situation
	switch roleToElect {
	case shared.President:
		situation = "President"
	case shared.Judge:
		situation = "Judge"
	default:
		situation = "Gifts"
	}

	var trustRank IslandTrustList
	for _, candidate := range candidateList {
		islandConf := IslandTrust{
			island: candidate,
			trust:  c.confidence(situation, candidate),
		}
		trustRank = append(trustRank, islandConf)
	}

	sort.Sort(trustRank)
	bordaList := make([]shared.ClientID, 0)

	for _, islandTrust := range trustRank {
		bordaList = append(bordaList, islandTrust.island)
	}

	return bordaList
}

//MonitorIIGORole decides whether to perform monitoring on a role
//COMPULOSRY: must be implemented
func (c *client) MonitorIIGORole(roleName shared.Role) bool {
	return false
}

//DecideIIGOMonitoringAnnouncement decides whether to share the result of monitoring a role and what result to share
//COMPULSORY: must be implemented
func (c *client) DecideIIGOMonitoringAnnouncement(monitoringResult bool) (resultToShare bool, announce bool) {
	resultToShare = false
	announce = false
	return
}

func (c *client) ResourceReport() shared.ResourcesReport {
	mood := c.MethodOfPlay()
	switch mood {
	case 2: // Free Rider
		return shared.ResourcesReport{
			ReportedAmount: 0.5 * c.gameState().ClientInfo.Resources,
			Reported:       true,
		}
	default: // Fair or Altruist
		return shared.ResourcesReport{
			ReportedAmount: c.gameState().ClientInfo.Resources,
			Reported:       true,
		}
	}
}
package api

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

func getDelegationResponse(p blockatlas.StakeAPI, address string) (blockatlas.DelegationResponse, error) {
	delegations, err := p.GetDelegations(address)
	if err != nil {
		return blockatlas.DelegationResponse{
			StakingResponse: getStakingResponse(p),
		}, errors.E("Unable to fetch delegations list", err)
	}
	balance, err := p.UndelegatedBalance(address)
	if err != nil {
		return blockatlas.DelegationResponse{
			StakingResponse: getStakingResponse(p),
		}, errors.E("Unable to fetch undelegated balance", err)
	}
	return blockatlas.DelegationResponse{
		Balance:         balance,
		Delegations:     delegations,
		Address:         address,
		StakingResponse: getStakingResponse(p),
	}, nil
}

func getStakingResponse(p blockatlas.StakeAPI) blockatlas.StakingResponse {
	c := p.Coin()
	return blockatlas.StakingResponse{
		Coin:    c.External(),
		Details: p.GetDetails(),
	}
}

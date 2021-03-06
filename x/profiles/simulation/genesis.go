package simulation

// DONTCOVER

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

// RandomizedGenState generates a random GenesisState for profile
func RandomizedGenState(simsState *module.SimulationState) {
	userRelationshipsMap := randomRelationships(simsState)

	profileGenesis := types.NewGenesisState(
		randomProfiles(simsState),
		types.NewParams(RandomMonikerParams(simsState.Rand), RandomDTagParams(simsState.Rand), RandomBioParams(simsState.Rand)),
		userRelationshipsMap,
	)

	fmt.Printf("Selected randomly generated profile parameters:\n%s\n%s\n%s\n",
		codec.MustMarshalJSONIndent(simsState.Cdc, profileGenesis.Params.MonikerParams),
		codec.MustMarshalJSONIndent(simsState.Cdc, profileGenesis.Params.DtagParams),
		codec.MustMarshalJSONIndent(simsState.Cdc, profileGenesis.Params.MaxBioLen),
	)

	simsState.GenState[types.ModuleName] = simsState.Cdc.MustMarshalJSON(profileGenesis)
}

// randomProfiles returns randomly generated genesis accounts
func randomProfiles(simState *module.SimulationState) (accounts types.Profiles) {
	accountsNumber := simState.Rand.Intn(50)

	accounts = make(types.Profiles, accountsNumber)
	for i := 0; i < accountsNumber; i++ {
		simAccount, _ := sim.RandomAcc(simState.Rand, simState.Accounts)
		accounts[i] = NewRandomProfile(simState.Rand, simAccount.Address)
	}

	return accounts
}

// randomRelationships returns randomly generated genesis relationships and their associated users - IDs map
func randomRelationships(simState *module.SimulationState) map[string][]sdk.AccAddress {
	relationshipsNumber := simState.Rand.Intn(sim.RandIntBetween(simState.Rand, 1, 100))
	usersRelationships := map[string][]sdk.AccAddress{}

	for index := 0; index < relationshipsNumber; index++ {
		sender, _ := sim.RandomAcc(simState.Rand, simState.Accounts)
		receiver, _ := sim.RandomAcc(simState.Rand, simState.Accounts)
		if !sender.Equals(receiver) {
			usersRelationships[sender.Address.String()] = []sdk.AccAddress{receiver.Address}
		}
	}

	return usersRelationships
}

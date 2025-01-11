package librarychain

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"librarychain/testutil/sample"
	librarychainsimulation "librarychain/x/librarychain/simulation"
	"librarychain/x/librarychain/types"
)

// avoid unused import issue
var (
	_ = librarychainsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateBook = "op_weight_msg_book"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateBook int = 100

	opWeightMsgUpdateBook = "op_weight_msg_book"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateBook int = 100

	opWeightMsgDeleteBook = "op_weight_msg_book"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteBook int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	librarychainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		BookList: []types.Book{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		BookCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&librarychainGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateBook int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateBook, &weightMsgCreateBook, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBook = defaultWeightMsgCreateBook
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateBook,
		librarychainsimulation.SimulateMsgCreateBook(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateBook int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateBook, &weightMsgUpdateBook, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateBook = defaultWeightMsgUpdateBook
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateBook,
		librarychainsimulation.SimulateMsgUpdateBook(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteBook int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteBook, &weightMsgDeleteBook, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteBook = defaultWeightMsgDeleteBook
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteBook,
		librarychainsimulation.SimulateMsgDeleteBook(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateBook,
			defaultWeightMsgCreateBook,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				librarychainsimulation.SimulateMsgCreateBook(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateBook,
			defaultWeightMsgUpdateBook,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				librarychainsimulation.SimulateMsgUpdateBook(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteBook,
			defaultWeightMsgDeleteBook,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				librarychainsimulation.SimulateMsgDeleteBook(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

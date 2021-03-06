package cmds

import (
	"github.com/spikeekips/mitum/launch"
	"github.com/spikeekips/mitum/util/hint"

	"github.com/soonkuk/mitum-blocksign/blocksign"
	"github.com/soonkuk/mitum-blocksign/digest"
	"github.com/spikeekips/mitum-currency/currency"
)

var (
	Hinters []hint.Hinter
	Types   []hint.Type
)

var types = []hint.Type{
	currency.KeyType,
	currency.KeysType,
	currency.NilFeeerType,
	currency.FixedFeeerType,
	currency.RatioFeeerType,
	currency.TransfersFactType,
	currency.TransfersType,
	currency.AccountType,
	currency.AmountStateType,
	currency.GenesisCurrenciesFactType,
	currency.GenesisCurrenciesType,
	currency.AmountType,
	currency.FeeOperationFactType,
	currency.FeeOperationType,
	currency.CurrencyDesignType,
	currency.CurrencyRegisterFactType,
	currency.CurrencyRegisterType,
	currency.CurrencyPolicyUpdaterFactType,
	currency.CurrencyPolicyUpdaterType,
	currency.CreateAccountsFactType,
	currency.CreateAccountsType,
	currency.CreateAccountsItemSingleAmountType,
	currency.TransfersItemMultiAmountsType,
	currency.CurrencyPolicyType,
	currency.AddressType,
	currency.CreateAccountsItemMultiAmountsType,
	currency.TransfersItemSingleAmountType,
	currency.KeyUpdaterFactType,
	currency.KeyUpdaterType,
	blocksign.CreateDocumentsItemSingleFileType,
	blocksign.CreateDocumentsFactType,
	blocksign.CreateDocumentsType,
	blocksign.SignItemSingleDocumentType,
	blocksign.SignDocumentsFactType,
	blocksign.SignDocumentsType,
	blocksign.DocumentDataType,
	blocksign.DocInfoType,
	blocksign.DocSignType,
	blocksign.DocumentInventoryType,
	digest.ProblemType,
	digest.NodeInfoType,
	digest.BaseHalType,
	digest.AccountValueType,
	digest.DocumentValueType,
	digest.OperationValueType,
}

var hinters = []hint.Hinter{
	currency.Account{},
	currency.Address(""),
	currency.AmountState{},
	currency.Amount{},
	currency.CreateAccountsFact{},
	currency.CreateAccountsItemMultiAmountsHinter,
	currency.CreateAccountsItemSingleAmountHinter,
	currency.CreateAccounts{},
	currency.CurrencyDesign{},
	currency.CurrencyPolicyUpdaterFact{},
	currency.CurrencyPolicyUpdater{},
	currency.CurrencyPolicy{},
	currency.CurrencyRegisterFact{},
	currency.CurrencyRegister{},
	currency.FeeOperationFact{},
	currency.FeeOperation{},
	currency.FixedFeeer{},
	currency.GenesisCurrenciesFact{},
	currency.GenesisCurrencies{},
	currency.KeyUpdaterFact{},
	currency.KeyUpdater{},
	currency.Keys{},
	currency.Key{},
	currency.NilFeeer{},
	currency.RatioFeeer{},
	currency.TransfersFact{},
	currency.TransfersItemMultiAmountsHinter,
	currency.TransfersItemSingleAmountHinter,
	currency.Transfers{},
	blocksign.CreateDocumentsFact{},
	blocksign.CreateDocumentsItemSingleFileHinter,
	blocksign.CreateDocuments{},
	blocksign.SignDocumentsFact{},
	blocksign.SignDocuments{},
	blocksign.SignItemSingleDocumentHinter,
	blocksign.DocumentData{},
	blocksign.DocInfo{},
	blocksign.DocSign{},
	blocksign.DocumentInventory{},
	digest.AccountValue{},
	digest.DocumentValue{},
	digest.BaseHal{},
	digest.NodeInfo{},
	digest.OperationValue{},
	digest.Problem{},
}

func init() {
	Hinters = make([]hint.Hinter, len(launch.EncoderHinters)+len(hinters))
	copy(Hinters, launch.EncoderHinters)
	copy(Hinters[len(launch.EncoderHinters):], hinters)

	Types = make([]hint.Type, len(launch.EncoderTypes)+len(types))
	copy(Types, launch.EncoderTypes)
	copy(Types[len(launch.EncoderTypes):], types)
}

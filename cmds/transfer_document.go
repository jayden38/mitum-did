package cmds

import (
	"github.com/pkg/errors"

	"github.com/spikeekips/mitum/base"
	"github.com/spikeekips/mitum/base/operation"
	"github.com/spikeekips/mitum/util"

	"github.com/soonkuk/mitum-blocksign/blocksign"
	currencycmds "github.com/spikeekips/mitum-currency/cmds"
)

type TransferDocumentCommand struct {
	*BaseCommand
	currencycmds.OperationFlags
	Sender   currencycmds.AddressFlag    `arg:"" name:"sender" help:"sender address" required:""`
	Currency currencycmds.CurrencyIDFlag `arg:"" name:"currency" help:"currency id" required:""`
	DocId    currencycmds.BigFlag        `arg:"" name:"documentid" help:"document id" required:""`
	Receiver currencycmds.AddressFlag    `arg:"" name:"reciever" help:"reciever address" required:""`
	Seal     currencycmds.FileLoad       `help:"seal" optional:""`
	sender   base.Address
	receiver base.Address
}

func NewTransferDocumentCommand() TransferDocumentCommand {
	return TransferDocumentCommand{
		BaseCommand: NewBaseCommand("transfer-document-operation"),
	}
}

func (cmd *TransferDocumentCommand) Run(version util.Version) error { // nolint:dupl
	if err := cmd.Initialize(cmd, version); err != nil {
		return errors.Errorf("failed to initialize command: %w", err)
	}

	if err := cmd.parseFlags(); err != nil {
		return err
	}

	var op operation.Operation
	if o, err := cmd.createOperation(); err != nil {
		return err
	} else {
		op = o
	}

	if sl, err := loadSealAndAddOperation(
		cmd.Seal.Bytes(),
		cmd.Privatekey,
		cmd.NetworkID.NetworkID(),
		op,
	); err != nil {
		return err
	} else {
		currencycmds.PrettyPrint(cmd.Out, cmd.Pretty, sl)
	}

	return nil
}

func (cmd *TransferDocumentCommand) parseFlags() error {
	if err := cmd.OperationFlags.IsValid(nil); err != nil {
		return err
	}

	if a, err := cmd.Sender.Encode(jenc); err != nil {
		return errors.Errorf("invalid sender format, %q: %w", cmd.Sender.String(), err)
	} else {
		cmd.sender = a
	}
	if a, err := cmd.Receiver.Encode(jenc); err != nil {
		return errors.Errorf("invalid receiver format, %q: %w", cmd.Receiver.String(), err)
	} else {
		cmd.receiver = a
	}

	return nil
}

func (cmd *TransferDocumentCommand) createOperation() (operation.Operation, error) { // nolint:dupl
	var items []blocksign.TransferDocumentsItem
	if i, err := loadOperations(cmd.Seal.Bytes(), cmd.NetworkID.NetworkID()); err != nil {
		return nil, err
	} else {
		for j := range i {
			if t, ok := i[j].(blocksign.TransferDocuments); ok {
				items = t.Fact().(blocksign.TransferDocumentsFact).Items()
			}
		}
	}

	item := blocksign.NewTransferDocumentsItemSingleFile(cmd.DocId.Big, cmd.sender, cmd.receiver, cmd.Currency.CID)

	if err := item.IsValid(nil); err != nil {
		return nil, err
	} else {
		items = append(items, item)
	}

	fact := blocksign.NewTransferDocumentsFact([]byte(cmd.Token), cmd.sender, items)

	var fs []operation.FactSign
	if sig, err := operation.NewFactSignature(cmd.Privatekey, fact, cmd.NetworkID.NetworkID()); err != nil {
		return nil, err
	} else {
		fs = append(fs, operation.NewBaseFactSign(cmd.Privatekey.Publickey(), sig))
	}

	if op, err := blocksign.NewTransferDocuments(fact, fs, cmd.Memo); err != nil {
		return nil, errors.Errorf("failed to create transfer-document operation: %w", err)
	} else {
		return op, nil
	}
}
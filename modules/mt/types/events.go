package types

// MT module event types
var (
	EventTypeIssueDenom    = "issue_denom"
	EventTypeTransfer      = "transfer_mt"
	EventTypeEditMT        = "edit_mt"
	EventTypeMintMT        = "mint_mt"
	EventTypeBurnMT        = "burn_mt"
	EventTypeTransferDenom = "transfer_denom"

	AttributeValueCategory = ModuleName

	AttributeKeySender    = "sender"
	AttributeKeyCreator   = "creator"
	AttributeKeyRecipient = "recipient"
	AttributeKeyOwner     = "owner"
	AttributeKeyTokenID   = "token_id"
	AttributeKeyDenomID   = "denom_id"
	AttributeKeySupply    = "supply"
	AttributeKeyDenomName = "denom_name"
)

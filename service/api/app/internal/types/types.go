// Code generated by goctl. DO NOT EDIT.
package types

type ReqGetStatus struct {
}

type RespGetStatus struct {
	Status    uint32 `json:"status"`
	NetworkId uint32 `json:"network_id"`
}

type ReqGetAccountStatusByPubKey struct {
	AccountPk string `form:"account_pk"`
}

type RespGetAccountStatusByPubKey struct {
	AccountStatus uint32 `json:"account_status"`
	AccountName   string `json:"account_name"`
	AccountIndex  uint32 `json:"account_index"`
}

type ReqGetAccountStatusByAccountName struct {
	AccountName string `form:"account_name"`
}

type RespGetAccountStatusByAccountName struct {
	AccountStatus uint32 `json:"account_status"`
	AccountIndex  uint32 `json:"account_index"`
	AccountPk     string `json:"account_pk"`
}

type Asset struct {
	AssetId uint32 `json:"asset_id"`
	Balance string `json:"balance"`
}

type ReqGetAccountInfoByAccountName struct {
	AccountName string `form:"account_name"`
}

type RespGetAccountInfoByAccountName struct {
	AccountIndex  uint32   `json:"account_index"`
	AccountPk     string   `json:"account_pk"`
	AssetsAccount []*Asset `json:"assets"`
}

type ReqGetBlanceByAssetIdAndAccountName struct {
	AssetId     uint32 `form:"asset_id"`
	AccountName string `form:"account_name"`
}

type RespGetBlanceInfoByAssetIdAndAccountName struct {
	Balance string `json:"balance_enc"`
}

type ReqGetAssetsByAccountName struct {
	AccountName string `form:"account_name"`
}

type RespGetAssetsByAccountName struct {
	Assets []*Asset `json:"assets"`
}

type ReqGetAccountLiquidityPairsByAccountIndex struct {
	AccountIndex uint32 `form:"account_index"`
}

type AccountLiquidityPairs struct {
	PairIndex   uint32 `json:"pair_index"`
	AssetAId    uint32 `json:"asset_a_id"`
	AssetAName  string `json:"asset_a_name"`
	AssetBId    uint32 `json:"asset_b_id"`
	AssetBName  string `json:"asset_b_name"`
	LpAmountEnc string `json:"lp_amount_enc"`
	CreatedAt   int64  `json:"created_at"`
}

type RespGetAccountLiquidityPairsByAccountIndex struct {
	Pairs []*AccountLiquidityPairs `json:"pairs"`
}

type ReqGetAssetsList struct {
}

type AssetInfo struct {
	AssetId       uint32 `json:"asset_id"`
	AssetName     string `json:"asset_name"`
	AssetDecimals uint32 `json:"asset_decimals"`
	AssetSymbol   string `json:"asset_symbol"`
	AssetAddress  string `json:"asset_address"`
}

type RespGetAssetsList struct {
	Assets []*AssetInfo `json:"assets"`
}

type ReqGetCurrencyPriceBySymbol struct {
	Symbol string `form:"symbol"`
}

type RespGetCurrencyPriceBySymbol struct {
	AssetId uint32  `json:"assetId"`
	Price   float64 `json:"price"`
}

type ReqGetCurrencyPrices struct {
}

type DataCurrencyPrices struct {
	Pair    string  `json:"pair"`
	AssetId uint32  `json:"assetId"`
	Price   float64 `json:"price"`
}

type RespGetCurrencyPrices struct {
	Data []*DataCurrencyPrices `json:"data"`
}

type ReqGetGasFee struct {
	AssetId uint32 `form:"asset_id"`
}

type RespGetGasFee struct {
	GasFee float64 `json:"gas_fee"`
}

type ReqGetWithdrawGasFee struct {
	AssetId         uint32 `form:"asset_id"`
	WithdrawAssetId uint32 `form:"withdraw_asset_id"`
	WithdrawAmount  uint64 `form:"withdraw_amount"`
}

type RespGetWithdrawGasFee struct {
	WithdrawGasFee float64 `json:"withdraw_gas_fee"`
}

type ReqGetSwapAmount struct {
	PairIndex   uint32 `form:"pair_index"`
	AssetId     uint32 `form:"asset_id"`
	AssetAmount string `form:"asset_amount"`
	IsFrom      bool   `form:"is_from"`
}

type RespGetSwapAmount struct {
	ResAssetAmount string `json:"res_asset_amount"`
	ResAssetId     uint32 `json:"res_asset_id"`
}

type ReqGetAvailablePairs struct {
}

type Pair struct {
	PairIndex    uint32 `json:"pair_index"`
	AssetAId     uint32 `json:"asset_a_id"`
	AssetAName   string `json:"asset_a_name"`
	AssetBId     uint32 `json:"asset_b_id"`
	AssetBName   string `json:"asset_b_name"`
	FeeRate      int64  `json:"fee_Rate"`
	TreasuryRate int64  `json:"treasury_rate"`
}

type RespGetAvailablePairs struct {
	Pairs []*Pair `json:"result"`
}

type ReqGetLPValue struct {
	PairIndex uint32 `form:"pair_index"`
	LpAmount  string `form:"lp_amount"`
}

type RespGetLPValue struct {
	AssetAId     uint32 `json:"asset_a_id"`
	AssetAName   string `json:"asset_a_name"`
	AssetAAmount string `json:"asset_a_amount"`
	AssetBid     uint32 `json:"asset_b_id"`
	AssetBName   string `json:"asset_b_name"`
	AssetBAmount string `json:"asset_b_amount"`
}

type ReqGetPairInfo struct {
	PairIndex uint32 `form:"pair_index"`
}

type RespGetPairInfo struct {
	AssetAId      uint32 `json:"asset_a_id"`
	AssetAAmount  string `json:"asset_a_amount"`
	AssetBId      uint32 `json:"asset_b_id"`
	AssetBAmount  string `json:"asset_b_amount"`
	TotalLpAmount string `json:"total_lp_amount"`
}

type TxDetail struct {
	AssetId        uint32 `json:"assetId"`
	AssetType      uint32 `json:"assetType"`
	AccountIndex   int32  `json:"accountIndex"`
	AccountName    string `json:"accountName"`
	AccountBalance string `json:"accountBalance"`
	AccountDelta   string `json:"accountDelta"`
}

type Tx struct {
	TxHash        string      `json:"tx_hash"`
	TxType        uint32      `json:"tx_type,range=[1:64]"`
	GasFeeAssetId uint32      `json:"gas_fee_asset_id"`
	GasFee        string      `json:"gas_fee"`
	NftIndex      uint32      `json:"nft_index"`
	PairIndex     uint32      `json:"pair_index"`
	AssetId       uint32      `json:"asset_id"`
	TxAmount      string      `json:"tx_amount"`
	NativeAddress string      `json:"native_adress"`
	TxDetails     []*TxDetail `json:"tx_details"`
	TxInfo        string      `json:"tx_info"`
	ExtraInfo     string      `json:"extra_info"`
	Memo          string      `json:"memo"`
	AccountIndex  uint32      `json:"account_index"`
	Nonce         uint32      `json:"nonce"`
	ExpiredAt     uint32      `json:"expire_at"`
	L2BlockHeight uint32      `json:"l2_block_height"`
	Status        uint32      `json:"status,options=0|1|2"`
	CreatedAt     uint32      `json:"created_at"`
	BlockID       uint32      `json:"block_id"`
}

type TxAccount struct {
	AccountIndex   uint32 `json:"account_index"`
	AccountName    string `json:"account_name"`
	AccountBalance string `json:"account_balance"`
	AccountDelta   string `json:"account_delta"`
}

type ReqGetTxsByAccountIndexAndTxType struct {
	AccountIndex   uint32 `form:"account_index"`
	TxType uint32 `form:"tx_type"`
	Offset uint32 `form:"offset"`
	Limit  uint32 `form:"limit"`
}

type RespGetTxsByAccountIndexAndTxType struct {
	Total uint32 `json:"total"`
	Txs   []*Tx  `json:"txs"`
}

type ReqGetTxsByAccountName struct {
	AccountName string `form:"account_name"`
	Offset      uint32 `form:"offset"`
	Limit       uint32 `form:"limit"`
}

type RespGetTxsByAccountName struct {
	Total uint32 `json:"total"`
	Txs   []*Tx  `json:"txs"`
}

type ReqGetTxsByPubKey struct {
	Pk     string `form:"pk"`
	Offset uint32 `form:"offset"`
	Limit  uint32 `form:"limit"`
}

type RespGetTxsByPubKey struct {
	Total uint32 `json:"total"`
	Txs   []*Tx  `json:"txs"`
}

type ReqGetTxByHash struct {
	TxHash string `form:"tx_hash"`
}

type RespGetTxByHash struct {
	Tx          Tx    `json:"result"`
	CommittedAt int64 `json:"committed_at"`
	VerifiedAt  int64 `json:"verified_at"`
	ExecutedAt  int64 `json:"executed_at"`
}

type ReqSendTx struct {
	TxType uint32 `form:"tx_type"`
	TxInfo string `form:"tx_info"`
}

type RespSendTx struct {
	TxId string `json:"tx_id"`
}

type ReqGetMempoolTxs struct {
	Offset uint32 `form:"offset"`
	Limit  uint32 `form:"limit"`
}

type RespGetMempoolTxs struct {
	Total      uint32 `json:"total"`
	MempoolTxs []*Tx  `json:"mempool_txs"`
}

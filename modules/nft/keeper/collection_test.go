package keeper_test

import (
	"github.com/irisnet/irismod/modules/nft/keeper"
	"github.com/irisnet/irismod/modules/nft/types"
)

func (suite *KeeperSuite) TestSetCollection() {
	nft := types.NewBaseNFT(tokenID, tokenNm, address, tokenURI, tokenURIHash, tokenData)
	// create a new NFT and add it to the collection created with the NFT mint
	nft2 := types.NewBaseNFT(tokenID2, tokenNm, address, tokenURI, tokenURIHash2, tokenData)

	denomE := types.Denom{
		Id:               denomID,
		Name:             denomNm,
		Schema:           schema,
		Creator:          address.String(),
		Symbol:           denomSymbol,
		MintRestricted:   true,
		UpdateRestricted: true,
	}

	collection2 := types.Collection{
		Denom: denomE,
		NFTs:  []types.BaseNFT{nft2, nft},
	}

	err := suite.keeper.SetCollection(suite.ctx, collection2)
	suite.Nil(err)

	collection2, err = suite.keeper.GetCollection(suite.ctx, denomID)
	suite.NoError(err)
	suite.Len(collection2.NFTs, 2)

	msg, fail := keeper.SupplyInvariant(suite.keeper)(suite.ctx)
	suite.False(fail, msg)
}

func (suite *KeeperSuite) TestGetCollection() {
	// MintNFT shouldn't fail when collection does not exist
	err := suite.keeper.MintNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenURIHash, tokenData, address)
	suite.NoError(err)

	// collection should exist
	collection, err := suite.keeper.GetCollection(suite.ctx, denomID)
	suite.NoError(err)
	suite.NotEmpty(collection)

	msg, fail := keeper.SupplyInvariant(suite.keeper)(suite.ctx)
	suite.False(fail, msg)
}

func (suite *KeeperSuite) TestGetCollections() {

	// MintNFT shouldn't fail when collection does not exist
	err := suite.keeper.MintNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenURIHash, tokenData, address)
	suite.NoError(err)

	msg, fail := keeper.SupplyInvariant(suite.keeper)(suite.ctx)
	suite.False(fail, msg)
}

func (suite *KeeperSuite) TestGetSupply() {
	// MintNFT shouldn't fail when collection does not exist
	err := suite.keeper.MintNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenURIHash, tokenData, address)
	suite.NoError(err)

	// MintNFT shouldn't fail when collection does not exist
	err = suite.keeper.MintNFT(suite.ctx, denomID, tokenID2, tokenNm2, tokenURI, tokenURIHash, tokenData, address2)
	suite.NoError(err)

	// MintNFT shouldn't fail when collection does not exist
	err = suite.keeper.MintNFT(suite.ctx, denomID2, tokenID, tokenNm2, tokenURI, tokenURIHash, tokenData, address2)
	suite.NoError(err)

	supply := suite.keeper.GetTotalSupply(suite.ctx, denomID)
	suite.Equal(uint64(2), supply)

	supply = suite.keeper.GetTotalSupply(suite.ctx, denomID2)
	suite.Equal(uint64(1), supply)

	supply = suite.keeper.GetTotalSupplyOfOwner(suite.ctx, denomID, address)
	suite.Equal(uint64(1), supply)

	supply = suite.keeper.GetTotalSupplyOfOwner(suite.ctx, denomID, address2)
	suite.Equal(uint64(1), supply)

	supply = suite.keeper.GetTotalSupply(suite.ctx, denomID)
	suite.Equal(uint64(2), supply)

	supply = suite.keeper.GetTotalSupply(suite.ctx, denomID2)
	suite.Equal(uint64(1), supply)

	//burn nft
	err = suite.keeper.BurnNFT(suite.ctx, denomID, tokenID, address)
	suite.NoError(err)

	supply = suite.keeper.GetTotalSupply(suite.ctx, denomID)
	suite.Equal(uint64(1), supply)

	supply = suite.keeper.GetTotalSupply(suite.ctx, denomID)
	suite.Equal(uint64(1), supply)

	//burn nft
	err = suite.keeper.BurnNFT(suite.ctx, denomID, tokenID2, address2)
	suite.NoError(err)

	supply = suite.keeper.GetTotalSupply(suite.ctx, denomID)
	suite.Equal(uint64(0), supply)

	supply = suite.keeper.GetTotalSupply(suite.ctx, denomID)
	suite.Equal(uint64(0), supply)
}

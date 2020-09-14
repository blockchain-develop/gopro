package _interface

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

/**
 * @dev ERC-721 non-fungible token standard.
 * See https://github.com/ethereum/EIPs/blob/master/EIPS/eip-721.md.
 */
type ERC721 interface {

	/**
	 * @dev Emits when ownership of any NFT changes by any mechanism. This event emits when NFTs are
	 * created (`from` == 0) and destroyed (`to` == 0). Exception: during contract creation, any
	 * number of NFTs may be created and assigned without emitting Transfer. At the time of any
	 * transfer, the approved address for that NFT (if any) is reset to none.
	 */
	emitTransfer(_from common.Address, _to common.Address, _tokenId *big.Int)

	/**
	 * @dev This emits when the approved address for an NFT is changed or reaffirmed. The zero
	 * address indicates there is no approved address. When a Transfer event emits, this also
	 * indicates that the approved address for that NFT (if any) is reset to none.
	 */
	emitApproval(_owner common.Address, _approved common.Address, _tokenId *big.Int)

	/**
	 * @dev This emits when an operator is enabled or disabled for an owner. The operator can manage
	 * all NFTs of the owner.
	 */
	emitApprovalForAll(_owner common.Address, _operator common.Address, _approved bool)

	/**
	 * @dev Transfers the ownership of an NFT from one address to another address.
	 * @notice Throws unless `msg.sender` is the current owner, an authorized operator, or the
	 * approved address for this NFT. Throws if `_from` is not the current owner. Throws if `_to` is
	 * the zero address. Throws if `_tokenId` is not a valid NFT. When transfer is complete, this
	 * function checks if `_to` is a smart contract (code size > 0). If so, it calls
	 * `onERC721Received` on `_to` and throws if the return value is not
	 * `bytes4(keccak256("onERC721Received(address,uint256,bytes)"))`.
	 * @param _from The current owner of the NFT.
	 * @param _to The new owner.
	 * @param _tokenId The NFT to transfer.
	 * @param _data Additional data with no specified format, sent in call to `_to`.
	 */
	safeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte)

	/**
	 * @dev Transfers the ownership of an NFT from one address to another address.
	 * @notice This works identically to the other function with an extra data parameter, except this
	 * function just sets data to ""
	 * @param _from The current owner of the NFT.
	 * @param _to The new owner.
	 * @param _tokenId The NFT to transfer.
	 */
	safeTransferFrom2(_from common.Address, _to common.Address, _tokenId *big.Int)

	/**
	 * @dev Throws unless `msg.sender` is the current owner, an authorized operator, or the approved
	 * address for this NFT. Throws if `_from` is not the current owner. Throws if `_to` is the zero
	 * address. Throws if `_tokenId` is not a valid NFT.
	 * @notice The caller is responsible to confirm that `_to` is capable of receiving NFTs or else
	 * they mayb be permanently lost.
	 * @param _from The current owner of the NFT.
	 * @param _to The new owner.
	 * @param _tokenId The NFT to transfer.
	 */
	transferFrom(_from common.Address, _to common.Address, _tokenId *big.Int)

	/**
	 * @dev Set or reaffirm the approved address for an NFT.
	 * @notice The zero address indicates there is no approved address. Throws unless `msg.sender` is
	 * the current NFT owner, or an authorized operator of the current owner.
	 * @param _approved The new approved NFT controller.
	 * @param _tokenId The NFT to approve.
	 */
	approve(_approved common.Address, _tokenId *big.Int)

	/**
	 * @dev Enables or disables approval for a third party ("operator") to manage all of
	 * `msg.sender`'s assets. It also emits the ApprovalForAll event.
	 * @notice The contract MUST allow multiple operators per owner.
	 * @param _operator Address to add to the set of authorized operators.
	 * @param _approved True if the operators is approved, false to revoke approval.
	 */
	setApprovalForAll(operator common.Address, _approved bool)

	/**
	 * @dev Returns the number of NFTs owned by `_owner`. NFTs assigned to the zero address are
	 * considered invalid, and this function throws for queries about the zero address.
	 * @param _owner Address for whom to query the balance.
	 * @return Balance of _owner.
	 */
	balanceOf(_owner common.Address)(*big.Int)

	/**
	 * @dev Returns the address of the owner of the NFT. NFTs assigned to zero address are considered
	 * invalid, and queries about them do throw.
	 * @param _tokenId The identifier for an NFT.
	 * @return Address of _tokenId owner.
	 */
	ownerOf(_tokenId *big.Int) (common.Address)

	/**
	 * @dev Get the approved address for a single NFT.
	 * @notice Throws if `_tokenId` is not a valid NFT.
	 * @param _tokenId The NFT to find the approved address for.
	 * @return Address that _tokenId is approved for.
	 */
	getApproved(_tokenId *big.Int) (common.Address)

	/**
	 * @dev Returns true if `_operator` is an approved operator for `_owner`, false otherwise.
	 * @param _owner The address that owns the NFTs.
	 * @param _operator The address that acts on behalf of the owner.
	 * @return True if approved for all, false otherwise.
	 */
	isApprovedForAll(_owner common.Address, _operator common.Address) (bool)
}

// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22;

interface IERC2981  {
    function royaltyInfo(
        uint256 _tokenId,
        uint256 _salePrice
    ) external view returns (
        address receiver,
        uint256 royaltyAmount
    );

}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

// Import Chainlink Aggregator Interface
import {AggregatorV3Interface} from "lib/chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";

contract Littcoin {
    string public name = "Litt Coin";
    string public symbol = "LIT";
    uint8 public decimals = 18;
    uint256 public tokensSold;

    // 15 USD per token, Chainlink price has 8 decimals
    uint256 public constant tokenPriceInUSD = 15 * 10**8;  // 15 USD per token

    // Chainlink ETH/USD Price Feed address (Mainnet example)
    AggregatorV3Interface internal priceFeed;

    uint256 private constant totalSupplyCoins = 10000000 * 10**18;
    mapping(address => uint256) public addressBalance;

    // Event declarations
    event Transfer(address indexed from, address indexed to, uint256 value);
    event TokensPurchased(address indexed buyer, uint256 amount);

    // Owner of the contract
    address public owner;

    // Constructor to initialize the price feed and allocate initial supply
    constructor(address _priceFeedAddress) {
        priceFeed = AggregatorV3Interface(_priceFeedAddress);
        owner = msg.sender;  // Set the owner of the contract

        // Assign total supply to the deployer
        addressBalance[msg.sender] = totalSupplyCoins;  
    }

    // Return the total supply of tokens
    function totalSupply() public pure returns (uint256) {
        return totalSupplyCoins;
    }

    // Get the balance of an account
    function balanceOf(address _account) public view returns (uint256) {
        return addressBalance[_account];
    }

    // Function to get the latest ETH price in USD
    function getLatestPrice() internal view returns (uint256) {
        (, int256 price, , ,) = priceFeed.latestRoundData();
        return uint256(price);  // Chainlink price has 8 decimals
    }

    // Buy tokens based on ETH sent
    function buyTokens(uint256 _numberOfTokens) public payable returns (bool) {
        // Get the current price of 1 ETH in USD
        uint256 ethPriceInUSD = getLatestPrice();

        // Calculate how much ETH the user needs to send to buy the specified number of tokens
        uint256 requiredETHAmount = (_numberOfTokens * tokenPriceInUSD) / ethPriceInUSD;

        // Ensure the buyer sent enough ETH
        require(msg.value >= requiredETHAmount, "Insufficient ETH sent for the specified number of tokens");

        // Check if there are enough tokens available
        uint256 tokensToBuy = _numberOfTokens;
        require(tokensToBuy <= totalSupplyCoins - tokensSold, "Not enough tokens available");

        // Update buyer's balance and contract's token supply
        addressBalance[msg.sender] += tokensToBuy;
        tokensSold += tokensToBuy;

        // Refund any excess ETH sent by the buyer
        uint256 excessETH = msg.value - requiredETHAmount;
        if (excessETH > 0) {
            payable(msg.sender).transfer(excessETH);
        }

        // Emit an event for the purchase
        emit TokensPurchased(msg.sender, tokensToBuy);

        return true;
    }

    // Transfer tokens from sender to recipient
    function transferTo(address _recipient, uint256 _amount) public returns (bool) {
        // Check if sender has enough balance
        require(addressBalance[msg.sender] >= _amount, "Insufficient balance");

        // Subtract the amount from sender's balance
        addressBalance[msg.sender] -= _amount;

        // Add the amount to recipient's balance
        addressBalance[_recipient] += _amount;

        // Emit a Transfer event
        emit Transfer(msg.sender, _recipient, _amount);

        return true;
    }

    // Function to withdraw collected ETH (only the contract owner can withdraw)
    function withdraw() public {
        require(msg.sender == owner, "Only the owner can withdraw");
        payable(msg.sender).transfer(address(this).balance);
    }

    // Function to calculate the number of tokens the user will get for a specific ETH value
    function convertETHToTokens(uint256 _ethAmount) public view returns (uint256) {
        // Get the current price of 1 ETH in USD
        uint256 ethPriceInUSD = getLatestPrice();  

        // Calculate the number of tokens the user will receive for the provided ETH amount
        uint256 tokensToReceive = (_ethAmount * ethPriceInUSD) / tokenPriceInUSD;

        return tokensToReceive;
    }
}

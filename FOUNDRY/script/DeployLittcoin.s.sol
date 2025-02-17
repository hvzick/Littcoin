// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import { Littcoin } from "../src/littcoin.sol";  // Make sure the path is correct

contract Deploy {
    Littcoin public littcoin;

    function run() external {
        littcoin = new Littcoin(0x779877A7B0D9E8603169DdbD7836e478b4624789);  // Pass constructor args if needed
    }
}

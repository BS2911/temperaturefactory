// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


contract Temperfactory is AccessControl,Ownable {
    uint256 public temperature;
    bytes32 public constant TEMPERFAC_MANAGER_ROLE = keccak256("TEMPERFAC_MANAGER_ROLE");
    bytes32 public constant TEMPERFAC_WORKER_ROLE = keccak256("TEMPERFAC_WORKER_ROLE");

    constructor() {
        _grantRole(TEMPERFAC_MANAGER_ROLE,msg.sender);
        _grantRole(TEMPERFAC_WORKER_ROLE,msg.sender);
    }

    function setTemperature(uint256 _value) public onlyRole(TEMPERFAC_WORKER_ROLE) {
        temperature=_value;
    }

    function getTemperature() external view returns (uint256) {
        return temperature;
    }

    //Official Management
    function setWorker(address _worker) public onlyRole(TEMPERFAC_MANAGER_ROLE){
        _grantRole(TEMPERFAC_WORKER_ROLE,_worker);
    }

    function setManager(address _manager) public onlyOwner {
        _grantRole(TEMPERFAC_WORKER_ROLE,_manager);
    }

}

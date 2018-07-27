pragma solidity ^0.4.19;

import "./DnomGenesisValidators.sol";

contract DnomGenesis {
    
    uint256 initializedTime = 0;
    
    uint16 genesisPublishStart = 0;
    
    uint16 genesisPublishEnd = 0;
    
    uint16 genesisPublishEndForValidators = 0;
    
    address genesisValidators;
    
    uint256 public ONE_DAY = 24 * 60 * 60 * 1000;
    
    constructor(address _genesisValidators, uint16 _genesisPublishStart, uint16 _genesisPublishEnd, uint16 _genesisPublishEndForValidators) public {
        require(_genesisPublishStart <= _genesisPublishEnd);
        initializedTime = block.timestamp;
        genesisPublishStart = _genesisPublishStart;
        genesisPublishEnd = _genesisPublishEnd;
        genesisValidators = _genesisValidators;
        genesisPublishEndForValidators = _genesisPublishEndForValidators;
    }

    struct Genesis {
        string url;
        uint256 time;
        string sha256hash;
        bool exists;
    }
    
    mapping (address => Genesis) publishedGenesis;

    address[] genesisList; 
    
    function publishGenesis(string url, string sha256hash) public {
        uint16 currentDay = (uint16 ((block.timestamp - initializedTime) / ONE_DAY)) + 1;
        require (currentDay >= genesisPublishStart && currentDay <= genesisPublishEnd);
        if (DnomGenesisValidators(genesisValidators).isValidator(msg.sender)) {
            require(currentDay <= genesisPublishEndForValidators);
        }
        Genesis memory genesis;
        genesis.url = url;
        genesis.sha256hash = sha256hash;
        genesis.time = block.timestamp;
        if (!publishedGenesis[msg.sender].exists) {
            genesis.exists = true;
            genesisList.push(msg.sender);
	    }
        publishedGenesis[msg.sender] = genesis;
    }

    function getTotalGenesis() public constant returns(uint256) {
        return genesisList.length;
    }
   
    function getGenesisAt(uint256 index) public constant returns(address ethAddr, string url, string sha256hash) {
        ethAddr = genesisList[index];
        url = publishedGenesis[ethAddr].url;
        sha256hash = publishedGenesis[ethAddr].sha256hash;
    }
    
    function getGenesisBy(address ethAddr) public constant returns(string url, string sha256hash) {
        url = publishedGenesis[ethAddr].url;
        sha256hash = publishedGenesis[ethAddr].sha256hash;
    }
}

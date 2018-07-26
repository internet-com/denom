pragma solidity ^0.4.19;

contract DnomGenesis {
    
    constructor() public {
    }

    struct Genesis {
        string url;
        uint256 time;
        bool exists;
    }
    
    mapping (address => Genesis) publishedGenesis;

    address[] genesisList; 
    
    function publishGenesis(string url) public {
        Genesis memory genesis;
        genesis.url = url;
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
   
    function getGenesisAt(uint256 index) public constant returns(address ethAddr, string url) {
        ethAddr = genesisList[index];
        url = publishedGenesis[ethAddr].url;
    }
    
    function getGenesisBy(address ethAddr) public constant returns(string url) {
        url = publishedGenesis[ethAddr].url;
    }
}

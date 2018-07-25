pragma solidity ^0.4.19;

contract DnomGenesis {
    
    constructor() {
    }

    struct Genesis {
        string url;
        string seed;
        uint256 time;
        string signature;
        bool exists;
    }
    
    mapping (address => Genesis) publishedGenesis;

    address[] genesisList; 
    
    function publishGenesis(string url, string seed, string signature) public {
        Genesis memory genesis;
        genesis.url = url;
        genesis.seed = seed;
        genesis.time = block.timestamp;
        genesis.signature = signature;
        if (!publishedGenesis[msg.sender].exists) {
            genesis.exists = true;
            genesisList.push(msg.sender);
	}
        publishedGenesis[msg.sender] = genesis;
    }

    function getTotalGenesis() public constant returns(uint256) {
        return genesisList.length;
    }
   
    function getGenesisAt(uint256 index) public constant returns(address ethAddr, string url, string seed, string signature) {
        ethAddr = genesisList[index];
        url = publishedGenesis[ethAddr].url;
        seed = publishedGenesis[ethAddr].seed;
        signature = publishedGenesis[ethAddr].signature; 
    }
}

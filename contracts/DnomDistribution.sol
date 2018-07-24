pragma solidity ^0.4.19;
contract DnomDistribution {
    
    uint64 public tokensDistributed;
    
    struct Claim {
        string denomPublicKey;
        string denomAddress;
        uint16 registeredDay;
        uint256 verificationFeePaid; // Optional verification fee
        bool claimed;
    }
    
    struct Domain {
        bool registered;
        string domainName;
        string owner;
        uint256 tokensAllocated; // Will be set if verification was successful
        mapping(string => Claim) claims; // List of claims for the domain
        string[] claimAddress;
    }
    
    mapping(string => Domain) domainsRegistered;
    
    string[] public domainsRegisteredByDay;
    
    address public owner;
    
    uint256 initializedTime = 0;
    
    uint256 public dayStart = 0;
    
    uint256 public dayEnd = 0;
    
    uint256 public ONE_DAY = 24 * 60 * 60 * 1000;
    
    string public genesisURL;
    
    uint256 public verificationFee;
    
    uint256 public maxRegistrationDays;
    
    modifier onlyOwner {
        if (msg.sender == owner) {
            _;
        }
    }
    
    constructor() public {
        initializedTime = block.timestamp;
        verificationFee = 1000000000000000000 / 10; //0.1 ETH
        maxRegistrationDays = 100;
    }
    
    function addDomainClaim(string domainName, Claim claim) private {
        if (msg.value >= 0) {
            claim.verificationFeePaid = msg.value;
        }
        domainsRegistered[domainName].claims[claim.denomAddress] = claim;
        domainsRegistered[domainName].claims[claim.denomAddress].claimed = true;
        domainsRegistered[domainName].claimAddress.push(claim.denomAddress);
    }
    
    function claimDomain(string domainName, string denomPublicKey, string denomAddress) public payable {
        uint16 currentDay = uint16 ((block.timestamp - initializedTime) / ONE_DAY);
        if ((msg.value > 0 && msg.value < verificationFee) || currentDay >= maxRegistrationDays) {
            revert();
        }
        Claim memory claim;
        claim.denomPublicKey = denomPublicKey;
        claim.denomAddress = denomAddress;
        claim.registeredDay = currentDay;
        if (domainsRegistered[domainName].registered) {
            if (domainsRegistered[domainName].claims[denomAddress].claimed) {
                if (msg.value >= 0) {
                    domainsRegistered[domainName].claims[denomAddress].verificationFeePaid += msg.value;
                }
            } else {
                addDomainClaim(domainName, claim);
            }
        } else {
            addDomainClaim(domainName, claim);
        }
    }
    
    function verifyDomain(string domainName, string denomAddress, uint256 tokensAllocated) public onlyOwner {
        
    }
    
    function publishGenesis(string url) public onlyOwner {
        genesisURL = url;
    }
    
}
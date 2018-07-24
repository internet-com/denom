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
        address owner;
        uint256 tokensAllocated; // Will be set if verification was successful
        mapping(address => Claim) claims; // List of claims for the domain
        address[] claimAddress;
    }
    
    mapping(string => Domain) domainsRegistered;
    
    string[] public registeredDomains;
    
    address public owner;
    
    uint256 initializedTime = 0;
    
    uint256 public dayStart = 0;
    
    uint256 public dayEnd = 0;
    
    uint256 public ONE_DAY = 24 * 60 * 60 * 1000;
    
    string public genesisURL;
    
    uint256 public verificationFee;
    
    uint256 public maxRegistrationDays;
    
    uint256 public waitPeriod;
    
    modifier onlyOwner {
        if (msg.sender == owner) {
            _;
        }
    }
    
    constructor(string denomPublicKey, string denomAddress) public {
        initializedTime = block.timestamp;
        verificationFee = 1000000000000000000 / 10; //0.1 ETH
        maxRegistrationDays = 100;
        waitPeriod = maxRegistrationDays + 7;
        owner = msg.sender;
        Claim memory claim;
        claim.denomPublicKey = denomPublicKey;
        claim.denomAddress = denomAddress;
        claim.registeredDay = 0;
        addDomainClaim("denom.org", claim);
    }
    
    function addDomainClaim(string domainName, Claim claim) private {
        if (msg.value > 0) {
            claim.verificationFeePaid = msg.value;
        }
        domainsRegistered[domainName].claims[msg.sender] = claim;
        domainsRegistered[domainName].claims[msg.sender].claimed = true;
        domainsRegistered[domainName].claimAddress.push(msg.sender);
        if (!domainsRegistered[domainName].registered) {
            domainsRegistered[domainName].registered = true;
            registeredDomains.push(domainName);
        }
    }
    
    function claimDomain(string domainName, string denomPublicKey, string denomAddress) public payable {
        uint16 currentDay = (uint16 ((block.timestamp - initializedTime) / ONE_DAY)) + 1;
        if ((msg.value > 0 && msg.value < verificationFee) || currentDay > maxRegistrationDays) {
            revert();
        }
        Claim memory claim;
        claim.denomPublicKey = denomPublicKey;
        claim.denomAddress = denomAddress;
        claim.registeredDay = currentDay;
        if (domainsRegistered[domainName].registered) {
            if (domainsRegistered[domainName].claims[msg.sender].claimed) {
                if (msg.value > 0) {
                    domainsRegistered[domainName].claims[msg.sender].verificationFeePaid += msg.value;
                }
            }
        }
        addDomainClaim(domainName, claim);
    }
    
    function verifyDomain(string domainName, address senderAddress, uint256 tokensAllocated) public onlyOwner {
        if (domainsRegistered[domainName].registered) {
            if (domainsRegistered[domainName].claims[senderAddress].claimed) {
                domainsRegistered[domainName].tokensAllocated = tokensAllocated;
                domainsRegistered[domainName].owner = senderAddress;
            }
        }
    }
    
    function cancelClaim(string domainName) public {
        uint16 currentDay = (uint16 ((block.timestamp - initializedTime) / ONE_DAY)) + 1;
        if (currentDay <= waitPeriod) {
            if (domainsRegistered[domainName].registered) {
                if (domainsRegistered[domainName].claims[msg.sender].claimed) {
                    domainsRegistered[domainName].claims[msg.sender].claimed = false;
                    if (domainsRegistered[domainName].claims[msg.sender].verificationFeePaid > 0) {
                        uint256 feePaid = domainsRegistered[domainName].claims[msg.sender].verificationFeePaid;
                        domainsRegistered[domainName].claims[msg.sender].verificationFeePaid = 0;
                        if (!msg.sender.send(feePaid)) {
                            domainsRegistered[domainName].claims[msg.sender].verificationFeePaid = feePaid;
                        }
                    }
                }
            }
        }
    }
    
    function getTotalDomains() public constant returns(uint256) {
        return registeredDomains.length;
    }
    
    function getDomainAt(uint256 index) public constant returns(string) {
        return registeredDomains[index];
    }
    
    function getTotalDomainClaims(string domainName) public constant returns(uint256) {
        return domainsRegistered[domainName].claimAddress.length;
    }
    
    function getDomainClaimAt(string domainName, uint256 index) public constant returns(address) {
        return domainsRegistered[domainName].claimAddress[index];
    }
    
    function getClaimDetails(string domainName, address claimAddress) public constant returns(string denomAddress, string denomPublicKey, uint16 registeredDay, uint256 verificationFeePaid) {
        denomAddress = domainsRegistered[domainName].claims[claimAddress].denomAddress;
        denomPublicKey = domainsRegistered[domainName].claims[claimAddress].denomPublicKey;
        registeredDay = domainsRegistered[domainName].claims[claimAddress].registeredDay;
        verificationFeePaid = domainsRegistered[domainName].claims[claimAddress].verificationFeePaid;
    }
    
    function getDomainOwner(string domainName) public constant returns(address claimAddress) {
        return domainsRegistered[domainName].owner;
    }
    
    function withdrawVerificationFee() public onlyOwner {
        uint16 currentDay = (uint16 ((block.timestamp - initializedTime) / ONE_DAY)) + 1;
        if (currentDay > waitPeriod) {
            msg.sender.transfer(address(this).balance);
        }
    }
    
    function publishGenesis(string url) public onlyOwner {
        genesisURL = url;
    }
    
}
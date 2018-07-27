pragma solidity ^0.4.19;
contract DnomDistribution {
    
    struct Claim {
        string denomAddress;
        uint16 registeredDay;
        bool claimed;
    }
    
    struct Domain {
        bool registered;
        string domainName;
        mapping(address => Claim) claims; // List of claims for the domain
        address[] claimAddress;
    }

    mapping(string => Domain) domainsRegistered;

    string[] public registeredDomains;
    
    uint256 initializedTime = 0;
    
    uint256 public ONE_DAY = 24 * 60 * 60 * 1000;
    
    uint256 public maxRegistrationDays;
    
    constructor(uint16 registrationDays) public {
        initializedTime = block.timestamp;
        maxRegistrationDays = registrationDays;
        Claim memory claim;
        claim.denomAddress = "0x0000"; // Change this
        claim.registeredDay = 0;
        addDomainClaim("denom.org", claim);
    }
    
    function addDomainClaim(string domainName, Claim claim) private {
        claim.claimed = true;
        domainsRegistered[domainName].claims[msg.sender] = claim;
        domainsRegistered[domainName].claimAddress.push(msg.sender);
        if (!domainsRegistered[domainName].registered) {
            domainsRegistered[domainName].registered = true;
            registeredDomains.push(domainName);
        }
    }
    
    function claimDomain(string domainName, string denomAddress) public {
        uint16 currentDay = (uint16 ((block.timestamp - initializedTime) / ONE_DAY)) + 1;
        require (currentDay <= maxRegistrationDays);
        Claim memory claim;
        claim.denomAddress = denomAddress;
        claim.registeredDay = currentDay;
        if (domainsRegistered[domainName].registered) {
            if (domainsRegistered[domainName].claims[msg.sender].claimed) {
                // Set denomAddress if it has changed
                return ;
            }
        }
        addDomainClaim(domainName, claim);
    }
    
    function cancelClaim(string domainName) public {
        uint16 currentDay = (uint16 ((block.timestamp - initializedTime) / ONE_DAY)) + 1;
        if (currentDay <= maxRegistrationDays) {
            if (domainsRegistered[domainName].registered) {
                domainsRegistered[domainName].claims[msg.sender].claimed = false;
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
    
    function getClaimDetails(string domainName, address claimAddress) public constant returns(string denomAddress, uint16 registeredDay) {
        denomAddress = domainsRegistered[domainName].claims[claimAddress].denomAddress;
        registeredDay = domainsRegistered[domainName].claims[claimAddress].registeredDay;
    }

}

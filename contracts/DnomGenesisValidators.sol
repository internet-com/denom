pragma solidity ^0.4.19;

contract DnomGenesisValidators {
    
    uint16 public delegationPeriodStart;
    
    uint16 public delegationPeriodEnd;
    
    uint256 initializedTime = 0;
    
    constructor(uint16 _delegationPeriodStart, uint16 _delegationPeriodEnd) public {
        require(_delegationPeriodStart <= _delegationPeriodEnd);
        initializedTime = block.timestamp;
        delegationPeriodStart = _delegationPeriodStart;
        delegationPeriodEnd = _delegationPeriodEnd;
    }

    struct Validator {
        string denomAddress;
        string denomPublicKey;
        string signature;
        string name;
        string website;
        bool exists;
    }
    
    struct Delegation {
        string signature;
        uint16 percentage;
        bool exists;
    }

    struct Delegator {
        string denomAddress;
        mapping (string => Delegation) delegation;
        string[] delegationList;
        uint16 percentageDelegated;
        bool exists;
    }

    mapping(address => Validator) validators;

    mapping(address => Delegator) delegators;

    address[] delegatorList;

    address[] validatorList;
    
    uint256 public ONE_DAY = 24 * 60 * 60 * 1000;
    
    

    function registerValidator(string validatorName, string website, string denomAddress, string denomPublicKey, string signature) public {
        uint16 currentDay = (uint16 ((block.timestamp - initializedTime) / ONE_DAY)) + 1;
        require(delegationPeriodStart >= currentDay && currentDay <= delegationPeriodEnd);
        Validator memory validator;
        validator.name = validatorName;
        validator.website = website;
        validator.denomAddress = denomAddress;
        validator.denomPublicKey = denomPublicKey;
        validator.signature = signature;
        if (!validators[msg.sender].exists) {
            validator.exists = true;
            validatorList.push(msg.sender);
        }
        validators[msg.sender] = validator;
    }

    function getTotalValidators() public constant returns(uint256) {
        return validatorList.length;
    }

    function getValidatorAt(uint256 index) public constant returns(address ethAddr, string validatorName, string website, string denomAddress, string denomPublicKey, string signature) {
        ethAddr = validatorList[index];
        validatorName = validators[ethAddr].name;
        website = validators[ethAddr].website;
        denomAddress = validators[ethAddr].denomAddress;
        denomPublicKey = validators[ethAddr].denomPublicKey;
        signature = validators[ethAddr].signature;
    }
    
    function getValidator(address ethAddr) public constant returns(string validatorName, string website, string denomAddress, string denomPublicKey, string signature) {
        validatorName = validators[ethAddr].name;
        website = validators[ethAddr].website;
        denomAddress = validators[ethAddr].denomAddress;
        denomPublicKey = validators[ethAddr].denomPublicKey;
        signature = validators[ethAddr].signature;
    }
    
    function isValidator(address ethAddr) public constant returns(bool) {
        return validators[ethAddr].exists;
    }

    function getTotalDelegators() public constant returns(uint256) {
        return delegatorList.length;
    }

    function getDelegatorAt(uint256 index) public constant returns(address ethAddr, string denomAddress, uint256 totalDelegations) {
        ethAddr = delegatorList[index];
        denomAddress = delegators[ethAddr].denomAddress;
        totalDelegations = delegators[ethAddr].delegationList.length;
    }
    
    function getDelegationAt(address ethAddr, uint256 index) public constant returns(string validatorDenomAddress, uint16 percentage, string signature, bool isValid) {
        validatorDenomAddress = delegators[ethAddr].delegationList[index];
        percentage = delegators[ethAddr].delegation[validatorDenomAddress].percentage;
        signature = delegators[ethAddr].delegation[validatorDenomAddress].signature;
        isValid = delegators[ethAddr].delegation[validatorDenomAddress].exists;
    }

    function delegateToValidator(string denomAddress, string validatorDenomAddress, uint16 percentage, string signature) public {
        uint16 currentDay = (uint16 ((block.timestamp - initializedTime) / ONE_DAY)) + 1;
        require(delegationPeriodStart >= currentDay && currentDay <= delegationPeriodEnd);
	require(delegators[msg.sender].percentageDelegated + percentage <= 100);
        if (!delegators[msg.sender].exists) {
            Delegator memory delegator;
            delegator.denomAddress = denomAddress;
            delegator.exists = true;
            delegatorList.push(msg.sender);
            delegators[msg.sender] = delegator;
        }
        if (!delegators[msg.sender].delegation[validatorDenomAddress].exists) {
            Delegation memory delegation;
            delegation.signature = signature;   
            delegation.percentage = percentage;
            delegation.exists = true;
            delegators[msg.sender].delegation[validatorDenomAddress] = delegation;
            delegators[msg.sender].percentageDelegated += percentage;
            delegators[msg.sender].delegationList.push(validatorDenomAddress);
        } else {
            delegators[msg.sender].percentageDelegated -= delegators[msg.sender].delegation[validatorDenomAddress].percentage;
            delegators[msg.sender].percentageDelegated += percentage;
            delegators[msg.sender].delegation[validatorDenomAddress].percentage = percentage;
            delegators[msg.sender].delegation[validatorDenomAddress].signature = signature;
        }
    }

    function invalidateDelegation(string validatorDenomAddress) public {
        uint16 currentDay = (uint16 ((block.timestamp - initializedTime) / ONE_DAY)) + 1;
        require(delegationPeriodStart >= currentDay && currentDay <= delegationPeriodEnd);
        delegators[msg.sender].delegation[validatorDenomAddress].exists = false;
        delegators[msg.sender].percentageDelegated -= delegators[msg.sender].delegation[validatorDenomAddress].percentage;
        delegators[msg.sender].delegation[validatorDenomAddress].percentage = 0;
    }

}

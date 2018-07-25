pragma solidity ^0.4.19;

contract GenesisDnomValidators {

    struct Validator {
        string denomAddress;
        string denomPublicKey;
        string signature;
        string name;
        string website;
        bool exists;
    }

    struct Delegator {
        string denomAddress;
        string validator;
        string signature;
        uint16 percentage;
        bool exists;
    }

    mapping(address => Validator) validators;

    mapping(address => Delegator) delegators;

    address[] delegatorList;

    address[] validatorList;

    function registerValidator(string validatorName, string website, string denomAddress, string denomPublicKey, string signature) public {
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

    function getTotalDelegators() public constant returns(uint256) {
        return delegatorList.length;
    }

    function getDelegatorAt(uint256 index) public constant returns(address ethAddr, string denomAddress, string validator, uint16 percentage, string signature) {
        ethAddr = delegatorList[index];
        denomAddress = delegators[ethAddr].denomAddress;
        validator = delegators[ethAddr].validator;
        percentage = delegators[ethAddr].percentage;
        signature = delegators[ethAddr].signature;
    }

    function delegateToValidator(string denomAddress, string validatorDenomAddress, uint16 percentage, string signature) public {
        Delegator memory delegator;
        delegator.denomAddress = denomAddress;
        delegator.validator = validatorDenomAddress;
        delegator.signature = signature;
        delegator.percentage = percentage;
        if (!delegators[msg.sender].exists) {
            delegator.exists = true;
            delegatorList.push(msg.sender);
        }
        delegators[msg.sender] = delegator;
    }

}

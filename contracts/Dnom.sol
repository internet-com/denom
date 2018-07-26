pragma solidity ^0.4.19;

import "./DnomDistribution.sol";
import "./DnomGenesis.sol";
import "./DnomGenesisValidators.sol";

contract Dnom {
    address public dnomDistributionContract;
    address public dnomGenesisContract;
    address public dnomGenesisValidatorContract;
    constructor(uint16 registrationPeriod, uint16 delegationPeriodStart, uint16 delegationPeriodEnd, uint16 genesisPeriodStart, uint16 genesisPeriodEnd) public {
        dnomDistributionContract = new DnomDistribution(registrationPeriod);
        dnomGenesisValidatorContract = new DnomGenesisValidators(delegationPeriodStart, delegationPeriodEnd);
        dnomGenesisContract = new DnomGenesis(genesisPeriodStart, genesisPeriodEnd);
    }
}

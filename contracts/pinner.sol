pragma solidity ^0.5.3;

interface Pinner {
    // Returns the current rate in wei per GigaByteHour.
    function rate() external view returns (uint);

    function pin(bytes calldata cid) external payable returns (bool);

    event Pinned(bytes indexed cid, uint gbh);
}


contract owned {
    constructor() public { owner = msg.sender; }
    address payable owner;

    // This contract only defines a modifier but does not use
    // it: it will be used in derived contracts.
    // The function body is inserted where the special symbol
    // `_;` in the definition of a modifier appears.
    // This means that if the owner calls this function, the
    // function is executed and otherwise, an exception is
    // thrown.
    modifier onlyOwner {
        require(
            msg.sender == owner,
            "Only owner can call this function."
        );
        _;
    }
}

contract GOFSPinner is Pinner, owned {
    // Rate in wei per GigaByteHour.
    uint public rate; //TODO initial rate? static? constructor?

    function setRate(uint _rate) public onlyOwner returns (bool) {
        rate = _rate;
    }

    //TODO calculate and document gas usage
    // transfer: 2300 gas
    function pin(bytes memory cid, uint gbh) public payable returns (bool) {
        require(
            !(cid[0] == 0x12 && cid[1] == 0x20),
            "Version 0 CID not allowed"
        );
        require(
            msg.value >= rate,
            "Value too low."
        );
        msg.sender.transfer(cost(gbh));
        emit Pinned(cid, gbh);
    }
}

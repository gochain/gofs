pragma solidity ^0.5.3;

interface Pinner {
    // Returns the current rate in wei per GigaByteHour.
    function rate() external view returns (uint);
    
    function pin(string calldata cid) external payable returns (bool);

    event Pinned(string indexed cid, uint gbh);
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
    uint public rate; //TODO initial rate? fixed? constructor?

    function setCost(uint _rate) public onlyOwner returns (bool) {
        rate = _rate;
    }

    function pin(string memory cid) public payable returns (bool) {
        require(
            msg.value >= rate,
            "Value too low."
        );
        //TODO reject zero? or <cost (minimum 1 GBH)?
        uint gbh = msg.value/rate;
        emit Pinned(cid, gbh);
    }
}

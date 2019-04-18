pragma solidity ^0.5.3;

interface Pinner {
    // Returns the current rate in wei per GigaByteHour.
    function rate() external view returns (uint);

    // Returns the number of the block when this contract was deployed.
    function deployed() external view returns (uint);

    // Pin a CID. gbh must be greater than 0. CID must not be version 0.
    function pin(bytes calldata cid, uint gbh) external payable returns (bool);

    event Pinned(address indexed user, bytes indexed cid);
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

//TODO killable, and then refund full amount when dead?
contract GOFSPinner is Pinner, owned {
    // Rate in wei per GigaByteHour.
    uint public rate;
    uint public deployed;

    constructor(uint _rate) public {
        rate = _rate;
        deployed = block.number;
    }

    function setRate(uint _rate) public onlyOwner returns (bool) {
        rate = _rate;
    }

    //TODO calculate and document gas usage
        // transfer: 2300 gas
    function pin(bytes memory cid, uint gbh) public payable returns (bool) {
        require(
            !(cid[0] == 0x12 && cid[1] == 0x20),
            "Version 0 CID not allowed."
        );
        require(
            gbh > 0,
            "Cannot purchase 0 storage."
        );
        uint cost = gbh*rate;
        require(
            msg.value >= cost,
            "Value too low."
        );
        // refund excess
        msg.sender.transfer(msg.value - cost);
        emit Pinned(msg.sender, cid);
    }
}

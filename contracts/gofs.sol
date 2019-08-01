pragma solidity ^0.5.3;

// The IGOFS interface defines the public functions for GOFS.
interface IGOFS {
    // Returns the current rate in attoGO per ByteHour.
    function rate() external view returns (uint);

    // Returns the number of the block when this contract was deployed.
    function deployed() external view returns (uint);

    // Pin a CID. Value must be greater than 0. CID must not be version 0.
    // Emits Pinned events.
    function pin(bytes calldata cid) external payable;

    // Get the address of the deposit wallet for a cid. Returns 0x0 if none exists.
    function wallet(bytes calldata cid) external view returns (address);

    // Create a deposit wallet for a cid. Returns false if one already exists. CID must not be version 0.
    // Emits CreatedWallet events.
    // Uses ~226935 gas.
    function newWallet(bytes calldata cid) external;

    // Emitted when a cid is pinned.
    event Pinned(address indexed user, bytes indexed cid, uint bh);
    // Emitted when a new wallet is created.
    event CreatedWallet(address indexed user, bytes indexed cid, address wallet);
}

// The Wallet contract contains only a fallback function to forward GO sent to GOFS for a CID.
contract Wallet {
    IGOFS gofs;
    bytes cid;

    constructor(address gofsAddr, bytes memory _cid) public {
        gofs = IGOFS(gofsAddr);
        cid = _cid;
    }

    function() external payable {
        gofs.pin.value(msg.value)(cid);
    }
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

    // Transfer ownership to a new account.
    function changeOwner(address payable newOwner) public onlyOwner {
        owner = newOwner;
    }
}

// The GOFS contract implements IGOFS and includes admin functions as well.
contract GOFS is IGOFS, owned {
    // Rate in attoGO per ByteHour.
    uint public rate;
    // Block number when deployed.
    uint public deployed;

    // Wallet contracts for CIDs.
    mapping(bytes=>address) wallets;

    constructor(uint _rate) public {
        rate = _rate;
        deployed = block.number;
    }

    // Set the storage rate in attoGO per ByteHour. Must be owner.
    function setRate(uint _rate) public onlyOwner{
        rate = _rate;
    }

    // Transfer the entire contract balance to the given address.
    function withdraw(address payable to) public onlyOwner {
        to.transfer(address(this).balance);
    }

    //TODO kill function? fallback to reject/refund when dead?

    /* IGOFS functions: */

    /* rate() and deployed() are auto-generated */

    function pin(bytes memory cid) public payable {
        require(
            !(cid[0] == 0x12 && cid[1] == 0x20),
            "Version 0 CID not allowed."
        );
        require(
            msg.value >= rate,
            "Cannot purchase 0 storage."
        );
        uint bh = msg.value/rate;
        emit Pinned(tx.origin, cid, bh);
    }

    function wallet(bytes memory cid) public view returns (address) {
        return wallets[cid];
    }

    function newWallet(bytes memory cid) public {
        require(
            !(cid[0] == 0x12 && cid[1] == 0x20),
            "Version 0 CID not allowed."
        );
        require(
            wallets[cid] == address(0),
            "Wallet already exists for cid"
        );
        Wallet w = new Wallet(address(this), cid);
        wallets[cid] = address(w);
        emit CreatedWallet(msg.sender, cid, address(w));
    }
}

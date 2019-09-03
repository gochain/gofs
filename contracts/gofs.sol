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
    // Uses <=300000 gas.
    function newWallet(bytes calldata cid) external;

    // Returns the CID which produces this keccak256 hash.
    // CID must have produced a CreatedWallet or Pinned event from this contract.
    function cidByHash(bytes32 hash) external view returns (bytes memory);

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

contract ProxyOwner {
    bytes32 private constant ownerPosition = keccak256("gochain.proxy.owner");

    modifier onlyOwner {
        require(
            msg.sender == owner(),
            "Only owner can call this function."
        );
        _;
    }

    function owner() public view returns (address addr) {
        bytes32 pos = ownerPosition;
        assembly {
            addr := sload(pos)
        }
    }

    // Transfer ownership to a new account.
    function changeOwner(address payable addr) public onlyOwner {
        bytes32 pos = ownerPosition;
        assembly {
            sstore(pos, addr)
        }
    }
}

// The GOFS contract implements IGOFS and includes admin functions as well.
contract GOFS is IGOFS, ProxyOwner {
    // Rate in attoGO per ByteHour.
    uint public rate;
    // Block number when deployed.
    uint public deployed;

    // Wallet contracts for CIDs.
    mapping(bytes=>address) wallets;
    // keccak256(CID) => CID
    mapping(bytes32=>bytes) public cidByHash;

    // Set the storage rate in attoGO per ByteHour. Must be owner.
    function setRate(uint _rate) public onlyOwner {
        if (deployed == 0) {
            deployed = block.number;
        }
        rate = _rate;
    }

    // Transfer the entire contract balance to the given address.
    function withdraw(address payable to) public onlyOwner {
        to.transfer(address(this).balance);
    }

    /* IGOFS functions: */

    /* rate() and deployed() are auto-generated */

    function pin(bytes memory cid) public payable {
        require(
            version(cid) != 0,
            "Version 0 CID not allowed."
        );
        require(
            msg.value >= rate,
            "Cannot purchase 0 storage."
        );
        uint bh = msg.value/rate;
        emit Pinned(tx.origin, cid, bh);
        _ensureHashStored(cid);
    }

    function wallet(bytes memory cid) public view returns (address) {
        return wallets[cid];
    }

    function version(bytes memory cid) internal pure returns (uint) {
        if (cid.length == 34 && cid[0] == 0x12 && cid[1] == 0x20) {
            return 0;
        }
        return 1;
    }

    function newWallet(bytes memory cid) public {
        require(
            version(cid) != 0,
            "Version 0 CID not allowed."
        );
        require(
            wallets[cid] == address(0),
            "Wallet already exists for cid"
        );
        Wallet w = new Wallet(address(this), cid);
        wallets[cid] = address(w);
        emit CreatedWallet(msg.sender, cid, address(w));
        _ensureHashStored(cid);
    }

    function _ensureHashStored(bytes memory cid) internal {
        bytes32 hash = keccak256(cid);
        if (cidByHash[hash].length == 0) {
            cidByHash[hash] = cid;
        }
    }
}

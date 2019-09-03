pragma solidity ^0.5.3;

/*
 * UpgradeableProxy is the base contract for all upgradeable contracts.
 * It implements proxy functionality, internal upgrade mechanisms, and internal
 * pause/resume mechanisms.
 *
 * Implementations must handle the specific upgrade & pause rules.
 */
contract UpgradeableProxy {
    event Upgraded(address indexed target);
    event Paused();
    event Resumed();

    bytes32 private constant targetPosition = keccak256("gochain.proxy.target");
    bytes32 private constant pausedPosition = keccak256("gochain.proxy.paused");

    /*
     * Initializes the starting target contract address. The placeholder
     * address is replaced during deployment to the correct address.
     */
    constructor() public {
        address initialTarget = 0xEEfFEEffeEffeeFFeeffeeffeEfFeEffeEFfEeff;
        _upgrade(initialTarget);
    }

    /*
     * Returns the contract address that is currently being proxied to.
     */
    function target() public view returns (address addr) {
        bytes32 pos = targetPosition;
        assembly {
            addr := sload(pos)
        }
    }

    /*
     * Abstract declaration of upgrade function.
     */
    function upgrade(address addr) public;

    /*
     * Updates the target contract address.
     */
    function _upgrade(address addr) internal {
        address current = target();
        require(current != addr);
        bytes32 pos = targetPosition;
        assembly {
            sstore(pos, addr)
        }
        emit Upgraded(addr);
    }

    /*
     * Returns whether the contract is currently paused.
     */
    function paused() public view returns (bool val) {
        bytes32 pos = pausedPosition;
        bytes32 val32 = 0;
        assembly {
            val32 := sload(pos)
        }
        val = val32 != 0;
    }

    /*
     * Abstract declaration of pause function.
     */
    function pause() public;

    /*
     * Abstract declaration of resume function.
     */
    function resume() public;

    /*
     * Marks the contract as paused.
     */
    function _pause() internal {
        bytes32 pos = pausedPosition;
        assembly {
            sstore(pos, 0x1)
        }
        emit Paused();
    }

    /*
     * Marks the contract as resumed (aka unpaused).
     */
    function _resume() internal {
        bytes32 pos = pausedPosition;
        assembly {
            sstore(pos, 0x0)
        }
        emit Resumed();
    }

    /*
     * Passthrough function for all function calls that cannot be found.
     * Functions are delegated to the target contract but maintain the local storage.
     */
    function() payable external {
        bool _paused = paused();
        require(!_paused);

        address _target = target();
        require(_target != address(0));

        assembly {
            let ptr := mload(0x40)
            calldatacopy(ptr, 0, calldatasize)
            let result := delegatecall(gas, _target, ptr, calldatasize, 0, 0)
            let size := returndatasize
            returndatacopy(ptr, 0, size)

            switch result
            case 0 { revert(ptr, size) }
            default { return(ptr, size) }
        }
    }
}

/*
 * OwnerUpgradeableProxy is an upgradeable proxy that only allows the contract
 * owner to upgrade and pause the proxy.
 */
contract OwnerUpgradeableProxy is UpgradeableProxy {
    bytes32 private constant ownerPosition = keccak256("gochain.proxy.owner");

    /*
     * Initializes the proxy and sets the owner.
     */
    constructor() public {
        _setOwner(msg.sender);
    }

    /*
     * Restricts a function to only allow execution by the proxy owner.
     */
    modifier ownerOnly() {
        require(msg.sender == owner());
        _;
    }

    /*
     * Returns the owner of the proxy contract.
     */
    function owner() public view returns (address addr) {
        bytes32 pos = ownerPosition;
        assembly {
            addr := sload(pos)
        }
    }

    /*
     * Sets the owner of the contract.
     */
    function _setOwner(address addr) internal {
        bytes32 pos = ownerPosition;
        assembly {
            sstore(pos, addr)
        }
    }

    /*
     * Upgrades the contract to a new target address. Only allowed by the owner.
     */
    function upgrade(address target) public ownerOnly {
        _upgrade(target);
    }

    /*
     * Pauses the contract and does not allow functions to be executed besides
     * declared functions directly on the proxy (e.g. upgrade(), resume()).
     */
    function pause() public ownerOnly {
        _pause();
    }

    /*
     * Resumes a previously paused contract.
     */
    function resume() public ownerOnly {
        _resume();
    }
}
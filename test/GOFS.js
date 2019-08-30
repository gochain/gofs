const GOFS = artifacts.require("GOFS");
const OwnerUpgradeableProxy = artifacts.require("OwnerUpgradeableProxy");
const truffleAssert = require('truffle-assertions');
const CID = require('cids');

contract("GOFS", async accounts => {
    const rate = 1000;
    const cid = new CID("zb2rhe5P4gXftAwvA4eXQ5HJwsER2owDyS9sKaQRRVQPn93bA");
    let gofs;

    beforeEach(async () => {
        let target = await GOFS.new();
        let proxy = await OwnerUpgradeableProxy.new();
        await proxy.upgrade(target.address);
        gofs = await GOFS.at(proxy.address);
        await gofs.setRate(rate);
    })

    it("pins", async () => {
        const bh = 100;
        const val = rate * bh;
        let tx = await gofs.pin(cid.buffer, {value: val});
        console.log("Gas used to call pin:", tx.receipt.gasUsed);
        truffleAssert.eventEmitted(tx, 'Pinned', (ev) => {
            return ev.user == accounts[0] &&
                ev.cid == web3.utils.keccak256(cid.buffer) &&
                ev.bh == bh;
        });
        assert.equal(await web3.eth.getBalance(gofs.address), val, "Unexpected contract balance");
        let addr = web3.utils.randomHex(20);
        await gofs.withdraw(addr);
        assert.equal(await web3.eth.getBalance(addr), val, "Unexpected account balance");
        assert.equal(await web3.eth.getBalance(gofs.address), 0, "Unexpected contract balance");
    })

    it("deposit wallets", async () => {
        assert.equal(await gofs.wallet(cid.buffer), 0, "Unexpected wallet contract address");

        let tx = await gofs.newWallet(cid.buffer);
        console.log("Gas used to create a new wallet:", tx.receipt.gasUsed);
        let hash = web3.utils.keccak256(cid.buffer);
        truffleAssert.eventEmitted(tx, 'CreatedWallet', (ev) => {
            return ev.user == accounts[0] &&
                ev.cid == hash;
        });
        let addr = await gofs.wallet(cid.buffer);
        assert.notEqual(addr, 0, "Expected a wallet contract address");
        try {
            await gofs.newWallet(cid.buffer);
            assert.fail("Unexpected recreate allowed");
        } catch(error) {
            assert.isTrue(error.toString().includes("Wallet already exists for cid"), "Unexpected error");
        }

        const bh = 100;
        let receipt = await web3.eth.sendTransaction({from: accounts[0], to: addr, value: bh * rate, gas: 300000});
        tx = await truffleAssert.createTransactionResult(gofs, receipt.transactionHash);
        truffleAssert.eventEmitted(tx, 'Pinned', (ev) => {
            return ev.user == accounts[0] &&
                ev.cid == hash &&
                ev.bh == bh;
        });
        assert.equal(await gofs.cidByHash(hash), '0x' + cid.buffer.toString('hex'), "Unexpected CID for hash");
    })

    it("owner", async () => {
        try {
            await gofs.setRate(200, {from:accounts[1]});
            assert.fail("Non-owner was able to set rate");
        } catch(error) {
            assert.isTrue(error.toString().includes("Only owner can call this function"));
        }
        try {
            await gofs.changeOwner(accounts[2], {from:accounts[1]});
            assert.fail("Non-owner was able to change owner");
        } catch(error) {
            assert.isTrue(error.toString().includes("Only owner can call this function"));
        }
        try {
            await gofs.withdraw(accounts[2], {from:accounts[1]});
            assert.fail("Non-owner was able to withdraw");
        } catch(error) {
            assert.isTrue(error.toString().includes("Only owner can call this function"));
        }

        await gofs.changeOwner(accounts[1]);
        await gofs.setRate(200, {from: accounts[1]});
        await gofs.withdraw(accounts[2], {from: accounts[1]});
        await gofs.changeOwner(accounts[0], {from: accounts[1]});
    })
})

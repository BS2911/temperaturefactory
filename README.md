
Rinkeby test network contract deployment address: https://rinkeby.etherscan.io/address/0x903d2ac7e2Ce844855a5e589ba6e9B4A29298C3B#code

- How to ensure we are decentralized and without a single point of failure?
  We can achieve decentralization by deploying contracts on the ETH chain and setting up multiple RPC nodes to submit RPC requests to miners, or if any node fails, the service can continue to operate without impact.

- How do we determine who can submit the temperature?
  By setting up contract owner, administrator, and staff identities in the contract, each identity has its own function and is isolated from each other to ensure that the tracking record temperature is set for the identity user to initiate the record.

- How can we make sure no one is submitting wrong values?
  I wrote a tool check via golang that can be processed by this method before writing data on the rpc chain,only the correct values can be written.

- How do we detect outliers?
  I wrote a tool check via golang that can be processed by this method before writing data on the rpc chain,only the correct values can be written.

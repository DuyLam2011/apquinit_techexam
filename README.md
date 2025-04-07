# apquinit_techexam

## Golang Backend
This application is a Golang backend service that interacts with the Etherscan.io API to provide the following features:

### Features
- Retrieve the last 10 transactions of a specified Ethereum wallet.
- Fetch the current Ethereum balance of a specified wallet.

### Prerequisites
- Go 1.18 or higher installed on your system.
- An Etherscan API key. You can obtain one by signing up at [Etherscan.io](https://etherscan.io/).

### Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/apquinit_techexam.git
    cd apquinit_techexam
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Set up environment variables:
    Create a `.env` file in the root directory and add your Etherscan API key:
    ```
    ETHERSCAN_API_KEY=your_api_key_here
    ```

### Usage
Run the application:
```bash
go run main.go
```

### API Endpoints
1. **Get Last 10 Transactions**
    - **Endpoint:** `/api/v1/transactions`
    - **Method:** `GET`
    - **Query Parameters:**
      - `wallet` (string): Ethereum wallet address.
    - **Response:** JSON array of the last 10 transactions.

2. **Get Wallet Balance**
    - **Endpoint:** `/api/v1/balance`
    - **Method:** `GET`
    - **Query Parameters:**
      - `wallet` (string): Ethereum wallet address.
    - **Response:** JSON object with the wallet balance.

### Example Requests
- Fetch last 10 transactions:
  ```bash
  curl "http://localhost:8080/api/v1/transactions?wallet=0xYourWalletAddress"
  ```

- Fetch wallet balance:
  ```bash
  curl "http://localhost:8080/api/v1/balance?wallet=0xYourWalletAddress"
  ```

### License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Frontend Next.js App

This is a Next.js frontend application that provides a user interface to connect to a MetaMask wallet and interact with a backend service to fetch Ethereum wallet information.

### Features
- Connect to a MetaMask wallet.
- Display the current Ethereum balance of the connected wallet.
- Fetch and display the last 10 transactions of the connected wallet.
- Show additional Ethereum network information such as gas price and block number.

### Prerequisites
- Node.js 16 or higher installed on your system.
- A MetaMask wallet installed in your browser.
- A running instance of the backend service.

### Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/apquinit_techexam.git
    cd apquinit_techexam/front-end
    ```

2. Install dependencies:
    ```bash
    npm install
    ```

3. Set up environment variables:
    Create a `.env.local` file in the `front-end` directory and add the following:
    ```
    NEXT_PUBLIC_BACKEND_URL=http://localhost:8080
    NEXT_PUBLIC_ETHERSCAN_API_KEY=your_etherscan_api_key_here
    ```

### Usage
Run the application:
```bash
npm run dev
```

Open your browser and navigate to `http://localhost:3000`.

### How It Works
1. Click the "Connect Wallet" button to open MetaMask and connect your wallet.
2. Once connected, the app will:
   - Display your wallet address and balance.
   - Fetch and display the last 10 transactions from the backend.
   - Show additional Ethereum network information.

### Example Backend Integration
Ensure the backend service is running and accessible at the URL specified in `NEXT_PUBLIC_BACKEND_URL`. The frontend will make API calls to the backend to fetch wallet data.

### License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Hardhat Smart Contract

This section provides instructions for deploying and interacting with a smart contract using Hardhat, Infura, and the Sepolia test network.

### Prerequisites
- Node.js 16 or higher installed on your system.
- Hardhat installed globally or locally in your project.
- An Infura account and project ID. You can sign up at [Infura.io](https://infura.io/).
- A Sepolia testnet wallet with test ETH. You can obtain test ETH from a [Sepolia faucet](https://faucet.sepolia.dev/).

### Installation
1. Initialize a new Hardhat project:
    ```bash
    mkdir hardhat-project
    cd hardhat-project
    npx hardhat
    ```

2. Install dependencies:
    ```bash
    npm install --save-dev @nomicfoundation/hardhat-toolbox dotenv
    ```

3. Set up environment variables:
    Create a `.env` file in the root directory and add the following:
    ```
    INFURA_PROJECT_ID=your_infura_project_id
    PRIVATE_KEY=your_sepolia_wallet_private_key
    ```

### Configuration
Update the `hardhat.config.js` file to include the Sepolia network:
```javascript
require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();

const { INFURA_PROJECT_ID, PRIVATE_KEY } = process.env;

module.exports = {
  solidity: "0.8.18",
  networks: {
    sepolia: {
      url: `https://sepolia.infura.io/v3/${INFURA_PROJECT_ID}`,
      accounts: [PRIVATE_KEY],
    },
  },
};
```

### Writing a Smart Contract
Create a new file `contracts/MyContract.sol`:
```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract MyContract {
    string public message;

    constructor(string memory _message) {
        message = _message;
    }

    function setMessage(string memory _message) public {
        message = _message;
    }
}
```

### Deployment Script
Create a new file `scripts/deploy.js`:
```javascript
const hre = require("hardhat");

async function main() {
  const MyContract = await hre.ethers.getContractFactory("MyContract");
  const myContract = await MyContract.deploy("Hello, Hardhat!");

  await myContract.deployed();

  console.log("MyContract deployed to:", myContract.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
```

### Deployment
Deploy the contract to the Sepolia network:
```bash
npx hardhat run scripts/deploy.js --network sepolia
```

### Interacting with the Contract
Use Hardhat's console to interact with the deployed contract:
```bash
npx hardhat console --network sepolia
```

Example interaction:
```javascript
const contract = await ethers.getContractAt("MyContract", "deployed_contract_address");
await contract.setMessage("New Message");
const message = await contract.message();
console.log(message);
```

### License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
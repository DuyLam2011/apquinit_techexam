const { ethers } = require("hardhat");

async function main() {
  const [deployer] = await ethers.getSigners();
  console.log("Deploying with:", deployer.address);

  const MyNFT = await ethers.getContractFactory("MyNFT");
  const myNFT = await MyNFT.deploy(deployer.address);
  await myNFT.deployed();

  console.log("MyNFT deployed to:", myNFT.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

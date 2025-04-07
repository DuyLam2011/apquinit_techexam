/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";

import { useState } from "react";
import { ethers } from "ethers";
import axios from "axios";

type Tx = {
  hash: string;
  from: string;
  to: string;
  value: string;
};

export default function Wallet() {
  const [account, setAccount] = useState<string | null>(null);
  const [balance, setBalance] = useState<string>("");
  const [txs, setTxs] = useState<Tx[]>([]);
  const [etherInfor, setEtherInfor] = useState<{ gas_price: string; block_number: number } | null>(null);
  const [error, setError] = useState<string | null>(null);

  const connectWallet = async () => {
    try {
      if (!(window as any).ethereum) throw new Error("MetaMask not installed");
      const provider = new ethers.BrowserProvider((window as any).ethereum);
      const accounts = await provider.send("eth_requestAccounts", []);
      const account = accounts[0];
      setAccount(account);

      const balance = await provider.getBalance(account);
      setBalance(ethers.formatEther(balance));

      // call backend
      const { data } = await axios.get(`http://localhost:8080/api/v1/eth/${account}`);
      setEtherInfor(data);

      // basic tx history from etherscan
      const txRes = await axios.get(`https://api.etherscan.io/api`, {
        params: {
          module: "account",
          action: "txlist",
          address: account,
          sort: "desc",
          apikey: process.env.NEXT_PUBLIC_ETHERSCAN_API_KEY,
        },
      });

      setTxs(txRes.data.result.slice(0, 10));
    } catch (err: any) {
      setError(err.message);
    }
  };

  return (
    <div className="p-4 max-w-xl mx-auto">
      <h1 className="text-2xl font-bold mb-4">Ethereum Wallet Info</h1>

      <button
        className="bg-blue-600 text-white px-4 py-2 rounded"
        onClick={connectWallet}
      >
        Connect Wallet
      </button>

      {error && <p className="text-red-600 mt-2">{error}</p>}

      {account && (
        <>
          <p className="mt-4">Connected: {account}</p>
          <p>Balance: {balance} ETH</p>

          {etherInfor && (
            <div className="mt-2 text-sm">
              <p>🔹 Gas Price: {etherInfor.gas_price}</p>
              <p>🔸 Block Number: {etherInfor.block_number}</p>
            </div>
          )}

          <h2 className="mt-6 text-xl font-semibold">Last 10 Transactions:</h2>
          <ul className="text-sm mt-2 space-y-1">
            {txs.map((tx) => (
              <li key={tx.hash}>
                <p>Hash: {tx.hash.slice(0, 16)}...</p>
                <p>
                  From: {tx.from.slice(0, 12)}... To: {tx.to.slice(0, 12)}...
                </p>
                <p>Value: {ethers.formatEther(tx.value)} ETH</p>
              </li>
            ))}
          </ul>
        </>
      )}
    </div>
  );
}

import json
import requests
from mnemonic import Mnemonic
from bip_utils import Bip39SeedGenerator, Bip44, Bip44Coins, Bip44Changes
from concurrent.futures import ThreadPoolExecutor

class WalletInfo:
    def __init__(self, mnemonic, wif, btc_addr, btc_balance=0):
        self.mnemonic = mnemonic
        self.wif = wif
        self.btc_addr = btc_addr
        self.btc_balance = btc_balance

    def to_dict(self):
        return self.__dict__

def generate_mnemonic():
    mnemo = Mnemonic("english")
    return mnemo.generate(strength=128)

def generate_btc_wallet(mnemonic):
    # Generate seed from mnemonic
    seed = Bip39SeedGenerator(mnemonic).Generate()

    # Generate Bip44 master key
    bip44_mst_ctx = Bip44.FromSeed(seed, Bip44Coins.BITCOIN)

    # Generate account, change, and address keys
    bip44_acc_ctx = bip44_mst_ctx.Purpose().Coin().Account(0)
    bip44_chg_ctx = bip44_acc_ctx.Change(Bip44Changes.CHAIN_EXT)
    bip44_addr_ctx = bip44_chg_ctx.AddressIndex(0)

    wif = bip44_addr_ctx.PrivateKey().ToWif()
    btc_addr = bip44_addr_ctx.PublicKey().ToAddress()

    return wif, btc_addr

def check_btc_balance(address):
    url = f"https://api.blockcypher.com/v1/btc/main/addrs/{address}/balance"
    response = requests.get(url)
    data = response.json()
    print(data)
    balance_satoshis = data.get('final_balance', 0)
    balance_btc = balance_satoshis / 1e8
    return balance_btc

def save_to_file(filename, wallets):
    with open(filename, 'a') as f:
        for wallet in wallets:
            json.dump(wallet.to_dict(), f, indent=2)
            f.write('\n')

def main():
    num_wallets_to_generate = 10

    wallets = []
    for _ in range(num_wallets_to_generate):
        mnemonic = generate_mnemonic()

        wif, btc_addr = generate_btc_wallet(mnemonic)

        wallets.append(WalletInfo(mnemonic, wif, btc_addr))

    with ThreadPoolExecutor(max_workers=10) as executor:
        futures = []
        for wallet in wallets:
            futures.append(executor.submit(check_btc_balance, wallet.btc_addr))

        for i, future in enumerate(futures):
            balance = future.result()
            wallets[i].btc_balance = balance

    save_to_file("all_wallets.json", wallets)

    wallets_with_balance = [wallet for wallet in wallets if wallet.btc_balance > 0]
    if wallets_with_balance:
        save_to_file("wallets_with_balance.json", wallets_with_balance)

    print(f"{num_wallets_to_generate} 个钱包的助记词和WIF私钥已保存到文件 all_wallets.json")
    if wallets_with_balance:
        print("有余额的钱包信息已保存到文件 wallets_with_balance.json")
    else:
        print("没有发现有余额的钱包")

if __name__ == "__main__":
    main()

import hashlib
import base58
import requests
from ecdsa import SigningKey, SECP256k1
from mnemonic import Mnemonic
import time

def generate_mnemonic():
    mnemo = Mnemonic("english")
    words = mnemo.generate(strength=128)
    return words

def mnemonic_to_seed(mnemonic):
    mnemo = Mnemonic("english")
    seed = mnemo.to_seed(mnemonic)
    return seed

def seed_to_private_key(seed):
    return SigningKey.from_string(seed[:32], curve=SECP256k1)

def private_key_to_wif(private_key):
    private_key_bytes = private_key.to_string()
    extended_key = b'\x80' + private_key_bytes
    sha256_1 = hashlib.sha256(extended_key).digest()
    sha256_2 = hashlib.sha256(sha256_1).digest()
    checksum = sha256_2[:4]
    wif = base58.b58encode(extended_key + checksum)
    return wif.decode('utf-8')

def private_key_to_public_key(private_key):
    public_key = private_key.get_verifying_key()
    return public_key.to_string()

def public_key_to_address(public_key):
    sha256 = hashlib.sha256(public_key).digest()
    ripemd160 = hashlib.new('ripemd160', sha256).digest()
    extended_ripemd160 = b'\x00' + ripemd160
    sha256_1 = hashlib.sha256(extended_ripemd160).digest()
    sha256_2 = hashlib.sha256(sha256_1).digest()
    checksum = sha256_2[:4]
    address_bytes = extended_ripemd160 + checksum
    return base58.b58encode(address_bytes).decode('utf-8')

def generate_wallet():
    mnemonic = generate_mnemonic()
    seed = mnemonic_to_seed(mnemonic)
    private_key = seed_to_private_key(seed)
    public_key = private_key_to_public_key(private_key)
    address = public_key_to_address(public_key)
    wif = private_key_to_wif(private_key)
    return {'mnemonic': mnemonic, 'wif': wif, 'address': address}

def check_balance(address):
    url = f"https://blockchain.info/q/addressbalance/{address}"
    response = requests.get(url)
    print(response)
    if response.status_code == 200:
        balance_satoshis = int(response.text)
        balance_btc = balance_satoshis / 1e8  # 将聪（satoshis）转换为比特币
        return balance_btc
    else:
        return None

def save_to_file(filename, wallets):
    with open(filename, 'a') as f:
        for i, wallet in enumerate(wallets):
            f.write(f"Wallet {i + 1}:\n")
            f.write(f"Mnemonic: {wallet['mnemonic']}\n")
            f.write(f"WIF: {wallet['wif']}\n")
            f.write(f"Address: {wallet['address']}\n\n")


while True:
    # 生成指定数量的钱包
    num_wallets_to_generate = 10
    wallets = [generate_wallet() for _ in range(num_wallets_to_generate)]
    # 检查每个钱包的余额
    wallets_with_balance = []
    for wallet in wallets:
        balance = check_balance(wallet['address'])
        if balance is not None and balance > 0:
            wallet['balance'] = balance
            wallets_with_balance.append(wallet)

    # 保存所有生成的钱包信息
    filename_all_wallets = "all_wallets.txt"
    save_to_file(filename_all_wallets, wallets)

    # 保存有余额的钱包信息
    filename_wallets_with_balance = "wallets_with_balance.txt"
    save_to_file(filename_wallets_with_balance, wallets_with_balance)

    if wallets_with_balance:
        print(f"有余额的钱包信息已保存到文件 {filename_wallets_with_balance}")
    else:
        print("0")
    time.sleep(0.01)




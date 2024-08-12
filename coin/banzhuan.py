import ccxt

def exec():

    print(ccxt.__all__)

    mexc = ccxt.mexc({
        'apiKey': 'mx0vglafpbId9gxCyV',
        'secret': 'e1f01d0bf22645789ff13e85a9371114',
    })

    gateio = ccxt.gateio({
        'apiKey': 'mx0vglafpbId9gxCyV',
        'secret': 'e1f01d0bf22645789ff13e85a9371114',
    })
 

    # 设置交易参数
    symbol = 'KAS/USDT'
    quantity = 10  # 购买或出售的数量

    # 在Mexc买入
    order = mexc.create_order(
        symbol=symbol,
        side='buy',
        type='market',
        amount=quantity
    )


    # 获取最新价格并执行交易
    binance_ticker = mexc.fetch_ticker(symbol)
    huobi_ticker = gateio.fetch_ticker(symbol)
    binance_price = binance_ticker['last']
    huobi_price = huobi_ticker['last']


if __name__ == '__main__':
    exec()

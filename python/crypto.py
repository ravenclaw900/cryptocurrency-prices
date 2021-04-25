import pip._vendor.requests as requests, json
cryptoResp = requests.get('https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum,dogecoin&vs_currencies=usd&include_24hr_change=true')

cryptoData = json.loads(cryptoResp.text)

print("Bitcoin is currently worth $" + str(cryptoData['bitcoin']['usd']) + ", and it has changed by " + str(round(cryptoData['bitcoin']['usd_24h_change'], 3)) + "% over the past 24 hours.")
print("Ethereum is currently worth $" + str(cryptoData['ethereum']['usd']) + ", and it has changed by " + str(round(cryptoData['ethereum']['usd_24h_change'], 3)) + "% over the past 24 hours.")
print("Dogecoin is currently worth $" + str(cryptoData['dogecoin']['usd']) + ", and it has changed by " + str(round(cryptoData['dogecoin']['usd_24h_change'], 3)) + "% over the past 24 hours.")

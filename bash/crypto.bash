#/bin/bash

cryptoData=$(curl -s -X GET "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin%2Cethereum%2Cdogecoin&vs_currencies=usd&include_24hr_change=true")

echo "Bitcoin is currently worth \$$(echo $cryptoData | jq '.bitcoin.usd'), and it has changed by $(printf "%.3f" $(echo $cryptoData | jq '.bitcoin.usd_24h_change'))% over the past 24 hours."
echo "Ethereum is currently worth \$$(echo $cryptoData | jq '.ethereum.usd'), and it has changed by $(printf "%.3f" $(echo $cryptoData | jq '.ethereum.usd_24h_change'))% over the past 24 hours."
echo "Dogecoin is currently worth \$$(echo $cryptoData | jq '.dogecoin.usd'), and it has changed by $(printf "%.3f" $(echo $cryptoData | jq '.dogecoin.usd_24h_change'))% over the past 24 hours."

var cryptoReq = new XMLHttpRequest();
cryptoReq.open( "GET", "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum,dogecoin&vs_currencies=usd&include_24hr_change=true", false );
cryptoReq.send()

var cryptoData = JSON.parse(cryptoReq.responseText)

document.getElementById("bitcoin").textContent = `Bitcoin is currently worth $${cryptoData.bitcoin.usd}, and it has changed by ${cryptoData.bitcoin.usd_24h_change.toFixed(3)}% over the past 24 hours.`
document.getElementById("ethereum").textContent = `Ethereum is currently worth $${cryptoData.ethereum.usd}, and it has changed by ${cryptoData.ethereum.usd_24h_change.toFixed(3)}% over the past 24 hours.`
document.getElementById("dogecoin").textContent = `Dogecoin is currently worth $${cryptoData.dogecoin.usd}, and it has changed by ${cryptoData.dogecoin.usd_24h_change.toFixed(3)}% over the past 24 hours.`

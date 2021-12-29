use std::collections::HashMap;

fn main() {
    let cryptoData = reqwest::blocking::get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum,dogecoin&vs_currencies=usd&include_24hr_change=true").unwrap()
        .json::<HashMap<String, HashMap<String, f32>>>().unwrap();
    println!("Bitcoin is currently worth ${}, and it has changed by {}% over the past 24 hours.", cryptoData["bitcoin"]["usd"], (cryptoData["bitcoin"]["usd_24h_change"] * 1000.0).round() / 1000.0);
    println!("Ethereum is currently worth ${}, and it has changed by {}% over the past 24 hours.", cryptoData["ethereum"]["usd"], (cryptoData["ethereum"]["usd_24h_change"] * 1000.0).round() / 1000.0);
    println!("Dogecoin is currently worth ${}, and it has changed by {}% over the past 24 hours.", cryptoData["dogecoin"]["usd"], (cryptoData["dogecoin"]["usd_24h_change"] * 1000.0).round() / 1000.0);
}
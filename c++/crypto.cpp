#include <string>
#include <iostream>
#include <iomanip> 
#include <curl/curl.h>
#include <nlohmann/json.hpp>

using namespace std;
using json = nlohmann::json;

size_t write_to_string(void *ptr, size_t size, size_t count, void *stream) {
  ((string*)stream)->append((char*)ptr, 0, size*count);
  return size*count;
}

int main(void) {
  CURL *curl;
  CURLcode cryptoReq;
  curl = curl_easy_init();
  if (curl) {
    curl_easy_setopt(curl, CURLOPT_URL, "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum,dogecoin&vs_currencies=usd&include_24hr_change=true");

    string cryptoResp;
    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_to_string);
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, &cryptoResp);

    cryptoReq = curl_easy_perform(curl);
    curl_easy_cleanup(curl);

    auto cryptoData = json::parse(cryptoResp);

    cout << "Bitcoin is currently worth $" << cryptoData["bitcoin"]["usd"] << ", and it has changed by " << setprecision(4) << (double)(cryptoData["bitcoin"]["usd_24h_change"]) << "% over the past 24 hours." << endl;
    cout << "Ethereum is currently worth $" << cryptoData["ethereum"]["usd"] << ", and it has changed by " << setprecision(4) << (double)(cryptoData["ethereum"]["usd_24h_change"]) << "% over the past 24 hours." << endl;
    cout << "Dogecoin is currently worth $" << cryptoData["dogecoin"]["usd"] << ", and it has changed by " << setprecision(4) << (double)(cryptoData["dogecoin"]["usd_24h_change"]) << "% over the past 24 hours." << endl;
  }
}
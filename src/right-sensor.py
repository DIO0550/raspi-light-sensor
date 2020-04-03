from gpiozero import MCP3008
from time import sleep
import urllib.request
import json

# 明るさの閾値
Bright = 26.0
# 値取得URL
FetchConferenceRoomAPI = ""
# 値更新URL
ConferenceRoomUpdateAPI = ""

adc = MCP3008(channel=0, device=0)

# 監視対象の部屋のID
Monitoring_ID = "7A"
# 前の値
# 値が変わっていない場合に、APIを無駄に呼び出さないようにする
before_is_use = False

# === GETで現在の値を取得する ===
method = "GET"
headers = {
    'Content-Type': 'application/json',
}
request = urllib.request.Request(FetchConferenceRoomAPI, method=method)

try:
    with urllib.request.urlopen(request) as response:
        body = json.load(response)
        for room_info in body:
            if room_info["name"] == Monitoring_ID:
                before_is_use = room_info["is_use"]
                print(room_info)
                break

except urllib.error.HTTPError as err:
    print(err.code)

except urllib.error.URLError as err:
    print(err.reason)


# === 5秒ごと明るさをチェックして、データを更新する === 
while True:
    brightness = (adc.value * 3.3) * 100
    is_use: bool = brightness > Bright
    if is_use == before_is_use:
       continue
    
    before_is_use = is_use
    param = {"name": "7A", "is_use": is_use}
    method = "PUT"
    headers = {
        'Content-Type': 'application/json',
    }
    json_data = json.dumps(param).encode('utf-8')
    request = urllib.request.Request(ConferenceRoomUpdateAPI, data=json_data, headers=headers, method=method)
    try:
        with urllib.request.urlopen(request) as response:
            body = json.load(response)
            print(body)
    except urllib.error.HTTPError as err:
        print(err.code)
    except urllib.error.URLError as err:
        print(err.reason)

    sleep(5)

from gpiozero import MCP3008
from time import sleep
import urllib.request
import json

Bright = 60.0
ConferenceRoomUpdateAPI = "http://localhost:8080/conferenceroom/update"
adc = MCP3008(channel=0, device=0)

while True:
    brightness = (adc.value * 3.3) * 100
    usage_situation: bool = brightness > Bright
    param = {"name": "7階D", "usage_situation": usage_situation}
    method = "POST"
    headers = {
        'Content-Type': 'application/json',
    }
    json_data = json.dumps(param).encode('utf-8')
    print(json_data) 
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
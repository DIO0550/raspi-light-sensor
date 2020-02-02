from gpiozero import MCP3008
from time import sleep
adc = MCP3008(channel=0, device=0)

while True:
    brightness = (adc.value * 3.3) * 100
    # TODO: 明るいかくらいかの判別
    sleep(5)
from gpiozero import MCP3008
from time import sleep
tmp = MCP3008(channel=0, device=0)

while True:
    temperature = (tmp.value * 3.3) * 100
    print (round(temperature,5))
    sleep(0.5)
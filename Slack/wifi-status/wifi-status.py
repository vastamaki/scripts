import os
import re
from slack import WebClient
from slack.errors import SlackApiError
from subprocess import check_output

home_wifi = "Wifi_1"
work_wifi = "Wifi_2"
token = "TOKEN_HERE"
client = WebClient(token)
wifi_interface = "wlp1s0"


def getSSID():
    scanoutput = str(check_output(["iwlist", wifi_interface, "scan"]))

    for line in scanoutput.split():
        if line.startswith("ESSID"):
            return line.split('"')[1]


def saveNewWifi(wifi_name):
    file = open("wifi.txt", "w")
    file.write(wifi_name)
    file.close()


def setNewStatus(emoji, message):
    client.users_profile_set(
        profile={
            "status_text": message,
            "status_emoji": emoji
        }
    )


wifi_name = getSSID()
old_wifi_name = open("wifi.txt", "r").read()

if wifi_name == old_wifi_name:
    print('Status change not required.')
else:
    saveNewWifi(wifi_name)
    if (wifi_name != old_wifi_name) and (wifi_name == home_wifi):
        print('Home wifi detected! Changing status to home :)')
        setNewStatus(":home:", "@ Home Office")
    elif (wifi_name != old_wifi_name) and (wifi_name == work_wifi):
        print('Work wifi detected! Changing status to office :)')
        setNewStatus(":office:", "@ Office")

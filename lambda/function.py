import json
import os
import urllib.parse
import requests
import pprint

print('Loading function')

def lambda_handler(event, context):
    # print("Received event: " + json.dumps(event, indent=2))

    # Get the object from the event
    bucket = event['Records'][0]['s3']['bucket']['name']
    key = urllib.parse.unquote_plus(event['Records'][0]['s3']['object']['key'], encoding='utf-8')

    # create dict ready for posting
    payload = {
        'bucket': bucket,
        'key':    key
    }

    pprint.pprint(payload)

    # post it
    url = os.getenv('MRU_API_URL')

    try:
        print("Posting to {}...".format(url))
        r = requests.post(url, json=payload)

    except Exception as e:
        print(e)
        print("Failed to post event to {}".format(url))
        raise e


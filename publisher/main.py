import os
import time
import urllib.request
import json

BASE_URL = os.getenv("BASE_URL", "http://localhost:10526")
SUBJECT = os.getenv("SUBJECT", "test-subject")
SCHEDULER_INTERVAL = int(os.getenv("SCHEDULER_INTERVAL", "10"))


def post_request(url, data):
    try:
        req = urllib.request.Request(
            url, data=data, method="POST", headers={"Content-Type": "application/json"}
        )
        with urllib.request.urlopen(req) as response:
            if response.status == 200:
                print(f"Request successful: {data}")
            else:
                print(f"Failed to send request. Status code: {response.status}")
    except Exception as e:
        print(f"Error sending request: {e}")


def publish_message():
    url = f"{BASE_URL}/queue/push?id={SUBJECT}"

    message = {"message": "Hello, World!"}
    data = json.dumps(message).encode("utf-8")

    post_request(url, data)


def simple_scheduler():
    while True:
        publish_message()
        time.sleep(SCHEDULER_INTERVAL)


if __name__ == "__main__":
    print("Starting message publisher...")
    simple_scheduler()

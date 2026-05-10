import os
import time
import urllib.request
import json

BASE_URL = os.getenv("BASE_URL", "http://localhost:10526")
SUBJECT = os.getenv("SUBJECT", "test-subject")
SCHEDULER_INTERVAL = int(os.getenv("SCHEDULER_INTERVAL", "10"))


def get_request(url):
    try:
        with urllib.request.urlopen(url) as response:
            if response.status == 200:
                data = json.loads(response.read())
                print(f"Consumed message: {data}")
            else:
                print(f"Failed to consume message. Status code: {response.status}")
    except Exception as e:
        print(f"Error consuming message: {e}")


def consume_message():
    url = f"{BASE_URL}/queue/pop?id={SUBJECT}"

    get_request(url)


def simple_scheduler():
    while True:
        consume_message()
        time.sleep(SCHEDULER_INTERVAL)


if __name__ == "__main__":
    print("Starting message consumer...")
    simple_scheduler()

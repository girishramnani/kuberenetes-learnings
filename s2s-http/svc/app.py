from flask import Flask, jsonify
import requests
import os
import json

app = Flask(__name__)


@app.route("/upcase/<word>")
def upcase(word):
    headers = {'content-type': 'application/json'}
    return requests.post("http://upcase:8080", data=json.dumps({"key": word}), headers=headers).json()


if __name__ == "__main__":
    port = int(os.environ.get("PORT", 5000))
    app.run(debug=True, host='0.0.0.0', port=port)

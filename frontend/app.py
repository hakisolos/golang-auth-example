from flask import Flask, render_template, request, jsonify
from flask_cors import CORS
import os

app = Flask(__name__, template_folder='.', static_folder='static')
CORS(app)

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/<path:page>')
def catch_all(page):
    return render_template('index.html')

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=3000, debug=False)

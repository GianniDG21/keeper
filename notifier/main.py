from Flask import Flask, request, jsonify
import logging

logging.basicConfig(level=logging.INFO, format='[NotifierService] %(asctime)s - %(message)s')

app = Flask(__name__)

@app.route('/notify', methods=['POST'])
def receive_notification():
    data = request.json
    
    if not data:
        logging.warning("Received empty notification")
        return jsonify({"error": "Empty notification"}), 400
    
    logging.info(f"Received notification: {data}")

    return jsonify({"status": "Notification received"}), 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001)
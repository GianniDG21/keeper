from flask import Flask, request, jsonify
import logging

# Configure logging with INFO level and custom format for the NotifierService
logging.basicConfig(level=logging.INFO, format='[NotifierService] %(asctime)s - %(message)s')

app = Flask(__name__)

@app.route('/notify', methods=['POST'])
def receive_notification():
    """
    Handle incoming notification requests.
    
    Expects JSON payload in request body.
    Returns success status or error if payload is invalid.
    """
    data = request.json
    
    # Validate request has JSON data
    if not data:
        logging.warning("Received empty notification")
        return jsonify({"error": "Empty notification"}), 400
    
    # Log the received notification for monitoring purposes
    logging.info(f"Received notification: {data}")

    return jsonify({"status": "Notification received"}), 200

if __name__ == '__main__':
    # Start Flask development server accessible from any host
    app.run(host='0.0.0.0', port=5001)
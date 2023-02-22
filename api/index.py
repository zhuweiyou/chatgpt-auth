from http.server import BaseHTTPRequestHandler
from OpenAIAuth import Authenticator


class handler(BaseHTTPRequestHandler):

    def do_GET(self):
        self.send_response(200)
        self.send_header('Content-type', 'application/json; charset=utf-8')
        self.end_headers()
        auth = Authenticator(
            email_address='dhfkvu@outlook.com', password='aMAfTEDw')
        auth.begin()
        self.wfile.write(auth.get_access_token().encode('utf-8'))
        return
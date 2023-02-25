from http.server import BaseHTTPRequestHandler
from OpenAIAuth import Authenticator


class handler(BaseHTTPRequestHandler):

    def do_POST(self):
        self.send_response(200)
        self.send_header('Content-type', 'text/plain')
        self.end_headers()
        auth = Authenticator(
            email_address=self.headers.get('X-Email'), password=self.headers.get('X-Password'))
        auth.begin()
        self.wfile.write(auth.get_access_token().encode('utf-8'))
        return

#!/usr/bin/env python3

# Allow direct execution
import os
import sys

import pytest

sys.path.insert(0, os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

import gzip
import http.client
import http.cookiejar
import http.server
import io
import pathlib
import random
import ssl
import tempfile
import threading
import time
import urllib.error
import urllib.request
import warnings
import zlib
from email.message import Message
from http.cookiejar import CookieJar

from test.helper import FakeYDL, http_server_port
from yt_dlp.cookies import YoutubeDLCookieJar
from yt_dlp.dependencies import brotli
from yt_dlp.networking import (
    HEADRequest,
    PUTRequest,
    Request,
    RequestDirector,
    RequestHandler,
    Response,
)
from yt_dlp.networking._urllib import UrllibRH
from yt_dlp.networking.exceptions import (
    CertificateVerifyError,
    HTTPError,
    IncompleteRead,
    NoSupportingHandlers,
    RequestError,
    SSLError,
    TransportError,
    UnsupportedRequest,
)
from yt_dlp.utils._utils import _YDLLogger as FakeLogger
from yt_dlp.utils.networking import HTTPHeaderDict

TEST_DIR = os.path.dirname(os.path.abspath(__file__))


def _build_proxy_handler(name):
    class HTTPTestRequestHandler(http.server.BaseHTTPRequestHandler):
        proxy_name = name

        def log_message(self, format, *args):
            pass

        def do_GET(self):
            self.send_response(200)
            self.send_header('Content-Type', 'text/plain; charset=utf-8')
            self.end_headers()
            self.wfile.write('{self.proxy_name}: {self.path}'.format(self=self).encode())
    return HTTPTestRequestHandler


class HTTPTestRequestHandler(http.server.BaseHTTPRequestHandler):
    protocol_version = 'HTTP/1.1'

    def log_message(self, format, *args):
        pass

    def _headers(self):
        payload = str(self.headers).encode()
        self.send_response(200)
        self.send_header('Content-Type', 'application/json')
        self.send_header('Content-Length', str(len(payload)))
        self.end_headers()
        self.wfile.write(payload)

    def _redirect(self):
        self.send_response(int(self.path[len('/redirect_'):]))
        self.send_header('Location', '/method')
        self.send_header('Content-Length', '0')
        self.end_headers()

    def _method(self, method, payload=None):
        self.send_response(200)
        self.send_header('Content-Length', str(len(payload or '')))
        self.send_header('Method', method)
        self.end_headers()
        if payload:
            self.wfile.write(payload)

    def _status(self, status):
        payload = f'<html>{status} NOT FOUND</html>'.encode()
        self.send_response(int(status))
        self.send_header('Content-Type', 'text/html; charset=utf-8')
        self.send_header('Content-Length', str(len(payload)))
        self.end_headers()
        self.wfile.write(payload)

    def _read_data(self):
        if 'Content-Length' in self.headers:
            return self.rfile.read(int(self.headers['Content-Length']))

    def do_POST(self):
        data = self._read_data() + str(self.headers).encode()
        if self.path.startswith('/redirect_'):
            self._redirect()
        elif self.path.startswith('/method'):
            self._method('POST', data)
        elif self.path.startswith('/headers'):
            self._headers()
        else:
            self._status(404)

    def do_HEAD(self):
        if self.path.startswith('/redirect_'):
            self._redirect()
        elif self.path.startswith('/method'):
            self._method('HEAD')
        else:
            self._status(404)

    def do_PUT(self):
        data = self._read_data() + str(self.headers).encode()
        if self.path.startswith('/redirect_'):
            self._redirect()
        elif self.path.startswith('/method'):
            self._method('PUT', data)
        else:
            self._status(404)

    def do_GET(self):
        if self.path == '/video.html':
            payload = b'<html><video src="/vid.mp4" /></html>'
            self.send_response(200)
            self.send_header('Content-Type', 'text/html; charset=utf-8')
            self.send_header('Content-Length', str(len(payload)))
            self.end_headers()
            self.wfile.write(payload)
        elif self.path == '/vid.mp4':
            payload = b'\x00\x00\x00\x00\x20\x66\x74[video]'
            self.send_response(200)
            self.send_header('Content-Type', 'video/mp4')
            self.send_header('Content-Length', str(len(payload)))
            self.end_headers()
            self.wfile.write(payload)
        elif self.path == '/%E4%B8%AD%E6%96%87.html':
            payload = b'<html><video src="/vid.mp4" /></html>'
            self.send_response(200)
            self.send_header('Content-Type', 'text/html; charset=utf-8')
            self.send_header('Content-Length', str(len(payload)))
            self.end_headers()
            self.wfile.write(payload)
        elif self.path == '/%c7%9f':
            payload = b'<html><video src="/vid.mp4" /></html>'
            self.send_response(200)
            self.send_header('Content-Type', 'text/html; charset=utf-8')
            self.send_header('Content-Length', str(len(payload)))
            self.end_headers()
            self.wfile.write(payload)
        elif self.path.startswith('/redirect_loop'):
            self.send_response(301)
            self.send_header('Location', self.path)
            self.send_header('Content-Length', '0')
            self.end_headers()
        elif self.path == '/redirect_dotsegments':
            self.send_response(301)
            # redirect to /headers but with dot segments before
            self.send_header('Location', '/a/b/./../../headers')
            self.send_header('Content-Length', '0')
            self.end_headers()
        elif self.path.startswith('/redirect_'):
            self._redirect()
        elif self.path.startswith('/method'):
            self._method('GET', str(self.headers).encode())
        elif self.path.startswith('/headers'):
            self._headers()
        elif self.path.startswith('/308-to-headers'):
            self.send_response(308)
            self.send_header('Location', '/headers')
            self.send_header('Content-Length', '0')
            self.end_headers()
        elif self.path == '/trailing_garbage':
            payload = b'<html><video src="/vid.mp4" /></html>'
            self.send_response(200)
            self.send_header('Content-Type', 'text/html; charset=utf-8')
            self.send_header('Content-Encoding', 'gzip')
            buf = io.BytesIO()
            with gzip.GzipFile(fileobj=buf, mode='wb') as f:
                f.write(payload)
            compressed = buf.getvalue() + b'trailing garbage'
            self.send_header('Content-Length', str(len(compressed)))
            self.end_headers()
            self.wfile.write(compressed)
        elif self.path == '/302-non-ascii-redirect':
            new_url = f'http://127.0.0.1:{http_server_port(self.server)}/中文.html'
            self.send_response(301)
            self.send_header('Location', new_url)
            self.send_header('Content-Length', '0')
            self.end_headers()
        elif self.path == '/content-encoding':
            encodings = self.headers.get('ytdl-encoding', '')
            payload = b'<html><video src="/vid.mp4" /></html>'
            for encoding in filter(None, (e.strip() for e in encodings.split(','))):
                if encoding == 'br' and brotli:
                    payload = brotli.compress(payload)
                elif encoding == 'gzip':
                    buf = io.BytesIO()
                    with gzip.GzipFile(fileobj=buf, mode='wb') as f:
                        f.write(payload)
                    payload = buf.getvalue()
                elif encoding == 'deflate':
                    payload = zlib.compress(payload)
                elif encoding == 'unsupported':
                    payload = b'raw'
                    break
                else:
                    self._status(415)
                    return
            self.send_response(200)
            self.send_header('Content-Encoding', encodings)
            self.send_header('Content-Length', str(len(payload)))
            self.end_headers()
            self.wfile.write(payload)
        elif self.path.startswith('/gen_'):
            payload = b'<html></html>'
            self.send_response(int(self.path[len('/gen_'):]))
            self.send_header('Content-Type', 'text/html; charset=utf-8')
            self.send_header('Content-Length', str(len(payload)))
            self.end_headers()
            self.wfile.write(payload)
        elif self.path.startswith('/incompleteread'):
            payload = b'<html></html>'
            self.send_response(200)
            self.send_header('Content-Type', 'text/html; charset=utf-8')
            self.send_header('Content-Length', '234234')
            self.end_headers()
            self.wfile.write(payload)
            self.finish()
        elif self.path.startswith('/timeout_'):
            time.sleep(int(self.path[len('/timeout_'):]))
            self._headers()
        elif self.path == '/source_address':
            payload = str(self.client_address[0]).encode()
            self.send_response(200)
            self.send_header('Content-Type', 'text/html; charset=utf-8')
            self.send_header('Content-Length', str(len(payload)))
            self.end_headers()
            self.wfile.write(payload)
            self.finish()
        else:
            self._status(404)

    def send_header(self, keyword, value):
        """
        Forcibly allow HTTP server to send non percent-encoded non-ASCII characters in headers.
        This is against what is defined in RFC 3986, however we need to test we support this
        since some sites incorrectly do this.
        """
        if keyword.lower() == 'connection':
            return super().send_header(keyword, value)

        if not hasattr(self, '_headers_buffer'):
            self._headers_buffer = []

        self._headers_buffer.append(f'{keyword}: {value}\r\n'.encode())


def validate_and_send(rh, req):
    rh.validate(req)
    return rh.send(req)


class TestRequestHandlerBase:
    @classmethod
    def setup_class(cls):
        cls.http_httpd = http.server.ThreadingHTTPServer(
            ('127.0.0.1', 0), HTTPTestRequestHandler)
        cls.http_port = http_server_port(cls.http_httpd)
        cls.http_server_thread = threading.Thread(target=cls.http_httpd.serve_forever)
        # FIXME: we should probably stop the http server thread after each test
        # See: https://github.com/yt-dlp/yt-dlp/pull/7094#discussion_r1199746041
        cls.http_server_thread.daemon = True
        cls.http_server_thread.start()

        # HTTPS server
        certfn = os.path.join(TEST_DIR, 'testcert.pem')
        cls.https_httpd = http.server.ThreadingHTTPServer(
            ('127.0.0.1', 0), HTTPTestRequestHandler)
        sslctx = ssl.SSLContext(ssl.PROTOCOL_TLS_SERVER)
        sslctx.load_cert_chain(certfn, None)
        cls.https_httpd.socket = sslctx.wrap_socket(cls.https_httpd.socket, server_side=True)
        cls.https_port = http_server_port(cls.https_httpd)
        cls.https_server_thread = threading.Thread(target=cls.https_httpd.serve_forever)
        cls.https_server_thread.daemon = True
        cls.https_server_thread.start()


class TestHTTPRequestHandler(TestRequestHandlerBase):
    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_verify_cert(self, handler):
        with handler() as rh:
            with pytest.raises(CertificateVerifyError):
                validate_and_send(rh, Request(f'https://127.0.0.1:{self.https_port}/headers'))

        with handler(verify=False) as rh:
            r = validate_and_send(rh, Request(f'https://127.0.0.1:{self.https_port}/headers'))
            assert r.status == 200
            r.close()

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_ssl_error(self, handler):
        # HTTPS server with too old TLS version
        # XXX: is there a better way to test this than to create a new server?
        https_httpd = http.server.ThreadingHTTPServer(
            ('127.0.0.1', 0), HTTPTestRequestHandler)
        sslctx = ssl.SSLContext(ssl.PROTOCOL_TLS_SERVER)
        https_httpd.socket = sslctx.wrap_socket(https_httpd.socket, server_side=True)
        https_port = http_server_port(https_httpd)
        https_server_thread = threading.Thread(target=https_httpd.serve_forever)
        https_server_thread.daemon = True
        https_server_thread.start()

        with handler(verify=False) as rh:
            with pytest.raises(SSLError, match='sslv3 alert handshake failure') as exc_info:
                validate_and_send(rh, Request(f'https://127.0.0.1:{https_port}/headers'))
            assert not issubclass(exc_info.type, CertificateVerifyError)

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_percent_encode(self, handler):
        with handler() as rh:
            # Unicode characters should be encoded with uppercase percent-encoding
            res = validate_and_send(rh, Request(f'http://127.0.0.1:{self.http_port}/中文.html'))
            assert res.status == 200
            res.close()
            # don't normalize existing percent encodings
            res = validate_and_send(rh, Request(f'http://127.0.0.1:{self.http_port}/%c7%9f'))
            assert res.status == 200
            res.close()

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_remove_dot_segments(self, handler):
        with handler() as rh:
            # This isn't a comprehensive test,
            # but it should be enough to check whether the handler is removing dot segments
            res = validate_and_send(rh, Request(f'http://127.0.0.1:{self.http_port}/a/b/./../../headers'))
            assert res.status == 200
            assert res.url == f'http://127.0.0.1:{self.http_port}/headers'
            res.close()

            res = validate_and_send(rh, Request(f'http://127.0.0.1:{self.http_port}/redirect_dotsegments'))
            assert res.status == 200
            assert res.url == f'http://127.0.0.1:{self.http_port}/headers'
            res.close()

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_unicode_path_redirection(self, handler):
        with handler() as rh:
            r = validate_and_send(rh, Request(f'http://127.0.0.1:{self.http_port}/302-non-ascii-redirect'))
            assert r.url == f'http://127.0.0.1:{self.http_port}/%E4%B8%AD%E6%96%87.html'
            r.close()

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_raise_http_error(self, handler):
        with handler() as rh:
            for bad_status in (400, 500, 599, 302):
                with pytest.raises(HTTPError):
                    validate_and_send(rh, Request('http://127.0.0.1:%d/gen_%d' % (self.http_port, bad_status)))

            # Should not raise an error
            validate_and_send(rh, Request('http://127.0.0.1:%d/gen_200' % self.http_port)).close()

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_response_url(self, handler):
        with handler() as rh:
            # Response url should be that of the last url in redirect chain
            res = validate_and_send(rh, Request(f'http://127.0.0.1:{self.http_port}/redirect_301'))
            assert res.url == f'http://127.0.0.1:{self.http_port}/method'
            res.close()
            res2 = validate_and_send(rh, Request(f'http://127.0.0.1:{self.http_port}/gen_200'))
            assert res2.url == f'http://127.0.0.1:{self.http_port}/gen_200'
            res2.close()

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_redirect(self, handler):
        with handler() as rh:
            def do_req(redirect_status, method, assert_no_content=False):
                data = b'testdata' if method in ('POST', 'PUT') else None
                res = validate_and_send(
                    rh, Request(f'http://127.0.0.1:{self.http_port}/redirect_{redirect_status}', method=method, data=data))

                headers = b''
                data_sent = b''
                if data is not None:
                    data_sent += res.read(len(data))
                    if data_sent != data:
                        headers += data_sent
                        data_sent = b''

                headers += res.read()

                if assert_no_content or data is None:
                    assert b'Content-Type' not in headers
                    assert b'Content-Length' not in headers
                else:
                    assert b'Content-Type' in headers
                    assert b'Content-Length' in headers

                return data_sent.decode(), res.headers.get('method', '')

            # A 303 must either use GET or HEAD for subsequent request
            assert do_req(303, 'POST', True) == ('', 'GET')
            assert do_req(303, 'HEAD') == ('', 'HEAD')

            assert do_req(303, 'PUT', True) == ('', 'GET')

            # 301 and 302 turn POST only into a GET
            assert do_req(301, 'POST', True) == ('', 'GET')
            assert do_req(301, 'HEAD') == ('', 'HEAD')
            assert do_req(302, 'POST', True) == ('', 'GET')
            assert do_req(302, 'HEAD') == ('', 'HEAD')

            assert do_req(301, 'PUT') == ('testdata', 'PUT')
            assert do_req(302, 'PUT') == ('testdata', 'PUT')

            # 307 and 308 should not change method
            for m in ('POST', 'PUT'):
                assert do_req(307, m) == ('testdata', m)
                assert do_req(308, m) == ('testdata', m)

            assert do_req(307, 'HEAD') == ('', 'HEAD')
            assert do_req(308, 'HEAD') == ('', 'HEAD')

            # These should not redirect and instead raise an HTTPError
            for code in (300, 304, 305, 306):
                with pytest.raises(HTTPError):
                    do_req(code, 'GET')

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_request_cookie_header(self, handler):
        # We should accept a Cookie header being passed as in normal headers and handle it appropriately.
        with handler() as rh:
            # Specified Cookie header should be used
            res = validate_and_send(
                rh, Request(
                    f'http://127.0.0.1:{self.http_port}/headers',
                    headers={'Cookie': 'test=test'})).read().decode()
            assert 'Cookie: test=test' in res

            # Specified Cookie header should be removed on any redirect
            res = validate_and_send(
                rh, Request(
                    f'http://127.0.0.1:{self.http_port}/308-to-headers',
                    headers={'Cookie': 'test=test'})).read().decode()
            assert 'Cookie: test=test' not in res

        # Specified Cookie header should override global cookiejar for that request
        cookiejar = YoutubeDLCookieJar()
        cookiejar.set_cookie(http.cookiejar.Cookie(
            version=0, name='test', value='ytdlp', port=None, port_specified=False,
            domain='127.0.0.1', domain_specified=True, domain_initial_dot=False, path='/',
            path_specified=True, secure=False, expires=None, discard=False, comment=None,
            comment_url=None, rest={}))

        with handler(cookiejar=cookiejar) as rh:
            data = validate_and_send(
                rh, Request(f'http://127.0.0.1:{self.http_port}/headers', headers={'cookie': 'test=test'})).read()
            assert b'Cookie: test=ytdlp' not in data
            assert b'Cookie: test=test' in data

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_redirect_loop(self, handler):
        with handler() as rh:
            with pytest.raises(HTTPError, match='redirect loop'):
                validate_and_send(rh, Request(f'http://127.0.0.1:{self.http_port}/redirect_loop'))

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_incompleteread(self, handler):
        with handler(timeout=2) as rh:
            with pytest.raises(IncompleteRead):
                validate_and_send(rh, Request('http://127.0.0.1:%d/incompleteread' % self.http_port)).read()

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_cookies(self, handler):
        cookiejar = YoutubeDLCookieJar()
        cookiejar.set_cookie(http.cookiejar.Cookie(
            0, 'test', 'ytdlp', None, False, '127.0.0.1', True,
            False, '/headers', True, False, None, False, None, None, {}))

        with handler(cookiejar=cookiejar) as rh:
            data = validate_and_send(rh, Request(f'http://127.0.0.1:{self.http_port}/headers')).read()
            assert b'Cookie: test=ytdlp' in data

        # Per request
        with handler() as rh:
            data = validate_and_send(
                rh, Request(f'http://127.0.0.1:{self.http_port}/headers', extensions={'cookiejar': cookiejar})).read()
            assert b'Cookie: test=ytdlp' in data

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_headers(self, handler):

        with handler(headers=HTTPHeaderDict({'test1': 'test', 'test2': 'test2'})) as rh:
            # Global Headers
            data = validate_and_send(rh, Request(f'http://127.0.0.1:{self.http_port}/headers')).read()
            assert b'Test1: test' in data

            # Per request headers, merged with global
            data = validate_and_send(rh, Request(
                f'http://127.0.0.1:{self.http_port}/headers', headers={'test2': 'changed', 'test3': 'test3'})).read()
            assert b'Test1: test' in data
            assert b'Test2: changed' in data
            assert b'Test2: test2' not in data
            assert b'Test3: test3' in data

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_timeout(self, handler):
        with handler() as rh:
            # Default timeout is 20 seconds, so this should go through
            validate_and_send(
                rh, Request(f'http://127.0.0.1:{self.http_port}/timeout_3'))

        with handler(timeout=0.5) as rh:
            with pytest.raises(TransportError):
                validate_and_send(
                    rh, Request(f'http://127.0.0.1:{self.http_port}/timeout_1'))

            # Per request timeout, should override handler timeout
            validate_and_send(
                rh, Request(f'http://127.0.0.1:{self.http_port}/timeout_1', extensions={'timeout': 4}))

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_source_address(self, handler):
        source_address = f'127.0.0.{random.randint(5, 255)}'
        with handler(source_address=source_address) as rh:
            data = validate_and_send(
                rh, Request(f'http://127.0.0.1:{self.http_port}/source_address')).read().decode()
            assert source_address == data

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_gzip_trailing_garbage(self, handler):
        with handler() as rh:
            data = validate_and_send(rh, Request(f'http://localhost:{self.http_port}/trailing_garbage')).read().decode()
            assert data == '<html><video src="/vid.mp4" /></html>'

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    @pytest.mark.skipif(not brotli, reason='brotli support is not installed')
    def test_brotli(self, handler):
        with handler() as rh:
            res = validate_and_send(
                rh, Request(
                    f'http://127.0.0.1:{self.http_port}/content-encoding',
                    headers={'ytdl-encoding': 'br'}))
            assert res.headers.get('Content-Encoding') == 'br'
            assert res.read() == b'<html><video src="/vid.mp4" /></html>'

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_deflate(self, handler):
        with handler() as rh:
            res = validate_and_send(
                rh, Request(
                    f'http://127.0.0.1:{self.http_port}/content-encoding',
                    headers={'ytdl-encoding': 'deflate'}))
            assert res.headers.get('Content-Encoding') == 'deflate'
            assert res.read() == b'<html><video src="/vid.mp4" /></html>'

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_gzip(self, handler):
        with handler() as rh:
            res = validate_and_send(
                rh, Request(
                    f'http://127.0.0.1:{self.http_port}/content-encoding',
                    headers={'ytdl-encoding': 'gzip'}))
            assert res.headers.get('Content-Encoding') == 'gzip'
            assert res.read() == b'<html><video src="/vid.mp4" /></html>'

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_multiple_encodings(self, handler):
        with handler() as rh:
            for pair in ('gzip,deflate', 'deflate, gzip', 'gzip, gzip', 'deflate, deflate'):
                res = validate_and_send(
                    rh, Request(
                        f'http://127.0.0.1:{self.http_port}/content-encoding',
                        headers={'ytdl-encoding': pair}))
                assert res.headers.get('Content-Encoding') == pair
                assert res.read() == b'<html><video src="/vid.mp4" /></html>'

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_unsupported_encoding(self, handler):
        with handler() as rh:
            res = validate_and_send(
                rh, Request(
                    f'http://127.0.0.1:{self.http_port}/content-encoding',
                    headers={'ytdl-encoding': 'unsupported'}))
            assert res.headers.get('Content-Encoding') == 'unsupported'
            assert res.read() == b'raw'

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_read(self, handler):
        with handler() as rh:
            res = validate_and_send(
                rh, Request(f'http://127.0.0.1:{self.http_port}/headers'))
            assert res.readable()
            assert res.read(1) == b'H'
            assert res.read(3) == b'ost'


class TestHTTPProxy(TestRequestHandlerBase):
    @classmethod
    def setup_class(cls):
        super().setup_class()
        # HTTP Proxy server
        cls.proxy = http.server.ThreadingHTTPServer(
            ('127.0.0.1', 0), _build_proxy_handler('normal'))
        cls.proxy_port = http_server_port(cls.proxy)
        cls.proxy_thread = threading.Thread(target=cls.proxy.serve_forever)
        cls.proxy_thread.daemon = True
        cls.proxy_thread.start()

        # Geo proxy server
        cls.geo_proxy = http.server.ThreadingHTTPServer(
            ('127.0.0.1', 0), _build_proxy_handler('geo'))
        cls.geo_port = http_server_port(cls.geo_proxy)
        cls.geo_proxy_thread = threading.Thread(target=cls.geo_proxy.serve_forever)
        cls.geo_proxy_thread.daemon = True
        cls.geo_proxy_thread.start()

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_http_proxy(self, handler):
        http_proxy = f'http://127.0.0.1:{self.proxy_port}'
        geo_proxy = f'http://127.0.0.1:{self.geo_port}'

        # Test global http proxy
        # Test per request http proxy
        # Test per request http proxy disables proxy
        url = 'http://foo.com/bar'

        # Global HTTP proxy
        with handler(proxies={'http': http_proxy}) as rh:
            res = validate_and_send(rh, Request(url)).read().decode()
            assert res == f'normal: {url}'

            # Per request proxy overrides global
            res = validate_and_send(rh, Request(url, proxies={'http': geo_proxy})).read().decode()
            assert res == f'geo: {url}'

            # and setting to None disables all proxies for that request
            real_url = f'http://127.0.0.1:{self.http_port}/headers'
            res = validate_and_send(
                rh, Request(real_url, proxies={'http': None})).read().decode()
            assert res != f'normal: {real_url}'
            assert 'Accept' in res

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_noproxy(self, handler):
        with handler(proxies={'proxy': f'http://127.0.0.1:{self.proxy_port}'}) as rh:
            # NO_PROXY
            for no_proxy in (f'127.0.0.1:{self.http_port}', '127.0.0.1', 'localhost'):
                nop_response = validate_and_send(
                    rh, Request(f'http://127.0.0.1:{self.http_port}/headers', proxies={'no': no_proxy})).read().decode(
                    'utf-8')
                assert 'Accept' in nop_response

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_allproxy(self, handler):
        url = 'http://foo.com/bar'
        with handler() as rh:
            response = validate_and_send(rh, Request(url, proxies={'all': f'http://127.0.0.1:{self.proxy_port}'})).read().decode(
                'utf-8')
            assert response == f'normal: {url}'

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_http_proxy_with_idn(self, handler):
        with handler(proxies={
            'http': f'http://127.0.0.1:{self.proxy_port}',
        }) as rh:
            url = 'http://中文.tw/'
            response = rh.send(Request(url)).read().decode()
            # b'xn--fiq228c' is '中文'.encode('idna')
            assert response == 'normal: http://xn--fiq228c.tw/'


class TestClientCertificate:

    @classmethod
    def setup_class(cls):
        certfn = os.path.join(TEST_DIR, 'testcert.pem')
        cls.certdir = os.path.join(TEST_DIR, 'testdata', 'certificate')
        cacertfn = os.path.join(cls.certdir, 'ca.crt')
        cls.httpd = http.server.ThreadingHTTPServer(('127.0.0.1', 0), HTTPTestRequestHandler)
        sslctx = ssl.SSLContext(ssl.PROTOCOL_TLS_SERVER)
        sslctx.verify_mode = ssl.CERT_REQUIRED
        sslctx.load_verify_locations(cafile=cacertfn)
        sslctx.load_cert_chain(certfn, None)
        cls.httpd.socket = sslctx.wrap_socket(cls.httpd.socket, server_side=True)
        cls.port = http_server_port(cls.httpd)
        cls.server_thread = threading.Thread(target=cls.httpd.serve_forever)
        cls.server_thread.daemon = True
        cls.server_thread.start()

    def _run_test(self, handler, **handler_kwargs):
        with handler(
            # Disable client-side validation of unacceptable self-signed testcert.pem
            # The test is of a check on the server side, so unaffected
            verify=False,
            **handler_kwargs,
        ) as rh:
            validate_and_send(rh, Request(f'https://127.0.0.1:{self.port}/video.html')).read().decode()

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_certificate_combined_nopass(self, handler):
        self._run_test(handler, client_cert={
            'client_certificate': os.path.join(self.certdir, 'clientwithkey.crt'),
        })

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_certificate_nocombined_nopass(self, handler):
        self._run_test(handler, client_cert={
            'client_certificate': os.path.join(self.certdir, 'client.crt'),
            'client_certificate_key': os.path.join(self.certdir, 'client.key'),
        })

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_certificate_combined_pass(self, handler):
        self._run_test(handler, client_cert={
            'client_certificate': os.path.join(self.certdir, 'clientwithencryptedkey.crt'),
            'client_certificate_password': 'foobar',
        })

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_certificate_nocombined_pass(self, handler):
        self._run_test(handler, client_cert={
            'client_certificate': os.path.join(self.certdir, 'client.crt'),
            'client_certificate_key': os.path.join(self.certdir, 'clientencrypted.key'),
            'client_certificate_password': 'foobar',
        })


class TestUrllibRequestHandler(TestRequestHandlerBase):
    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_file_urls(self, handler):
        # See https://github.com/ytdl-org/youtube-dl/issues/8227
        tf = tempfile.NamedTemporaryFile(delete=False)
        tf.write(b'foobar')
        tf.close()
        req = Request(pathlib.Path(tf.name).as_uri())
        with handler() as rh:
            with pytest.raises(UnsupportedRequest):
                rh.validate(req)

            # Test that urllib never loaded FileHandler
            with pytest.raises(TransportError):
                rh.send(req)

        with handler(enable_file_urls=True) as rh:
            res = validate_and_send(rh, req)
            assert res.read() == b'foobar'
            res.close()

        os.unlink(tf.name)

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_http_error_returns_content(self, handler):
        # urllib HTTPError will try close the underlying response if reference to the HTTPError object is lost
        def get_response():
            with handler() as rh:
                # headers url
                try:
                    validate_and_send(rh, Request(f'http://127.0.0.1:{self.http_port}/gen_404'))
                except HTTPError as e:
                    return e.response

        assert get_response().read() == b'<html></html>'

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_verify_cert_error_text(self, handler):
        # Check the output of the error message
        with handler() as rh:
            with pytest.raises(
                CertificateVerifyError,
                match=r'\[SSL: CERTIFICATE_VERIFY_FAILED\] certificate verify failed: self.signed certificate'
            ):
                validate_and_send(rh, Request(f'https://127.0.0.1:{self.https_port}/headers'))

    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    @pytest.mark.parametrize('req,match,version_check', [
        # https://github.com/python/cpython/blob/987b712b4aeeece336eed24fcc87a950a756c3e2/Lib/http/client.py#L1256
        # bpo-39603: Check implemented in 3.7.9+, 3.8.5+
        (
            Request('http://127.0.0.1', method='GET\n'),
            'method can\'t contain control characters',
            lambda v: v < (3, 7, 9) or (3, 8, 0) <= v < (3, 8, 5)
        ),
        # https://github.com/python/cpython/blob/987b712b4aeeece336eed24fcc87a950a756c3e2/Lib/http/client.py#L1265
        # bpo-38576: Check implemented in 3.7.8+, 3.8.3+
        (
            Request('http://127.0.0. 1', method='GET'),
            'URL can\'t contain control characters',
            lambda v: v < (3, 7, 8) or (3, 8, 0) <= v < (3, 8, 3)
        ),
        # https://github.com/python/cpython/blob/987b712b4aeeece336eed24fcc87a950a756c3e2/Lib/http/client.py#L1288C31-L1288C50
        (Request('http://127.0.0.1', headers={'foo\n': 'bar'}), 'Invalid header name', None),
    ])
    def test_httplib_validation_errors(self, handler, req, match, version_check):
        if version_check and version_check(sys.version_info):
            pytest.skip(f'Python {sys.version} version does not have the required validation for this test.')

        with handler() as rh:
            with pytest.raises(RequestError, match=match) as exc_info:
                validate_and_send(rh, req)
            assert not isinstance(exc_info.value, TransportError)


def run_validation(handler, error, req, **handler_kwargs):
    with handler(**handler_kwargs) as rh:
        if error:
            with pytest.raises(error):
                rh.validate(req)
        else:
            rh.validate(req)


class TestRequestHandlerValidation:

    class ValidationRH(RequestHandler):
        def _send(self, request):
            raise RequestError('test')

    class NoCheckRH(ValidationRH):
        _SUPPORTED_FEATURES = None
        _SUPPORTED_PROXY_SCHEMES = None
        _SUPPORTED_URL_SCHEMES = None

        def _check_extensions(self, extensions):
            extensions.clear()

    class HTTPSupportedRH(ValidationRH):
        _SUPPORTED_URL_SCHEMES = ('http',)

    URL_SCHEME_TESTS = [
        # scheme, expected to fail, handler kwargs
        ('Urllib', [
            ('http', False, {}),
            ('https', False, {}),
            ('data', False, {}),
            ('ftp', False, {}),
            ('file', UnsupportedRequest, {}),
            ('file', False, {'enable_file_urls': True}),
        ]),
        (NoCheckRH, [('http', False, {})]),
        (ValidationRH, [('http', UnsupportedRequest, {})])
    ]

    PROXY_SCHEME_TESTS = [
        # scheme, expected to fail
        ('Urllib', [
            ('http', False),
            ('https', UnsupportedRequest),
            ('socks4', False),
            ('socks4a', False),
            ('socks5', False),
            ('socks5h', False),
            ('socks', UnsupportedRequest),
        ]),
        (NoCheckRH, [('http', False)]),
        (HTTPSupportedRH, [('http', UnsupportedRequest)]),
    ]

    PROXY_KEY_TESTS = [
        # key, expected to fail
        ('Urllib', [
            ('all', False),
            ('unrelated', False),
        ]),
        (NoCheckRH, [('all', False)]),
        (HTTPSupportedRH, [('all', UnsupportedRequest)]),
        (HTTPSupportedRH, [('no', UnsupportedRequest)]),
    ]

    EXTENSION_TESTS = [
        ('Urllib', [
            ({'cookiejar': 'notacookiejar'}, AssertionError),
            ({'cookiejar': YoutubeDLCookieJar()}, False),
            ({'cookiejar': CookieJar()}, AssertionError),
            ({'timeout': 1}, False),
            ({'timeout': 'notatimeout'}, AssertionError),
            ({'unsupported': 'value'}, UnsupportedRequest),
        ]),
        (NoCheckRH, [
            ({'cookiejar': 'notacookiejar'}, False),
            ({'somerandom': 'test'}, False),  # but any extension is allowed through
        ]),
    ]

    @pytest.mark.parametrize('handler,scheme,fail,handler_kwargs', [
        (handler_tests[0], scheme, fail, handler_kwargs)
        for handler_tests in URL_SCHEME_TESTS
        for scheme, fail, handler_kwargs in handler_tests[1]

    ], indirect=['handler'])
    def test_url_scheme(self, handler, scheme, fail, handler_kwargs):
        run_validation(handler, fail, Request(f'{scheme}://'), **(handler_kwargs or {}))

    @pytest.mark.parametrize('handler,fail', [('Urllib', False)], indirect=['handler'])
    def test_no_proxy(self, handler, fail):
        run_validation(handler, fail, Request('http://', proxies={'no': '127.0.0.1,github.com'}))
        run_validation(handler, fail, Request('http://'), proxies={'no': '127.0.0.1,github.com'})

    @pytest.mark.parametrize('handler,proxy_key,fail', [
        (handler_tests[0], proxy_key, fail)
        for handler_tests in PROXY_KEY_TESTS
        for proxy_key, fail in handler_tests[1]
    ], indirect=['handler'])
    def test_proxy_key(self, handler, proxy_key, fail):
        run_validation(handler, fail, Request('http://', proxies={proxy_key: 'http://example.com'}))
        run_validation(handler, fail, Request('http://'), proxies={proxy_key: 'http://example.com'})

    @pytest.mark.parametrize('handler,scheme,fail', [
        (handler_tests[0], scheme, fail)
        for handler_tests in PROXY_SCHEME_TESTS
        for scheme, fail in handler_tests[1]
    ], indirect=['handler'])
    def test_proxy_scheme(self, handler, scheme, fail):
        run_validation(handler, fail, Request('http://', proxies={'http': f'{scheme}://example.com'}))
        run_validation(handler, fail, Request('http://'), proxies={'http': f'{scheme}://example.com'})

    @pytest.mark.parametrize('handler', ['Urllib', HTTPSupportedRH], indirect=True)
    def test_empty_proxy(self, handler):
        run_validation(handler, False, Request('http://', proxies={'http': None}))
        run_validation(handler, False, Request('http://'), proxies={'http': None})

    @pytest.mark.parametrize('proxy_url', ['//example.com', 'example.com', '127.0.0.1', '/a/b/c'])
    @pytest.mark.parametrize('handler', ['Urllib'], indirect=True)
    def test_invalid_proxy_url(self, handler, proxy_url):
        run_validation(handler, UnsupportedRequest, Request('http://', proxies={'http': proxy_url}))

    @pytest.mark.parametrize('handler,extensions,fail', [
        (handler_tests[0], extensions, fail)
        for handler_tests in EXTENSION_TESTS
        for extensions, fail in handler_tests[1]
    ], indirect=['handler'])
    def test_extension(self, handler, extensions, fail):
        run_validation(
            handler, fail, Request('http://', extensions=extensions))

    def test_invalid_request_type(self):
        rh = self.ValidationRH(logger=FakeLogger())
        for method in (rh.validate, rh.send):
            with pytest.raises(TypeError, match='Expected an instance of Request'):
                method('not a request')


class FakeResponse(Response):
    def __init__(self, request):
        # XXX: we could make request part of standard response interface
        self.request = request
        super().__init__(fp=io.BytesIO(b''), headers={}, url=request.url)


class FakeRH(RequestHandler):

    def _validate(self, request):
        return

    def _send(self, request: Request):
        if request.url.startswith('ssl://'):
            raise SSLError(request.url[len('ssl://'):])
        return FakeResponse(request)


class FakeRHYDL(FakeYDL):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self._request_director = self.build_request_director([FakeRH])


class TestRequestDirector:

    def test_handler_operations(self):
        director = RequestDirector(logger=FakeLogger())
        handler = FakeRH(logger=FakeLogger())
        director.add_handler(handler)
        assert director.handlers.get(FakeRH.RH_KEY) is handler

        # Handler should overwrite
        handler2 = FakeRH(logger=FakeLogger())
        director.add_handler(handler2)
        assert director.handlers.get(FakeRH.RH_KEY) is not handler
        assert director.handlers.get(FakeRH.RH_KEY) is handler2
        assert len(director.handlers) == 1

        class AnotherFakeRH(FakeRH):
            pass
        director.add_handler(AnotherFakeRH(logger=FakeLogger()))
        assert len(director.handlers) == 2
        assert director.handlers.get(AnotherFakeRH.RH_KEY).RH_KEY == AnotherFakeRH.RH_KEY

        director.handlers.pop(FakeRH.RH_KEY, None)
        assert director.handlers.get(FakeRH.RH_KEY) is None
        assert len(director.handlers) == 1

        # RequestErrors should passthrough
        with pytest.raises(SSLError):
            director.send(Request('ssl://something'))

    def test_send(self):
        director = RequestDirector(logger=FakeLogger())
        with pytest.raises(RequestError):
            director.send(Request('any://'))
        director.add_handler(FakeRH(logger=FakeLogger()))
        assert isinstance(director.send(Request('http://')), FakeResponse)

    def test_unsupported_handlers(self):
        class SupportedRH(RequestHandler):
            _SUPPORTED_URL_SCHEMES = ['http']

            def _send(self, request: Request):
                return Response(fp=io.BytesIO(b'supported'), headers={}, url=request.url)

        director = RequestDirector(logger=FakeLogger())
        director.add_handler(SupportedRH(logger=FakeLogger()))
        director.add_handler(FakeRH(logger=FakeLogger()))

        # First should take preference
        assert director.send(Request('http://')).read() == b'supported'
        assert director.send(Request('any://')).read() == b''

        director.handlers.pop(FakeRH.RH_KEY)
        with pytest.raises(NoSupportingHandlers):
            director.send(Request('any://'))

    def test_unexpected_error(self):
        director = RequestDirector(logger=FakeLogger())

        class UnexpectedRH(FakeRH):
            def _send(self, request: Request):
                raise TypeError('something')

        director.add_handler(UnexpectedRH(logger=FakeLogger))
        with pytest.raises(NoSupportingHandlers, match=r'1 unexpected error'):
            director.send(Request('any://'))

        director.handlers.clear()
        assert len(director.handlers) == 0

        # Should not be fatal
        director.add_handler(FakeRH(logger=FakeLogger()))
        director.add_handler(UnexpectedRH(logger=FakeLogger))
        assert director.send(Request('any://'))

    def test_preference(self):
        director = RequestDirector(logger=FakeLogger())
        director.add_handler(FakeRH(logger=FakeLogger()))

        class SomeRH(RequestHandler):
            _SUPPORTED_URL_SCHEMES = ['http']

            def _send(self, request: Request):
                return Response(fp=io.BytesIO(b'supported'), headers={}, url=request.url)

        def some_preference(rh, request):
            return (0 if not isinstance(rh, SomeRH)
                    else 100 if 'prefer' in request.headers
                    else -1)

        director.add_handler(SomeRH(logger=FakeLogger()))
        director.preferences.add(some_preference)

        assert director.send(Request('http://')).read() == b''
        assert director.send(Request('http://', headers={'prefer': '1'})).read() == b'supported'


# XXX: do we want to move this to test_YoutubeDL.py?
class TestYoutubeDLNetworking:

    @staticmethod
    def build_handler(ydl, handler: RequestHandler = FakeRH):
        return ydl.build_request_director([handler]).handlers.get(handler.RH_KEY)

    def test_compat_opener(self):
        with FakeYDL() as ydl:
            with warnings.catch_warnings():
                warnings.simplefilter('ignore', category=DeprecationWarning)
                assert isinstance(ydl._opener, urllib.request.OpenerDirector)

    @pytest.mark.parametrize('proxy,expected', [
        ('http://127.0.0.1:8080', {'all': 'http://127.0.0.1:8080'}),
        ('', {'all': '__noproxy__'}),
        (None, {'http': 'http://127.0.0.1:8081', 'https': 'http://127.0.0.1:8081'})  # env, set https
    ])
    def test_proxy(self, proxy, expected):
        old_http_proxy = os.environ.get('HTTP_PROXY')
        try:
            os.environ['HTTP_PROXY'] = 'http://127.0.0.1:8081'  # ensure that provided proxies override env
            with FakeYDL({'proxy': proxy}) as ydl:
                assert ydl.proxies == expected
        finally:
            if old_http_proxy:
                os.environ['HTTP_PROXY'] = old_http_proxy

    def test_compat_request(self):
        with FakeRHYDL() as ydl:
            assert ydl.urlopen('test://')
            urllib_req = urllib.request.Request('http://foo.bar', data=b'test', method='PUT', headers={'X-Test': '1'})
            urllib_req.add_unredirected_header('Cookie', 'bob=bob')
            urllib_req.timeout = 2
            with warnings.catch_warnings():
                warnings.simplefilter('ignore', category=DeprecationWarning)
                req = ydl.urlopen(urllib_req).request
                assert req.url == urllib_req.get_full_url()
                assert req.data == urllib_req.data
                assert req.method == urllib_req.get_method()
                assert 'X-Test' in req.headers
                assert 'Cookie' in req.headers
                assert req.extensions.get('timeout') == 2

            with pytest.raises(AssertionError):
                ydl.urlopen(None)

    def test_extract_basic_auth(self):
        with FakeRHYDL() as ydl:
            res = ydl.urlopen(Request('http://user:pass@foo.bar'))
            assert res.request.headers['Authorization'] == 'Basic dXNlcjpwYXNz'

    def test_sanitize_url(self):
        with FakeRHYDL() as ydl:
            res = ydl.urlopen(Request('httpss://foo.bar'))
            assert res.request.url == 'https://foo.bar'

    def test_file_urls_error(self):
        # use urllib handler
        with FakeYDL() as ydl:
            with pytest.raises(RequestError, match=r'file:// URLs are disabled by default'):
                ydl.urlopen('file://')

    def test_legacy_server_connect_error(self):
        with FakeRHYDL() as ydl:
            for error in ('UNSAFE_LEGACY_RENEGOTIATION_DISABLED', 'SSLV3_ALERT_HANDSHAKE_FAILURE'):
                with pytest.raises(RequestError, match=r'Try using --legacy-server-connect'):
                    ydl.urlopen(f'ssl://{error}')

            with pytest.raises(SSLError, match='testerror'):
                ydl.urlopen('ssl://testerror')

    @pytest.mark.parametrize('proxy_key,proxy_url,expected', [
        ('http', '__noproxy__', None),
        ('no', '127.0.0.1,foo.bar', '127.0.0.1,foo.bar'),
        ('https', 'example.com', 'http://example.com'),
        ('https', '//example.com', 'http://example.com'),
        ('https', 'socks5://example.com', 'socks5h://example.com'),
        ('http', 'socks://example.com', 'socks4://example.com'),
        ('http', 'socks4://example.com', 'socks4://example.com'),
        ('unrelated', '/bad/proxy', '/bad/proxy'),  # clean_proxies should ignore bad proxies
    ])
    def test_clean_proxy(self, proxy_key, proxy_url, expected):
        # proxies should be cleaned in urlopen()
        with FakeRHYDL() as ydl:
            req = ydl.urlopen(Request('test://', proxies={proxy_key: proxy_url})).request
            assert req.proxies[proxy_key] == expected

        # and should also be cleaned when building the handler
        env_key = f'{proxy_key.upper()}_PROXY'
        old_env_proxy = os.environ.get(env_key)
        try:
            os.environ[env_key] = proxy_url  # ensure that provided proxies override env
            with FakeYDL() as ydl:
                rh = self.build_handler(ydl)
                assert rh.proxies[proxy_key] == expected
        finally:
            if old_env_proxy:
                os.environ[env_key] = old_env_proxy

    def test_clean_proxy_header(self):
        with FakeRHYDL() as ydl:
            req = ydl.urlopen(Request('test://', headers={'ytdl-request-proxy': '//foo.bar'})).request
            assert 'ytdl-request-proxy' not in req.headers
            assert req.proxies == {'all': 'http://foo.bar'}

        with FakeYDL({'http_headers': {'ytdl-request-proxy': '//foo.bar'}}) as ydl:
            rh = self.build_handler(ydl)
            assert 'ytdl-request-proxy' not in rh.headers
            assert rh.proxies == {'all': 'http://foo.bar'}

    def test_clean_header(self):
        with FakeRHYDL() as ydl:
            res = ydl.urlopen(Request('test://', headers={'Youtubedl-no-compression': True}))
            assert 'Youtubedl-no-compression' not in res.request.headers
            assert res.request.headers.get('Accept-Encoding') == 'identity'

        with FakeYDL({'http_headers': {'Youtubedl-no-compression': True}}) as ydl:
            rh = self.build_handler(ydl)
            assert 'Youtubedl-no-compression' not in rh.headers
            assert rh.headers.get('Accept-Encoding') == 'identity'

    def test_build_handler_params(self):
        with FakeYDL({
            'http_headers': {'test': 'testtest'},
            'socket_timeout': 2,
            'proxy': 'http://127.0.0.1:8080',
            'source_address': '127.0.0.45',
            'debug_printtraffic': True,
            'compat_opts': ['no-certifi'],
            'nocheckcertificate': True,
            'legacyserverconnect': True,
        }) as ydl:
            rh = self.build_handler(ydl)
            assert rh.headers.get('test') == 'testtest'
            assert 'Accept' in rh.headers  # ensure std_headers are still there
            assert rh.timeout == 2
            assert rh.proxies.get('all') == 'http://127.0.0.1:8080'
            assert rh.source_address == '127.0.0.45'
            assert rh.verbose is True
            assert rh.prefer_system_certs is True
            assert rh.verify is False
            assert rh.legacy_ssl_support is True

    @pytest.mark.parametrize('ydl_params', [
        {'client_certificate': 'fakecert.crt'},
        {'client_certificate': 'fakecert.crt', 'client_certificate_key': 'fakekey.key'},
        {'client_certificate': 'fakecert.crt', 'client_certificate_key': 'fakekey.key', 'client_certificate_password': 'foobar'},
        {'client_certificate_key': 'fakekey.key', 'client_certificate_password': 'foobar'},
    ])
    def test_client_certificate(self, ydl_params):
        with FakeYDL(ydl_params) as ydl:
            rh = self.build_handler(ydl)
            assert rh._client_cert == ydl_params  # XXX: Too bound to implementation

    def test_urllib_file_urls(self):
        with FakeYDL({'enable_file_urls': False}) as ydl:
            rh = self.build_handler(ydl, UrllibRH)
            assert rh.enable_file_urls is False

        with FakeYDL({'enable_file_urls': True}) as ydl:
            rh = self.build_handler(ydl, UrllibRH)
            assert rh.enable_file_urls is True


class TestRequest:

    def test_query(self):
        req = Request('http://example.com?q=something', query={'v': 'xyz'})
        assert req.url == 'http://example.com?q=something&v=xyz'

        req.update(query={'v': '123'})
        assert req.url == 'http://example.com?q=something&v=123'
        req.update(url='http://example.com', query={'v': 'xyz'})
        assert req.url == 'http://example.com?v=xyz'

    def test_method(self):
        req = Request('http://example.com')
        assert req.method == 'GET'
        req.data = b'test'
        assert req.method == 'POST'
        req.data = None
        assert req.method == 'GET'
        req.data = b'test2'
        req.method = 'PUT'
        assert req.method == 'PUT'
        req.data = None
        assert req.method == 'PUT'
        with pytest.raises(TypeError):
            req.method = 1

    def test_request_helpers(self):
        assert HEADRequest('http://example.com').method == 'HEAD'
        assert PUTRequest('http://example.com').method == 'PUT'

    def test_headers(self):
        req = Request('http://example.com', headers={'tesT': 'test'})
        assert req.headers == HTTPHeaderDict({'test': 'test'})
        req.update(headers={'teSt2': 'test2'})
        assert req.headers == HTTPHeaderDict({'test': 'test', 'test2': 'test2'})

        req.headers = new_headers = HTTPHeaderDict({'test': 'test'})
        assert req.headers == HTTPHeaderDict({'test': 'test'})
        assert req.headers is new_headers

        # test converts dict to case insensitive dict
        req.headers = new_headers = {'test2': 'test2'}
        assert isinstance(req.headers, HTTPHeaderDict)
        assert req.headers is not new_headers

        with pytest.raises(TypeError):
            req.headers = None

    def test_data_type(self):
        req = Request('http://example.com')
        assert req.data is None
        # test bytes is allowed
        req.data = b'test'
        assert req.data == b'test'
        # test iterable of bytes is allowed
        i = [b'test', b'test2']
        req.data = i
        assert req.data == i

        # test file-like object is allowed
        f = io.BytesIO(b'test')
        req.data = f
        assert req.data == f

        # common mistake: test str not allowed
        with pytest.raises(TypeError):
            req.data = 'test'
        assert req.data != 'test'

        # common mistake: test dict is not allowed
        with pytest.raises(TypeError):
            req.data = {'test': 'test'}
        assert req.data != {'test': 'test'}

    def test_content_length_header(self):
        req = Request('http://example.com', headers={'Content-Length': '0'}, data=b'')
        assert req.headers.get('Content-Length') == '0'

        req.data = b'test'
        assert 'Content-Length' not in req.headers

        req = Request('http://example.com', headers={'Content-Length': '10'})
        assert 'Content-Length' not in req.headers

    def test_content_type_header(self):
        req = Request('http://example.com', headers={'Content-Type': 'test'}, data=b'test')
        assert req.headers.get('Content-Type') == 'test'
        req.data = b'test2'
        assert req.headers.get('Content-Type') == 'test'
        req.data = None
        assert 'Content-Type' not in req.headers
        req.data = b'test3'
        assert req.headers.get('Content-Type') == 'application/x-www-form-urlencoded'

    def test_update_req(self):
        req = Request('http://example.com')
        assert req.data is None
        assert req.method == 'GET'
        assert 'Content-Type' not in req.headers
        # Test that zero-byte payloads will be sent
        req.update(data=b'')
        assert req.data == b''
        assert req.method == 'POST'
        assert req.headers.get('Content-Type') == 'application/x-www-form-urlencoded'

    def test_proxies(self):
        req = Request(url='http://example.com', proxies={'http': 'http://127.0.0.1:8080'})
        assert req.proxies == {'http': 'http://127.0.0.1:8080'}

    def test_extensions(self):
        req = Request(url='http://example.com', extensions={'timeout': 2})
        assert req.extensions == {'timeout': 2}

    def test_copy(self):
        req = Request(
            url='http://example.com',
            extensions={'cookiejar': CookieJar()},
            headers={'Accept-Encoding': 'br'},
            proxies={'http': 'http://127.0.0.1'},
            data=[b'123']
        )
        req_copy = req.copy()
        assert req_copy is not req
        assert req_copy.url == req.url
        assert req_copy.headers == req.headers
        assert req_copy.headers is not req.headers
        assert req_copy.proxies == req.proxies
        assert req_copy.proxies is not req.proxies

        # Data is not able to be copied
        assert req_copy.data == req.data
        assert req_copy.data is req.data

        # Shallow copy extensions
        assert req_copy.extensions is not req.extensions
        assert req_copy.extensions['cookiejar'] == req.extensions['cookiejar']

        # Subclasses are copied by default
        class AnotherRequest(Request):
            pass

        req = AnotherRequest(url='http://127.0.0.1')
        assert isinstance(req.copy(), AnotherRequest)

    def test_url(self):
        req = Request(url='https://фtest.example.com/ some spaceв?ä=c',)
        assert req.url == 'https://xn--test-z6d.example.com/%20some%20space%D0%B2?%C3%A4=c'

        assert Request(url='//example.com').url == 'http://example.com'

        with pytest.raises(TypeError):
            Request(url='https://').url = None


class TestResponse:

    @pytest.mark.parametrize('reason,status,expected', [
        ('custom', 200, 'custom'),
        (None, 404, 'Not Found'),  # fallback status
        ('', 403, 'Forbidden'),
        (None, 999, None)
    ])
    def test_reason(self, reason, status, expected):
        res = Response(io.BytesIO(b''), url='test://', headers={}, status=status, reason=reason)
        assert res.reason == expected

    def test_headers(self):
        headers = Message()
        headers.add_header('Test', 'test')
        headers.add_header('Test', 'test2')
        headers.add_header('content-encoding', 'br')
        res = Response(io.BytesIO(b''), headers=headers, url='test://')
        assert res.headers.get_all('test') == ['test', 'test2']
        assert 'Content-Encoding' in res.headers

    def test_get_header(self):
        headers = Message()
        headers.add_header('Set-Cookie', 'cookie1')
        headers.add_header('Set-cookie', 'cookie2')
        headers.add_header('Test', 'test')
        headers.add_header('Test', 'test2')
        res = Response(io.BytesIO(b''), headers=headers, url='test://')
        assert res.get_header('test') == 'test, test2'
        assert res.get_header('set-Cookie') == 'cookie1'
        assert res.get_header('notexist', 'default') == 'default'

    def test_compat(self):
        res = Response(io.BytesIO(b''), url='test://', status=404, headers={'test': 'test'})
        with warnings.catch_warnings():
            warnings.simplefilter('ignore', category=DeprecationWarning)
            assert res.code == res.getcode() == res.status
            assert res.geturl() == res.url
            assert res.info() is res.headers
            assert res.getheader('test') == res.get_header('test')

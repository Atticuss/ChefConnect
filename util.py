from flask import jsonify

from werkzeug.routing import BaseConverter

_conn = None


def init_conn(c):
    global _conn
    _conn = c


def get_conn():
    return _conn


def api_response(resp, err):
    if err == None:
        return jsonify(resp)
    else:
        return jsonify(err)


def init_db():
    return

# https://stackoverflow.com/a/5872904
class RegexConverter(BaseConverter):
    def __init__(self, url_map, *items):
        super(RegexConverter, self).__init__(url_map)
        self.regex = items[0]


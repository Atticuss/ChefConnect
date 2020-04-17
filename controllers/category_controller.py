import util
import json


def get_categories():
    conn = util.get_conn()
    txn = conn.txn(read_only=True)

    api_resp = None
    err = None

    try:
        query = """
            {
                categories(func: has(category_name)) {
                uid
                category_name
                }
            }
        """
        resp = txn.query(query)
        api_resp = json.loads(resp.json)
    finally:
        txn.discard()

    return api_resp, err


def get_category(category_id):
    conn = util.get_conn()
    txn = conn.txn(read_only=True)

    api_resp = None
    err = None

    try:
        variables = {"$id": category_id}
        query = """
            query all($id: string) {
                category(func: uid($id)) {
                    uid
                    category_name
                }
            }
        """

        resp = txn.query(query, variables=variables)
        api_resp = json.loads(resp.json)
    finally:
        txn.discard()

    return api_resp, err


def create_category(data):
    conn = util.get_conn()
    txn = conn.txn()

    api_resp = None
    err = None

    try:
        p = {"category_name": data["name"]}

        resp = txn.mutate(set_obj=p)
        api_resp = {"uid": resp.uids["blank-0"], "category_name": data["name"]}
        txn.commit()
    finally:
        txn.discard()

    return api_resp, err


def update_category(category_id, data):
    err = {"error": "Not implemented"}
    return None, err


def get_category_types():
    conn = util.get_conn()
    txn = conn.txt(read_only=True)

    api_resp = None
    err = None

    #TODO: this query needs to be udpated
    try:
        variables = {"$id": category_id}
        query = """
            query all($id: string) {
                category(func: uid($id)) {
                    uid
                    category_name
                }
            }
        """

        resp = txn.query(query, variables=variables)
        api_resp = json.loads(resp.json)
    finally:
        txn.discard()

    return api_resp, err

def create_category_type(data):
    err = {"error": "Not implemented"}
    return None, err
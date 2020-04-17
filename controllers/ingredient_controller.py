import json
import util


def get_ingredients():
    conn = util.get_conn()
    txn = conn.txn(read_only=True)

    api_resp = None
    err = None

    try:
        query = """
      {
        ingredients(func: has(ingredient_name)) {
          uid
          ingredient_name
        }
      }
    """
        resp = txn.query(query)
        api_resp = json.loads(resp.json)
    finally:
        txn.discard()

    return api_resp, err


def get_ingredient(ingredient_id):
    conn = util.get_conn()
    txn = conn.txn(read_only=True)

    api_resp = None
    err = None

    try:
        variables = {"$id": ingredient_id}
        query = """
            query all($id: string) {
                ingredient(func: uid($id)) {
                    uid
                    ingredient_name
                }
            }
        """

        resp = txn.query(query, variables=variables)
        api_resp = json.loads(resp.json)
    finally:
        txn.discard()

    return api_resp, err


def create_ingredient(data):
    conn = util.get_conn()
    txn = conn.txn()

    api_resp = None
    err = None

    try:
        p = {"ingredient_name": data["name"]}

        if "type" in data.keys():
            p["ingredient_type"] = data["type"]

        resp = txn.mutate(set_obj=p)
        api_resp = {"uid": resp.uids["blank-0"]}
        txn.commit()
    finally:
        txn.discard()

    return api_resp, err


def update_ingredient(data):
    err = {"error": "Not implemented"}
    return None, err

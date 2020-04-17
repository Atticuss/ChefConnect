import json
import util

from urllib.parse import urlparse


def get_recipes():
    conn = util.get_conn()
    txn = conn.txn(read_only=True)

    api_resp = None
    err = None

    try:
        query = """
      {
        recipes(func: has(recipe_name)) {
          uid
          recipe_name
        }
      }
    """
        resp = txn.query(query)
        api_resp = json.loads(resp.json)
    finally:
        txn.discard()

    return api_resp, err


def get_recipe(recipe_id):
    conn = util.get_conn()
    txn = conn.txn(read_only=True)

    api_resp = None
    err = None

    try:
        variables = {"$id": recipe_id}
        query = """
            query all($id: string) {
                recipe(func: uid($id)) {
                    uid
                    recipe_name
                }
            }
        """

        resp = txn.query(query, variables=variables)
        api_resp = json.loads(resp.json)
    finally:
        txn.discard()

    return api_resp, err


def create_recipe(data):
    conn = util.get_conn()
    txn = conn.txn()

    api_resp = None
    err = None

    try:
        p = {
            "recipe_name": data["name"],
            "directions": data["directions"],
            "has_ingredient": [],
        }

        for ingredient in data["ingredients"]:
            p["has_ingredient"].append(
                {"uid": ingredient["id"], "has_ingredient|amount": ingredient["amount"]}
            )

        if "categories" in data.keys():
            p["in_category"] = data["categories"]

        if "url" in data.keys():
            p["url"] = data["url"]
            p["domain"] = urlparse(data["url"]).hostname
        else:
            p["is_homebrewed"] = True

        resp = txn.mutate(set_obj=p)
        api_resp = {"uid": resp.uids["blank-0"]}
        txn.commit()
    finally:
        txn.discard()

    return api_resp, err


def update_recipe(recipe_id, data):
    err = {"error": "Not implemented"}
    return None, err


def search_recipes(data):
    err = {"error": "Not implemented"}
    return None, err


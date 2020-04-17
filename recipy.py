import sqlite3, sys, pydgraph

from flask import (
    Flask,
    jsonify,
    request,
    make_response,
    send_file,
    Response,
    send_from_directory,
    render_template,
)

import util
from controllers import recipe_controller, category_controller, ingredient_controller

app = Flask(__name__)
app.url_map.converters["regex"] = util.RegexConverter

dgraph_server = "127.0.0.1:9080"
client_stub = pydgraph.DgraphClientStub(dgraph_server)
dgraph_client = pydgraph.DgraphClient(client_stub)
util.init_conn(dgraph_client)

"""
Static content
"""


@app.route("/")
def send_index():
    return send_from_directory("static", "index.html")


@app.route("/<path:path>")
def send_static(path):
    print("searching for path: %s" % path)
    if len(path.split(".")) < 2:
        path += ".html"
        return send_from_directory("static", "index.html")
    return send_from_directory("static", path)


"""
Recipe routes
"""


@app.route("/api/recipes", methods=["GET", "POST"])
def recipes():
    data = request.get_json()

    if request.method == "GET":
        resp, err = recipe_controller.get_recipes()
    else:
        resp, err = recipe_controller.create_recipe(data)

    return util.api_response(resp, err)


@app.route("/api/recipes/<regex('0x[a-fA-F0-9]*'):dgid>", methods=["GET", "PUT"])
def recipe(dgid):
    data = request.get_json()

    if request.method == "GET":
        resp, err = recipe_controller.get_recipe(dgid)
    else:
        resp, err = recipe_controller.update_recipe(dgid, data)
    return util.api_response(resp, err)


@app.route("/api/recipes/search", methods=["POST"])
def recipe_search():
    data = request.get_json()
    resp, err = recipe_controller.search_recipes(data)
    return util.api_response(resp, err)


"""
Ingredient routes
"""


@app.route("/api/ingredients", methods=["GET", "POST"])
def ingredients():
    data = request.get_json()

    if request.method == "GET":
        resp, err = ingredient_controller.get_ingredients()
    else:
        resp, err = ingredient_controller.create_ingredient(data)

    return util.api_response(resp, err)


@app.route("/api/ingredients/<regex('0x[a-fA-F0-9]*'):dgid>", methods=["GET", "PUT"])
def ingredient(dgid):
    data = request.get_json()

    if request.method == "GET":
        resp, err = ingredient_controller.get_ingredient(dgid)
    else:
        resp, err = ingredient_controller.update_ingredient(dgid, data)


"""
Category routes
"""


@app.route("/api/categories", methods=["GET", "POST"])
def categories():
    data = request.get_json()

    if request.method == "GET":
        resp, err = category_controller.get_categories()
    else:
        resp, err = category_controller.create_category(data)

    return util.api_response(resp, err)


@app.route("/api/categories/<regex('0x[a-fA-F0-9]*'):dgid>", methods=["GET", "PUT"])
def get_category(dgid):
    data = request.get_json()

    if request.method == "GET":
        resp, err = category_controller.get_category(dgid)
    else:
        resp, err = category_controller.update_category(dgid, data)
    return util.api_response(resp, err)


@app.route("/api/categories/types", methods=["GET", "POST"])
def category_types():
    data = request.get_json()

    if request.method == "GET":
        resp, err = category_controller.get_category_types()
    else:
        resp, err = category_controller.create_category_type(data)

    return util.api_response(resp, err)


"""
Display routes
"""


@app.route("/api/display/view/<recipe_id>", methods=["GET"])
def view_recipe(recipe_id):
    # render full screen template with recipe id passed in
    pass


@app.route("/api/display/set/<recipe_id>", methods=["GET"])
def set_view(recipe_id):
    resp, err = display_controller.set_view(recipe_id)
    return util.api_response(resp, err)


@app.route("/api/display/scroll/<direction>", methods=["GET"])
def scroll(direction):
    pass


if __name__ == "__main__":
    with app.app_context():
        if len(sys.argv) > 1 and sys.argv[1] == "--build-db":
            util.init_db()
            print("[*] Database built")
        else:
            app.run(debug=True, host="0.0.0.0", port=4000)


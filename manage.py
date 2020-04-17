import sys
import json
import pydgraph

dgraph_server = "127.0.0.1:9080"
client_stub = pydgraph.DgraphClientStub(dgraph_server)
dgraph_client = pydgraph.DgraphClient(client_stub)


def build_schema():
    schema = """
    recipe_name: string .
    url: string @index(exact) .
    domain: string @index(exact) .
    directions: string .
    has_ingredient: uid @reverse .
    in_category: uid @reverse .
    category_name: string @index(term) .
    ingredient_name: string @index(term) .
    ingredient_type: string .
    cooking_time: string .
    master_recipe: uid @reverse .
    star_rating: int @index(int) .
    is_favorited: bool @index(bool) .
    has_parent_dir: string @index(exact) .
    directory_name: string @index(exact) .
  """

    op = pydgraph.Operation(schema=schema)
    dgraph_client.alter(op)


def clear():
    op = pydgraph.Operation(drop_all=True)
    dgraph_client.alter(op)


def load_test_data():
    # txn = dgraph_client.txn()
    nquads = """
        _:egg <ingredient_name> "Egg" .
        _:mushroom <ingredient_name> "Mushroom" .
        _:spinach <ingredient_name> "Spinach" .
        _:milk <ingredient_name> "Milk" .
        _:tomato_sauce <ingredient_name> "Tomato Sauce" .
        _:pepper <ingredient_name> "Black Pepper" .
        _:pepper <ingredient_type> "Spice" .
        _:veggie_burger <ingredient_name> "Veggie Burger" .
        _:sliced_bread <ingredient_name> "Sliced Bread" .
        _:breakfast <category_name> "Breakfast" .

        _:egg_scramble <recipe_name> "Egg Scramble" .
        _:egg_scramble <has_ingredient> _:egg (amount="3") .
        _:egg_scramble <has_ingredient> _:mushroom (amount="1/3C") .
        _:egg_scramble <has_ingredient> _:spinach (amount="1/2C") .
        _:egg_scramble <has_ingredient> _:milk (amount="1tbs") .
        _:egg_scramble <has_ingredient> _:tomato_sauce (amount="3tbs") .
        _:egg_scramble <in_category> _:breakfast .
        _:egg_scramble <directions> "Cook mushrooms in a pan over medium heat for five minutes. In a separate bowl, whisk the eggs, milk, and tomato sauce together. When the mushrooms are done, add the eggs. When the eggs are almost fully cooked, add the spinach. Turn the heat down to low and cover to let the spinach wilt." .

        _:egg_burger <recipe_name> "Egg Burger" .
        _:egg_burger <has_ingredient> _:egg (amount="3") .
        _:egg_burger <has_ingredient> _:spinach (amount="handful") .
        _:egg_burger <has_ingredient> _:veggie_burger (amount="1") .
        _:egg_burger <has_ingredient> _:sliced_bread (amount="2") .
        _:egg_burger <in_category> _:breakfast .
        _:egg_burger <directions> "Cook the veggie burger as directed on the packaging. When done, cook the eggs over easy. When done, stack everything together on the bread." .
    """

    for quad in nquads.splitlines():
        quad = quad.strip()

        if len(quad) == 0:
            continue

        txn = dgraph_client.txn()
        res = txn.mutate(set_nquads=quad)
        txn.commit()


if __name__ == "__main__":
    clear()
    build_schema()
    load_test_data()

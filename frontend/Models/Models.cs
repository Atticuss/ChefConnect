using System.Collections.Generic;

namespace frontend.Models{
    public class Ingredient
    {
        public int id { get; set; }

        public string name { get; set; }
    }

    public class QueryIngredients
    {
        public List<Ingredient> ingredients { get; set; }
    }
}


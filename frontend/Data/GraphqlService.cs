using GraphQL;
using GraphQL.Client.Http;
using GraphQL.Client.Serializer.Newtonsoft;
using System.Threading.Tasks;
using System.Collections.Generic;
using frontend.Models;

namespace frontend.Data
{
    public class GraphqlService
    {
        private readonly GraphQL.Client.Http.GraphQLHttpClient _graphqlClient = InitClient();

        // nuking the User-Agent header due to the CORS config of Dgraph Cloud. by default, a
        // custom User-Agent header is set by the `HttpClient`. this causes `user-agent` to be
        // added to the Access-Control-Request-Headers header in the OPTIONS request, but the
        // the server resp does not include it in Access-Control-Allow-Headers.
        private static GraphQL.Client.Http.GraphQLHttpClient InitClient()
        {
            GraphQL.Client.Http.GraphQLHttpClient client = new GraphQLHttpClient("https://blue-surf-510082.us-east-1.aws.cloud.dgraph.io/graphql", new NewtonsoftJsonSerializer());
            client.HttpClient.DefaultRequestHeaders.Remove("User-Agent");
            return client;
        }

        private readonly GraphQLRequest _fetchIngredientsQuery = new GraphQLRequest
        {
            Query = @"
            query FetchIngredients {
                ingredients {
                    name
                    id
                }
            }
        ",
            OperationName = "FetchIngredients"
        };

        //public async Task<ManyIngredients> FetchIngredients()
        public async Task<List<Ingredient>> FetchIngredients()
        {
            var fetchQuery = await _graphqlClient.SendQueryAsync<QueryIngredients>(_fetchIngredientsQuery);
            return fetchQuery.Data.ingredients;
        }

         public async Task InsertIngredient(string ingredientName)
        {
            var createIngredientMutationString = new GraphQLRequest
            {
                Query = @"            
                 mutation AddIngredient($ingredientName : String) {
                   addIngredient(input: {
                       name: $ingredientName
                     }) 
                    {
                      id
                      name
                     }
                 }
             ",
                OperationName = "AddIngredient",
                Variables = new
                {
                    ingredientName
                }
            };

            await _graphqlClient.SendMutationAsync<Ingredient>(createIngredientMutationString);
        }
    }
}